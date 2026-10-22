# ibm_dl_gateway_macsec_config

Schema validation for the ibm_dl_gateway_macsec_config resource.

This rule validates `ibm_dl_gateway_macsec_config` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/gateways/{id}/macsec:patch` (`direct_link`).

## Configuration

```hcl
rule "ibm_dl_gateway_macsec_config" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `active` | yes | API: Sets the MACsec feature to be active (true) or inactive (false) for a direct link. |
| `cipher_suite` | no | — |
| `confidentiality_offset` | no | — |
| `created_at` | no | — |
| `gateway` | yes | — |
| `key_server_priority` | no | — |
| `security_policy` | yes | one of: `must_secure`, `should_secure` |
| `status` | no | — |
| `status_reasons` | no | — |
| `updated_at` | no | — |
| `window_size` | no | range 0–65535 |
| `caks` (block) | yes | nested block |
| `caks.name` | yes | — |
| `caks.session` | yes | one of: `primary`, `fallback` |
| `caks.key` (block) | yes | nested block |
| `caks.key.crn` | yes | crn format |
| `sak_rekey` (block) | yes | nested block |
| `sak_rekey.interval` | no | range 60–65535 |
| `sak_rekey.mode` | yes | one of: `packet_number_rollover` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_gateway_macsec_config_invalid_active` | `active` |
| `ibm_dl_gateway_macsec_config_invalid_gateway` | `gateway` |
| `ibm_dl_gateway_macsec_config_invalid_security_policy` | `security_policy` |
| `ibm_dl_gateway_macsec_config_invalid_window_size` | `window_size` |
| `ibm_dl_gateway_macsec_config_invalid_caks__name` | `caks.name` |
| `ibm_dl_gateway_macsec_config_invalid_caks__session` | `caks.session` |
| `ibm_dl_gateway_macsec_config_invalid_caks__key__crn` | `caks.key.crn` |
| `ibm_dl_gateway_macsec_config_invalid_sak_rekey__interval` | `sak_rekey.interval` |
| `ibm_dl_gateway_macsec_config_invalid_sak_rekey__mode` | `sak_rekey.mode` |

## Example

### Valid

```hcl
resource "ibm_dl_gateway_macsec_config" "test" {
  active = true
  gateway = "example"
  security_policy = "must_secure"
  caks {
    name = "example"
    session = "primary"
    key {
      crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
    }
  }
  sak_rekey {
    mode = "packet_number_rollover"
  }
}
```
