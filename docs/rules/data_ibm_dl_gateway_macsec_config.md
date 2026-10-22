# ibm_dl_gateway_macsec_config_data

Schema validation for the ibm_dl_gateway_macsec_config data source.

This rule validates `ibm_dl_gateway_macsec_config` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `GatewayBfdConfigActionTemplate` (`direct_link`).

## Configuration

```hcl
rule "ibm_dl_gateway_macsec_config_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `active` | no | — |
| `cipher_suite` | no | — |
| `confidentiality_offset` | no | — |
| `created_at` | no | — |
| `gateway` | yes | — |
| `key_server_priority` | no | — |
| `sak_rekey` | no | — |
| `security_policy` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `updated_at` | no | — |
| `window_size` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_gateway_macsec_config_data_invalid_gateway` | `gateway` |

## Example

### Valid

```hcl
data "ibm_dl_gateway_macsec_config" "test" {
  gateway = "example"
}
```
