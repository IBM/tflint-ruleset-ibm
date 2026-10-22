# ibm_backup_recovery_connector_access_token

Schema validation for the ibm_backup_recovery_connector_access_token resource.

This rule validates `ibm_backup_recovery_connector_access_token` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `TokenRequestParams` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery_connector_access_token" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_token` | no | — |
| `domain` | no | API: Specifies the domain the user is logging in to. For a local user the domain is LOCAL. For LDAP/AD user, the domain will map to a LDAP connection string. A user is uniquely identified by a combination of username and domain. LOCAL is the default domain. |
| `password` | yes | API: Specifies the password of the IBM user account. |
| `privileges` | no | — |
| `token_type` | no | — |
| `username` | yes | API: Specifies the login name of the IBM user. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_connector_access_token_invalid_domain` | `domain` |
| `ibm_backup_recovery_connector_access_token_invalid_password` | `password` |
| `ibm_backup_recovery_connector_access_token_invalid_username` | `username` |

## Example

### Valid

```hcl
variable "password" {
  sensitive = true
}

resource "ibm_backup_recovery_connector_access_token" "test" {
  password = var.password
  username = "example"
}
```
