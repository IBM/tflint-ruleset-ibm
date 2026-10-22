# ibm_is_backup_policy_plan_data

Reads ibm_is_backup_policy_plan. The unique user-defined name for this backup policy plan.

This rule validates `ibm_is_backup_policy_plan` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_backup_policy_plan_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `active` | no | — |
| `attach_user_tags` | no | — |
| `backup_policy_id` | yes | — |
| `clone_policy` | no | — |
| `copy_user_tags` | no | — |
| `created_at` | no | — |
| `cron_spec` | no | — |
| `deletion_trigger` | no | — |
| `href` | no | — |
| `identifier` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `remote_region_policy` | no | — |
| `resource_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_backup_policy_plan_data_invalid_backup_policy_id` | `backup_policy_id` |
| `ibm_is_backup_policy_plan_data_invalid_identifier` | `identifier` |
| `ibm_is_backup_policy_plan_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_backup_policy_plan" "test" {
  backup_policy_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
