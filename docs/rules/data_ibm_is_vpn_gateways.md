# ibm_is_vpn_gateways_data

Schema validation for the ibm_is_vpn_gateways data source.

This rule validates `ibm_is_vpn_gateways` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerTrustedPlatformModulePatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpn_gateways_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `mode` | no | one of: `disabled`, `tpm_2`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `resource_group` | no | — |
| `vpn_gateways` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpn_gateways_data_invalid_mode` | `mode` |
| `ibm_is_vpn_gateways_data_invalid_resource_group` | `resource_group` |

## Example

### Valid

```hcl
data "ibm_is_vpn_gateways" "test" {
}
```
