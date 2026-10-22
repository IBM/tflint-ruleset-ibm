# ibm_dl_routers_data

Schema validation for the ibm_dl_routers data source.

This rule validates `ibm_dl_routers` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `GatewayTypeDedicatedTemplate` (`direct_link`).

## Configuration

```hcl
rule "ibm_dl_routers_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cross_connect_routers` | no | — |
| `location_name` | yes | API: Gateway location |
| `offering_type` | yes | one of: `dedicated` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_routers_data_invalid_location_name` | `location_name` |
| `ibm_dl_routers_data_invalid_offering_type` | `offering_type` |

## Example

### Valid

```hcl
data "ibm_dl_routers" "test" {
  location_name = "example"
  offering_type = "dedicated"
}
```
