# ibm_backup_recovery_protection_policy_data

Reads ibm_backup_recovery_protection_policy. Specifies the name of the Protection Policy.

This rule validates `ibm_backup_recovery_protection_policy` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_protection_policy_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `backup_policy` | no | — |
| `blackout_window` | no | — |
| `cascaded_targets_config` | no | — |
| `data_lock` | no | — |
| `description` | no | — |
| `extended_retention` | no | — |
| `is_cbs_enabled` | no | — |
| `is_replicated` | no | — |
| `is_usable` | no | — |
| `last_modification_time_usecs` | no | — |
| `name` | no | — |
| `num_protected_objects` | no | — |
| `num_protection_groups` | no | — |
| `policy_id` | no | — |
| `protection_policy_id` | yes | — |
| `remote_target_policy` | no | — |
| `request_initiator_type` | no | — |
| `retry_options` | no | — |
| `template_id` | no | — |
| `version` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_protection_policy_data_invalid_protection_policy_id` | `protection_policy_id` |
| `ibm_backup_recovery_protection_policy_data_invalid_request_initiator_type` | `request_initiator_type` |
| `ibm_backup_recovery_protection_policy_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_protection_policy" "test" {
  protection_policy_id = "01234567-89ab-cdef-0123-456789abcdef"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
