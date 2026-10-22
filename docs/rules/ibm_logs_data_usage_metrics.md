# ibm_logs_data_usage_metrics

Schema validation for the ibm_logs_data_usage_metrics resource.

This rule validates `ibm_logs_data_usage_metrics` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/data_usage:put` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_data_usage_metrics" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `enabled` | yes | API: Whether metrics export is enabled or not. When 'true' metrics are exported. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_data_usage_metrics_invalid_enabled` | `enabled` |

## Example

### Valid

```hcl
resource "ibm_logs_data_usage_metrics" "test" {
  enabled = true
}
```
