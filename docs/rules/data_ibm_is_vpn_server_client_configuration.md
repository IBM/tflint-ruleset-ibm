# ibm_is_vpn_server_client_configuration_data

Schema validation for the ibm_is_vpn_server_client_configuration data source.

This rule validates `ibm_is_vpn_server_client_configuration` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_vpn_server_client_configuration_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `file_path` | no | — |
| `vpn_server` | yes | — |
| `vpn_server_client_configuration` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_server_client_configuration_data_invalid_file_path` | `file_path` |
| `ibm_is_vpn_server_client_configuration_data_invalid_vpn_server` | `vpn_server` |

## Example

### Valid

```hcl
data "ibm_is_vpn_server_client_configuration" "test" {
  vpn_server = "example"
}
```
