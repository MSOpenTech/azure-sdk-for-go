package main

// To run this package...
// go run gen.go -- --sdk 3.14.16

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	do "gopkg.in/godo.v2"
)

type service struct {
	Name      string
	Fullname  string
	Namespace string
	TaskName  string
	Tag       string
	Input     string
	Output    string
}

const (
	testsSubDir = "tests"
)

type mapping struct {
	PlaneInput  string
	PlaneOutput string
	Services    []service
}

var (
	gopath          = os.Getenv("GOPATH")
	sdkVersion      string
	autorestDir     string
	swaggersDir     string
	testGen         bool
	deps            do.S
	services        = []*service{}
	servicesMapping = []mapping{
		{
			PlaneOutput: "arm",
			PlaneInput:  "resource-manager",
			Services: []service{
				{Name: "advisor"},
				{Name: "analysisservices"},
				// {
				// Autorest Bug
				// Name: "apimanagement",
				// },
				{Name: "appinsights"},
				{Name: "authorization"},
				{Name: "automation"},
				{Name: "batch"},
				{Name: "billing"},
				{Name: "cdn"},
				// {
				// bug in AutoRest (duplicated files)
				// Name: "cognitiveservices",
				// },
				{Name: "commerce"},
				{Name: "compute"},
				{
					Name:  "containerservice",
					Input: "compute",
				},
				{Name: "consumption"},
				{Name: "containerregistry"},
				{
					Name: "customer-insights",
				},
				{
					Name:   "account",
					Input:  "datalake-analytics",
					Output: "datalake-analytics/account",
				},
				{
					Name:   "account",
					Input:  "datalake-store",
					Output: "datalake-store/account",
				},
				{
					Name: "devtestlabs",
				},
				{
					Name: "dns",
				},
				{
					Name: "eventhub",
				},
				{
					Name: "hdinsight",
				},
				{
					Name: "intune",
				},
				{
					Name: "iothub",
				},
				{
					Name: "keyvault",
				},
				{
					Name: "logic",
				},
				{
					Name:   "commitmentplans",
					Input:  "machinelearning",
					Output: "machinelearning/commitmentPlans",
				},
				{
					Name:   "webservices",
					Input:  "machinelearning",
					Output: "machinelearning/webservice",
				},
				{
					Name: "mediaservices",
				},
				{
					Name: "mobileengagement",
				},
				{
					Name: "monitor",
				},
				{
					Name: "network",
				},
				{
					Name: "notificationhubs",
				},
				{
					// bug in the Go generator https://github.com/Azure/autorest/issues/2219
					Name: "operationalinsights",
				},
				{
					Name: "powerbiembedded",
				},
				{
					// bug in the go generator
					Name: "recoveryservices",
				},
				{
					// When using the readme.md, there is an exception in the modeler
					Name: "recoveryservicesbackup",
				},
				{
					Name: "recoveryservicessiterecovery",
				},
				{
					Name: "redis",
				},
				{
					Name: "relay",
				},
				{
					Name: "resourcehealth",
				},
				{
					Name:   "features",
					Input:  "resources",
					Output: "resources/features",
				},
				{
					Name:   "links",
					Input:  "resources",
					Output: "resources/links",
				},
				{
					Name:   "locks",
					Input:  "resources",
					Output: "resources/locks",
				},
				{
					Name:   "managedapplications",
					Input:  "resources",
					Output: "resources/managedapplications",
				},
				{
					Name:   "policy",
					Input:  "resources",
					Output: "resources/policy",
				},
				{
					Name:   "resources",
					Input:  "resources",
					Output: "resources/resources",
				},
				{
					Name:   "subscriptions",
					Input:  "resources",
					Output: "resources/subscriptions",
				},
				{
					Name: "scheduler",
				},
				{
					Name: "search",
				},
				{
					Name: "servermanagement",
				},
				{
					Name: "service-map",
				},
				{
					Name: "servicebus",
				},
				{
					Name: "servicefabric",
				},
				{
					Name: "sql",
				},
				{
					Name: "storage",
				},
				{
					Name: "storageimportexport",
				},
				{
					Name: "storsimple8000series",
				},
				{
					Name: "streamanalytics",
				},
				// {
				// error in the modeler
				// 	Name:    "timeseriesinsights",
				// },
				{
					Name: "trafficmanager",
				},
				{
					Name: "web",
				},
			},
		},
		{
			PlaneOutput: "dataplane",
			PlaneInput:  "data-plane",
			Services: []service{
				{
					Name: "keyvault",
				},
			},
		},
		{
			PlaneInput: "data-plane",
			Services: []service{
				{
					Name:   "datalake-store",
					Output: "datalake-store/filesystem",
				},
			},
		},
		{
			PlaneOutput: "arm",
			PlaneInput:  "data-plane",
			Services: []service{
				{
					Name: "graphrbac",
				},
			},
		},
	}
)

func main() {
	for _, swaggerGroup := range servicesMapping {
		swg := swaggerGroup
		for _, service := range swg.Services {
			s := service
			initAndAddService(&s, swg.PlaneInput, swg.PlaneOutput)
		}
	}
	do.Godo(tasks)
}

func initAndAddService(service *service, planeInput, planeOutput string) {
	if service.Input == "" {
		service.Input = service.Name
	}
	service.Input = filepath.Join(service.Input, planeInput, "readme.md")
	if service.Output == "" {
		service.Output = service.Name
	}
	service.TaskName = fmt.Sprintf("%s>%s", planeOutput, strings.Join(strings.Split(service.Output, "/"), ">"))
	service.Fullname = filepath.Join(planeOutput, service.Output)
	service.Namespace = filepath.Join("github.com", "Azure", "azure-sdk-for-go", service.Fullname)
	service.Output = filepath.Join(gopath, "src", service.Namespace)

	services = append(services, service)
	deps = append(deps, service.TaskName)
}

func tasks(p *do.Project) {
	p.Task("default", do.S{"setvars", "generate:all", "management"}, nil)
	p.Task("setvars", nil, setVars)
	p.Use("generate", generateTasks)
	p.Use("gofmt", formatTasks)
	p.Use("gobuild", buildTasks)
	p.Use("golint", lintTasks)
	p.Use("govet", vetTasks)
	p.Use("delete", deleteTasks)
	p.Task("management", do.S{"setvars"}, managementVersion)
}

func setVars(c *do.Context) {
	if gopath == "" {
		panic("Gopath not set\n")
	}

	sdkVersion = c.Args.MustString("s", "sdk", "version")
	autorestDir = c.Args.MayString("", "a", "ar", "autorest")
	swaggersDir = c.Args.MayString("C:/", "w", "sw", "swagger")
	testGen = c.Args.MayBool(false, "t", "testgen")
}

func generateTasks(p *do.Project) {
	addTasks(generate, p)
}

func generate(service *service) {
	codegen := "--go"
	if testGen {
		codegen = "--go.namespace=TestGen"
		service.Fullname = strings.Join([]string{service.Fullname, testsSubDir}, string(os.PathSeparator))
		service.Output = filepath.Join(service.Output, testsSubDir)
	}

	fmt.Printf("Working on %s...\n\n", service.Fullname)

	delete(service)

	fmt.Printf("Generating on %s...\n\n", service.Fullname)

	execCommand := "autorest"
	commandArgs := []string{
		fmt.Sprintf("https://raw.githubusercontent.com/Azure/azure-rest-api-specs/current/specification/%s", service.Input),
		codegen,
		"--license-header=MICROSOFT_APACHE",
		fmt.Sprintf("--namespace=%s", service.Name),
		fmt.Sprintf("--output-folder=%s", service.Output),
		fmt.Sprintf("--package-version=%s", sdkVersion),
	}
	if testGen {
		commandArgs = append([]string{"-LEGACY"}, commandArgs...)
	}

	// default to the current directory
	workingDir := ""

	if autorestDir != "" {
		// if an AutoRest directory was specified then assume
		// the caller wants to use a locally-built version.
		execCommand = "gulp"
		commandArgs = append([]string{"autorest"}, commandArgs...)
		workingDir = filepath.Join(autorestDir, "autorest")
	}

	autorest := exec.Command(execCommand, commandArgs...)
	autorest.Dir = workingDir

	fmt.Println(commandArgs)

	if err := runner(autorest); err != nil {
		panic(fmt.Errorf("Autorest error: %s", err))
	}

	format(service)
	build(service)
	lint(service)
	vet(service)
}

func deleteTasks(p *do.Project) {
	addTasks(format, p)
}

func delete(service *service) {
	fmt.Printf("Deleting %s...\n\n", service.Fullname)
	if service.Fullname == "" || service.Output == gopath+"/src/github.com/Azure/azure-sdk-for-go" {
		panic(fmt.Sprintf("attempted to remove %s", service.Output))
	}
	err := os.RemoveAll(service.Output)
	if err != nil {
		panic(fmt.Sprintf("Error deleting %s : %s\n", service.Output, err))
	}
}

func formatTasks(p *do.Project) {
	addTasks(format, p)
}

func format(service *service) {
	fmt.Printf("Formatting %s...\n\n", service.Fullname)
	gofmt := exec.Command("gofmt", "-w", service.Output)
	err := runner(gofmt)
	if err != nil {
		panic(fmt.Errorf("gofmt error: %s", err))
	}
}

func buildTasks(p *do.Project) {
	addTasks(build, p)
}

func build(service *service) {
	fmt.Printf("Building %s...\n\n", service.Fullname)
	gobuild := exec.Command("go", "build", service.Namespace)
	err := runner(gobuild)
	if err != nil {
		panic(fmt.Errorf("go build error: %s", err))
	}
}

func lintTasks(p *do.Project) {
	addTasks(lint, p)
}

func lint(service *service) {
	fmt.Printf("Linting %s...\n\n", service.Fullname)
	golint := exec.Command(filepath.Join(gopath, "bin", "golint"), service.Namespace)
	err := runner(golint)
	if err != nil {
		panic(fmt.Errorf("golint error: %s", err))
	}
}

func vetTasks(p *do.Project) {
	addTasks(vet, p)
}

func vet(service *service) {
	fmt.Printf("Vetting %s...\n\n", service.Fullname)
	govet := exec.Command("go", "vet", service.Namespace)
	err := runner(govet)
	if err != nil {
		panic(fmt.Errorf("go vet error: %s", err))
	}
}

func managementVersion(c *do.Context) {
	version("management")
}

func version(packageName string) {
	versionFile := filepath.Join(packageName, "version.go")
	os.Remove(versionFile)
	template := `package %s

var (
	sdkVersion = "%s"
)
`
	data := []byte(fmt.Sprintf(template, packageName, sdkVersion))
	ioutil.WriteFile(versionFile, data, 0644)
}

func addTasks(fn func(*service), p *do.Project) {
	for _, service := range services {
		s := service
		p.Task(s.TaskName, nil, func(c *do.Context) {
			fn(s)
		})
	}
	p.Task("all", deps, nil)
}

func runner(cmd *exec.Cmd) error {
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	if stdout.Len() > 0 {
		fmt.Println(stdout.String())
	}
	if stderr.Len() > 0 {
		fmt.Println(stderr.String())
	}
	return err
}
