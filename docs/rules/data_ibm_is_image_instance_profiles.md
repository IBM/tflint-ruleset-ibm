# ibm_is_image_instance_profiles_data

Schema validation for the ibm_is_image_instance_profiles data source.

This rule validates `ibm_is_image_instance_profiles` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_image_instance_profiles_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `identifier` | yes | — |
| `instance_profiles` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_image_instance_profiles_data_invalid_identifier` | `identifier` |

## Example

### Valid

```hcl
data "ibm_is_image_instance_profiles" "test" {
  identifier = "example"
}
```
