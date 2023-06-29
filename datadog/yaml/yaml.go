package yaml

import (
	yaml3 "gopkg.in/yaml.v3"
)

type TestYamlConfig struct {
	Kind        string   `yaml:"kind"`
	Metadata    struct {
		Name string `yaml:"name"`
		DisplayName string `yaml:"displayName"`
	} `yaml:"metadata"`
	Spec        struct {
		Indicator struct {
			Spec struct {
				RatioMetric struct {
					Good struct {
						MetricSource struct {
							Spec struct {
								Query string `yaml:"query"`
							} `yaml:"spec"`
						} `yaml:"metricSource"`
					} `yaml:"good"`
					Total struct {
						MetricSource struct {
							Spec struct {
								Query string `yaml:"query"`
							} `yaml:"spec"`
						} `yaml:"metricSource"`
					} `yaml:"total"`
				} `yaml:"ratioMetric"`
			} `yaml:"spec"`
		} `yaml:"indicator"`
		TimeWindow []Window `yaml:"timeWindow"`
		Objectives []Object `yaml:"objectives"`
	} `yaml:"spec"`
}

type Object struct {
	DisplayName string     `yaml:"displayName"`
	Target      float64    `yaml:"target"`
}

type Window struct {
	Duration string `yaml:"duration"`
	IsRolling bool  `yaml:"isRolling"`
}

func NewTestYamlConfig() TestYamlConfig {
	return TestYamlConfig{}
}

func (y TestYamlConfig) Parse(src []byte) (TestYamlConfig, error) {
	err := yaml3.Unmarshal(src, &y)
	if err != nil {
		return y, err
	}
	
	return y, nil
} 
