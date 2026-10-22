# ibm_is_virtual_endpoint_gateways_data

Reads ibm_is_virtual_endpoint_gateways. The name of the endpoint gateway group

This rule validates `ibm_is_virtual_endpoint_gateways` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `FloatingIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_virtual_endpoint_gateways_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `virtual_endpoint_gateways` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_virtual_endpoint_gateways_data_invalid_name` | `name` |
| `ibm_is_virtual_endpoint_gateways_data_invalid_resource_group` | `resource_group` |

## Example

### Valid

```hcl
data "ibm_is_virtual_endpoint_gateways" "test" {
}
```
