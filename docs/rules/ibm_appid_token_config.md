# ibm_appid_token_config

Schema validation for the ibm_appid_token_config resource.

This rule validates `ibm_appid_token_config` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `leaf-unique` (`app-id-management`).

## Configuration

```hcl
rule "ibm_appid_token_config" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_token_expires_in` | no | — |
| `anonymous_access_enabled` | no | — |
| `anonymous_token_expires_in` | no | — |
| `refresh_token_enabled` | no | — |
| `refresh_token_expires_in` | no | — |
| `tenant_id` | yes | — |
| `access_token_claim` (block) | no | nested block |
| `access_token_claim.destination_claim` | no | — |
| `access_token_claim.source` | yes | one of: `saml`, `cloud_directory`, `appid_custom`, `facebook`, `google`, `ibmid`, `attributes`, `roles` |
| `access_token_claim.source_claim` | no | — |
| `id_token_claim` (block) | no | nested block |
| `id_token_claim.destination_claim` | no | — |
| `id_token_claim.source` | yes | one of: `saml`, `cloud_directory`, `appid_custom`, `facebook`, `google`, `ibmid`, `attributes`, `roles` |
| `id_token_claim.source_claim` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_token_config_invalid_access_token_expires_in` | `access_token_expires_in` |
| `ibm_appid_token_config_invalid_anonymous_access_enabled` | `anonymous_access_enabled` |
| `ibm_appid_token_config_invalid_anonymous_token_expires_in` | `anonymous_token_expires_in` |
| `ibm_appid_token_config_invalid_refresh_token_enabled` | `refresh_token_enabled` |
| `ibm_appid_token_config_invalid_refresh_token_expires_in` | `refresh_token_expires_in` |
| `ibm_appid_token_config_invalid_tenant_id` | `tenant_id` |
| `ibm_appid_token_config_invalid_access_token_claim__destination_claim` | `access_token_claim.destination_claim` |
| `ibm_appid_token_config_invalid_access_token_claim__source` | `access_token_claim.source` |
| `ibm_appid_token_config_invalid_access_token_claim__source_claim` | `access_token_claim.source_claim` |
| `ibm_appid_token_config_invalid_id_token_claim__destination_claim` | `id_token_claim.destination_claim` |
| `ibm_appid_token_config_invalid_id_token_claim__source` | `id_token_claim.source` |
| `ibm_appid_token_config_invalid_id_token_claim__source_claim` | `id_token_claim.source_claim` |

## Example

### Valid

```hcl
resource "ibm_appid_token_config" "test" {
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
