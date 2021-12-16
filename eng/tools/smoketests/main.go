package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var smoketestModFile string
var smoketestDir string
var exampleFuncs []string
var envVars []string

func handle(e error) {
	if e != nil {
		panic(e)
	}
}

func getVersion() string {
	v := runtime.Version()
	if strings.Contains(v, "go") {
		v = strings.TrimLeft(v, "go")

		return fmt.Sprintf("go %s", v)
	}

	// Default, go is not from a tag
	return "go 1.17"
}

func inIgnoredDirectories(path string) bool {
	if strings.Contains(path, "internal") {
		return true
	}
	if strings.Contains(path, "samples") {
		return true
	}
	if strings.Contains(path, "smoketests") {
		return true
	}
	if strings.Contains(path, "template") {
		return true
	}

	return false
}

func findModuleDirectories(root string) []string {
	var ret []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		handle(err)
		if strings.Contains(info.Name(), "go.mod") && !inIgnoredDirectories(path) {
			path = strings.ReplaceAll(path, "\\", "/")
			path = strings.ReplaceAll(path, "/go.mod", "")
			parts := strings.Split(path, "/sdk/")
			formatted := fmt.Sprintf("github.com/Azure/azure-sdk-for-go/sdk/%s", parts[1])
			ret = append(ret, formatted)
		}
		return nil
	})
	handle(err)

	return ret
}

func getAllTags() []string {
	result, err := exec.Command("git", "tag", "-l").Output()
	handle(err)
	res := bytes.NewBuffer(result).String()
	return strings.Split(res, "\n")
}

type Module struct {
	Name    string
	Version string
	Replace string
}

type SemVer struct {
	Major, Minor, Patch int
}

func (s SemVer) Newer(s2 SemVer) bool {
	if s.Major > s2.Major {
		return true
	} else if s.Major == s2.Major && s.Minor > s2.Minor {
		return true
	} else if s.Major == s2.Major && s.Minor == s2.Minor && s.Patch > s2.Patch {
		return true
	}
	return false
}

func (s SemVer) String() string {
	return fmt.Sprintf("v%d.%d.%d", s.Major, s.Minor, s.Patch)
}

func toInt(a string) int {
	r, err := strconv.Atoi(a)
	handle(err)
	return r
}

func NewSemVerFromTag(s string) SemVer {
	path := strings.Split(s, "/")
	versionStr := path[len(path)-1]
	versionStr = strings.TrimLeft(versionStr, "v")
	parts := strings.Split(versionStr, ".")
	return SemVer{
		Major: toInt(parts[0]),
		Minor: toInt(parts[1]),
		Patch: toInt(parts[2]),
	}
}

func findLatestTag(p string, tags []string) (string, error) {
	var v SemVer
	for i, tag := range tags {
		if strings.Contains(tag, p) {
			v = NewSemVerFromTag(tag)
			for strings.Contains(tags[i+1], p) {
				newV := NewSemVerFromTag(tags[i+1])
				if newV.Newer(v) {
					v = newV
				}
				i += 1
			}
			return v.String(), nil
		}
	}
	return "", fmt.Errorf("could not find a version for module %s", p)
}

func matchModulesAndTags(goModFiles []string, tags []string) []Module {
	var m []Module

	for _, goModFile := range goModFiles {
		packagePath := strings.Split(goModFile, "github.com/Azure/azure-sdk-for-go/sdk/")
		relativePackagePath := packagePath[1]
		version, err := findLatestTag(relativePackagePath, tags)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			m = append(m, Module{
				Name:    goModFile,
				Replace: fmt.Sprintf("../%s", relativePackagePath),
				Version: version,
			})
		}
	}

	return m
}

// GetTopLevel runs "git rev-parse --show-toplevel" to get the an absolute path to the current repo
func GetTopLevel() string {
	topLevel, err := exec.Command("git", "rev-parse", "--show-toplevel").CombinedOutput()
	handle(err)
	return strings.ReplaceAll(bytes.NewBuffer(topLevel).String(), "\n", "")
}

// BuildModFile creates a go.mod file and adds replace directives for the appropriate modules.
// If serviceDirectory is a blank string it replaces all modules, otherwise it only replaces matching modules
func BuildModFile(modules []Module, serviceDirectory string) error {
	f, err := os.OpenFile(smoketestModFile, os.O_RDWR, 0666)
	handle(err)
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/sdk/smoketests\n\n%s\n\n", getVersion()))
	if err != nil {
		return err
	}

	replaceString := "replace %s => %s\n"
	if serviceDirectory == "notset" {
		for _, module := range modules {
			s := fmt.Sprintf(replaceString, module.Name, module.Replace)
			_, err = f.Write([]byte(s))
			handle(err)
		}
	} else {
		fmt.Printf("Replace directive for %s\n", serviceDirectory)
		for _, module := range modules {
			if strings.Contains(module.Name, serviceDirectory) {
				s := fmt.Sprintf(replaceString, module.Name, module.Replace)
				_, err = f.Write([]byte(s))
				handle(err)
			}
		}
	}

	_, err = f.WriteString("\n\nrequire (\n")

	if err != nil {
		return err
	}

	requireString := "\t%s %s\n"
	for _, module := range modules {
		s := fmt.Sprintf(requireString, module.Name, module.Version)
		_, err = f.Write([]byte(s))
		handle(err)
	}

	_, err = f.WriteString(")")
	handle(err)
	return nil
}

// FindExampleFiles finds all files that are named "example_*.go".
// If serviceDirectory
func FindExampleFiles(root, serviceDirectory string) ([]string, error) {
	var ret []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		handle(err)
		if strings.HasPrefix(info.Name(), "example_") && !inIgnoredDirectories(path) && strings.HasSuffix(info.Name(), ".go") {
			path = strings.ReplaceAll(path, "\\", "/")
			if serviceDirectory == "" || strings.Contains(path, serviceDirectory) {
				ret = append(ret, path)
			}

		}
		return nil
	})

	return ret, err
}

// copyFile copies the contents from src to dest. Creating the dest file first
func copyFile(src, dest string) error {
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	stringData := bytes.NewBuffer(data).String()

	functionRegex := regexp.MustCompile(`(?m)^func\sExampleNew(.*)\(\)\s{`)
	functions := functionRegex.FindAllString(stringData, -1)
	for i, l := range functions {
		l = strings.ReplaceAll(l, "()", "")
		parts := strings.Split(l, " ")
		functions[i] = parts[1]
	}

	// do a name change for clashing
	var newNames []string
	for _, f := range functions {
		newNames = append(newNames, fmt.Sprintf("%s%d", f, rand.Intn(100000)))
	}

	for i := 0; i < len(functions); i++ {
		stringData = strings.Replace(stringData, functions[i], newNames[i], 1)
		exampleFuncs = append(exampleFuncs, newNames[i])
	}

	err = ioutil.WriteFile(dest, []byte(stringData), 0644)
	return err
}

// CopyExampleFiles copies all the example files to the destination directory.
// This creates a hash of the fileName for the destination path
func CopyExampleFiles(exFiles []string, dest string) error {
	fmt.Printf("Copying %d example files to %s\n", len(exFiles), dest)

	for _, exFile := range exFiles {
		newFileName := strings.ReplaceAll(exFile[10:], "/", "_")
		newFileName = strings.ReplaceAll(newFileName, " ", "")
		destinationPath := filepath.Join(dest, fmt.Sprintf("%s.go", newFileName))

		err := copyFile(exFile, destinationPath)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReplacePackageStatement replaces all "package ***" with a common "package main" statement
func ReplacePackageStatement(root string) error {
	fmt.Println("Fixing package names in", root)
	packageName := "package main"
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".go") {
			handle(err)
			data, err := ioutil.ReadFile(path)
			handle(err)

			datastring := bytes.NewBuffer(data).String()

			m := regexp.MustCompile("(?m)^package (.*)$")
			datastring = m.ReplaceAllString(datastring, packageName)

			err = ioutil.WriteFile(path, []byte(datastring), 0666)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func BuildMainFile(root string, c ConfigFile) error {
	mainFile := filepath.Join(root, "main.go")
	f, err := os.Create(mainFile)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}

	// Write the main.go file

	src := "package main\nfunc main() {"

	fmt.Println(envVars)
	for _, envVar := range envVars {
		src += fmt.Sprintf(`os.Setenv("%s", "%s")`, envVar, FindEnvVarFromConfig(c, envVar))
		src += "\n"
	}

	for _, exampleFunc := range exampleFuncs {
		src += fmt.Sprintf("%s()\n", exampleFunc)
	}

	src += "}"

	err = ioutil.WriteFile(mainFile, []byte(src), 0666)
	return err
}

func FindEnvVars(root string) error {
	fmt.Println("Find all environment variables using `os.Getenv` or `os.LookupEnv`")

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
			// Find Env Vars
			searchFile(path)
		}
		return nil
	})
	return err
}

func searchFile(path string) error {
	fmt.Println("searching file")
	var envVarRegex *regexp.Regexp
	if runtime.GOOS == "windows" {
		envVarRegex = regexp.MustCompile(`(?m)os.LookupEnv(.*)\r\n`)
	} else {
		envVarRegex = regexp.MustCompile(`(?m)os.LookupEnv(.*)\n`)
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	stringData := bytes.NewBuffer(data).String()

	LookupEnvs := envVarRegex.FindAllString(stringData, -1)
	envVars = append(envVars, trimLookupEnvs(LookupEnvs)...)

	if runtime.GOOS == "windows" {
		envVarRegex = regexp.MustCompile(`(?m)os.Getenv(.*)\r\n`)
	} else {
		envVarRegex = regexp.MustCompile(`(?m)os.Getenv(.*)\n`)
	}

	Getenvs := envVarRegex.FindAllString(stringData, -1)
	envVars = append(envVars, trimGetenvs(Getenvs)...)
	return nil
}

func FindEnvVarFromConfig(c ConfigFile, envVar string) string {
	for _, p := range c.Packages {
		for key, value := range p.EnvironmentVariables {
			if key == envVar {
				return value
			}
		}
	}

	return ""
}

func trimLookupEnvs(values []string) []string {
	pseudoSet := make(map[string]struct{})

	for _, value := range values {
		value = strings.TrimSpace(value)
		value = strings.TrimPrefix(value, `os.LookupEnv("`)
		value = strings.TrimSuffix(value, `")`)
		pseudoSet[value] = struct{}{}
	}

	var ret []string
	for v := range pseudoSet {
		ret = append(ret, v)
	}

	return ret
}

func trimGetenvs(values []string) []string {
	var ret []string

	for _, value := range values {
		value = strings.TrimPrefix(value, `os.Getenv("`)
		value = strings.TrimSuffix(value, `")`)
		ret = append(ret, value)
	}

	return ret
}

func LoadEngConfig(rootDirectory string) ConfigFile {
	fileName := filepath.Join(rootDirectory, "eng", "config.json")
	buffer, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var p ConfigFile
	err = json.Unmarshal(buffer, &p)
	handle(err)
	return p
}

func main() {
	serviceDirectory := flag.String("serviceDirectory", "", "pass in a single service directory for nightly run")
	flag.Parse()

	fmt.Println("Running smoketest")

	rootDirectory := GetTopLevel()

	configFile := LoadEngConfig(rootDirectory)
	fmt.Println(configFile)

	absSDKPath, err := filepath.Abs(fmt.Sprintf("%s/sdk", rootDirectory))
	handle(err)
	fmt.Println("Root SDK directory: ", absSDKPath)

	smoketestDir = filepath.Join(absSDKPath, "smoketests")
	fmt.Println("Smoke test directory: ", smoketestDir)

	smoketestModFile = filepath.Join(smoketestDir, "go.mod")

	exampleFiles, err := FindExampleFiles(absSDKPath, *serviceDirectory)
	handle(err)
	fmt.Printf("Found %d example files for smoke tests\n", len(exampleFiles))
	for _, e := range exampleFiles {
		fmt.Println(e)
	}

	moduleDirectories := findModuleDirectories(absSDKPath)
	fmt.Printf("Found %d modules\n", len(moduleDirectories))

	allTags := getAllTags()
	fmt.Printf("Found %d tags\n", len(allTags))

	modules := matchModulesAndTags(moduleDirectories, allTags)
	_ = modules

	err = CopyExampleFiles(exampleFiles, smoketestDir)
	handle(err)

	err = BuildModFile(modules, *serviceDirectory)
	handle(err)

	err = ReplacePackageStatement(smoketestDir)
	handle(err)

	err = FindEnvVars(smoketestDir)
	handle(err)

	err = BuildMainFile(smoketestDir, configFile)
	handle(err)

}
