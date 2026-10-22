# ibm_is_backup_policy_data

Reads ibm_is_backup_policy. The unique user-defined name for this backup policy.

This rule validates `ibm_is_backup_policy` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_backup_policy_data" {
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
| `identifier` | no | — |
| `included_content` | no | — |
| `last_job_completed_at` | no | — |
| `lifecycle_state` | no | — |
| `match_resource_type` | no | — |
| `match_resource_types` | no | — |
| `match_user_tags` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `plans` | no | — |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `scope` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_backup_policy_data_invalid_identifier` | `identifier` |
| `ibm_is_backup_policy_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_backup_policy" "test" {
}
```
