# ibm_logs_router_tenants_data

Reads ibm_logs_router_tenants. Optional: The name of a tenant.

This rule validates `ibm_logs_router_tenants` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/tenants/{tenant_id}:patch` (`logs-router-service-api-logs-router-v1`).

## Configuration

```hcl
rule "ibm_logs_router_tenants_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | yes | pattern `^[a-zA-Z][a-zA-Z0-9-]*$` |
| `region` | yes | — |
| `tenants` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_router_tenants_data_invalid_name` | `name` |
| `ibm_logs_router_tenants_data_invalid_region` | `region` |

## Example

### Valid

```hcl
data "ibm_logs_router_tenants" "test" {
  name = "example-name"
  region = "example"
}
```
