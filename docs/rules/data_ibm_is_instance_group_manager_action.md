# ibm_is_instance_group_manager_action_data

Reads ibm_is_instance_group_manager_action. instance group manager action name

This rule validates `ibm_is_instance_group_manager_action` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_group_manager_action_data" {
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
| `cron_spec` | no | — |
| `instance_group` | yes | — |
| `instance_group_manager` | yes | — |
| `last_applied_at` | no | — |
| `max_membership_count` | no | — |
| `membership_count` | no | — |
| `min_membership_count` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `next_run_at` | no | — |
| `resource_type` | no | — |
| `run_at` | no | — |
| `status` | no | — |
| `target_manager` | no | — |
| `target_manager_name` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_group_manager_action_data_invalid_instance_group` | `instance_group` |
| `ibm_is_instance_group_manager_action_data_invalid_instance_group_manager` | `instance_group_manager` |
| `ibm_is_instance_group_manager_action_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_instance_group_manager_action" "test" {
  instance_group = "example"
  instance_group_manager = "example"
  name = "example-name"
}
```
