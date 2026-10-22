# ibm_is_instance_group_manager_policies_data

Schema validation for the ibm_is_instance_group_manager_policies data source.

This rule validates `ibm_is_instance_group_manager_policies` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Instance` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_group_manager_policies_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_group` | yes | API: The instance group associated with this membership. |
| `instance_group_manager` | yes | — |
| `instance_group_manager_policies` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_group_manager_policies_data_invalid_instance_group` | `instance_group` |
| `ibm_is_instance_group_manager_policies_data_invalid_instance_group_manager` | `instance_group_manager` |

## Example

### Valid

```hcl
data "ibm_is_instance_group_manager_policies" "test" {
  instance_group = "example"
  instance_group_manager = "example"
}
```
