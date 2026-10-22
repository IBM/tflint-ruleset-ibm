# ibm_firewall_policy

Schema validation for the ibm_firewall_policy resource.

This rule validates `ibm_firewall_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_firewall_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `firewall_id` | yes | — |
| `tags` | no | — |
| `rules` (block) | yes | nested block |
| `rules.action` | yes | — |
| `rules.dst_ip_address` | yes | — |
| `rules.dst_ip_cidr` | yes | cidr format |
| `rules.dst_port_range_end` | no | — |
| `rules.dst_port_range_start` | no | — |
| `rules.notes` | no | — |
| `rules.protocol` | yes | — |
| `rules.src_ip_address` | yes | — |
| `rules.src_ip_cidr` | yes | cidr format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_firewall_policy_invalid_firewall_id` | `firewall_id` |
| `ibm_firewall_policy_invalid_tags` | `tags` |
| `ibm_firewall_policy_invalid_rules__action` | `rules.action` |
| `ibm_firewall_policy_invalid_rules__dst_ip_address` | `rules.dst_ip_address` |
| `ibm_firewall_policy_invalid_rules__dst_ip_cidr` | `rules.dst_ip_cidr` |
| `ibm_firewall_policy_invalid_rules__dst_port_range_end` | `rules.dst_port_range_end` |
| `ibm_firewall_policy_invalid_rules__dst_port_range_start` | `rules.dst_port_range_start` |
| `ibm_firewall_policy_invalid_rules__notes` | `rules.notes` |
| `ibm_firewall_policy_invalid_rules__protocol` | `rules.protocol` |
| `ibm_firewall_policy_invalid_rules__src_ip_address` | `rules.src_ip_address` |
| `ibm_firewall_policy_invalid_rules__src_ip_cidr` | `rules.src_ip_cidr` |

## Example

### Valid

```hcl
resource "ibm_firewall_policy" "test" {
  firewall_id = 1
  rules {
    action = "example"
    dst_ip_address = "example"
    dst_ip_cidr = "10.0.0.0/24"
    protocol = "example"
    src_ip_address = "example"
    src_ip_cidr = "10.0.0.0/24"
  }
}
```
