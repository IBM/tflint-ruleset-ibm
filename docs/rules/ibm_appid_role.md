# ibm_appid_role

Manages ibm_appid_role. Unique role name

This rule validates `ibm_appid_role` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/management/v4/{tenantId}/roles:post` (`app-id-management`).

## Configuration

```hcl
rule "ibm_appid_role" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | API: The role description. |
| `name` | yes | API: The role name. |
| `role_id` | no | — |
| `tenant_id` | yes | — |
| `access` (block) | no | nested block |
| `access.application_id` | yes | API: The application ID. |
| `access.scopes` | yes | API: The scopes. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_role_invalid_description` | `description` |
| `ibm_appid_role_invalid_name` | `name` |
| `ibm_appid_role_invalid_tenant_id` | `tenant_id` |
| `ibm_appid_role_invalid_access__application_id` | `access.application_id` |
| `ibm_appid_role_invalid_access__scopes` | `access.scopes` |

## Example

### Valid

```hcl
resource "ibm_appid_role" "test" {
  name = "aaaaaaa"
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
