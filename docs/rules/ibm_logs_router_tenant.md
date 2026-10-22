# ibm_logs_router_tenant

Manages ibm_logs_router_tenant. The name for this tenant.

This rule validates `ibm_logs_router_tenant` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/tenants/{tenant_id}/targets/{target_id}:patch` (`logs-router-service-api-logs-router-v1`).

## Configuration

```hcl
rule "ibm_logs_router_tenant" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `crn` | no | — |
| `etag` | no | — |
| `name` | yes | pattern `[a-z,A-Z,0-9,-,.]` |
| `region` | yes | pattern `^[a-zA-Z0-9 -._:]+$` |
| `updated_at` | no | — |
| `write_status` | no | — |
| `targets` (block) | yes | nested block |
| `targets.log_sink_crn` | no | crn format |
| `targets.name` | no | pattern `^[a-zA-Z][a-zA-Z0-9-]*$` |
| `targets.type` | no | — |
| `targets.parameters` (block) | yes | nested block |
| `targets.parameters.host` | yes | pattern `^[a-zA-Z0-9-.:/]+$` |
| `targets.parameters.port` | yes | range 1–65535; port format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_router_tenant_invalid_name` | `name` |
| `ibm_logs_router_tenant_invalid_region` | `region` |
| `ibm_logs_router_tenant_invalid_targets__log_sink_crn` | `targets.log_sink_crn` |
| `ibm_logs_router_tenant_invalid_targets__name` | `targets.name` |
| `ibm_logs_router_tenant_invalid_targets__type` | `targets.type` |
| `ibm_logs_router_tenant_invalid_targets__parameters__host` | `targets.parameters.host` |
| `ibm_logs_router_tenant_invalid_targets__parameters__port` | `targets.parameters.port` |

## Example

### Valid

```hcl
resource "ibm_logs_router_tenant" "test" {
  name = "example-name"
  region = "example-name"
  targets {
    parameters {
      host = "example-name"
      port = 80
    }
  }
}
```
