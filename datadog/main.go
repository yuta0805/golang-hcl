package main

import (
	"io/ioutil"
	"log"

	"golang-hcl/datadog/yaml"
	"golang-hcl/datadog/hcl"
)

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
/*
type T struct {
	A string
	B struct {
			RenamedC int   `yaml:"c"`
			D        []int `yaml:",flow"`
	}
}
*/


func main() {
	path := "./sample.yaml"
	src, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	testYaml := yaml.NewTestYamlConfig()
	testYaml, err = testYaml.Parse(src)
	if err != nil {
		log.Fatal(err)
	}
	
	// yamlをhclに変換

	file := hcl.NewHCLFile()
	blocks := file.Create(testYaml)
	ok := file.Append(blocks)
	if ok.Err != nil {
		log.Fatal(ok.Err.Error())
	}
}
