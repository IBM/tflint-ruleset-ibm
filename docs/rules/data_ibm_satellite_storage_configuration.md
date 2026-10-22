# ibm_satellite_storage_configuration_data

Schema validation for the ibm_satellite_storage_configuration data source.

This rule validates `ibm_satellite_storage_configuration` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_storage_configuration_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `config_name` | yes | — |
| `config_version` | no | — |
| `location` | yes | — |
| `storage_class_parameters` | no | — |
| `storage_template_name` | no | — |
| `storage_template_version` | no | — |
| `user_config_parameters` | no | — |
| `uuid` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_storage_configuration_data_invalid_config_name` | `config_name` |
| `ibm_satellite_storage_configuration_data_invalid_location` | `location` |

## Example

### Valid

```hcl
data "ibm_satellite_storage_configuration" "test" {
  config_name = "example"
  location = "example"
}
```
