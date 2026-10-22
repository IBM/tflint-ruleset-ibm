# ibm_appid_password_regex

Schema validation for the ibm_appid_password_regex resource.

This rule validates `ibm_appid_password_regex` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `passwordRegexConfigParams` (`app-id-management`).

## Configuration

```hcl
rule "ibm_appid_password_regex" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `base64_encoded_regex` | no | — |
| `error_message` | no | API: The error message. |
| `regex` | yes | API: The regular expression that is used by App ID for password strength validation. |
| `tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_password_regex_invalid_error_message` | `error_message` |
| `ibm_appid_password_regex_invalid_regex` | `regex` |
| `ibm_appid_password_regex_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
resource "ibm_appid_password_regex" "test" {
  regex = "example"
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
