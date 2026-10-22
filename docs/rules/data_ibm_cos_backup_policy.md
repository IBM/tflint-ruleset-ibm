# ibm_cos_backup_policy_data

Schema validation for the ibm_cos_backup_policy data source.

This rule validates `ibm_cos_backup_policy` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/backup_vaults/{backup_vault_name}:patch` (`cos-cos-configuration`).

## Configuration

```hcl
rule "ibm_cos_backup_policy_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `backup_type` | no | — |
| `bucket_name` | yes | — |
| `initial_delete_after_days` | no | — |
| `policy_id` | yes | pattern `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `policy_name` | no | — |
| `target_backup_vault_crn` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cos_backup_policy_data_invalid_bucket_name` | `bucket_name` |
| `ibm_cos_backup_policy_data_invalid_policy_id` | `policy_id` |

## Example

### Valid

```hcl
data "ibm_cos_backup_policy" "test" {
  bucket_name = "example"
  policy_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
