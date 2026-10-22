# ibm_logs_router_route

Manages ibm_logs_router_route. The name of the route.

This rule validates `ibm_logs_router_route` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/tenants/{tenant_id}:patch` (`logs-router-service-api-logs-router-v1`).

## Configuration

```hcl
rule "ibm_logs_router_route" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `crn` | no | — |
| `managed_by` | no | one of: `account`, `enterprise` |
| `name` | yes | pattern `^[a-zA-Z0-9 \-._:]+$` |
| `updated_at` | no | — |
| `rules` (block) | yes | nested block |
| `rules.action` | no | — |
| `rules.inclusion_filters` (block) | no | nested block |
| `rules.inclusion_filters.operand` | yes | — |
| `rules.inclusion_filters.operator` | yes | — |
| `rules.inclusion_filters.values` | yes | — |
| `rules.targets` (block) | yes | nested block |
| `rules.targets.id` | yes | uuid format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_router_route_invalid_managed_by` | `managed_by` |
| `ibm_logs_router_route_invalid_name` | `name` |
| `ibm_logs_router_route_invalid_rules__action` | `rules.action` |
| `ibm_logs_router_route_invalid_rules__inclusion_filters__operand` | `rules.inclusion_filters.operand` |
| `ibm_logs_router_route_invalid_rules__inclusion_filters__operator` | `rules.inclusion_filters.operator` |
| `ibm_logs_router_route_invalid_rules__inclusion_filters__values` | `rules.inclusion_filters.values` |
| `ibm_logs_router_route_invalid_rules__targets__id` | `rules.targets.id` |

## Example

### Valid

```hcl
resource "ibm_logs_router_route" "test" {
  name = "example-name"
  rules {
    targets {
      id = "01234567-89ab-cdef-0123-456789abcdef"
    }
  }
}
```
