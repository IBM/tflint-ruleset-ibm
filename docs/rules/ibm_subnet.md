# ibm_subnet

Schema validation for the ibm_subnet resource.

This rule validates `ibm_subnet` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_subnet" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `capacity` | yes | — |
| `endpoint_ip` | no | — |
| `ip_version` | no | — |
| `notes` | no | — |
| `private` | no | — |
| `subnet_cidr` | no | — |
| `tags` | no | — |
| `type` | yes | — |
| `vlan_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_subnet_invalid_capacity` | `capacity` |
| `ibm_subnet_invalid_endpoint_ip` | `endpoint_ip` |
| `ibm_subnet_invalid_ip_version` | `ip_version` |
| `ibm_subnet_invalid_notes` | `notes` |
| `ibm_subnet_invalid_private` | `private` |
| `ibm_subnet_invalid_tags` | `tags` |
| `ibm_subnet_invalid_type` | `type` |
| `ibm_subnet_invalid_vlan_id` | `vlan_id` |

## Example

### Valid

```hcl
resource "ibm_subnet" "test" {
  capacity = 1
  type = "example"
}
```
