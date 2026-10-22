# ibm_network_public_ip

Schema validation for the ibm_network_public_ip resource.

This rule validates `ibm_network_public_ip` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_network_public_ip" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `ip_address` | no | — |
| `notes` | no | — |
| `routes_to` | yes | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_network_public_ip_invalid_notes` | `notes` |
| `ibm_network_public_ip_invalid_routes_to` | `routes_to` |
| `ibm_network_public_ip_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_network_public_ip" "test" {
  routes_to = "example"
}
```
