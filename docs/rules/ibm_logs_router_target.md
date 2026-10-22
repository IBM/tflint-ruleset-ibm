# ibm_logs_router_target

Manages ibm_logs_router_target. The name of the target resource.

This rule validates `ibm_logs_router_target` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Target` (`logs-router-service-api-logs-router-v1`).

## Configuration

```hcl
rule "ibm_logs_router_target" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `crn` | no | — |
| `destination_crn` | yes | pattern `^[a-zA-Z0-9 \-._:\/]+$` |
| `managed_by` | no | one of: `account`, `enterprise` |
| `name` | yes | pattern `^[a-zA-Z0-9 \-._:]+$` |
| `region` | no | pattern `^[a-zA-Z0-9 \-._:]+$` |
| `target_type` | no | — |
| `updated_at` | no | — |
| `write_status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_router_target_invalid_destination_crn` | `destination_crn` |
| `ibm_logs_router_target_invalid_managed_by` | `managed_by` |
| `ibm_logs_router_target_invalid_name` | `name` |
| `ibm_logs_router_target_invalid_region` | `region` |

## Example

### Valid

```hcl
resource "ibm_logs_router_target" "test" {
  destination_crn = "example-name"
  name = "example-name"
}
```
