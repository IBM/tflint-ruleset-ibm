# ibm_backup_recovery_protection_policies_data

Schema validation for the ibm_backup_recovery_protection_policies data source.

This rule validates `ibm_backup_recovery_protection_policies` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_protection_policies_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `exclude_linked_policies` | no | — |
| `ids` | no | — |
| `include_replicated_policies` | no | — |
| `include_stats` | no | — |
| `policies` | no | — |
| `policy_names` | no | — |
| `request_initiator_type` | no | — |
| `types` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_protection_policies_data_invalid_exclude_linked_policies` | `exclude_linked_policies` |
| `ibm_backup_recovery_protection_policies_data_invalid_ids` | `ids` |
| `ibm_backup_recovery_protection_policies_data_invalid_include_replicated_policies` | `include_replicated_policies` |
| `ibm_backup_recovery_protection_policies_data_invalid_include_stats` | `include_stats` |
| `ibm_backup_recovery_protection_policies_data_invalid_policy_names` | `policy_names` |
| `ibm_backup_recovery_protection_policies_data_invalid_request_initiator_type` | `request_initiator_type` |
| `ibm_backup_recovery_protection_policies_data_invalid_types` | `types` |
| `ibm_backup_recovery_protection_policies_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_protection_policies" "test" {
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
