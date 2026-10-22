# ibm_is_vpn_server_client_data

Schema validation for the ibm_is_vpn_server_client data source.

This rule validates `ibm_is_vpn_server_client` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_vpn_server_client_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `client_ip` | no | — |
| `common_name` | no | — |
| `created_at` | no | — |
| `disconnected_at` | no | — |
| `href` | no | — |
| `identifier` | yes | — |
| `remote_ip` | no | — |
| `remote_port` | no | — |
| `resource_type` | no | — |
| `status` | no | — |
| `username` | no | — |
| `vpn_server` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_server_client_data_invalid_identifier` | `identifier` |
| `ibm_is_vpn_server_client_data_invalid_vpn_server` | `vpn_server` |

## Example

### Valid

```hcl
data "ibm_is_vpn_server_client" "test" {
  identifier = "example"
  vpn_server = "example"
}
```
