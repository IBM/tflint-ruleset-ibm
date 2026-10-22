# ibm_cd_tekton_pipeline_property_data

Reads ibm_cd_tekton_pipeline_property. Property name.

This rule validates `ibm_cd_tekton_pipeline_property` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Property` (`tekton-pipeline`).

## Configuration

```hcl
rule "ibm_cd_tekton_pipeline_property_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `enum` | no | — |
| `href` | no | — |
| `locked` | no | — |
| `name` | no | — |
| `path` | no | — |
| `pipeline_id` | yes | pattern `^[-0-9a-z]+$` |
| `property_name` | yes | — |
| `type` | no | — |
| `value` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_tekton_pipeline_property_data_invalid_pipeline_id` | `pipeline_id` |
| `ibm_cd_tekton_pipeline_property_data_invalid_property_name` | `property_name` |

## Example

### Valid

```hcl
data "ibm_cd_tekton_pipeline_property" "test" {
  pipeline_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  property_name = "example"
}
```
