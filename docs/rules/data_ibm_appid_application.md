# ibm_appid_application_data

Reads ibm_appid_application. The application name

This rule validates `ibm_appid_application` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_appid_application_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `client_id` | yes | — |
| `discovery_endpoint` | no | — |
| `name` | no | — |
| `oauth_server_url` | no | — |
| `profiles_url` | no | — |
| `secret` | no | — |
| `tenant_id` | yes | — |
| `type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_application_data_invalid_client_id` | `client_id` |
| `ibm_appid_application_data_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
data "ibm_appid_application" "test" {
  client_id = "01234567-89ab-cdef-0123-456789abcdef"
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
