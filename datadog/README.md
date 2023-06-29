```
 Terraform will perform the following actions:
datadog-slo    # datadog_service_level_objective.slo (slo) will be created
               + resource "datadog_service_level_objective" "slo" {
                   + description       = "管理者としてpoint計算ができること"
                   + id                = (known after apply)
                   + name              = "point計算"
                   + target_threshold  = (known after apply)
                   + timeframe         = (known after apply)
                   + type              = "metric"
                   + warning_threshold = (known after apply)

                   + query {
                       + denominator = "sum:hoge.sli.hr.nencho.good{*}.as_count()"
                       + numerator   = "sum:hoge.sli.hr.nencho{*}.as_count()"
                     }

                   + thresholds {
                       + target          = 99.9
                       + target_display  = (known after apply)
                       + timeframe       = "30d"
                       + warning         = 99.8
                       + warning_display = (known after apply)
                     }
                 }

             Plan: 1 to add, 0 to change, 0 to destroy
```
