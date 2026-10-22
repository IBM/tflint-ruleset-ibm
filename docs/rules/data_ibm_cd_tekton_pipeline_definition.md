# ibm_cd_tekton_pipeline_definition_data

Schema validation for the ibm_cd_tekton_pipeline_definition data source.

This rule validates `ibm_cd_tekton_pipeline_definition` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Definition` (`tekton-pipeline`).

## Configuration

```hcl
rule "ibm_cd_tekton_pipeline_definition_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `definition_id` | yes | pattern `^[-0-9a-z]+$` |
| `href` | no | — |
| `pipeline_id` | yes | pattern `^[-0-9a-z]+$` |
| `source` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_tekton_pipeline_definition_data_invalid_definition_id` | `definition_id` |
| `ibm_cd_tekton_pipeline_definition_data_invalid_pipeline_id` | `pipeline_id` |

## Example

### Valid

```hcl
data "ibm_cd_tekton_pipeline_definition" "test" {
  definition_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  pipeline_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
}
```
