# ibm_is_instance_group_manager_policy

Manages ibm_is_instance_group_manager_policy. instance group manager policy name

This rule validates `ibm_is_instance_group_manager_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `InstanceGroupManagerPolicyPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_group_manager_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_group` | yes | — |
| `instance_group_manager` | yes | — |
| `metric_type` | yes | one of: `cpu`, `memory`, `network_in`, `network_out`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `metric_value` | yes | range 0–9007199254740991 |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `policy_id` | no | — |
| `policy_type` | yes | one of: `target`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_group_manager_policy_invalid_instance_group` | `instance_group` |
| `ibm_is_instance_group_manager_policy_invalid_instance_group_manager` | `instance_group_manager` |
| `ibm_is_instance_group_manager_policy_invalid_metric_type` | `metric_type` |
| `ibm_is_instance_group_manager_policy_invalid_metric_value` | `metric_value` |
| `ibm_is_instance_group_manager_policy_invalid_name` | `name` |
| `ibm_is_instance_group_manager_policy_invalid_policy_type` | `policy_type` |

## Example

### Valid

```hcl
resource "ibm_is_instance_group_manager_policy" "test" {
  instance_group = "example"
  instance_group_manager = "example"
  metric_type = "cpu"
  metric_value = 0
  policy_type = "target"
}
```
