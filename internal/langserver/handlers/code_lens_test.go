package handlers

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/hashicorp/go-version"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/ms-henglu/azurerm-restapi-lsp/internal/langserver"
	"github.com/ms-henglu/azurerm-restapi-lsp/internal/langserver/session"
	"github.com/ms-henglu/azurerm-restapi-lsp/internal/lsp"
	"github.com/ms-henglu/azurerm-restapi-lsp/internal/terraform/exec"
	"github.com/stretchr/testify/mock"
)

func TestCodeLens_withoutInitialization(t *testing.T) {
	ls := langserver.NewLangServerMock(t, NewMockSession(nil))
	stop := ls.Start(t)
	defer stop()

	ls.CallAndExpectError(t, &langserver.CallRequest{
		Method: "textDocument/codeLens",
		ReqParams: fmt.Sprintf(`{
			"textDocument": {
				"uri": "%s/main.tf"
			}
		}`, TempDir(t).URI())}, session.SessionNotInitialized.Err())
}

func TestCodeLens_withoutOptIn(t *testing.T) {
	tmpDir := TempDir(t)
	InitPluginCache(t, tmpDir.Dir())

	var testSchema tfjson.ProviderSchemas
	err := json.Unmarshal([]byte(testModuleSchemaOutput), &testSchema)
	if err != nil {
		t.Fatal(err)
	}

	ls := langserver.NewLangServerMock(t, NewMockSession(nil))
	stop := ls.Start(t)
	defer stop()

	ls.Call(t, &langserver.CallRequest{
		Method: "initialize",
		ReqParams: fmt.Sprintf(`{
		"capabilities": {},
		"rootUri": %q,
		"processId": 12345
	}`, tmpDir.URI())})
	ls.Notify(t, &langserver.CallRequest{
		Method:    "initialized",
		ReqParams: "{}",
	})
	ls.Call(t, &langserver.CallRequest{
		Method: "textDocument/didOpen",
		ReqParams: fmt.Sprintf(`{
		"textDocument": {
			"version": 0,
			"languageId": "terraform",
			"text": "provider \"test\" {\n\n}\n",
			"uri": "%s/main.tf"
		}
	}`, tmpDir.URI())})
	ls.CallAndExpectResponse(t, &langserver.CallRequest{
		Method: "textDocument/codeLens",
		ReqParams: fmt.Sprintf(`{
			"textDocument": {
				"uri": "%s/main.tf"
			}
		}`, TempDir(t).URI()),
	}, `{
				"jsonrpc": "2.0",
				"id": 3,
				"result": []
	}`)
}

func TestCodeLens_referenceCount(t *testing.T) {
	tmpDir := TempDir(t)
	InitPluginCache(t, tmpDir.Dir())

	var testSchema tfjson.ProviderSchemas
	err := json.Unmarshal([]byte(testModuleSchemaOutput), &testSchema)
	if err != nil {
		t.Fatal(err)
	}

	ls := langserver.NewLangServerMock(t, NewMockSession(&MockSessionInput{
		TerraformCalls: &exec.TerraformMockCalls{
			PerWorkDir: map[string][]*mock.Call{
				tmpDir.Dir(): {
					{
						Method:        "Version",
						Repeatability: 1,
						Arguments: []interface{}{
							mock.AnythingOfType(""),
						},
						ReturnArguments: []interface{}{
							version.Must(version.NewVersion("0.12.0")),
							nil,
							nil,
						},
					},
					{
						Method:        "GetExecPath",
						Repeatability: 1,
						ReturnArguments: []interface{}{
							"",
						},
					},
					{
						Method:        "ProviderSchemas",
						Repeatability: 1,
						Arguments: []interface{}{
							mock.AnythingOfType(""),
						},
						ReturnArguments: []interface{}{
							&testSchema,
							nil,
						},
					},
				},
			},
		}}))
	stop := ls.Start(t)
	defer stop()

	ls.Call(t, &langserver.CallRequest{
		Method: "initialize",
		ReqParams: fmt.Sprintf(`{
		"capabilities": {
			"experimental": {
				"showReferencesCommandId": "test.id"
			}
		},
		"rootUri": %q,
		"processId": 12345
	}`, tmpDir.URI())})
	ls.Notify(t, &langserver.CallRequest{
		Method:    "initialized",
		ReqParams: "{}",
	})
	ls.Call(t, &langserver.CallRequest{
		Method: "textDocument/didOpen",
		ReqParams: fmt.Sprintf(`{
		"textDocument": {
			"version": 0,
			"languageId": "terraform",
			"text": %q,
			"uri": "%s/main.tf"
		}
	}`, `variable "test" {
}
output "test" {
	value = var.test
}
`, tmpDir.URI())})
	ls.CallAndExpectResponse(t, &langserver.CallRequest{
		Method: "textDocument/codeLens",
		ReqParams: fmt.Sprintf(`{
			"textDocument": {
				"uri": "%s/main.tf"
			}
		}`, TempDir(t).URI()),
	}, `{
				"jsonrpc": "2.0",
				"id": 3,
				"result": [
					{
						"range": {
							"start": {
								"line": 0,
								"character": 0
							},
							"end": {
								"line": 1,
								"character": 1
							}
						},
						"command": {
							"title": "1 reference",
							"command": "test.id",
							"arguments": [
								{
									"line": 0,
									"character": 7
								},
								{
									"includeDeclaration": false
								}
							]
						}
					}
				]
	}`)
}

func TestCodeLens_referenceCount_crossModule(t *testing.T) {
	rootModPath, err := filepath.Abs(filepath.Join("testdata", "single-submodule"))
	if err != nil {
		t.Fatal(err)
	}

	submodPath := filepath.Join(rootModPath, "application")

	rootModUri := lsp.FileHandlerFromDirPath(rootModPath)
	submodUri := lsp.FileHandlerFromDirPath(submodPath)

	var testSchema tfjson.ProviderSchemas
	err = json.Unmarshal([]byte(testModuleSchemaOutput), &testSchema)
	if err != nil {
		t.Fatal(err)
	}

	ls := langserver.NewLangServerMock(t, NewMockSession(&MockSessionInput{
		TerraformCalls: &exec.TerraformMockCalls{
			PerWorkDir: map[string][]*mock.Call{
				submodPath:  validTfMockCalls(),
				rootModPath: validTfMockCalls(),
			},
		}}))
	stop := ls.Start(t)
	defer stop()

	ls.Call(t, &langserver.CallRequest{
		Method: "initialize",
		ReqParams: fmt.Sprintf(`{
		"capabilities": {
			"experimental": {
				"showReferencesCommandId": "test.id"
			}
		},
		"rootUri": %q,
		"processId": 12345
	}`, rootModUri.URI())})
	ls.Notify(t, &langserver.CallRequest{
		Method:    "initialized",
		ReqParams: "{}",
	})
	ls.Call(t, &langserver.CallRequest{
		Method: "textDocument/didOpen",
		ReqParams: fmt.Sprintf(`{
		"textDocument": {
			"version": 0,
			"languageId": "terraform",
			"text": %q,
			"uri": "%s/main.tf"
		}
	}`, `variable "environment_name" {
  type = string
}

variable "app_prefix" {
  type = string
}

variable "instances" {
  type = number
}
`, submodUri.URI())})
	// TODO remove once we support synchronous dependent tasks
	// See https://github.com/ms-henglu/azurerm-restapi-lsp/issues/719
	time.Sleep(2 * time.Second)
	ls.CallAndExpectResponse(t, &langserver.CallRequest{
		Method: "textDocument/codeLens",
		ReqParams: fmt.Sprintf(`{
			"textDocument": {
				"uri": "%s/main.tf"
			}
		}`, submodUri.URI()),
	}, `{
			"jsonrpc": "2.0",
			"id": 3,
			"result": [
				{
					"range": {
						"start": {
							"line": 0,
							"character": 0
						},
						"end": {
							"line": 2,
							"character": 1
						}
					},
					"command": {
						"title": "1 reference",
						"command": "test.id",
						"arguments": [
							{
								"line": 0,
								"character": 13
							},
							{
								"includeDeclaration": false
							}
						]
					}
				},
				{
					"range": {
						"start": {
							"line": 4,
							"character": 0
						},
						"end": {
							"line": 6,
							"character": 1
						}
					},
					"command": {
						"title": "1 reference",
						"command": "test.id",
						"arguments": [
							{
								"line": 4,
								"character": 10
							},
							{
								"includeDeclaration": false
							}
						]
					}
				},
				{
					"range": {
						"start": {
							"line": 8,
							"character": 0
						},
						"end": {
							"line": 10,
							"character": 1
						}
					},
					"command": {
						"title": "1 reference",
						"command": "test.id",
						"arguments": [
							{
								"line": 8,
								"character": 10
							},
							{
								"includeDeclaration": false
							}
						]
					}
				}
			]
	}`)
}