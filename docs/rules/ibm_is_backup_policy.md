# ibm_is_backup_policy

Manages ibm_is_backup_policy. The user-defined name for this backup policy.

This rule validates `ibm_is_backup_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BackupPolicyPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_backup_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `crn` | no | — |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `href` | no | — |
| `included_content` | no | one of: `boot_volume`, `data_volumes`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `last_job_completed_at` | no | — |
| `lifecycle_state` | no | — |
| `match_resource_type` | no | one of: `instance`, `volume`, `share`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `match_resource_types` | no | — |
| `match_user_tags` | yes | pattern `^[A-Za-z0-9:_ .-]+$` |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_type` | no | — |
| `version` | no | — |
| `scope` (block) | no | nested block |
| `scope.crn` | no | crn format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_backup_policy_invalid_included_content` | `included_content` |
| `ibm_is_backup_policy_invalid_match_resource_type` | `match_resource_type` |
| `ibm_is_backup_policy_invalid_match_resource_types` | `match_resource_types` |
| `ibm_is_backup_policy_invalid_match_user_tags` | `match_user_tags` |
| `ibm_is_backup_policy_invalid_name` | `name` |
| `ibm_is_backup_policy_invalid_resource_group` | `resource_group` |
| `ibm_is_backup_policy_invalid_scope__crn` | `scope.crn` |

## Example

### Valid

```hcl
resource "ibm_is_backup_policy" "test" {
  match_user_tags = ["example-name"]
  name = "example-name"
}
```
