# ibm_is_network_acl

Manages ibm_is_network_acl. Network ACL name

This rule validates `ibm_is_network_acl` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `NetworkACLRulePatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_network_acl" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `crn` | no | — |
| `incremental_rule_update` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `vpc` | no | — |
| `rules` (block) | no | nested block |
| `rules.action` | yes | one of: `allow`, `deny`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `rules.code` | no | range 0–255 |
| `rules.destination` | yes | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))?$` |
| `rules.direction` | yes | one of: `inbound`, `outbound`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `rules.name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `rules.port_max` | no | range 1–65535 |
| `rules.port_min` | no | range 1–65535 |
| `rules.protocol` | no | — |
| `rules.source` | yes | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))?$` |
| `rules.source_port_max` | no | range 1–65535 |
| `rules.source_port_min` | no | range 1–65535 |
| `rules.type` | no | range 0–254 |
| `rules.icmp` (block) | no | nested block |
| `rules.icmp.code` | no | range 0–255 |
| `rules.icmp.type` | no | range 0–254 |
| `rules.tcp` (block) | no | nested block |
| `rules.tcp.port_max` | no | range 1–65535 |
| `rules.tcp.port_min` | no | range 1–65535 |
| `rules.tcp.source_port_max` | no | range 1–65535 |
| `rules.tcp.source_port_min` | no | range 1–65535 |
| `rules.udp` (block) | no | nested block |
| `rules.udp.port_max` | no | range 1–65535 |
| `rules.udp.port_min` | no | range 1–65535 |
| `rules.udp.source_port_max` | no | range 1–65535 |
| `rules.udp.source_port_min` | no | range 1–65535 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_network_acl_invalid_access_tags` | `access_tags` |
| `ibm_is_network_acl_invalid_incremental_rule_update` | `incremental_rule_update` |
| `ibm_is_network_acl_invalid_name` | `name` |
| `ibm_is_network_acl_invalid_resource_group` | `resource_group` |
| `ibm_is_network_acl_invalid_tags` | `tags` |
| `ibm_is_network_acl_invalid_vpc` | `vpc` |
| `ibm_is_network_acl_invalid_rules__action` | `rules.action` |
| `ibm_is_network_acl_invalid_rules__code` | `rules.code` |
| `ibm_is_network_acl_invalid_rules__destination` | `rules.destination` |
| `ibm_is_network_acl_invalid_rules__direction` | `rules.direction` |
| `ibm_is_network_acl_invalid_rules__name` | `rules.name` |
| `ibm_is_network_acl_invalid_rules__port_max` | `rules.port_max` |
| `ibm_is_network_acl_invalid_rules__port_min` | `rules.port_min` |
| `ibm_is_network_acl_invalid_rules__protocol` | `rules.protocol` |
| `ibm_is_network_acl_invalid_rules__source` | `rules.source` |
| `ibm_is_network_acl_invalid_rules__source_port_max` | `rules.source_port_max` |
| `ibm_is_network_acl_invalid_rules__source_port_min` | `rules.source_port_min` |
| `ibm_is_network_acl_invalid_rules__type` | `rules.type` |
| `ibm_is_network_acl_invalid_rules__icmp__code` | `rules.icmp.code` |
| `ibm_is_network_acl_invalid_rules__icmp__type` | `rules.icmp.type` |
| `ibm_is_network_acl_invalid_rules__tcp__port_max` | `rules.tcp.port_max` |
| `ibm_is_network_acl_invalid_rules__tcp__port_min` | `rules.tcp.port_min` |
| `ibm_is_network_acl_invalid_rules__tcp__source_port_max` | `rules.tcp.source_port_max` |
| `ibm_is_network_acl_invalid_rules__tcp__source_port_min` | `rules.tcp.source_port_min` |
| `ibm_is_network_acl_invalid_rules__udp__port_max` | `rules.udp.port_max` |
| `ibm_is_network_acl_invalid_rules__udp__port_min` | `rules.udp.port_min` |
| `ibm_is_network_acl_invalid_rules__udp__source_port_max` | `rules.udp.source_port_max` |
| `ibm_is_network_acl_invalid_rules__udp__source_port_min` | `rules.udp.source_port_min` |

## Example

### Valid

```hcl
resource "ibm_is_network_acl" "test" {
}
```
