# ibm_is_backup_policy_plan

Manages ibm_is_backup_policy_plan. The user-defined name for this backup policy plan.

This rule validates `ibm_is_backup_policy_plan` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BackupPolicyPlanPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_backup_policy_plan" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `active` | no | API: Indicates whether the plan is active. |
| `attach_user_tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `backup_policy_id` | yes | — |
| `backup_policy_plan_id` | no | — |
| `copy_user_tags` | no | API: Indicates whether to copy the source's user tags to the created backups (snapshots). |
| `created_at` | no | — |
| `cron_spec` | yes | pattern `^((((\d+,)+\d+|([\d\*]+(\/|-)\d+)|\d+|\*) ?){5,7})$` |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_type` | no | — |
| `version` | no | — |
| `clone_policy` (block) | no | nested block |
| `clone_policy.max_snapshots` | no | range 1–1000 |
| `clone_policy.zones` | no | API: Identifies a zone by a unique property. |
| `deletion_trigger` (block) | no | nested block |
| `deletion_trigger.delete_after` | no | range 1–1000 |
| `deletion_trigger.delete_over_count` | no | min 1 |
| `remote_region_policy` (block) | no | nested block |
| `remote_region_policy.delete_over_count` | no | min 1 |
| `remote_region_policy.encryption_key` | no | API: The root key to use to rewrap the data encryption key for the backup (snapshot). If unspecified, the source's `encryption_key` will be used. The specified key may be in a different account, subject to IAM policies. |
| `remote_region_policy.region` | yes | API: The region this backup policy plan will create backups in. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_backup_policy_plan_invalid_active` | `active` |
| `ibm_is_backup_policy_plan_invalid_attach_user_tags` | `attach_user_tags` |
| `ibm_is_backup_policy_plan_invalid_backup_policy_id` | `backup_policy_id` |
| `ibm_is_backup_policy_plan_invalid_copy_user_tags` | `copy_user_tags` |
| `ibm_is_backup_policy_plan_invalid_cron_spec` | `cron_spec` |
| `ibm_is_backup_policy_plan_invalid_name` | `name` |
| `ibm_is_backup_policy_plan_invalid_clone_policy__max_snapshots` | `clone_policy.max_snapshots` |
| `ibm_is_backup_policy_plan_invalid_clone_policy__zones` | `clone_policy.zones` |
| `ibm_is_backup_policy_plan_invalid_deletion_trigger__delete_after` | `deletion_trigger.delete_after` |
| `ibm_is_backup_policy_plan_invalid_deletion_trigger__delete_over_count` | `deletion_trigger.delete_over_count` |
| `ibm_is_backup_policy_plan_invalid_remote_region_policy__delete_over_count` | `remote_region_policy.delete_over_count` |
| `ibm_is_backup_policy_plan_invalid_remote_region_policy__encryption_key` | `remote_region_policy.encryption_key` |
| `ibm_is_backup_policy_plan_invalid_remote_region_policy__region` | `remote_region_policy.region` |

## Example

### Valid

```hcl
resource "ibm_is_backup_policy_plan" "test" {
  backup_policy_id = "01234567-89ab-cdef-0123-456789abcdef"
  cron_spec = "0 0 * * *"
}
```
