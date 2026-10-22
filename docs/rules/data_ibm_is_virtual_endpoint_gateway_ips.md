# ibm_is_virtual_endpoint_gateway_ips_data

Schema validation for the ibm_is_virtual_endpoint_gateway_ips data source.

This rule validates `ibm_is_virtual_endpoint_gateway_ips` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_virtual_endpoint_gateway_ips_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `gateway` | yes | — |
| `ips` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_virtual_endpoint_gateway_ips_data_invalid_gateway` | `gateway` |

## Example

### Valid

```hcl
data "ibm_is_virtual_endpoint_gateway_ips" "test" {
  gateway = "example"
}
```
