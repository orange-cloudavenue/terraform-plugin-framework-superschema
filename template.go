//go:build ignore

/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

// go generate
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"html/template"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
)

type templateInfos struct {
	TypeName string
}

//go:embed type_attribute.go.tmpl
var templateTypeAttribute string

//go:embed supertype_attribute.go.tmpl
var templateSuperTypeAttribute string

var templateFuncs = template.FuncMap{
	"trimSuffix": func(s, suffix string) string {
		return strings.TrimSuffix(s, suffix)
	},
}

func main() {
	fmt.Println("generating types files...")
	tA := []string{"string", "bool", "float64", "int32", "int64", "list", "list_nested", "object", "single_nested", "set", "set_nested", "number", "map", "map_nested"}

	for _, t := range tA {
		infos := templateInfos{
			TypeName: strcase.ToCamel(t),
		}

		tmplType, err := template.New("template").Funcs(templateFuncs).Parse(templateTypeAttribute)
		if err != nil {
			fmt.Printf("error from template parse : %v\n", err)
			os.Exit(1)
		}

		tmplSuperType, err := template.New("template").Funcs(templateFuncs).Parse(templateSuperTypeAttribute)
		if err != nil {
			fmt.Printf("error from template parse : %v\n", err)
			os.Exit(1)
		}

		var (
			tpl      bytes.Buffer
			tplSuper bytes.Buffer
		)

		if err := tmplType.Execute(&tpl, infos); err != nil {
			fmt.Printf("error from template execute : %v\n", err)
			os.Exit(1)
		}

		if err := tmplSuperType.Execute(&tplSuper, infos); err != nil {
			fmt.Printf("error from template execute : %v\n", err)
			os.Exit(1)
		}

		// format the code
		formattedTmplType, errFormat := format.Source(tpl.Bytes())
		if errFormat != nil {
			fmt.Printf("error from format : %v\n", errFormat)
			os.Exit(1)
		}

		formattedTmplSuperType, errFormat := format.Source(tplSuper.Bytes())
		if errFormat != nil {
			fmt.Printf("error from format : %v\n", errFormat)
			os.Exit(1)
		}

		if err := os.WriteFile(t+"_attribute.go", formattedTmplType, 0644); err != nil {
			fmt.Printf("write to file error : %v\n", err)
		}

		if err := os.WriteFile("super"+t+"_attribute.go", formattedTmplSuperType, 0644); err != nil {
			fmt.Printf("write to file error : %v\n", err)
		}
	}
}
