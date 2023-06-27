package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type Resource struct {
	Type      string               `hcl:"type"`
	TypeName  []string             `hcl:"label"`
	Attribute map[string]cty.Value `hcl:"attr"`
}

type Variable struct {
	Type 	  string             `hcl:"type"`
	TypeName  string             `hcl:"label"`
	AttributeType hclwrite.Tokens `hcl:"type,attr"`
	Description string		  `hcl:"description,attr"`
	Default 	  hclwrite.Tokens `hcl:"default,attr"`
}

func main() {
	// fileのparse
	root := "./"
	fmt.Println(root)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if info.IsDir() || !strings.HasSuffix(info.Name(), ".tf") {
			return nil
		}

		scr, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		// hclwrite.ParseConfig(fileの中身,　ファイル名,　位置情報)
		// hclwrite.ParseConfigは内部で hclsyntax.ParseConfigを呼び出している。
		//  hclsyntax.ParseConfigはhcl.File{}を返している。
		// 返された、hcl.File{}は、hclwrite.File{}に変換されて返されてくる。
		// data型が違うので(hcl.Fileとhclwrite.Fileは別物) hclwriteは書き込み専用と考えても良いかも。
		file, err := hclwrite.ParseConfig(scr, path, hcl.Pos{Line: 1, Column: 1})

		body := file.Body()
		for _, block := range body.Blocks() {
			switch block.Type() {
			case "variable":
				fmt.Println("variableです")
			case "resource":
				fmt.Println("resourceです")
			case "data":
				fmt.Println("dataです")
			case "provider":
				fmt.Println("providerです")
			case "module":
				fmt.Println("moduleです")
			case "output":
				fmt.Println("outputです")
			}
			if block.Type() == "resource" {
				fmt.Println("resourceです")
				// block.Type()でresourceとかのタイプを取得できる
				fmt.Printf("block type is %s\n", block.Type())
				// block.Labels()でresourceに付与したlabel("aws_vpc" "main")を取得できる
				for _, label := range block.Labels() {
					fmt.Printf("block label is %s\n", label)
				}
				
				//  bodyの取得
				body := block.Body()
				// bodyの中のattributesを取得
				for k, attr := range body.Attributes() {
					fmt.Println(k) // attribute名 (cidr_blockとか)
					fmt.Println(string(attr.Expr().BuildTokens(nil).Bytes()))
				}
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}
