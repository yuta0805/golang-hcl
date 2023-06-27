package main

import (
	"log"
	"os"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func main() {
	file := hclwrite.NewEmptyFile()

	// File.Bodyでhclファイルの中のbodyを取得できる
	rootBody := file.Body()
	// まっさらなhclファイル上にblockを追加
	vpcBlock := hclwrite.NewBlock("resource", []string{"aws_vpc", "main"})
	// 作成したblockの中にattributeを追加
	vpcBody := vpcBlock.Body()
	vpcBody.SetAttributeValue("cidr_block", cty.StringVal("10.0.0.0/16"))
	vpcBody.SetAttributeValue("instance_tenancy", cty.StringVal("default"))
	vpcBody.SetAttributeValue("tags", cty.ObjectVal(map[string]cty.Value{
		"Name": cty.StringVal("main-vpc"),
		"ENV":  cty.StringVal("dev"),
	}))

	// rootBodyにvpcBlockを追加
	rootBody.AppendBlock(vpcBlock)

	f, err := os.Create("write.tf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.Write(file.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}
