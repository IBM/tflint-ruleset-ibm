# ibm_is_private_path_service_gateways_data

Schema validation for the ibm_is_private_path_service_gateways data source.

This rule validates `ibm_is_private_path_service_gateways` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_private_path_service_gateways_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `private_path_service_gateways` | no | — |

## Example

### Valid

```hcl
data "ibm_is_private_path_service_gateways" "test" {
}
```
