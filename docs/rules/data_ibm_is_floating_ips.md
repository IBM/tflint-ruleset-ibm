# ibm_is_floating_ips_data

Reads ibm_is_floating_ips. The unique user-defined name for this floating IP.

This rule validates `ibm_is_floating_ips` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `FloatingIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_floating_ips_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `floating_ips` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_floating_ips_data_invalid_name` | `name` |
| `ibm_is_floating_ips_data_invalid_resource_group` | `resource_group` |

## Example

### Valid

```hcl
data "ibm_is_floating_ips" "test" {
}
```
