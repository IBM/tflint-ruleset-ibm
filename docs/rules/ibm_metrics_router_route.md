# ibm_metrics_router_route

Manages ibm_metrics_router_route. The name of the route.

This rule validates `ibm_metrics_router_route` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Route` (`metrics-router-metrics-router-v3`).

## Configuration

```hcl
rule "ibm_metrics_router_route" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `crn` | no | — |
| `managed_by` | no | one of: `enterprise`, `account` |
| `name` | yes | pattern `^[a-zA-Z0-9 \-._:]+$` |
| `updated_at` | no | — |
| `rules` (block) | yes | nested block |
| `rules.action` | no | one of: `send`, `drop` |
| `rules.inclusion_filters` (block) | no | nested block |
| `rules.inclusion_filters.operand` | yes | one of: `location`, `service_name`, `service_instance`, `resource_type`, `resource` |
| `rules.inclusion_filters.operator` | yes | one of: `is`, `in` |
| `rules.inclusion_filters.values` | yes | API: The provided string values of the operand to be compared with. |
| `rules.targets` (block) | no | nested block |
| `rules.targets.id` | yes | pattern `^[a-zA-Z0-9 \-._:]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_metrics_router_route_invalid_managed_by` | `managed_by` |
| `ibm_metrics_router_route_invalid_name` | `name` |
| `ibm_metrics_router_route_invalid_rules__action` | `rules.action` |
| `ibm_metrics_router_route_invalid_rules__inclusion_filters__operand` | `rules.inclusion_filters.operand` |
| `ibm_metrics_router_route_invalid_rules__inclusion_filters__operator` | `rules.inclusion_filters.operator` |
| `ibm_metrics_router_route_invalid_rules__inclusion_filters__values` | `rules.inclusion_filters.values` |
| `ibm_metrics_router_route_invalid_rules__targets__id` | `rules.targets.id` |

## Example

### Valid

```hcl
resource "ibm_metrics_router_route" "test" {
  name = "example-name"
  rules {
  }
}
```
