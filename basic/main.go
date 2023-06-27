package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// 構造体(スキーマ)を用意しておく(decode用)
type Config struct {
	Version string `hcl:"version"`
	Pods    []Pod  `hcl:"pod,block"`
}

type Pod struct {
	ApiVersion string `hcl:"api_version"`
	Kind       string `hcl:"kind"`
	Name       string `hcl:"name"`
	Namespace  string `hcl:"namespace"`
}

func main() {
	// parseをcreateする
	parser := hclparse.NewParser()
	// 読みたいファイルparseする
	// hcl.File型帰ってくる
	file, diag := parser.ParseHCLFile("./basic.hcl")
	if diag != nil {
		log.Fatalf("Failed to parse: %v", diag.Error())
	}

	var config Config
	// gohcl.DecodeBody(ファイルの中身, ファイル名, 構造体のポインタ)
	// 構造体のポインタにファイルの中身をdecodeしてくれる(構造体のフィールドとparseした内容mappingしてくれる)
	configDiags := gohcl.DecodeBody(file.Body, nil, &config)
	if configDiags.HasErrors() {
		log.Fatalf("Failed to decodeです: %v", configDiags.Error())
	}

	// hclwrite.NewEmptyFile()で空のhclファイルを作成
	newFile := hclwrite.NewEmptyFile()
	// gohcl.EncodeIntoBody(構造体のポインタ, hclwrite.File.Body)
	// 構造体のfieldをhclwrite.File.Bodyにencodeしてくれる(fileへ書き込み)
	gohcl.EncodeIntoBody(&config, newFile.Body())
	fmt.Println(string(newFile.Bytes()))
	content := newFile.Bytes()

	f, err := os.Create("basic.tf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 作成したファイルに書き込み
	_, err = f.Write(content)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Success")
	
	for _, pod := range config.Pods {
		fmt.Println(pod)
	}
}
