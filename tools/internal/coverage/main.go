package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	coverageXmlFile = "coverage.xml"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

var coverageFiles []string

func filterFiles(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if strings.Contains(path, "coverage.xml") {
		coverageFiles = append(coverageFiles, path)
	}
	return nil
}

func FindCoverageFiles(p string) {
	err := filepath.Walk(".", filterFiles)
	check(err)
}

func main() {
	coverageFiles = make([]string, 0)
	rootPath, err := filepath.Abs(".")
	check(err)
	fmt.Println(rootPath)
	FindCoverageFiles(rootPath)
	fmt.Println(coverageFiles)

	coverageGoal := flag.Float64("coverage-goal", 0.80, "The goal coverage. This script will fail if coverage is below.")
	packagePath := flag.String("package-path", "", "The path to a package from sdk/...")

	flag.Parse()

	if *packagePath == "" {
		fmt.Println("Path was not provided, please provide a path to a package.")
		os.Exit(1)
	}

	fmt.Printf("Checking coverage for package located at %v\n", rootPath)
	fmt.Printf("Failing if the coverage is below %.2f\n", *coverageGoal)

	// Need to add a step to find all coverage files

	// xmlFilePath := filepath.Join(rootPath, *packagePath, coverageXmlFile)
	for _, coverageFiles := range coverageFiles {
		xmlFile, err := os.Open(coverageFiles)
		check(err)
		defer xmlFile.Close()

		byteValue, err := ioutil.ReadAll(xmlFile)
		check(err)

		re := regexp.MustCompile(`<coverage line-rate=\"\d.\d+\"`)
		coverageValue := re.FindAll(byteValue, -1)
		if coverageValue == nil {
			log.Fatalf("Could not match regexp to coverage.xml file.")
		}

		fmt.Printf("%q\n", re.FindAll(byteValue, -1))

		for _, value := range coverageValue {
			parts := strings.Split(string(value), "=")
			coverageNumber := parts[1]

			coverageNumber = coverageNumber[1 : len(coverageNumber)-1]
			coverageFloat, err := strconv.ParseFloat(coverageNumber, 32)
			check(err)

			fmt.Printf("Found a coverage of %.4f for package %v\n", coverageFloat, *packagePath)
			if coverageFloat < *coverageGoal {
				fmt.Printf("Coverage is lower than expected. Got %.4f, expected %.4f\n", coverageFloat, *coverageGoal)
				os.Exit(1)
			}
		}

	}
}
