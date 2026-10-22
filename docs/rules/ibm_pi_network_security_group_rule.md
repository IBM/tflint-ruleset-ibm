# ibm_pi_network_security_group_rule

Manages ibm_pi_network_security_group_rule. The name of the network security group.

This rule validates `ibm_pi_network_security_group_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Network` (`power-cloud`).

## Configuration

```hcl
rule "ibm_pi_network_security_group_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `default` | no | — |
| `members` | no | — |
| `name` | no | — |
| `network_security_group_id` | no | — |
| `pi_action` | no | one of: `allow`, `deny` |
| `pi_cloud_instance_id` | yes | — |
| `pi_network_security_group_id` | yes | — |
| `pi_network_security_group_rule_id` | no | — |
| `rules` | no | — |
| `user_tags` | no | — |
| `pi_destination_port` (block) | no | nested block |
| `pi_destination_port.maximum` | no | max 65535 |
| `pi_destination_port.minimum` | no | min 1 |
| `pi_destination_ports` (block) | no | nested block |
| `pi_destination_ports.maximum` | no | max 65535 |
| `pi_destination_ports.minimum` | no | min 1 |
| `pi_protocol` (block) | no | nested block |
| `pi_protocol.icmp_type` | no | one of: `all`, `destination-unreach`, `echo`, `echo-reply`, `source-quench`, `time-exceeded` |
| `pi_protocol.type` | yes | one of: `all`, `icmp`, `tcp`, `udp` |
| `pi_protocol.tcp_flags` (block) | no | nested block |
| `pi_protocol.tcp_flags.flag` | yes | one of: `syn`, `ack`, `fin`, `rst` |
| `pi_remote` (block) | no | nested block |
| `pi_remote.id` | no | API: Unique identifier for the trusted profile |
| `pi_remote.type` | no | one of: `default-network-address-group`, `network-address-group`, `network-security-group` |
| `pi_source_port` (block) | no | nested block |
| `pi_source_port.maximum` | no | max 65535 |
| `pi_source_port.minimum` | no | min 1 |
| `pi_source_ports` (block) | no | nested block |
| `pi_source_ports.maximum` | no | max 65535 |
| `pi_source_ports.minimum` | no | min 1 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_network_security_group_rule_invalid_pi_action` | `pi_action` |
| `ibm_pi_network_security_group_rule_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_network_security_group_rule_invalid_pi_network_security_group_id` | `pi_network_security_group_id` |
| `ibm_pi_network_security_group_rule_invalid_pi_network_security_group_rule_id` | `pi_network_security_group_rule_id` |
| `ibm_pi_network_security_group_rule_invalid_pi_destination_port__maximum` | `pi_destination_port.maximum` |
| `ibm_pi_network_security_group_rule_invalid_pi_destination_port__minimum` | `pi_destination_port.minimum` |
| `ibm_pi_network_security_group_rule_invalid_pi_destination_ports__maximum` | `pi_destination_ports.maximum` |
| `ibm_pi_network_security_group_rule_invalid_pi_destination_ports__minimum` | `pi_destination_ports.minimum` |
| `ibm_pi_network_security_group_rule_invalid_pi_protocol__icmp_type` | `pi_protocol.icmp_type` |
| `ibm_pi_network_security_group_rule_invalid_pi_protocol__type` | `pi_protocol.type` |
| `ibm_pi_network_security_group_rule_invalid_pi_protocol__tcp_flags__flag` | `pi_protocol.tcp_flags.flag` |
| `ibm_pi_network_security_group_rule_invalid_pi_remote__id` | `pi_remote.id` |
| `ibm_pi_network_security_group_rule_invalid_pi_remote__type` | `pi_remote.type` |
| `ibm_pi_network_security_group_rule_invalid_pi_source_port__maximum` | `pi_source_port.maximum` |
| `ibm_pi_network_security_group_rule_invalid_pi_source_port__minimum` | `pi_source_port.minimum` |
| `ibm_pi_network_security_group_rule_invalid_pi_source_ports__maximum` | `pi_source_ports.maximum` |
| `ibm_pi_network_security_group_rule_invalid_pi_source_ports__minimum` | `pi_source_ports.minimum` |

## Example

### Valid

```hcl
resource "ibm_pi_network_security_group_rule" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_network_security_group_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
