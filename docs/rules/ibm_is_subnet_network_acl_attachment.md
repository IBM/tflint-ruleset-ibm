# ibm_is_subnet_network_acl_attachment

Manages ibm_is_subnet_network_acl_attachment. Network ACL name

This rule validates `ibm_is_subnet_network_acl_attachment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ClusterNetworkInterfacePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_subnet_network_acl_attachment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `name` | no | — |
| `network_acl` | yes | — |
| `resource_group` | no | — |
| `rules` | no | — |
| `subnet` | yes | API: Identifies a cluster network subnet by a unique property. |
| `vpc` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_subnet_network_acl_attachment_invalid_network_acl` | `network_acl` |
| `ibm_is_subnet_network_acl_attachment_invalid_subnet` | `subnet` |

## Example

### Valid

```hcl
resource "ibm_is_subnet_network_acl_attachment" "test" {
  network_acl = "example"
  subnet = "example"
}
```
