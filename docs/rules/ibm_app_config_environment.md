# ibm_app_config_environment

Manages ibm_app_config_environment. Environment name.

This rule validates `ibm_app_config_environment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Environment` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_environment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `color_code` | no | API: Color code to distinguish the environment. The Hex code for the color. For example `#FF0000` for `red`. |
| `created_time` | no | — |
| `description` | no | API: Environment description |
| `environment_id` | yes | API: Environment id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `guid` | yes | — |
| `href` | no | — |
| `name` | yes | API: Environment name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `tags` | no | pattern `^[a-zA-Z0-9_\. ,\-]+(:[a-zA-Z0-9_\. ,\-]+)*$` |
| `updated_time` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_environment_invalid_color_code` | `color_code` |
| `ibm_app_config_environment_invalid_description` | `description` |
| `ibm_app_config_environment_invalid_environment_id` | `environment_id` |
| `ibm_app_config_environment_invalid_guid` | `guid` |
| `ibm_app_config_environment_invalid_name` | `name` |
| `ibm_app_config_environment_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_app_config_environment" "test" {
  environment_id = "01234567-89ab-cdef-0123-456789abcdef"
  guid = "example"
  name = "aaaaaaa"
}
```
