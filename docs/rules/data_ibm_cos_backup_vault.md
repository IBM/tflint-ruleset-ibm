# ibm_cos_backup_vault_data

Schema validation for the ibm_cos_backup_vault data source.

This rule validates `ibm_cos_backup_vault` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:post:/backup_vaults/{backup_vault_name}/restores` (`cos-cos-configuration`).

## Configuration

```hcl
rule "ibm_cos_backup_vault_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `activity_tracking_management_events` | no | — |
| `backup_vault_name` | yes | API: The name given to a BackupVault. BackupVault names must be between 3 and 63 characters long must be made of lowercase letters, numbers, dots (periods), and dashes (hyphens). BackupVault names must begin and end with a lowercase letter or number. BackupVault names canâ€™t contain consecutive dots or dashes. BackupVault names that resemble IP addresses are not allowed. Bucket and BackupVault names e |
| `kms_key_crn` | no | — |
| `metrics_monitoring_usage_metrics` | no | — |
| `region` | yes | pattern `^[\w\-]*$` |
| `service_instance_id` | yes | pattern `^[0-9a-f]{32}$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cos_backup_vault_data_invalid_backup_vault_name` | `backup_vault_name` |
| `ibm_cos_backup_vault_data_invalid_region` | `region` |
| `ibm_cos_backup_vault_data_invalid_service_instance_id` | `service_instance_id` |

## Example

### Valid

```hcl
data "ibm_cos_backup_vault" "test" {
  backup_vault_name = "aaaaaaa"
  region = "example-name"
  service_instance_id = "0123456789abcdef0123456789abcdef"
}
```
