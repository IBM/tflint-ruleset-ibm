# ibm_appid_languages

Schema validation for the ibm_appid_languages resource.

This rule validates `ibm_appid_languages` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/management/v4/{tenantId}/config/ui/languages:put` (`app-id-management`).

## Configuration

```hcl
rule "ibm_appid_languages" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `languages` | yes | API: The list of languages that can be used to customize email templates for CLoud Directory. |
| `tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_languages_invalid_languages` | `languages` |
| `ibm_appid_languages_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
resource "ibm_appid_languages" "test" {
  languages = "example"
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
