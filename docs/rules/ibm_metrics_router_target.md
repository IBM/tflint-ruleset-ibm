# ibm_metrics_router_target

Manages ibm_metrics_router_target. The name of the target.

This rule validates `ibm_metrics_router_target` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `TargetPrototype` (`metrics-router-metrics-router-v3`).

## Configuration

```hcl
rule "ibm_metrics_router_target" {
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

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_metrics_router_target_invalid_destination_crn` | `destination_crn` |
| `ibm_metrics_router_target_invalid_managed_by` | `managed_by` |
| `ibm_metrics_router_target_invalid_name` | `name` |
| `ibm_metrics_router_target_invalid_region` | `region` |

## Example

### Valid

```hcl
resource "ibm_metrics_router_target" "test" {
  destination_crn = "example-name"
  name = "example-name"
}
```
