# ibm_scc_profiles_data

Schema validation for the ibm_scc_profiles data source.

This rule validates `ibm_scc_profiles` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `leaf-unique` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_profiles_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_id` | yes | — |
| `profile_type` | no | one of: `custom`, `predefined` |
| `profiles` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_profiles_data_invalid_instance_id` | `instance_id` |
| `ibm_scc_profiles_data_invalid_profile_type` | `profile_type` |

## Example

### Valid

```hcl
data "ibm_scc_profiles" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
