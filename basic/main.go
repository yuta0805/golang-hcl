package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
	"github.com/hashicorp/hcl2/hclwrite"
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
	parser := hclparse.NewParser()
	file, diag := parser.ParseHCLFile("./basic.hcl")
	if diag != nil {
		log.Fatalf("Failed to parse: %v", diag.Error())
	}

	var config Config
	configDiags := gohcl.DecodeBody(file.Body, nil, &config)
	if configDiags.HasErrors() {
		log.Fatalf("Failed to decodeです: %v", configDiags.Error())
	}

	newFile := hclwrite.NewEmptyFile()
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
