# ibm_cos_backup_policy

Schema validation for the ibm_cos_backup_policy resource.

This rule validates `ibm_cos_backup_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BackupPolicyPrototype` (`cos-cos-configuration`).

## Configuration

```hcl
rule "ibm_cos_backup_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `backup_type` | yes | one of: `continuous` |
| `bucket_crn` | yes | crn format |
| `initial_delete_after_days` | yes | — |
| `policy_id` | no | — |
| `policy_name` | yes | pattern `^[\w\d\-\_\.]*$` |
| `target_backup_vault_crn` | yes | crn format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cos_backup_policy_invalid_backup_type` | `backup_type` |
| `ibm_cos_backup_policy_invalid_bucket_crn` | `bucket_crn` |
| `ibm_cos_backup_policy_invalid_initial_delete_after_days` | `initial_delete_after_days` |
| `ibm_cos_backup_policy_invalid_policy_name` | `policy_name` |
| `ibm_cos_backup_policy_invalid_target_backup_vault_crn` | `target_backup_vault_crn` |

## Example

### Valid

```hcl
resource "ibm_cos_backup_policy" "test" {
  backup_type = "continuous"
  bucket_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  initial_delete_after_days = 1
  policy_name = "example-name"
  target_backup_vault_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
