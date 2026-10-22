# ibm_is_instance_group_manager_action

Manages ibm_is_instance_group_manager_action. instance group manager action name

This rule validates `ibm_is_instance_group_manager_action` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `InstanceGroupManagerPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_group_manager_action" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action_id` | no | — |
| `action_type` | no | — |
| `auto_delete` | no | — |
| `auto_delete_timeout` | no | — |
| `created_at` | no | — |
| `cron_spec` | no | pattern `^((((\d+,)+\d+|([\d\*]+(\/|-)\d+)|\d+|\*) ?){5,7})$` |
| `instance_group` | yes | — |
| `instance_group_manager` | yes | — |
| `last_applied_at` | no | — |
| `max_membership_count` | no | range 1–1000 |
| `membership_count` | no | range 0–100 |
| `min_membership_count` | no | range 1–1000 |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `next_run_at` | no | — |
| `resource_type` | no | — |
| `run_at` | no | pattern `^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[\+-]\d{2}:\d{2})?)$` |
| `status` | no | — |
| `target_manager` | no | — |
| `target_manager_name` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_group_manager_action_invalid_cron_spec` | `cron_spec` |
| `ibm_is_instance_group_manager_action_invalid_instance_group` | `instance_group` |
| `ibm_is_instance_group_manager_action_invalid_instance_group_manager` | `instance_group_manager` |
| `ibm_is_instance_group_manager_action_invalid_max_membership_count` | `max_membership_count` |
| `ibm_is_instance_group_manager_action_invalid_membership_count` | `membership_count` |
| `ibm_is_instance_group_manager_action_invalid_min_membership_count` | `min_membership_count` |
| `ibm_is_instance_group_manager_action_invalid_name` | `name` |
| `ibm_is_instance_group_manager_action_invalid_run_at` | `run_at` |
| `ibm_is_instance_group_manager_action_invalid_target_manager` | `target_manager` |

## Example

### Valid

```hcl
resource "ibm_is_instance_group_manager_action" "test" {
  instance_group = "example"
  instance_group_manager = "example"
}
```
