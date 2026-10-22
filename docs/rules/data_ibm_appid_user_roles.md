# ibm_appid_user_roles_data

Schema validation for the ibm_appid_user_roles data source.

This rule validates `ibm_appid_user_roles` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `getTemplate` (`app-id-management`).

## Configuration

```hcl
rule "ibm_appid_user_roles_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `roles` | no | — |
| `subject` | yes | API: The subject of the email. |
| `tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_user_roles_data_invalid_subject` | `subject` |
| `ibm_appid_user_roles_data_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
data "ibm_appid_user_roles" "test" {
  subject = "example"
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
