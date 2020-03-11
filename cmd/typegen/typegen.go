// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/willbeason/typegen/pkg/jsonschema"
	"github.com/willbeason/typegen/pkg/languages/ts"
	"io/ioutil"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"os"
	"strings"
)

var filters []string

func init() {
	mainCmd.Flags().StringSliceVar(&filters, "definitions", []string{"*"},
		`Comma-delimited list of swagger Definitions to generate types for. Includes transitive dependencies.
Defaults to all Definitions if unset. Use '*' for wildcard.`)
}

var mainCmd = &cobra.Command{
	Use:  "typegen [SWAGGER_PATH] [OUTPUT_PATH]",
	Long: "Generate TypeScript types from a swagger.json.",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		swaggerPath := args[0]
		bytes, err := ioutil.ReadFile(swaggerPath)
		if err != nil {
			return errors.Wrap(err, "unable to find Swagger file")
		}

		swagger := v1.JSONSchemaProps{}
		err = json.Unmarshal(bytes, &swagger)
		if err != nil {
			return errors.Wrap(err, "unable to parse Swagger file as JSONSchemaProps")
		}

		for fullName, props := range swagger.Definitions {
			idx := strings.LastIndex(fullName, ".")

			d, err := jsonschema.Parse(fullName[idx+1:], []string{fullName[:idx]}, props)
			if err != nil {
				return err
			}
			fmt.Printf("%s\n\n", ts.Print(d))
		}

		return nil
	},
}

//func printTS(outPath string, refObjects map[definition.Ref]definition.Object, definitions []definition.Definition) error {
//	pkgs := make(map[string][]definition.Definition)
//	for _, d := range definitions {
//		pkg := d.Metadata().Package
//		pkgs[pkg] = append(pkgs[pkg], d)
//	}
//
//	pkgs = swagger.FilterDefinitions(filters, pkgs)
//
//	lang := language.TypeScript{
//		RefObjects: refObjects,
//	}
//	for pkg, defs := range pkgs {
//		var contents []string
//		header := lang.PrintHeader(defs)
//		if header != "" {
//			contents = append(contents, header)
//		}
//		sort.Slice(defs, func(i, j int) bool {
//			return definition.FullName(defs[i]) < definition.FullName(defs[j])
//		})
//		for _, d := range defs {
//			contents = append(contents, lang.PrintDefinition(d))
//		}
//
//		err := ioutil.WriteFile(filepath.Join(outPath, pkg+".ts"), []byte(strings.Join(contents, "\n\n")), 0644)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

func main() {
	err := mainCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
