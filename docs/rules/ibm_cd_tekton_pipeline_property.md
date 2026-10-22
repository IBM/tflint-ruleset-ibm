# ibm_cd_tekton_pipeline_property

Manages ibm_cd_tekton_pipeline_property. Property name.

This rule validates `ibm_cd_tekton_pipeline_property` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Property` (`tekton-pipeline`).

## Configuration

```hcl
rule "ibm_cd_tekton_pipeline_property" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `enum` | no | pattern `^[-0-9a-zA-Z_.]{1,253}$` |
| `href` | no | — |
| `locked` | no | API: When true, this property cannot be overridden by a trigger property or at runtime. Attempting to override it will result in run requests being rejected. The default is false. |
| `name` | yes | pattern `^[-0-9a-zA-Z_.]{1,253}$` |
| `path` | no | pattern `^[-0-9a-zA-Z_.]*$` |
| `pipeline_id` | yes | pattern `^[-0-9a-z]+$` |
| `type` | yes | one of: `appconfig`, `integration`, `secure`, `single_select`, `text` |
| `value` | no | API: Property value. Any string value is valid |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_tekton_pipeline_property_invalid_enum` | `enum` |
| `ibm_cd_tekton_pipeline_property_invalid_locked` | `locked` |
| `ibm_cd_tekton_pipeline_property_invalid_name` | `name` |
| `ibm_cd_tekton_pipeline_property_invalid_path` | `path` |
| `ibm_cd_tekton_pipeline_property_invalid_pipeline_id` | `pipeline_id` |
| `ibm_cd_tekton_pipeline_property_invalid_type` | `type` |
| `ibm_cd_tekton_pipeline_property_invalid_value` | `value` |

## Example

### Valid

```hcl
resource "ibm_cd_tekton_pipeline_property" "test" {
  name = "example-name"
  pipeline_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  type = "appconfig"
}
```
