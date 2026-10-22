# ibm_backup_recovery_source_registration_data

Reads ibm_backup_recovery_source_registration. The user specified name for this source.

This rule validates `ibm_backup_recovery_source_registration` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_source_registration_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `advanced_configs` | no | — |
| `authentication_status` | no | — |
| `connection_id` | no | — |
| `connections` | no | — |
| `connector_group_id` | no | — |
| `data_source_connection_id` | no | — |
| `environment` | no | — |
| `external_metadata` | no | — |
| `kubernetes_params` | no | — |
| `last_refreshed_time_msecs` | no | — |
| `name` | no | — |
| `physical_params` | no | — |
| `registration_time_msecs` | no | — |
| `request_initiator_type` | no | — |
| `source_id` | no | — |
| `source_info` | no | — |
| `source_registration_id` | yes | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_source_registration_data_invalid_request_initiator_type` | `request_initiator_type` |
| `ibm_backup_recovery_source_registration_data_invalid_source_registration_id` | `source_registration_id` |
| `ibm_backup_recovery_source_registration_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_source_registration" "test" {
  source_registration_id = 1
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
