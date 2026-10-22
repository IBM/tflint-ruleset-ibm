# ibm_cis_tls_settings

Schema validation for the ibm_cis_tls_settings resource.

This rule validates `ibm_cis_tls_settings` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cis_tls_settings" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `min_tls_version` | no | one of: `1.1`, `1.2`, `1.3`, `1.4` |
| `tls_1_3` | no | one of: `on`, `off`, `zrt` |
| `universal_ssl` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_tls_settings_invalid_cis_id` | `cis_id` |
| `ibm_cis_tls_settings_invalid_domain_id` | `domain_id` |
| `ibm_cis_tls_settings_invalid_min_tls_version` | `min_tls_version` |
| `ibm_cis_tls_settings_invalid_tls_1_3` | `tls_1_3` |
| `ibm_cis_tls_settings_invalid_universal_ssl` | `universal_ssl` |

## Example

### Valid

```hcl
resource "ibm_cis_tls_settings" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
