# ibm_is_lb

Manages ibm_is_lb. Load Balancer name

This rule validates `ibm_is_lb` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_lb" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_mode` | no | — |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `advanced_health_checks_supported` | no | — |
| `asymmetric_routing_supported` | no | — |
| `attached_load_balancer_pool_members` | no | — |
| `availability` | no | — |
| `crn` | no | — |
| `failsafe_policy_actions` | no | — |
| `fqdn_pool_members_supported` | no | — |
| `hostname` | no | — |
| `instance_groups_supported` | no | — |
| `logging` | no | API: The logging configuration to use for this load balancer. See [VPC Datapath Logging](https://cloud.ibm.com/docs/vpc?topic=vpc-datapath-logging) on the logging format, fields and permitted values. If unspecified, `datapath.active` will be `false`. To activate logging, the load balancer profile must support the specified logging type. |
| `mtls_supported` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `operating_status` | no | — |
| `private_ip` | no | — |
| `private_ips` | no | — |
| `profile` | no | one of: `network-fixed`, `network-private-path` |
| `public_ips` | no | — |
| `resource_controller_url` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `route_mode` | no | API: Indicates whether [route mode](https://cloud.ibm.com/docs/vpc?topic=vpc-nlb-vnf&interface=ui) is enabled for this load balancer. At present, public load balancers are not supported with route mode enabled. |
| `security_group_supported` | no | — |
| `security_groups` | no | API: Identifies a security group by a unique property. |
| `source_ip_session_persistence_supported` | no | — |
| `status` | no | — |
| `subnets` | yes | API: Identifies a subnet by a unique property. |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `type` | no | one of: `public`, `private`, `private_path`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `udp_supported` | no | — |
| `version` | no | — |
| `dns` (block) | no | nested block |
| `dns.instance_crn` | yes | crn format |
| `dns.zone_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_invalid_access_tags` | `access_tags` |
| `ibm_is_lb_invalid_logging` | `logging` |
| `ibm_is_lb_invalid_name` | `name` |
| `ibm_is_lb_invalid_profile` | `profile` |
| `ibm_is_lb_invalid_resource_group` | `resource_group` |
| `ibm_is_lb_invalid_route_mode` | `route_mode` |
| `ibm_is_lb_invalid_security_groups` | `security_groups` |
| `ibm_is_lb_invalid_subnets` | `subnets` |
| `ibm_is_lb_invalid_tags` | `tags` |
| `ibm_is_lb_invalid_type` | `type` |
| `ibm_is_lb_invalid_dns__instance_crn` | `dns.instance_crn` |
| `ibm_is_lb_invalid_dns__zone_id` | `dns.zone_id` |

## Example

### Valid

```hcl
resource "ibm_is_lb" "test" {
  name = "example-name"
  subnets = "example"
}
```
