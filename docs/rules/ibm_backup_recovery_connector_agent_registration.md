# ibm_backup_recovery_connector_agent_registration

Schema validation for the ibm_backup_recovery_connector_agent_registration resource.

This rule validates `ibm_backup_recovery_connector_agent_registration` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_connector_agent_registration" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `connection_name` | yes | — |
| `join_existing_connection` | no | — |
| `registration_status` | no | — |
| `registration_token` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_connector_agent_registration_invalid_connection_name` | `connection_name` |
| `ibm_backup_recovery_connector_agent_registration_invalid_join_existing_connection` | `join_existing_connection` |
| `ibm_backup_recovery_connector_agent_registration_invalid_registration_token` | `registration_token` |

## Example

### Valid

```hcl
variable "registration_token" {
  sensitive = true
}

resource "ibm_backup_recovery_connector_agent_registration" "test" {
  connection_name = "example"
  registration_token = var.registration_token
}
```
