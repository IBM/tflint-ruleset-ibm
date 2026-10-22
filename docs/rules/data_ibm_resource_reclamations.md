# ibm_resource_reclamations_data

Schema validation for the ibm_resource_reclamations data source.

This rule validates `ibm_resource_reclamations` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_resource_reclamations_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `reclamations` | no | — |

## Example

### Valid

```hcl
data "ibm_resource_reclamations" "test" {
}
```
