# ibm_is_lb_pool_data

Reads ibm_is_lb_pool. The user-defined name for this load balancer pool.

This rule validates `ibm_is_lb_pool` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_lb_pool_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `algorithm` | no | — |
| `client_authentication` | no | — |
| `created_at` | no | — |
| `failsafe_policy` | no | — |
| `health_monitor` | no | — |
| `href` | no | — |
| `identifier` | no | — |
| `instance_group` | no | — |
| `lb` | yes | — |
| `members` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `protocol` | no | — |
| `provisioning_status` | no | — |
| `proxy_protocol` | no | — |
| `server_authentication` | no | — |
| `session_persistence` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_pool_data_invalid_identifier` | `identifier` |
| `ibm_is_lb_pool_data_invalid_lb` | `lb` |
| `ibm_is_lb_pool_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_lb_pool" "test" {
  lb = "example"
}
```
