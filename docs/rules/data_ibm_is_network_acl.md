# ibm_is_network_acl_data

Reads ibm_is_network_acl. The network acl name.

This rule validates `ibm_is_network_acl` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SubnetPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_network_acl_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `href` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `network_acl` | no | API: Identifies a network ACL by a unique property. |
| `resource_group` | no | — |
| `rules` | no | — |
| `subnets` | no | — |
| `vpc` | no | — |
| `vpc_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_network_acl_data_invalid_name` | `name` |
| `ibm_is_network_acl_data_invalid_network_acl` | `network_acl` |
| `ibm_is_network_acl_data_invalid_vpc_name` | `vpc_name` |

## Example

### Valid

```hcl
data "ibm_is_network_acl" "test" {
}
```
