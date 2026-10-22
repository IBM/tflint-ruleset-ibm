# ibm_is_instance_group_manager

Manages ibm_is_instance_group_manager. instance group manager name

This rule validates `ibm_is_instance_group_manager` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `InstanceGroupManagerAutoScalePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_group_manager" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `actions` | no | — |
| `aggregation_window` | no | range 90–600 |
| `cooldown` | no | range 120–3600 |
| `enable_manager` | no | — |
| `instance_group` | yes | — |
| `manager_id` | no | — |
| `manager_type` | no | one of: `autoscale`, `scheduled`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `max_membership_count` | no | range 1–1000 |
| `min_membership_count` | no | range 1–1000 |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `policies` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_group_manager_invalid_aggregation_window` | `aggregation_window` |
| `ibm_is_instance_group_manager_invalid_cooldown` | `cooldown` |
| `ibm_is_instance_group_manager_invalid_enable_manager` | `enable_manager` |
| `ibm_is_instance_group_manager_invalid_instance_group` | `instance_group` |
| `ibm_is_instance_group_manager_invalid_manager_type` | `manager_type` |
| `ibm_is_instance_group_manager_invalid_max_membership_count` | `max_membership_count` |
| `ibm_is_instance_group_manager_invalid_min_membership_count` | `min_membership_count` |
| `ibm_is_instance_group_manager_invalid_name` | `name` |

## Example

### Valid

```hcl
resource "ibm_is_instance_group_manager" "test" {
  instance_group = "example"
}
```
