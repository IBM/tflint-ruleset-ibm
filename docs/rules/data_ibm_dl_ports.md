# ibm_dl_ports_data

Schema validation for the ibm_dl_ports data source.

This rule validates `ibm_dl_ports` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `GatewayTypeDedicatedTemplate` (`direct_link`).

## Configuration

```hcl
rule "ibm_dl_ports_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `location_name` | no | API: Gateway location |
| `ports` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_ports_data_invalid_location_name` | `location_name` |

## Example

### Valid

```hcl
data "ibm_dl_ports" "test" {
}
```
