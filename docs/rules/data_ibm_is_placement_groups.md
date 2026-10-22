# ibm_is_placement_groups_data

Schema validation for the ibm_is_placement_groups data source.

This rule validates `ibm_is_placement_groups` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_placement_groups_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `placement_groups` | no | — |
| `total_count` | no | — |

## Example

### Valid

```hcl
data "ibm_is_placement_groups" "test" {
}
```
