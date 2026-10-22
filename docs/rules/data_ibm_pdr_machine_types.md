# ibm_pdr_machine_types_data

Schema validation for the ibm_pdr_machine_types data source.

This rule validates `ibm_pdr_machine_types` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `leaf-unique` (`dr-automation-powervs`).

## Configuration

```hcl
rule "ibm_pdr_machine_types_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | — |
| `instance_id` | yes | — |
| `primary_workspace_name` | yes | pattern `^[A-Za-z0-9._:-]+$` |
| `standby_workspace_name` | no | pattern `^[A-Za-z0-9._:-]+$` |
| `workspaces` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pdr_machine_types_data_invalid_accept_language` | `accept_language` |
| `ibm_pdr_machine_types_data_invalid_instance_id` | `instance_id` |
| `ibm_pdr_machine_types_data_invalid_primary_workspace_name` | `primary_workspace_name` |
| `ibm_pdr_machine_types_data_invalid_standby_workspace_name` | `standby_workspace_name` |

## Example

### Valid

```hcl
data "ibm_pdr_machine_types" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  primary_workspace_name = "example-name"
}
```
