# ibm_appid_application_roles

Schema validation for the ibm_appid_application_roles resource.

This rule validates `ibm_appid_application_roles` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `getUserRolesResponse` (`app-id-management`).

## Configuration

```hcl
rule "ibm_appid_application_roles" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `client_id` | yes | — |
| `roles` | yes | API: The user's roles. |
| `tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_application_roles_invalid_client_id` | `client_id` |
| `ibm_appid_application_roles_invalid_roles` | `roles` |
| `ibm_appid_application_roles_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
resource "ibm_appid_application_roles" "test" {
  client_id = "01234567-89ab-cdef-0123-456789abcdef"
  roles = "example"
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
