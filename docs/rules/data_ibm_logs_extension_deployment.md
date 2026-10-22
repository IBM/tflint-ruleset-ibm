# ibm_logs_extension_deployment_data

Schema validation for the ibm_logs_extension_deployment data source.

This rule validates `ibm_logs_extension_deployment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_extension_deployment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `applications` | no | — |
| `item_ids` | no | — |
| `logs_extension_id` | yes | — |
| `subsystems` | no | — |
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_extension_deployment_data_invalid_logs_extension_id` | `logs_extension_id` |

## Example

### Valid

```hcl
data "ibm_logs_extension_deployment" "test" {
  logs_extension_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
