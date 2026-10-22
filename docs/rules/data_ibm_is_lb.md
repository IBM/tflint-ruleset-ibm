# ibm_is_lb_data

Reads ibm_is_lb. Load Balancer name

This rule validates `ibm_is_lb` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_lb_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_mode` | no | — |
| `access_tags` | no | — |
| `advanced_health_checks_supported` | no | — |
| `asymmetric_routing_supported` | no | — |
| `attached_load_balancer_pool_members` | no | — |
| `availability` | no | — |
| `crn` | no | — |
| `dns` | no | — |
| `failsafe_policy_actions` | no | — |
| `fqdn_pool_members_supported` | no | — |
| `hostname` | no | — |
| `instance_groups_supported` | no | — |
| `listeners` | no | — |
| `logging` | no | — |
| `mtls_supported` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `operating_status` | no | — |
| `pools` | no | — |
| `private_ip` | no | — |
| `private_ips` | no | — |
| `profile` | no | — |
| `public_ips` | no | — |
| `resource_controller_url` | no | — |
| `resource_group` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `route_mode` | no | — |
| `security_group_supported` | no | — |
| `security_groups` | no | — |
| `source_ip_session_persistence_supported` | no | — |
| `status` | no | — |
| `subnets` | no | — |
| `tags` | no | — |
| `type` | no | — |
| `udp_supported` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_lb" "test" {
  name = "example-name"
}
```
