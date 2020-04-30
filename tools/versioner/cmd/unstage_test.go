// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/tools/internal/files"
)

const versionGoFormat = `package foo

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
	return "Azure-SDK-For-Go/" + Version() + " foo/2019-04-01"
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
	return "%s"
}

// tag: %s
`

func Test_findLatestMajorVersion(t *testing.T) {
	testData := []struct {
		stage    string
		expected string
	}{
		{
			stage:    "../../testdata/scenarioa/foo/stage",
			expected: filepath.Join("..", "..", "testdata", "scenarioa", "foo"),
		},
		{
			stage:    "../../testdata/scenariod/foo/stage",
			expected: filepath.Join("..", "..", "testdata", "scenariod", "foo"),
		},
		{
			stage:    "../../testdata/scenariof/foo/stage",
			expected: filepath.Join("..", "..", "testdata", "scenariof", "foo"),
		},
	}

	for _, c := range testData {
		t.Logf("Testing '%s'", c.stage)
		lmv, err := findLatestMajorVersion(c.stage)
		if err != nil {
			t.Fatalf("unexpected error: %+v", err)
		}
		if lmv != c.expected {
			t.Fatalf("expected %s but got %s", c.expected, lmv)
		}
	}
}

type byteBufferSeeker struct {
	buf *bytes.Buffer
}

func (b byteBufferSeeker) Read(p []byte) (int, error) {
	return b.buf.Read(p)
}

func (b byteBufferSeeker) Write(p []byte) (int, error) {
	return b.buf.Write(p)
}

func (b byteBufferSeeker) Seek(offset int64, whence int) (int64, error) {
	if offset != 0 && whence != 0 {
		panic("seek only supports 0, 0")
	}
	b.buf.Reset()
	return 0, nil
}

func TestUpdateGoMod(t *testing.T) {
	testData := []struct {
		name     string
		content  string
		ver      string
		expected string
	}{
		{
			name:     "from v1 to v1",
			content:  fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/services/foo/mgmt/2019-05-01/foo\n\n%s\n", goVersion),
			ver:      "",
			expected: fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/services/foo/mgmt/2019-05-01/foo\n\n%s\n", goVersion),
		},
		{
			name:     "from v1 to v2",
			content:  fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/services/foo/mgmt/2019-05-01/foo\n\n%s\n", goVersion),
			ver:      "v2",
			expected: fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/services/foo/mgmt/2019-05-01/foo/v2\n\n%s\n", goVersion),
		},
		{
			name:     "from v2 to v2",
			content:  fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/services/foo/mgmt/2019-05-01/foo/v2\n\n%s\n", goVersion),
			ver:      "v2",
			expected: fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/services/foo/mgmt/2019-05-01/foo/v2\n\n%s\n", goVersion),
		},
		{
			name:     "from v2 to v3",
			content:  fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/services/foo/mgmt/2019-05-01/foo/v2\n\n%s\n", goVersion),
			ver:      "v3",
			expected: fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/services/foo/mgmt/2019-05-01/foo/v3\n\n%s\n", goVersion),
		},
	}

	for _, c := range testData {
		t.Logf("Testing %s", c.name)
		buf := byteBufferSeeker{
			buf: bytes.NewBuffer([]byte(c.content)),
		}
		err := updateGoMod(buf, c.ver)
		if err != nil {
			t.Fatalf("updateGoMod failed: %+v", err)
		}
		if !strings.EqualFold(buf.buf.String(), c.expected) {
			t.Fatalf("expected '%s' but got '%s'", c.expected, buf.buf.String())
		}
	}
}

func verifyVersion(t *testing.T, path, version, tag string) {
	b, err := ioutil.ReadFile(filepath.Join(path, versionFilename))
	if err != nil {
		t.Fatalf("failed to read version.go file: %v", err)
	}
	expected := fmt.Sprintf(versionGoFormat, version, tag)
	if !fileContentEquals(expected, string(b)) {
		t.Fatalf("bad version.go file, expected '%s' got '%s'", expected, string(b))
	}
}

func verifyGoMod(t *testing.T, path, expected string) {
	b, err := ioutil.ReadFile(filepath.Join(path, goModFilename))
	if err != nil {
		t.Fatalf("failed to read go.mod file: %v", err)
	}
	if !fileContentEquals(expected, string(b)) {
		t.Fatalf("bad go.mod file, expected '%s' got '%s'", expected, string(b))
	}
}

func verifyChangelog(t *testing.T, path string) {
	if exists, err := files.Exists(filepath.Join(path, "CHANGELOG.md")); err != nil {
		t.Fatalf("unexpected error: %+v", err)
	} else if !exists {
		t.Fatalf("expected changelog in %s", path)
	}
}

func verifyNoChangelog(t *testing.T, path string) {
	if exists, err := files.Exists(filepath.Join(path, "CHANGELOG.md")); err != nil {
		t.Fatalf("unexpected error: %+v", err)
	} else if exists {
		t.Fatalf("unexpected changelog in %s", path)
	}
}

func verifyGoVet(t *testing.T, root string) {
	root, err := filepath.Abs(root)
	if err != nil {
		t.Fatalf("failed to get absolute path: %v", err)
	}
	folders := make([]string, 0)
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			pkg := path[strings.Index(path, "github.com"):]
			folders = append(folders, pkg)
			return nil
		}
		return nil
	})
	if err != nil {
		t.Fatalf("failed to list all sub-folders in root '%s': %v", root, err)
	}
	for _, folder := range folders {
		c := exec.Command("go", "vet", folder)
		if output, err := c.CombinedOutput(); err != nil {
			t.Fatalf("vet failed: %s", string(output))
		}
	}
}

func fileContentEquals(expected, content string) bool {
	replacedContent := strings.ReplaceAll(content, "\r\n", "\n")
	return strings.EqualFold(replacedContent, expected)
}

var defaultVersionSetting = &VersionSetting{
	InitialVersion:        startingModVer,
	InitialVersionPreview: startingModVerPreview,
}

func TestExecuteUnstage(t *testing.T) {
	type expected struct {
		dest      string
		tag       string
		mod       string
		version   string
		changelog bool
	}
	testData := []struct {
		name        string
		getTagsHook TagsHookFunc
		stage       string
		expected
	}{
		{
			name: "scenario a: minor version v1",
			getTagsHook: func(root, prefix string) ([]string, error) {
				// root doesn't matter
				if !strings.HasSuffix(prefix, "/testdata/scenarioa/foo") {
					return nil, fmt.Errorf("bad prefix '%s'", prefix)
				}
				return []string{
					"tools/testdata/scenarioa/foo/v1.0.0",
				}, nil
			},
			stage: "../../testdata/scenarioa/foo/stage",
			expected: expected{
				dest:      "../../testdata/scenarioa/foo",
				tag:       "tools/testdata/scenarioa/foo/v1.1.0",
				mod:       fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenarioa/foo\n\n%s\n", goVersion),
				version:   "v1.1.0",
				changelog: true,
			},
		},
		{
			name: "scenario b: major version v2",
			getTagsHook: func(root, prefix string) ([]string, error) {
				// root doesn't matter
				if !strings.HasSuffix(prefix, "/testdata/scenariob/foo") {
					return nil, fmt.Errorf("bad prefix '%s'", prefix)
				}
				return []string{
					"tools/testdata/scenariob/foo/v1.0.0",
					"tools/testdata/scenariob/foo/v1.1.0",
				}, nil
			},
			stage: "../../testdata/scenariob/foo/stage",
			expected: expected{
				dest:      "../../testdata/scenariob/foo",
				tag:       "tools/testdata/scenariob/foo/v2.0.0",
				mod:       fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenariob/foo/v2\n\n%s\n", goVersion),
				version:   "v2.0.0",
				changelog: true,
			},
		},
		{
			name: "scenario c: patch version v1",
			getTagsHook: func(root, prefix string) ([]string, error) {
				// root doesn't matter
				if !strings.HasSuffix(prefix, "/testdata/scenarioc/foo") {
					return nil, fmt.Errorf("bad prefix '%s'", prefix)
				}
				return []string{
					"tools/testdata/scenarioc/foo/v1.0.0",
				}, nil
			},
			stage: "../../testdata/scenarioc/foo/stage",
			expected: expected{
				dest:      "../../testdata/scenarioc/foo",
				tag:       "tools/testdata/scenarioc/foo/v1.0.1",
				mod:       fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenarioc/foo\n\n%s\n", goVersion),
				version:   "v1.0.1",
				changelog: true,
			},
		},
		{
			name: "scenario d: major version v3",
			getTagsHook: func(root, prefix string) ([]string, error) {
				if !strings.HasSuffix(prefix, "/testdata/scenariod/foo") {
					return nil, fmt.Errorf("bad prefix '%s'", prefix)
				}
				return []string{
					"tools/testdata/scenariod/foo/v1.0.0",
					"tools/testdata/scenariod/foo/v1.0.1",
					"tools/testdata/scenariod/foo/v1.1.0",
					"tools/testdata/scenariod/foo/v1.2.0",
					"tools/testdata/scenariod/foo/v2.0.0",
					"tools/testdata/scenariod/foo/v2.1.0",
					"tools/testdata/scenariod/foo/v2.1.1",
				}, nil
			},
			stage: "../../testdata/scenariod/foo/stage",
			expected: expected{
				dest:      "../../testdata/scenariod/foo",
				tag:       "tools/testdata/scenariod/foo/v3.0.0",
				mod:       fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenariod/foo/v3\n\n%s\n", goVersion),
				version:   "v3.0.0",
				changelog: true,
			},
		},
		{
			name: "scenario e: update from v2.1.0 to v2.2.0",
			getTagsHook: func(root, prefix string) ([]string, error) {
				// root doesn't matter
				if !strings.HasSuffix(prefix, "/testdata/scenarioe/foo") {
					return nil, fmt.Errorf("bad prefix '%s'", prefix)
				}
				return []string{
					"tools/testdata/scenarioe/foo/v1.0.0",
					"tools/testdata/scenarioe/foo/v1.1.0",
					"tools/testdata/scenarioe/foo/v2.0.0",
					"tools/testdata/scenarioe/foo/v2.1.0",
				}, nil
			},
			stage: "../../testdata/scenarioe/foo/stage",
			expected: expected{
				dest:      "../../testdata/scenarioe/foo",
				tag:       "tools/testdata/scenarioe/foo/v2.2.0",
				mod:       fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenarioe/foo/v2\n\n%s\n", goVersion),
				version:   "v2.2.0",
				changelog: true,
			},
		},
		{
			name: "scenario f: new module",
			getTagsHook: func(root, prefix string) ([]string, error) {
				// root doesn't matter
				if !strings.HasSuffix(prefix, "/testdata/scenariof/foo") {
					return nil, fmt.Errorf("bad prefix '%s'", prefix)
				}
				return []string{}, nil
			},
			stage: "../../testdata/scenariof/foo/stage",
			expected: expected{
				dest:      "../../testdata/scenariof/foo",
				tag:       "tools/testdata/scenariof/foo/v1.0.0",
				mod:       fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenariof/foo\n\n%s\n", goVersion),
				version:   "v1.0.0",
				changelog: false,
			},
		},
		{
			name: "scenario g: new management plane major v2",
			getTagsHook: func(root string, prefix string) ([]string, error) {
				if !strings.HasSuffix(prefix, "testdata/scenariog/foo/mgmt/2019-10-11/foo") {
					return nil, fmt.Errorf("bad prefix '%s'", prefix)
				}
				return []string{
					"tools/testdata/scenariog/foo/mgmt/2019-10-11/foo/v1.0.0",
				}, nil
			},
			stage: "../../testdata/scenariog/foo/mgmt/2019-10-11/foo/stage",
			expected: expected{
				dest:      "../../testdata/scenariog/foo/mgmt/2019-10-11/foo",
				tag:       "tools/testdata/scenariog/foo/mgmt/2019-10-11/foo/v2.0.0",
				mod:       fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenariog/foo/mgmt/2019-10-11/foo/v2\n\n%s\n", goVersion),
				version:   "v2.0.0",
				changelog: true,
			},
		},
		{
			name: "scenario h: major version v2 (import need changes)",
			getTagsHook: func(root string, prefix string) ([]string, error) {
				if !strings.HasSuffix(prefix, "testdata/scenarioh/foo/mgmt/2019-10-11/foo") {
					return nil, fmt.Errorf("bad prefix '%s'", prefix)
				}
				return []string{
					"tools/testdata/scenarioh/foo/mgmt/2019-10-11/foo/v1.0.0",
					"tools/testdata/scenarioh/foo/mgmt/2019-10-11/foo/v1.1.0",
					"tools/testdata/scenarioh/foo/mgmt/2019-10-11/foo/v1.2.0",
				}, nil
			},
			stage: "../../testdata/scenarioh/foo/mgmt/2019-10-11/foo/stage",
			expected: expected{
				dest:      "../../testdata/scenarioh/foo/mgmt/2019-10-11/foo",
				tag:       "tools/testdata/scenarioh/foo/mgmt/2019-10-11/foo/v2.0.0",
				mod:       fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenarioh/foo/mgmt/2019-10-11/foo/v2\n\n%s\n", goVersion),
				version:   "v2.0.0",
				changelog: true,
			},
		},
		{
			name: "scenario i: identical, should have no update",
			getTagsHook: func(root string, prefix string) ([]string, error) {
				if !strings.HasSuffix(prefix, "testdata/scenarioi/foo/mgmt/2019-10-23/foo") {
					return nil, fmt.Errorf("bad prefix '%s'", prefix)
				}
				return []string{
					"tools/testdata/scenarioi/foo/mgmt/2019-10-23/foo/v1.0.0",
					"tools/testdata/scenarioi/foo/mgmt/2019-10-23/foo/v1.0.1",
					"tools/testdata/scenarioi/foo/mgmt/2019-10-23/foo/v1.1.0",
					"tools/testdata/scenarioi/foo/mgmt/2019-10-23/foo/v1.1.1",
					"tools/testdata/scenarioi/foo/mgmt/2019-10-23/foo/v1.1.2",
				}, nil
			},
			stage: "../../testdata/scenarioi/foo/mgmt/2019-10-23/foo/stage",
			expected: expected{
				dest:      "../../testdata/scenarioi/foo/mgmt/2019-10-23/foo",
				tag:       "tools/testdata/scenarioi/foo/mgmt/2019-10-23/foo/v1.1.2",
				mod:       fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenarioi/foo/mgmt/2019-10-23/foo\n\n%s\n", goVersion),
				version:   "v1.1.2",
				changelog: false,
			},
		},
		{
			name: "scenario j: identical v2, should have no update",
			getTagsHook: func(root string, prefix string) ([]string, error) {
				if !strings.HasSuffix(prefix, "testdata/scenarioj/foo/mgmt/2019-10-23/foo") {
					return nil, fmt.Errorf("bad prefix '%s'", prefix)
				}
				return []string{
					"tools/testdata/scenarioj/foo/mgmt/2019-10-23/foo/v1.0.0",
					"tools/testdata/scenarioj/foo/mgmt/2019-10-23/foo/v1.0.1",
					"tools/testdata/scenarioj/foo/mgmt/2019-10-23/foo/v1.1.0",
					"tools/testdata/scenarioj/foo/mgmt/2019-10-23/foo/v1.1.1",
					"tools/testdata/scenarioj/foo/mgmt/2019-10-23/foo/v1.1.2",
					"tools/testdata/scenarioj/foo/mgmt/2019-10-23/foo/v2.0.0",
				}, nil
			},
			stage: "../../testdata/scenarioj/foo/mgmt/2019-10-23/foo/stage",
			expected: expected{
				dest:      "../../testdata/scenarioj/foo/mgmt/2019-10-23/foo",
				tag:       "tools/testdata/scenarioj/foo/mgmt/2019-10-23/foo/v2.0.0",
				mod:       fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenarioj/foo/mgmt/2019-10-23/foo/v2\n\n%s\n", goVersion),
				version:   "v2.0.0",
				changelog: false,
			},
		},
		{
			name: "scenario k: preview with addition changes",
			getTagsHook: func(root string, prefix string) ([]string, error) {
				if !strings.HasSuffix(prefix, "testdata/scenariok/foo/mgmt/2019-11-01-preview/foo") {
					return nil, fmt.Errorf("bad prefix '%s'", prefix)
				}
				return []string{
					"tools/testdata/scenariok/foo/mgmt/2019-11-01-preview/foo/v0.0.0",
					"tools/testdata/scenariok/foo/mgmt/2019-11-01-preview/foo/v0.0.1",
					"tools/testdata/scenariok/foo/mgmt/2019-11-01-preview/foo/v0.1.0",
					"tools/testdata/scenariok/foo/mgmt/2019-11-01-preview/foo/v0.1.1",
					"tools/testdata/scenariok/foo/mgmt/2019-11-01-preview/foo/v0.1.2",
				}, nil
			},
			stage: "../../testdata/scenariok/foo/mgmt/2019-11-01-preview/foo/stage",
			expected: expected{
				dest:      "../../testdata/scenariok/foo/mgmt/2019-11-01-preview/foo",
				tag:       "tools/testdata/scenariok/foo/mgmt/2019-11-01-preview/foo/v0.2.0",
				mod:       fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenariok/foo/mgmt/2019-11-01-preview/foo\n\n%s\n", goVersion),
				version:   "v0.2.0",
				changelog: true,
			},
		},
		{
			name: "scenario l: breaking changes in preview package",
			getTagsHook: func(root string, prefix string) ([]string, error) {
				if !strings.HasSuffix(prefix, "testdata/scenariol/foo/mgmt/2019-11-01-preview/foo") {
					return nil, fmt.Errorf("bad prefix '%s'", prefix)
				}
				return []string{
					"tools/testdata/scenariol/foo/mgmt/2019-11-01-preview/foo/v0.0.0",
					"tools/testdata/scenariol/foo/mgmt/2019-11-01-preview/foo/v0.0.1",
					"tools/testdata/scenariol/foo/mgmt/2019-11-01-preview/foo/v0.1.0",
					"tools/testdata/scenariol/foo/mgmt/2019-11-01-preview/foo/v0.1.1",
					"tools/testdata/scenariol/foo/mgmt/2019-11-01-preview/foo/v0.1.2",
				}, nil
			},
			stage: "../../testdata/scenariol/foo/mgmt/2019-11-01-preview/foo/stage",
			expected: expected{
				dest:      "../../testdata/scenariol/foo/mgmt/2019-11-01-preview/foo",
				tag:       "tools/testdata/scenariol/foo/mgmt/2019-11-01-preview/foo/v0.2.0",
				mod:       fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenariol/foo/mgmt/2019-11-01-preview/foo\n\n%s\n", goVersion),
				version:   "v0.2.0",
				changelog: true,
			},
		},
		{
			name: "scenario m: new preview module",
			getTagsHook: func(root string, prefix string) ([]string, error) {
				if !strings.HasSuffix(prefix, "testdata/scenariom/foo/mgmt/2019-11-01-preview/foo") {
					return nil, fmt.Errorf("bad prefix '%s'", prefix)
				}
				return []string{}, nil
			},
			stage: "../../testdata/scenariom/foo/mgmt/2019-11-01-preview/foo/stage",
			expected: expected{
				dest:      "../../testdata/scenariom/foo/mgmt/2019-11-01-preview/foo",
				tag:       "tools/testdata/scenariom/foo/mgmt/2019-11-01-preview/foo/v0.0.0",
				mod:       fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenariom/foo/mgmt/2019-11-01-preview/foo\n\n%s\n", goVersion),
				version:   "v0.0.0",
				changelog: false,
			},
		},
		{
			name: "scenario n: patch version for preview module",
			getTagsHook: func(root string, prefix string) ([]string, error) {
				if !strings.HasSuffix(prefix, "testdata/scenarion/foo/mgmt/2019-11-01-preview/foo") {
					return nil, fmt.Errorf("bad prefix '%s'", prefix)
				}
				return []string{
					"tools/testdata/scenarion/foo/mgmt/2019-11-01-preview/foo/v0.0.0",
					"tools/testdata/scenarion/foo/mgmt/2019-11-01-preview/foo/v0.0.1",
					"tools/testdata/scenarion/foo/mgmt/2019-11-01-preview/foo/v0.1.0",
					"tools/testdata/scenarion/foo/mgmt/2019-11-01-preview/foo/v0.1.1",
					"tools/testdata/scenarion/foo/mgmt/2019-11-01-preview/foo/v0.1.2",
				}, nil
			},
			stage: "../../testdata/scenarion/foo/mgmt/2019-11-01-preview/foo/stage",
			expected: expected{
				dest:      "../../testdata/scenarion/foo/mgmt/2019-11-01-preview/foo",
				tag:       "tools/testdata/scenarion/foo/mgmt/2019-11-01-preview/foo/v0.1.3",
				mod:       fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenarion/foo/mgmt/2019-11-01-preview/foo\n\n%s\n", goVersion),
				version:   "v0.1.3",
				changelog: true,
			},
		},
		{
			name: "scenario o: no change in preview module",
			getTagsHook: func(root string, prefix string) ([]string, error) {
				if !strings.HasSuffix(prefix, "testdata/scenarioo/foo/mgmt/2019-11-01-preview/foo") {
					return nil, fmt.Errorf("bad prefix '%s'", prefix)
				}
				return []string{
					"tools/testdata/scenarioo/foo/mgmt/2019-11-01-preview/foo/v0.0.0",
					"tools/testdata/scenarioo/foo/mgmt/2019-11-01-preview/foo/v0.0.1",
					"tools/testdata/scenarioo/foo/mgmt/2019-11-01-preview/foo/v0.1.0",
					"tools/testdata/scenarioo/foo/mgmt/2019-11-01-preview/foo/v0.1.1",
					"tools/testdata/scenarioo/foo/mgmt/2019-11-01-preview/foo/v0.1.2",
				}, nil
			},
			stage: "../../testdata/scenarioo/foo/mgmt/2019-11-01-preview/foo/stage",
			expected: expected{
				dest:      "../../testdata/scenarioo/foo/mgmt/2019-11-01-preview/foo",
				tag:       "tools/testdata/scenarioo/foo/mgmt/2019-11-01-preview/foo/v0.1.2",
				mod:       fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenarioo/foo/mgmt/2019-11-01-preview/foo\n\n%s\n", goVersion),
				version:   "v0.1.2",
				changelog: false,
			},
		},
	}

	cleanTestData()
	defer cleanTestData()

	const repoRoot = "github.com/Azure/azure-sdk-for-go"

	for _, c := range testData {
		t.Logf("Testing %s", c.name)
		dest, tag, err := ExecuteUnstage(c.stage, repoRoot, defaultVersionSetting, c.getTagsHook)
		if err != nil {
			t.Fatalf("unexpected error: %+v", err)
		}
		if _, err := os.Stat(c.stage); err == nil || !os.IsNotExist(err) {
			t.Fatalf("stage directory is still there")
		}
		expectedDest, err := filepath.Abs(c.expected.dest)
		if err != nil {
			t.Fatalf("unexpected error: %+v", err)
		}
		if dest != expectedDest {
			t.Fatalf("bad dest: expected '%s' but got '%s'", expectedDest, dest)
		}
		if c.expected.tag != "" && tag != c.expected.tag {
			t.Fatalf("bad tag: expected '%s' but got '%s'", c.expected.tag, tag)
		}
		verifyGoMod(t, c.expected.dest, c.expected.mod)
		if c.expected.tag != "" {
			verifyVersion(t, c.expected.dest, c.expected.version, c.expected.tag)
		}
		if c.expected.changelog {
			verifyChangelog(t, c.expected.dest)
		} else {
			verifyNoChangelog(t, c.expected.dest)
		}
		verifyGoVet(t, filepath.Dir(c.stage))
	}
}

func TestCheckIdentical(t *testing.T) {
	testData := []struct {
		name      string
		baseline  string
		stage     string
		identical bool
	}{
		{
			name:      "scenario a",
			baseline:  "../../testdata/scenarioa/foo",
			stage:     "../../testdata/scenarioa/foo/stage",
			identical: false,
		},
		{
			name:      "scenario i",
			baseline:  "../../testdata/scenarioi/foo/mgmt/2019-10-23/foo",
			stage:     "../../testdata/scenarioi/foo/mgmt/2019-10-23/foo/stage",
			identical: true,
		},
		{
			name:      "scenario j",
			baseline:  "../../testdata/scenarioj/foo/mgmt/2019-10-23/foo",
			stage:     "../../testdata/scenarioj/foo/mgmt/2019-10-23/foo/stage",
			identical: true,
		},
	}

	cleanTestData()

	for _, c := range testData {
		t.Logf("Testing %s", c.name)
		baseline, err := filepath.Abs(c.baseline)
		if err != nil {
			t.Fatalf("unexpected error: %+v", err)
		}
		stage, err := filepath.Abs(c.stage)
		if err != nil {
			t.Fatalf("unexpected error: %+v", err)
		}
		identical, err := checkIdentical(baseline, stage)
		if err != nil {
			t.Fatalf("unexpected error: %+v", err)
		}
		if identical != c.identical {
			t.Fatalf("expected %v but got %v", c.identical, identical)
		}
	}
}

func cleanTestData() {
	cmd := exec.Command("git", "clean", "-xdf", "../../testdata")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(string(output))
	}
	cmd = exec.Command("git", "checkout", "--", "../../testdata")
	output, err = cmd.CombinedOutput()
	if err != nil {
		panic(string(output))
	}
}
