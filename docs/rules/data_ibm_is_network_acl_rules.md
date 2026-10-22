# ibm_is_network_acl_rules_data

Schema validation for the ibm_is_network_acl_rules data source.

This rule validates `ibm_is_network_acl_rules` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SubnetPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_network_acl_rules_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `direction` | no | one of: `inbound`, `outbound`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `network_acl` | yes | API: Identifies a network ACL by a unique property. |
| `rules` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_network_acl_rules_data_invalid_direction` | `direction` |
| `ibm_is_network_acl_rules_data_invalid_network_acl` | `network_acl` |

## Example

### Valid

```hcl
data "ibm_is_network_acl_rules" "test" {
  network_acl = "example"
}
```
