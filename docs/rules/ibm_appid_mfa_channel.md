# ibm_appid_mfa_channel

Schema validation for the ibm_appid_mfa_channel resource.

This rule validates `ibm_appid_mfa_channel` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `introspection-response` (`app-id-auth`).

## Configuration

```hcl
rule "ibm_appid_mfa_channel" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `active` | yes | one of: `email`, `sms` |
| `tenant_id` | yes | — |
| `sms_config` (block) | no | nested block |
| `sms_config.from` | yes | — |
| `sms_config.key` | yes | — |
| `sms_config.secret` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_mfa_channel_invalid_active` | `active` |
| `ibm_appid_mfa_channel_invalid_tenant_id` | `tenant_id` |
| `ibm_appid_mfa_channel_invalid_sms_config__from` | `sms_config.from` |
| `ibm_appid_mfa_channel_invalid_sms_config__key` | `sms_config.key` |
| `ibm_appid_mfa_channel_invalid_sms_config__secret` | `sms_config.secret` |

## Example

### Valid

```hcl
resource "ibm_appid_mfa_channel" "test" {
  active = "email"
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
