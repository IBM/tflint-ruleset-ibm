# ibm_appid_application

Manages ibm_appid_application. The application name to be registered.

This rule validates `ibm_appid_application` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/management/v4/{tenantId}/applications:post` (`app-id-management`).

## Configuration

```hcl
rule "ibm_appid_application" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `client_id` | no | — |
| `discovery_endpoint` | no | — |
| `name` | yes | API: The application name to be registered. Application name cannot exceed 50 characters. |
| `oauth_server_url` | no | — |
| `profiles_url` | no | — |
| `secret` | no | — |
| `tenant_id` | yes | — |
| `type` | no | one of: `regularwebapp`, `singlepageapp` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_application_invalid_name` | `name` |
| `ibm_appid_application_invalid_tenant_id` | `tenant_id` |
| `ibm_appid_application_invalid_type` | `type` |

## Example

### Valid

```hcl
resource "ibm_appid_application" "test" {
  name = "aaaaaaa"
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
