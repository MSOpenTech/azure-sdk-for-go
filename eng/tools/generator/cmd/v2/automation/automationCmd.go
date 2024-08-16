// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package automation

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/cmd/v2/automation/pipeline"
	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/cmd/v2/common"
	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/repo"
	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/typespec"
	"github.com/Azure/azure-sdk-for-go/eng/tools/internal/utils"
	"github.com/spf13/cobra"
)

// Command returns the automation v2 command. Note that this command is designed to run in the root directory of
// azure-sdk-for-go. It does not work if you are running this tool in somewhere else
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "automation-v2 <generate input filepath> <generate output filepath> [goVersion]",
		Args: cobra.RangeArgs(2, 3),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			log.SetFlags(0)          // remove the time stamp prefix
			log.SetOutput(os.Stdout) // set the output to stdout
			cmd.SetErrPrefix("[ERROR]")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			goVersion := "1.18"
			if len(args) == 3 {
				goVersion = args[2]
			}
			if err := execute(args[0], args[1], goVersion); err != nil {
				return errors.New(logError(err))
			}
			return nil
		},
		SilenceUsage: true, // this command is used for a pipeline, the usage should never show
	}

	return cmd
}

func execute(inputPath, outputPath, goVersion string) error {
	log.Printf("Reading generate input file from '%s'...", inputPath)
	input, err := pipeline.ReadInput(inputPath)
	if err != nil {
		return fmt.Errorf("cannot read generate input: %+v", err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Printf("Using current directory as SDK root: %s", cwd)

	ctx := automationContext{
		sdkRoot:    utils.NormalizePath(cwd),
		specRoot:   input.SpecFolder,
		commitHash: input.HeadSha,
		goVersion:  goVersion,
	}
	output, err := ctx.generate(input)
	if output != nil && len(output.Packages) != 0 {
		log.Printf("Writing output to file '%s'...", outputPath)
		if err := pipeline.WriteOutput(outputPath, output); err != nil {
			return fmt.Errorf("cannot write generate output: %+v", err)
		}
	}
	if err != nil {
		return err
	}

	return nil
}

type automationContext struct {
	sdkRoot    string
	specRoot   string
	commitHash string
	goVersion  string
}

// TODO -- support dry run
func (ctx *automationContext) generate(input *pipeline.GenerateInput) (*pipeline.GenerateOutput, error) {
	if input.DryRun {
		return nil, fmt.Errorf("dry run not supported yet")
	}

	// iterate over all the readme
	results := make([]pipeline.PackageResult, 0)
	errorBuilder := generateErrorBuilder{}

	// create sdk repo ref
	sdkRepo, err := repo.OpenSDKRepository(ctx.sdkRoot)
	if err != nil {
		return nil, fmt.Errorf("failed to get sdk repo: %+v", err)
	}

	// typespec
	// Generated by tsp only when tspconfig.yaml exists and has typespec-go option
	for _, tspProjectFolder := range input.RelatedTypeSpecProjectFolder {
		tspconfigPath := filepath.Join(input.SpecFolder, tspProjectFolder, "tspconfig.yaml")
		tsc, err := typespec.ParseTypeSpecConfig(tspconfigPath)
		if err != nil {
			errorBuilder.add(fmt.Errorf("failed to parse %s: %+v\nInvalid tspconfig.yaml provided and refer to the sample file: https://aka.ms/azsdk/tspconfig-sample-mpg to fix the content", tspconfigPath, err))
			continue
		}

		if ok := tsc.ExistEmitOption(string(typespec.TypeSpec_GO)); ok {
			log.Printf("Start to process typespec project: %s", tspProjectFolder)
			generateCtx := common.GenerateContext{
				SDKPath:        sdkRepo.Root(),
				SDKRepo:        &sdkRepo,
				SpecPath:       ctx.specRoot,
				SpecCommitHash: ctx.commitHash,
				SpecRepoURL:    input.RepoHTTPSURL,
				TypeSpecConfig: tsc,
			}

			module, err := tsc.GetModuleName()
			if err != nil {
				errorBuilder.add(err)
				continue
			}

			namespaceResult, err := generateCtx.GenerateForTypeSpec(&common.GenerateParam{
				RPName:              module[0],
				NamespaceName:       module[1],
				SkipGenerateExample: true,
				GoVersion:           ctx.goVersion,
				TspClientOptions:    []string{"--debug"},
			})
			if err != nil {
				errorBuilder.add(err)
				continue
			} else {
				content := namespaceResult.ChangelogMD
				breaking := namespaceResult.Changelog.HasBreakingChanges()
				breakingChangeItems := namespaceResult.Changelog.GetBreakingChangeItems()

				srcFolder := filepath.Join(sdkRepo.Root(), "sdk", "resourcemanager", namespaceResult.RPName, namespaceResult.PackageName)
				apiViewArtifact := filepath.Join(sdkRepo.Root(), "sdk", "resourcemanager", namespaceResult.RPName, namespaceResult.PackageName+".gosource")
				err := zipDirectory(srcFolder, apiViewArtifact)
				if err != nil {
					fmt.Println(err)
				}

				results = append(results, pipeline.PackageResult{
					Version:         namespaceResult.Version,
					PackageName:     fmt.Sprintf("sdk/resourcemanager/%s/%s", namespaceResult.RPName, namespaceResult.PackageName),
					Path:            []string{fmt.Sprintf("sdk/resourcemanager/%s/%s", namespaceResult.RPName, namespaceResult.PackageName)},
					PackageFolder:   fmt.Sprintf("sdk/resourcemanager/%s/%s", namespaceResult.RPName, namespaceResult.PackageName),
					TypespecProject: []string{tspProjectFolder},
					Changelog: &pipeline.Changelog{
						Content:             &content,
						HasBreakingChange:   &breaking,
						BreakingChangeItems: &breakingChangeItems,
					},
					APIViewArtifact: fmt.Sprintf("sdk/resourcemanager/%s/%s", namespaceResult.RPName, namespaceResult.PackageName+".gosource"),
					Language:        "Go",
				})

				log.Printf("Finish processing typespec file: %s", tspconfigPath)
			}
		} else {
			errorBuilder.add(fmt.Errorf("`@azure-tools/typespec-go` option not found in %s, it is required, please refer to `https://aka.ms/azsdk/tspconfig-sample-mpg` to configure it", tspconfigPath))
		}
	}

	// autorest
	if input.RelatedReadmeMdFile != "" {
		input.RelatedReadmeMdFiles = append(input.RelatedReadmeMdFiles, input.RelatedReadmeMdFile)
	}

	for _, readme := range input.RelatedReadmeMdFiles {
		log.Printf("Start to process autorest project: %s", readme)

		sepStrs := strings.Split(readme, "/")
		for i, sepStr := range sepStrs {
			if sepStr == "resource-manager" {
				readme = strings.Join(sepStrs[i-1:], "/")
				if i > 1 {
					ctx.specRoot = input.SpecFolder + "/" + strings.Join(sepStrs[:i-1], "/")
				}
				break
			}
		}

		generateCtx := common.GenerateContext{
			SDKPath:  sdkRepo.Root(),
			SDKRepo:  &sdkRepo,
			SpecPath: ctx.specRoot,
		}

		namespaceResults, errors := generateCtx.GenerateForAutomation(readme, input.RepoHTTPSURL, ctx.goVersion)
		if len(errors) != 0 {
			errorBuilder.add(errors...)
			continue
		}

		for _, namespaceResult := range namespaceResults {
			content := namespaceResult.ChangelogMD
			breaking := namespaceResult.Changelog.HasBreakingChanges()
			breakingChangeItems := namespaceResult.Changelog.GetBreakingChangeItems()

			srcFolder := filepath.Join(sdkRepo.Root(), "sdk", "resourcemanager", namespaceResult.RPName, namespaceResult.PackageName)
			apiViewArtifact := filepath.Join(sdkRepo.Root(), "sdk", "resourcemanager", namespaceResult.RPName, namespaceResult.PackageName+".gosource")
			err := zipDirectory(srcFolder, apiViewArtifact)
			if err != nil {
				fmt.Println(err)
			}

			results = append(results, pipeline.PackageResult{
				Version:       namespaceResult.Version,
				PackageName:   fmt.Sprintf("sdk/resourcemanager/%s/%s", namespaceResult.RPName, namespaceResult.PackageName),
				Path:          []string{fmt.Sprintf("sdk/resourcemanager/%s/%s", namespaceResult.RPName, namespaceResult.PackageName)},
				PackageFolder: fmt.Sprintf("sdk/resourcemanager/%s/%s", namespaceResult.RPName, namespaceResult.PackageName),
				ReadmeMd:      []string{readme},
				Changelog: &pipeline.Changelog{
					Content:             &content,
					HasBreakingChange:   &breaking,
					BreakingChangeItems: &breakingChangeItems,
				},
				APIViewArtifact: fmt.Sprintf("sdk/resourcemanager/%s/%s", namespaceResult.RPName, namespaceResult.PackageName+".gosource"),
				Language:        "Go",
			})
		}
		log.Printf("Finish processing readme file: %s", readme)
	}

	return &pipeline.GenerateOutput{
		Packages: results,
	}, errorBuilder.build()
}

type generateErrorBuilder struct {
	errors []error
}

func (b *generateErrorBuilder) add(err ...error) {
	b.errors = append(b.errors, err...)
}

func (b *generateErrorBuilder) build() error {
	if len(b.errors) == 0 {
		return nil
	}
	var messages []string
	for _, err := range b.errors {
		messages = append(messages, err.Error())
	}
	return fmt.Errorf("total %d error(s): \n%s\n%s", len(b.errors), strings.Join(messages, "\n"), `Refer to the detail errors for potential fixes.
If the issue persists, contact the Go language support channel at https://aka.ms/azsdk/go-lang-teams-channel and include this spec pull request.`)
}

func logError(err error) string {
	buidler := strings.Builder{}
	for i, line := range strings.Split(err.Error(), "\n") {
		if l := strings.TrimSpace(line); l != "" {
			if i == 0 {
				buidler.WriteString(fmt.Sprintf("%s\n", l))
				continue
			}
			buidler.WriteString(fmt.Sprintf("[ERROR] %s\n", l))
		}
	}

	return buidler.String()
}

func zipDirectory(srcFolder, dstZip string) error {
	outFile, err := os.Create(dstZip)
	if err != nil {
		return err
	}
	w := zip.NewWriter(outFile)
	srcFolder = strings.TrimSuffix(srcFolder, string(os.PathSeparator))
	err = filepath.Walk(srcFolder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Method = zip.Deflate
		header.Name, err = filepath.Rel(filepath.Dir(srcFolder), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += string(os.PathSeparator)
		}
		hw, err := w.CreateHeader(header)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		_, err = io.Copy(hw, f)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	err = outFile.Close()
	if err != nil {
		return err
	}
	return nil
}
