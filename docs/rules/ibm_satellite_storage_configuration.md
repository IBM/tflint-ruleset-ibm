# ibm_satellite_storage_configuration

Schema validation for the ibm_satellite_storage_configuration resource.

This rule validates `ibm_satellite_storage_configuration` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_storage_configuration" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `config_name` | yes | — |
| `config_version` | no | — |
| `delete_assignments` | no | — |
| `location` | yes | — |
| `storage_class_parameters` | no | — |
| `storage_template_name` | yes | — |
| `storage_template_version` | yes | — |
| `update_assignments` | no | — |
| `user_config_parameters` | yes | — |
| `user_secret_parameters` | yes | — |
| `uuid` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_storage_configuration_invalid_config_name` | `config_name` |
| `ibm_satellite_storage_configuration_invalid_delete_assignments` | `delete_assignments` |
| `ibm_satellite_storage_configuration_invalid_location` | `location` |
| `ibm_satellite_storage_configuration_invalid_storage_class_parameters` | `storage_class_parameters` |
| `ibm_satellite_storage_configuration_invalid_storage_template_name` | `storage_template_name` |
| `ibm_satellite_storage_configuration_invalid_storage_template_version` | `storage_template_version` |
| `ibm_satellite_storage_configuration_invalid_update_assignments` | `update_assignments` |
| `ibm_satellite_storage_configuration_invalid_user_config_parameters` | `user_config_parameters` |
| `ibm_satellite_storage_configuration_invalid_user_secret_parameters` | `user_secret_parameters` |

## Example

### Valid

```hcl
variable "user_secret_parameters" {
  sensitive = true
}

resource "ibm_satellite_storage_configuration" "test" {
  config_name = "example"
  location = "example"
  storage_template_name = "example"
  storage_template_version = "example"
  user_config_parameters = "example"
  user_secret_parameters = var.user_secret_parameters
}
```
