# ibm_iam_idp

Manages ibm_iam_idp. Speaking name of the Identity Provider.

This rule validates `ibm_iam_idp` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Idp` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_idp" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | yes | — |
| `active` | no | — |
| `created_at` | no | — |
| `entity_tag` | no | — |
| `idp_id` | no | — |
| `modified_at` | no | — |
| `name` | yes | — |
| `type` | yes | one of: `saml`, `appid`, `ldap` |
| `properties` (block) | no | nested block |
| `properties.idp` (block) | no | nested block |
| `properties.idp.entity_id` | no | — |
| `properties.idp.logout_url` | no | — |
| `properties.idp.redirect_binding_url` | no | — |
| `properties.idp.want_request_signed` | no | — |
| `properties.idp.xml_import` | no | — |
| `properties.sp` (block) | no | nested block |
| `properties.sp.encrypt_response` | no | — |
| `properties.sp.idp_initiated_login_enabled` | no | — |
| `properties.sp.idp_initiated_urls` | no | — |
| `properties.sp.logout_url_enabled_when_available` | no | — |
| `properties.sp.want_assertion_signed` | no | — |
| `properties.sp.want_response_signed` | no | — |
| `secrets` (block) | no | nested block |
| `secrets.idp` (block) | no | nested block |
| `secrets.idp.xml_import` | no | — |
| `secrets.sp` (block) | no | nested block |
| `share_scope` (block) | no | nested block |
| `share_scope.id` | no | — |
| `share_scope.type` | no | one of: `account`, `enterprise` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_idp_invalid_account_id` | `account_id` |
| `ibm_iam_idp_invalid_active` | `active` |
| `ibm_iam_idp_invalid_name` | `name` |
| `ibm_iam_idp_invalid_type` | `type` |
| `ibm_iam_idp_invalid_properties__idp__entity_id` | `properties.idp.entity_id` |
| `ibm_iam_idp_invalid_properties__idp__logout_url` | `properties.idp.logout_url` |
| `ibm_iam_idp_invalid_properties__idp__redirect_binding_url` | `properties.idp.redirect_binding_url` |
| `ibm_iam_idp_invalid_properties__idp__want_request_signed` | `properties.idp.want_request_signed` |
| `ibm_iam_idp_invalid_properties__idp__xml_import` | `properties.idp.xml_import` |
| `ibm_iam_idp_invalid_properties__sp__encrypt_response` | `properties.sp.encrypt_response` |
| `ibm_iam_idp_invalid_properties__sp__idp_initiated_login_enabled` | `properties.sp.idp_initiated_login_enabled` |
| `ibm_iam_idp_invalid_properties__sp__idp_initiated_urls` | `properties.sp.idp_initiated_urls` |
| `ibm_iam_idp_invalid_properties__sp__logout_url_enabled_when_available` | `properties.sp.logout_url_enabled_when_available` |
| `ibm_iam_idp_invalid_properties__sp__want_assertion_signed` | `properties.sp.want_assertion_signed` |
| `ibm_iam_idp_invalid_properties__sp__want_response_signed` | `properties.sp.want_response_signed` |
| `ibm_iam_idp_invalid_secrets__idp__xml_import` | `secrets.idp.xml_import` |
| `ibm_iam_idp_invalid_share_scope__id` | `share_scope.id` |
| `ibm_iam_idp_invalid_share_scope__type` | `share_scope.type` |

## Example

### Valid

```hcl
resource "ibm_iam_idp" "test" {
  account_id = "01234567-89ab-cdef-0123-456789abcdef"
  name = "example"
  type = "saml"
}
```
