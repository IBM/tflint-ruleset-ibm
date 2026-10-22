# ibm_schematics_output_data

Schema validation for the ibm_schematics_output data source.

This rule validates `ibm_schematics_output` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `JobStatusTemplate` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_output_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `location` | no | — |
| `output_json` | no | — |
| `output_values` | no | — |
| `resource_controller_url` | no | — |
| `template_id` | yes | API: Template Id |
| `workspace_id` | yes | API: Workspace id |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_output_data_invalid_location` | `location` |
| `ibm_schematics_output_data_invalid_output_json` | `output_json` |
| `ibm_schematics_output_data_invalid_template_id` | `template_id` |
| `ibm_schematics_output_data_invalid_workspace_id` | `workspace_id` |

## Example

### Valid

```hcl
data "ibm_schematics_output" "test" {
  template_id = "01234567-89ab-cdef-0123-456789abcdef"
  workspace_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
