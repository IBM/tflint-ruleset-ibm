# ibm_is_instance_group_manager_policy_data

Reads ibm_is_instance_group_manager_policy. The name of the instance group manager policy

This rule validates `ibm_is_instance_group_manager_policy` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_group_manager_policy_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_group` | yes | — |
| `instance_group_manager` | yes | — |
| `metric_type` | no | — |
| `metric_value` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `policy_id` | no | — |
| `policy_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_group_manager_policy_data_invalid_instance_group` | `instance_group` |
| `ibm_is_instance_group_manager_policy_data_invalid_instance_group_manager` | `instance_group_manager` |
| `ibm_is_instance_group_manager_policy_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_instance_group_manager_policy" "test" {
  instance_group = "example"
  instance_group_manager = "example"
  name = "example-name"
}
```
