# ibm_cd_tekton_pipeline_definition

Schema validation for the ibm_cd_tekton_pipeline_definition resource.

This rule validates `ibm_cd_tekton_pipeline_definition` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DefinitionSourceProperties` (`tekton-pipeline`).

## Configuration

```hcl
rule "ibm_cd_tekton_pipeline_definition" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `definition_id` | no | — |
| `href` | no | — |
| `pipeline_id` | yes | pattern `^[-0-9a-z]+$` |
| `source` (block) | yes | nested block |
| `source.type` | yes | — |
| `source.properties` (block) | yes | nested block |
| `source.properties.branch` | no | pattern `^[-0-9a-zA-Z_.]{1,253}$` |
| `source.properties.path` | yes | pattern `^[-0-9a-zA-Z_.]{1,253}$` |
| `source.properties.tag` | no | pattern `^[-0-9a-zA-Z_]{1,253}$` |
| `source.properties.url` | yes | pattern `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$` |
| `source.properties.tool` (block) | no | nested block |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_tekton_pipeline_definition_invalid_pipeline_id` | `pipeline_id` |
| `ibm_cd_tekton_pipeline_definition_invalid_source__type` | `source.type` |
| `ibm_cd_tekton_pipeline_definition_invalid_source__properties__branch` | `source.properties.branch` |
| `ibm_cd_tekton_pipeline_definition_invalid_source__properties__path` | `source.properties.path` |
| `ibm_cd_tekton_pipeline_definition_invalid_source__properties__tag` | `source.properties.tag` |
| `ibm_cd_tekton_pipeline_definition_invalid_source__properties__url` | `source.properties.url` |

## Example

### Valid

```hcl
resource "ibm_cd_tekton_pipeline_definition" "test" {
  pipeline_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  source {
    type = "example"
    properties {
      path = "example-name"
      url = "http://example.com"
    }
  }
}
```
