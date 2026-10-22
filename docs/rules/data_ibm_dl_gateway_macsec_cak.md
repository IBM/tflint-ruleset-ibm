# ibm_dl_gateway_macsec_cak_data

Reads ibm_dl_gateway_macsec_cak. The name identifies the connectivity association key (CAK) within the MACsec key chain.

This rule validates `ibm_dl_gateway_macsec_cak` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dl_gateway_macsec_cak_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cak_id` | yes | — |
| `created_at` | no | — |
| `gateway` | yes | — |
| `key` | no | — |
| `name` | no | — |
| `session` | no | — |
| `status` | no | — |
| `updated_at` | no | — |
| `active_delta` (block) | no | nested block |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_gateway_macsec_cak_data_invalid_cak_id` | `cak_id` |
| `ibm_dl_gateway_macsec_cak_data_invalid_gateway` | `gateway` |

## Example

### Valid

```hcl
data "ibm_dl_gateway_macsec_cak" "test" {
  cak_id = "01234567-89ab-cdef-0123-456789abcdef"
  gateway = "example"
}
```
