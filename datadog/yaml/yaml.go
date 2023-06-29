package yaml

import (
	yaml3 "gopkg.in/yaml.v3"
)

// type TestYamlConfig struct {
// 	Kind     string `yaml:"kind"`
// 	Metadata struct {
// 		Name        string `yaml:"name"`
// 		DisplayName string `yaml:"displayName"`
// 	} `yaml:"metadata"`
// 	Spec struct {
// 		Indicator struct {
// 			Metadata struct {
// 				Name string `yaml:"name"`
// 			} `yaml:"metadata"`
// 			Spec struct {
// 				Description string `yaml:"description"`
// 				RatioMetric struct {
// 					Counter bool `yaml:"counter"`
// 					Good    struct {
// 						MetricSource struct {
// 							MetricSourceRef string `yaml:"metricSourceRef"`
// 							Type            string `yaml:"type"`
// 							Spec            struct {
// 								Query string `yaml:"query"`
// 							} `yaml:"spec"`
// 						} `yaml:"metricSource"`
// 					} `yaml:"good"`
// 					Total struct {
// 						MetricSource struct {
// 							MetricSourceRef string `yaml:"metricSourceRef"`
// 							Type            string `yaml:"type"`
// 							Spec            struct {
// 								Query string `yaml:"query"`
// 							} `yaml:"spec"`
// 						} `yaml:"metricSource"`
// 					} `yaml:"total"`
// 				} `yaml:"ratioMetric"`
// 			} `yaml:"spec"`
// 		} `yaml:"indicator"`
// 		TimeWindow []struct {
// 			Duration  string `yaml:"duration"`
// 			IsRolling bool   `yaml:"isRolling"`
// 		} `yaml:"timeWindow"`
// 		BudgetingMethod string `yaml:"budgetingMethod"`
// 		Objectives      []struct {
// 			DisplayName string  `yaml:"displayName"`
// 			Target      float64 `yaml:"target"`
// 		} `yaml:"objectives"`
// 	} `yaml:"spec"`
// }

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
