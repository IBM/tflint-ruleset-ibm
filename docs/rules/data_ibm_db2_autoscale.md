# ibm_db2_autoscale_data

Schema validation for the ibm_db2_autoscale data source.

This rule validates `ibm_db2_autoscale` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SuccessTakeoverDRResponse` (`db2-on-cloud-db2-on-cloud-v4`).

## Configuration

```hcl
rule "ibm_db2_autoscale_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `auto_scaling_allow_plan_limit` | no | — |
| `auto_scaling_enabled` | no | — |
| `auto_scaling_max_storage` | no | — |
| `auto_scaling_over_time_period` | no | — |
| `auto_scaling_pause_limit` | no | — |
| `auto_scaling_threshold` | no | — |
| `deployment_id` | yes | — |
| `storage_unit` | no | — |
| `storage_utilization_percentage` | no | — |
| `support_auto_scaling` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_db2_autoscale_data_invalid_deployment_id` | `deployment_id` |

## Example

### Valid

```hcl
data "ibm_db2_autoscale" "test" {
  deployment_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
