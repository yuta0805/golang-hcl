package hcl

import (
	"errors"
	"fmt"
	"golang-hcl/datadog/yaml"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

var (
	target int64
	warning int64
	timewindow string
)
type HCLFile struct {
	File       *hclwrite.File
}

type HCLFileErr struct {
	Err        error
}

func NewHCLFile() HCLFile {
	return HCLFile{
		File: hclwrite.NewEmptyFile(),
	}
}

func (f HCLFile) Create(yaml yaml.TestYamlConfig) []*hclwrite.Block {
	sloBlock := hclwrite.NewBlock("resource", []string{"datadog_service_level_objective", "slo"})
	sloBody := sloBlock.Body()
	sloBody.SetAttributeValue("name", cty.StringVal(yaml.Metadata.Name))
	sloBody.SetAttributeValue("type", cty.StringVal("metric"))
	sloBody.SetAttributeValue("description", cty.StringVal(yaml.Metadata.DisplayName))
	sloBody.AppendNewline()
	sloBody.SetAttributeValue("query", cty.ObjectVal(map[string]cty.Value{
		"denominator": cty.StringVal(yaml.Spec.Indicator.Spec.RatioMetric.Total.MetricSource.Spec.Query),
		"numerator":   cty.StringVal(yaml.Spec.Indicator.Spec.RatioMetric.Good.MetricSource.Spec.Query),
	}))
	sloBlock.Body().AppendNewline()

	for _, object := range yaml.Spec.Objectives {
		if object.DisplayName == "Target" {
			target = int64(object.Target * 100)
		}

		if object.DisplayName == "Warning" {
			warning = int64(object.Target * 100)
		}
	}

	for _, window := range yaml.Spec.TimeWindow {
		timewindow = window.Duration
	}

	sloBody.SetAttributeValue("thresholds", cty.ObjectVal(map[string]cty.Value{
		"target":  cty.NumberIntVal(target),
		"warning": cty.NumberIntVal(warning),
		"timefram": cty.StringVal(timewindow),
	}))

	providerBlock := f.createProvider()

	var blocks []*hclwrite.Block
	blocks = append(blocks, providerBlock)
	blocks = append(blocks, sloBlock)
	
	return blocks
	// rootBody.AppendBlock(sloBlock)
	// fmt.Println(string(f.File.Bytes()))
}

func (f HCLFile) createProvider() *hclwrite.Block {
	providerBlock := hclwrite.NewBlock("provider", []string{"datadog"})
	providerBody := providerBlock.Body()
	providerBody.SetAttributeValue("api_key", cty.StringVal("hogehoge"))
	providerBody.SetAttributeValue("app_key", cty.StringVal("hogehoge"))
	
	return providerBlock
}

func (f HCLFile) Append(blocks []*hclwrite.Block) HCLFileErr {
	if len(blocks) == 0 {
		err := HCLFileErr{
			Err: errors.New("blocks is empty"),
		}

		return err 
	}

	rootBody := f.File.Body()
	for _, block := range blocks {
		rootBody.AppendBlock(block)
		rootBody.AppendNewline()
	}
	fmt.Println(string(f.File.Bytes()))

	return HCLFileErr{}
}
