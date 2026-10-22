# ibm_appid_idp_saml

Schema validation for the ibm_appid_idp_saml resource.

This rule validates `ibm_appid_idp_saml` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `leaf-unique` (`app-id-management`).

## Configuration

```hcl
rule "ibm_appid_idp_saml" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `is_active` | yes | — |
| `tenant_id` | yes | — |
| `config` (block) | no | nested block |
| `config.certificates` | yes | — |
| `config.display_name` | no | — |
| `config.encrypt_response` | no | — |
| `config.entity_id` | yes | — |
| `config.include_scoping` | no | — |
| `config.sign_in_url` | yes | — |
| `config.sign_request` | no | — |
| `config.authn_context` (block) | no | nested block |
| `config.authn_context.class` | no | — |
| `config.authn_context.comparison` | no | one of: `exact`, `maximum`, `minimum`, `better` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_idp_saml_invalid_is_active` | `is_active` |
| `ibm_appid_idp_saml_invalid_tenant_id` | `tenant_id` |
| `ibm_appid_idp_saml_invalid_config__certificates` | `config.certificates` |
| `ibm_appid_idp_saml_invalid_config__display_name` | `config.display_name` |
| `ibm_appid_idp_saml_invalid_config__encrypt_response` | `config.encrypt_response` |
| `ibm_appid_idp_saml_invalid_config__entity_id` | `config.entity_id` |
| `ibm_appid_idp_saml_invalid_config__include_scoping` | `config.include_scoping` |
| `ibm_appid_idp_saml_invalid_config__sign_in_url` | `config.sign_in_url` |
| `ibm_appid_idp_saml_invalid_config__sign_request` | `config.sign_request` |
| `ibm_appid_idp_saml_invalid_config__authn_context__class` | `config.authn_context.class` |
| `ibm_appid_idp_saml_invalid_config__authn_context__comparison` | `config.authn_context.comparison` |

## Example

### Valid

```hcl
resource "ibm_appid_idp_saml" "test" {
  is_active = true
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
