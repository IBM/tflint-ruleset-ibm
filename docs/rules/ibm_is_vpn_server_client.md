# ibm_is_vpn_server_client

Schema validation for the ibm_is_vpn_server_client resource.

This rule validates `ibm_is_vpn_server_client` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_vpn_server_client" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `delete` | no | — |
| `description` | no | — |
| `status_code` | no | — |
| `vpn_client` | yes | — |
| `vpn_server` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_server_client_invalid_delete` | `delete` |
| `ibm_is_vpn_server_client_invalid_vpn_client` | `vpn_client` |
| `ibm_is_vpn_server_client_invalid_vpn_server` | `vpn_server` |

## Example

### Valid

```hcl
resource "ibm_is_vpn_server_client" "test" {
  vpn_client = "example"
  vpn_server = "example"
}
```
