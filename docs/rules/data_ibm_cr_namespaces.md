# ibm_cr_namespaces_data

Schema validation for the ibm_cr_namespaces data source.

This rule validates `ibm_cr_namespaces` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cr_namespaces_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `namespaces` | no | — |

## Example

### Valid

```hcl
data "ibm_cr_namespaces" "test" {
}
```
