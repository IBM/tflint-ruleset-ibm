# ibm_logs_extension_deployment

Schema validation for the ibm_logs_extension_deployment resource.

This rule validates `ibm_logs_extension_deployment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Extension deployment` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_extension_deployment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `applications` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `item_ids` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `logs_extension_id` | yes | — |
| `subsystems` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `version` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_extension_deployment_invalid_applications` | `applications` |
| `ibm_logs_extension_deployment_invalid_item_ids` | `item_ids` |
| `ibm_logs_extension_deployment_invalid_logs_extension_id` | `logs_extension_id` |
| `ibm_logs_extension_deployment_invalid_subsystems` | `subsystems` |
| `ibm_logs_extension_deployment_invalid_version` | `version` |

## Example

### Valid

```hcl
resource "ibm_logs_extension_deployment" "test" {
  item_ids = "example-name"
  logs_extension_id = "01234567-89ab-cdef-0123-456789abcdef"
  version = "example-name"
}
```
