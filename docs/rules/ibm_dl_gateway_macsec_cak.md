# ibm_dl_gateway_macsec_cak

Manages ibm_dl_gateway_macsec_cak. The name identifies the connectivity association key (CAK) within the MACsec key chain.

This rule validates `ibm_dl_gateway_macsec_cak` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `GatewayMacsecCakPrototype` (`direct_link`).

## Configuration

```hcl
rule "ibm_dl_gateway_macsec_cak" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cak` | no | — |
| `cak_id` | no | — |
| `gateway` | yes | — |
| `name` | yes | pattern `^([0-9a-fA-F]{2}){1,32}$` |
| `session` | yes | one of: `primary`, `fallback` |
| `key` (block) | yes | nested block |
| `key.crn` | yes | crn format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_gateway_macsec_cak_invalid_gateway` | `gateway` |
| `ibm_dl_gateway_macsec_cak_invalid_name` | `name` |
| `ibm_dl_gateway_macsec_cak_invalid_session` | `session` |
| `ibm_dl_gateway_macsec_cak_invalid_key__crn` | `key.crn` |

## Example

### Valid

```hcl
resource "ibm_dl_gateway_macsec_cak" "test" {
  gateway = "example"
  name = "10"
  session = "primary"
  key {
    crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  }
}
```
