// package main

// import (
// 	"fmt"
// 	// "io/ioutil"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"strings"

// 	// hcl "github.com/hashicorp/hcl/v2"
// 	// "github.com/hashicorp/hcl/v2/hclsyntax"
// 	"github.com/hashicorp/hcl/v2/hclwrite"
// 	"github.com/hashicorp/hcl2/hclparse"
// 	// "github.com/hashicorp/hcl2/gohcl"
// 	// "github.com/zclconf/go-cty/cty"
// )

// // Variable is terraform module variable
// type Variable struct {
// 	Name        string          `hcl:"name,label"`
// 	Type        hclwrite.Tokens `hcl:"type,attr"`
// 	Description string          `hcl:"description,attr"`
// 	Default     hclwrite.Tokens `hcl:"default,attr"`
// }

// // Resource is terraform module resource
// type Resource struct {
// 	Type string `hcl:"type,label"`
// 	Name string `hcl:"name,label"`
// }

// // Parse parses a local terraform module and returns module structs
// func main() {
// 	// root := "./terraform"
// 	root := "./"
// 	paser := hclparse.NewParser()
// 	file, diafs := paser.ParseJSONFile("test.json")
// 	if diafs.HasErrors() {
// 		log.Fatal(diafs)
// 	}

// 	// body := file.Body.(*hclsyntax.Body)
// 	// for _, block := range body.Blocks {
// 	// 	fmt.Println(block)
// 	// }
// 	fmt.Println(file.Body.MissingItemRange().Start)
// 	// var fooInstance foo
//     // decodeDiags := gohcl.DecodeBody(file.Body, nil, &fooInstance)
//     // if decodeDiags.HasErrors() {
//     //     log.Fatal(decodeDiags.Error())
//     // }

// 	// fmt.Println(file)

// 	// var variables []*Variable
// 	// var resources []*Resource

// 	err := filepath.Walk(root,
// 		func(path string, info os.FileInfo, err error) error {
// 			if err != nil {
// 				return err
// 			}

// 			// if info.IsDir() || !strings.HasSuffix(info.Name(), ".tf") {
// 			// 	return nil
// 			// }

// 			if info.IsDir() || !strings.HasSuffix(info.Name(), ".hcl") {
// 				return nil
// 			}

// 			// src, err := ioutil.ReadFile(path)
// 			// if err != nil {
// 			// 	return err
// 			// }

// 			// file, diags := hclwrite.ParseConfig(src, path, hcl.InitialPos)
// 			// if diags.HasErrors() {
// 			// 	return diags
// 			// }

// 			// body := file.Body()
// 			// for _, block := range body.Blocks() {
// 			// 	fmt.Println(block.Type())
// 			// 	switch block.Type() {
// 			// 	case "variable":
// 			// 		variables = append(variables, parseVariable(block))
// 			// 	// case "output":
// 			// 	// 	outputs = append(outputs, p.parseOutput(block))
// 			// 	case "resource":
// 			// 		resources = append(resources, parseResource(block))
// 			// 	}
// 			// }

// 			return nil
// 		})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// return &Module{
// 	// 	Source:    p.Source,
// 	// 	Variables: variables,
// 	// 	Outputs:   outputs,
// 	// 	Resources: resources,
// 	// }, nil

// }

// // func parseVariable(block *hclwrite.Block) *Variable {
// // 	variable := Variable{
// // 		Name:    block.Labels()[0],
// // 		Default: hclwrite.TokensForValue(cty.StringVal("")),
// // 	}
// // 	body := block.Body()
// // 	for k, v := range body.Attributes() {
// // 		switch k {
// // 		case "type":
// // 			var typeTokens hclwrite.Tokens
// // 			for _, t := range v.Expr().BuildTokens(nil) {
// // 				if t.Type != hclsyntax.TokenNewline {
// // 					typeTokens = append(typeTokens, t)
// // 				}
// // 			}
// // 			variable.Type = typeTokens
// // 		case "default":
// // 			variable.Default = v.Expr().BuildTokens(nil)
// // 		case "description":
// // 			description := string(v.Expr().BuildTokens(nil).Bytes())
// // 			variable.Description = description[2 : len(description)-1]
// // 		}
// // 	}
// // 	return &variable
// // }

// // func parseResource(block *hclwrite.Block) *Resource {
// // 	resource := Resource{
// // 		Type: block.Labels()[0],
// // 		Name: block.Labels()[1],
// // 	}

// // 	fmt.Println(resource.Name)
// // 	return &resource
// // }

package main

import (
)

func main() {
	// parser := hclparse.NewParser()

	// file, diag := parser.ParseJSONFile("test.json")
	// if diag != nil {
	// 	log.Fatal(diag)
	// }	
	// newFile, err := os.Create("test.hcl")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = newFile.Write(file.Bytes)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%s", file.Bytes)
}
