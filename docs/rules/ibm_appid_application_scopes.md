# ibm_appid_application_scopes

Schema validation for the ibm_appid_application_scopes resource.

This rule validates `ibm_appid_application_scopes` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/management/v4/{tenantId}/applications/{clientId}/scopes:put` (`app-id-management`).

## Configuration

```hcl
rule "ibm_appid_application_scopes" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `client_id` | yes | — |
| `scopes` | yes | API: The scopes. |
| `tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_application_scopes_invalid_client_id` | `client_id` |
| `ibm_appid_application_scopes_invalid_scopes` | `scopes` |
| `ibm_appid_application_scopes_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
resource "ibm_appid_application_scopes" "test" {
  client_id = "01234567-89ab-cdef-0123-456789abcdef"
  scopes = "aaaaaaa"
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
