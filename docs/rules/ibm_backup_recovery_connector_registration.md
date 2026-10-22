# ibm_backup_recovery_connector_registration

Schema validation for the ibm_backup_recovery_connector_registration resource.

This rule validates `ibm_backup_recovery_connector_registration` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_connector_registration" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_token` | yes | — |
| `connector_id` | no | — |
| `registration_token` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_connector_registration_invalid_access_token` | `access_token` |
| `ibm_backup_recovery_connector_registration_invalid_connector_id` | `connector_id` |
| `ibm_backup_recovery_connector_registration_invalid_registration_token` | `registration_token` |

## Example

### Valid

```hcl
variable "access_token" {
  sensitive = true
}

variable "registration_token" {
  sensitive = true
}

resource "ibm_backup_recovery_connector_registration" "test" {
  access_token = var.access_token
  registration_token = var.registration_token
}
```
