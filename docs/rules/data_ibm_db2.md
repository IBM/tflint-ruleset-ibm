# ibm_db2_data

Reads ibm_db2. Resource instance name for example, myobjectstorage

This rule validates `ibm_db2` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_db2_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `backup_location` | no | — |
| `crn` | no | — |
| `disk_encryption_crn` | no | crn format |
| `disk_encryption_instance_crn` | no | crn format |
| `extensions` | no | — |
| `guid` | no | — |
| `high_availability` | no | — |
| `identifier` | no | — |
| `instance_type` | no | — |
| `location` | no | — |
| `name` | no | — |
| `onetime_credentials` | no | — |
| `oracle_compatibility` | no | — |
| `parameters_json` | no | — |
| `plan` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `service` | no | — |
| `status` | no | — |
| `subscription_id` | no | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_db2_data_invalid_backup_location` | `backup_location` |
| `ibm_db2_data_invalid_disk_encryption_crn` | `disk_encryption_crn` |
| `ibm_db2_data_invalid_disk_encryption_instance_crn` | `disk_encryption_instance_crn` |
| `ibm_db2_data_invalid_high_availability` | `high_availability` |
| `ibm_db2_data_invalid_identifier` | `identifier` |
| `ibm_db2_data_invalid_instance_type` | `instance_type` |
| `ibm_db2_data_invalid_location` | `location` |
| `ibm_db2_data_invalid_name` | `name` |
| `ibm_db2_data_invalid_oracle_compatibility` | `oracle_compatibility` |
| `ibm_db2_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_db2_data_invalid_service` | `service` |
| `ibm_db2_data_invalid_subscription_id` | `subscription_id` |

## Example

### Valid

```hcl
data "ibm_db2" "test" {
}
```
