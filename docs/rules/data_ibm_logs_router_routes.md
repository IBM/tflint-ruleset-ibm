# ibm_logs_router_routes_data

Reads ibm_logs_router_routes. The name of the route.

This rule validates `ibm_logs_router_routes` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/tenants/{tenant_id}:patch` (`logs-router-service-api-logs-router-v1`).

## Configuration

```hcl
rule "ibm_logs_router_routes_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | no | pattern `^[a-zA-Z][a-zA-Z0-9-]*$` |
| `routes` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_router_routes_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_logs_router_routes" "test" {
}
```
