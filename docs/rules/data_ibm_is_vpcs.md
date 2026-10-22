# ibm_is_vpcs_data

Schema validation for the ibm_is_vpcs data source.

This rule validates `ibm_is_vpcs` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VPCPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpcs_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `classic_access` | no | API: Indicates whether this VPC will be connected to Classic Infrastructure. If true, this VPC's resources will have private network connectivity to the account's Classic Infrastructure resources. Only one VPC, per region, may be connected in this way. This value is set at creation and subsequently immutable. Setting this property to `true` is supported only for accounts that have been [granted approva |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `vpcs` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpcs_data_invalid_classic_access` | `classic_access` |
| `ibm_is_vpcs_data_invalid_resource_group` | `resource_group` |

## Example

### Valid

```hcl
data "ibm_is_vpcs" "test" {
}
```
