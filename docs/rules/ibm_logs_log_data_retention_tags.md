# ibm_logs_log_data_retention_tags

Schema validation for the ibm_logs_log_data_retention_tags resource.

This rule validates `ibm_logs_log_data_retention_tags` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/log_data_retention_tags:put` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_log_data_retention_tags" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `tags` | yes | pattern `^[a-zA-Z0-9_-]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_log_data_retention_tags_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_logs_log_data_retention_tags" "test" {
  tags = ["example-name"]
}
```
