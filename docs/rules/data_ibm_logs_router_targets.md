# ibm_logs_router_targets_data

Reads ibm_logs_router_targets. The name of the target resource.

This rule validates `ibm_logs_router_targets` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/tenants/{tenant_id}:patch` (`logs-router-service-api-logs-router-v1`).

## Configuration

```hcl
rule "ibm_logs_router_targets_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | no | pattern `^[a-zA-Z][a-zA-Z0-9-]*$` |
| `region` | no | — |
| `targets` | no | — |
| `tenant_id` | no | uuid format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_router_targets_data_invalid_name` | `name` |
| `ibm_logs_router_targets_data_invalid_region` | `region` |
| `ibm_logs_router_targets_data_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
data "ibm_logs_router_targets" "test" {
}
```
