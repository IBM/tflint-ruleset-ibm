# ibm_dl_provider_ports_data

Schema validation for the ibm_dl_provider_ports data source.

This rule validates `ibm_dl_provider_ports` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dl_provider_ports_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `ports` | no | — |

## Example

### Valid

```hcl
data "ibm_dl_provider_ports" "test" {
}
```
