# ibm_vmaas_transit_gateway_connection_data

Schema validation for the ibm_vmaas_transit_gateway_connection data source.

This rule validates `ibm_vmaas_transit_gateway_connection` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_vmaas_transit_gateway_connection_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `connections` | no | — |
| `edge_id` | yes | pattern `^[A-Za-z0-9_-]{1,128}$` |
| `region` | no | — |
| `status` | no | — |
| `vdc_id` | yes | pattern `^[A-Za-z0-9_-]{1,128}$` |
| `vmaas_transit_gateway_connection_id` | yes | pattern `^[A-Za-z0-9_-]{1,128}$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_vmaas_transit_gateway_connection_data_invalid_edge_id` | `edge_id` |
| `ibm_vmaas_transit_gateway_connection_data_invalid_vdc_id` | `vdc_id` |
| `ibm_vmaas_transit_gateway_connection_data_invalid_vmaas_transit_gateway_connection_id` | `vmaas_transit_gateway_connection_id` |

## Example

### Valid

```hcl
data "ibm_vmaas_transit_gateway_connection" "test" {
  edge_id = "example-name"
  vdc_id = "example-name"
  vmaas_transit_gateway_connection_id = "example-name"
}
```
