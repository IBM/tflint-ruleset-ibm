# ibm_scc_control_libraries_data

Schema validation for the ibm_scc_control_libraries data source.

This rule validates `ibm_scc_control_libraries` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/instances/{instance_id}/v3/control_libraries:post` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_control_libraries_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `control_libraries` | no | — |
| `control_library_type` | no | one of: `custom` |
| `instance_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_control_libraries_data_invalid_control_library_type` | `control_library_type` |
| `ibm_scc_control_libraries_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_scc_control_libraries" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
