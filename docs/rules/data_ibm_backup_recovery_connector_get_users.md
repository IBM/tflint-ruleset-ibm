# ibm_backup_recovery_connector_get_users_data

Schema validation for the ibm_backup_recovery_connector_get_users data source.

This rule validates `ibm_backup_recovery_connector_get_users` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_connector_get_users_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `all_under_hierarchy` | no | — |
| `domain` | no | — |
| `email_addresses` | no | — |
| `partial_match` | no | — |
| `session_name_cookie` | yes | — |
| `tenant_ids` | no | — |
| `usernames` | no | — |
| `users` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_connector_get_users_data_invalid_all_under_hierarchy` | `all_under_hierarchy` |
| `ibm_backup_recovery_connector_get_users_data_invalid_domain` | `domain` |
| `ibm_backup_recovery_connector_get_users_data_invalid_email_addresses` | `email_addresses` |
| `ibm_backup_recovery_connector_get_users_data_invalid_partial_match` | `partial_match` |
| `ibm_backup_recovery_connector_get_users_data_invalid_session_name_cookie` | `session_name_cookie` |
| `ibm_backup_recovery_connector_get_users_data_invalid_tenant_ids` | `tenant_ids` |
| `ibm_backup_recovery_connector_get_users_data_invalid_usernames` | `usernames` |

## Example

### Valid

```hcl
variable "session_name_cookie" {
  sensitive = true
}

data "ibm_backup_recovery_connector_get_users" "test" {
  session_name_cookie = var.session_name_cookie
}
```
