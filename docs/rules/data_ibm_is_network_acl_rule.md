# ibm_is_network_acl_rule_data

Reads ibm_is_network_acl_rule. The user-defined name for this rule

This rule validates `ibm_is_network_acl_rule` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_network_acl_rule_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | no | — |
| `before` | no | — |
| `destination` | no | — |
| `direction` | no | — |
| `href` | no | — |
| `icmp` | no | — |
| `ip_version` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `network_acl` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `protocol` | no | — |
| `rule_id` | no | — |
| `source` | no | — |
| `tcp` | no | — |
| `udp` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_network_acl_rule_data_invalid_name` | `name` |
| `ibm_is_network_acl_rule_data_invalid_network_acl` | `network_acl` |

## Example

### Valid

```hcl
data "ibm_is_network_acl_rule" "test" {
  name = "example-name"
  network_acl = "example-name"
}
```
