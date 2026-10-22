# ibm_is_lb_profile_data

Reads ibm_is_lb_profile. The name for this load balancer profile

This rule validates `ibm_is_lb_profile` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_lb_profile_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_modes` | no | — |
| `advanced_health_checks_supported` | no | — |
| `asymmetric_routing_supported` | no | — |
| `asymmetric_routing_supported_type` | no | — |
| `failsafe_policy_actions` | no | — |
| `family` | no | — |
| `fqdn_pool_members_supported` | no | — |
| `href` | no | — |
| `mtls_supported` | no | — |
| `mtls_supported_type` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `route_mode_supported` | no | — |
| `route_mode_type` | no | — |
| `targetable_load_balancer_profiles` | no | — |
| `targetable_resource_types` | no | — |
| `udp_supported` | no | — |
| `udp_supported_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_profile_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_lb_profile" "test" {
  name = "example-name"
}
```
