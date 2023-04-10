package main

import (
	"bytes"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	const openAPIFile = "dataplatform.yaml"
	swagger, err := openapi3.NewLoader().LoadFromFile(openAPIFile)
	if err != nil {
		panic(err)
	}

	operationID := "clustersNodepoolsDelete"
	operation, _, method, err := findOperation(swagger, operationID)
	if err != nil {
		panic(err)
	}

	flags, err := extractFlags(swagger, operation)
	if err != nil {
		panic(err)
	}

	log.Printf("Extracted flags: %+v\n", flags)

	command := CLICommand{
		FunctionName:     pascalCase(operationID),
		Namespace:        "dataplatform",
		Resource:         "nodepool",
		Verb:             strings.ToLower(method),
		Aliases:          "{\"del\", \"d\"}",
		ShortDesc:        operation.Summary,
		RequiredFlagSets: "[]string{constants.ArgAll}, []string{constants.FlagClusterId}, []string{constants.ArgAll, constants.FlagClusterId}",
		InitClient:       "true",
		Flags:            flags,
	}

	command.Example = fmt.Sprintf("ionosctl %s %s %s --cluster-id <cluster-id>", command.Namespace, command.Resource, command.Verb)

	tmpl, err := template.New("cli-command").Parse(cliCommandTemplate)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, command)
	if err != nil {
		panic(err)
	}

	// Write the output to a file or print it
	f, err := os.Create("generated_command.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
}

type CLICommand struct {
	FunctionName     string
	Namespace        string
	Resource         string
	Verb             string
	Aliases          string
	ShortDesc        string
	Example          string
	RequiredFlagSets string
	InitClient       string
	Flags            []Flag
}

type Flag struct {
	Name        string
	ShortName   string
	Type        string
	Default     string
	Description string
	Required    bool
}

func findOperation(swagger *openapi3.T, operationID string) (*openapi3.Operation, string, string, error) {
	for path, pathItem := range swagger.Paths {
		for method, operation := range pathItem.Operations() {
			if operation.OperationID == operationID {
				return operation, path, method, nil
			}
		}
	}
	return nil, "", "", fmt.Errorf("operation not found: %s", operationID)
}

func extractFlags(swagger *openapi3.T, operation *openapi3.Operation) ([]Flag, error) {
	var flags []Flag

	j, _ := operation.MarshalJSON()
	log.Printf("Operation: %s\n", j)
	// Iterate through operation parameters
	for _, paramRef := range operation.Parameters {
		log.Printf("Looking at parameter: %+v\n", paramRef.Value.Description)

		param := paramRef.Value

		// Extract parameter properties
		flag := Flag{
			Name:        param.Name,
			ShortName:   "", // You can provide a custom mapping for short names or leave it empty
			Type:        flagTypeFromSchema(param.Schema.Value),
			Default:     flagDefaultFromSchema(param.Schema.Value),
			Description: param.Description,
			Required:    param.Required,
		}
		flags = append(flags, flag)
	}

	// Iterate through request body properties
	if operation.RequestBody != nil {
		content := operation.RequestBody.Value.Content["application/json"]
		if content != nil {
			for propName, prop := range content.Schema.Value.Properties {
				// Extract property properties
				flag := Flag{
					Name:        propName,
					ShortName:   "", // You can provide a custom mapping for short names or leave it empty
					Type:        flagTypeFromSchema(prop.Value),
					Default:     flagDefaultFromSchema(prop.Value),
					Description: prop.Value.Description,
					Required:    slices.Contains(content.Schema.Value.Required, propName),
				}
				flags = append(flags, flag)
			}
		}
	}

	return flags, nil
}

func flagTypeFromSchema(schema *openapi3.Schema) string {
	switch schema.Type {
	case "integer":
		return "Int"
	case "number":
		return "Float64"
	case "string":
		return "String"
	case "boolean":
		return "Bool"
	case "array":
		return "StringSlice"
	case "object":
		return "StringToString/*Check me!*/"
	default:
		return "String" // Default to string for unknown types
	}
}

func flagDefaultFromSchema(schema *openapi3.Schema) string {
	switch schema.Type {
	case "integer":
		return "0"
	case "number":
		return "0.0"
	case "string":
		return `""`
	case "boolean":
		return "false"
	case "array":
		return "[]string{}"
	case "object":
		return "map[string]string{}"
	default:
		return `""`
	}
}

func pascalCase(s string) string {
	words := strings.Split(s, "_")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}

const cliCommandTemplate = `func {{.FunctionName}}Cmd() *core.Command {
	cmd := core.NewCommand(context.TODO(), nil, core.CommandBuilder{
		Namespace: "{{.Namespace}}",
		Resource:  "{{.Resource}}",
		Verb:      "{{.Verb}}",
		Aliases:   []string{{.Aliases}},
		ShortDesc: "{{.ShortDesc}}",
		Example:   "{{.Example}}",
		PreCmdRun: func(c *core.PreCommandConfig) error {
			return core.CheckRequiredFlagsSets(c.Command, c.NS, []string{{.RequiredFlagSets}})
		},
		CmdRun: func(c *core.CommandConfig) error {
			// Implement the actual command logic here
		},
		InitClient: {{.InitClient}},
	})

	// TODO: Check me! Did I successfully add all flags for {{.FunctionName}}?
	{{range .Flags}}
	cmd.Add{{.Type}}Flag("{{.Name}}", "{{.ShortName}}", {{.Default}}, "{{.Description}}"{{if .Required}}, core.RequiredFlagOption(){{end}})
	{{end}}


	return cmd
}
`
