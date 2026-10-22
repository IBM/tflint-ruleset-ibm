# ibm_is_zone_data

Schema validation for the ibm_is_zone data source.

This rule validates `ibm_is_zone` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Zone` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_zone_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `data_center` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `region` | yes | API: The region this zone resides in |
| `status` | no | — |
| `universal_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_zone_data_invalid_name` | `name` |
| `ibm_is_zone_data_invalid_region` | `region` |

## Example

### Valid

```hcl
data "ibm_is_zone" "test" {
  name = "example-name"
  region = "example"
}
```
