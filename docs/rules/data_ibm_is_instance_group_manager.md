# ibm_is_instance_group_manager_data

Reads ibm_is_instance_group_manager. Name of the instance group manager.

This rule validates `ibm_is_instance_group_manager` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_group_manager_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `actions` | no | — |
| `aggregation_window` | no | — |
| `cooldown` | no | — |
| `instance_group` | yes | — |
| `manager_id` | no | — |
| `manager_type` | no | — |
| `max_membership_count` | no | — |
| `min_membership_count` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `policies` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_group_manager_data_invalid_instance_group` | `instance_group` |
| `ibm_is_instance_group_manager_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_instance_group_manager" "test" {
  instance_group = "example"
  name = "example-name"
}
```
