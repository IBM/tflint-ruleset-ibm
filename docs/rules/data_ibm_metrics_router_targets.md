# ibm_metrics_router_targets_data

Reads ibm_metrics_router_targets. The name of the target resource.

This rule validates `ibm_metrics_router_targets` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/targets/{id}:patch` (`metrics-router-metrics-router-v3`).

## Configuration

```hcl
rule "ibm_metrics_router_targets_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | no | pattern `^[a-zA-Z0-9 \-._:]+$` |
| `targets` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_metrics_router_targets_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_metrics_router_targets" "test" {
}
```
