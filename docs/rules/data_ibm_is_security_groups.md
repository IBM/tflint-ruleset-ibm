# ibm_is_security_groups_data

Schema validation for the ibm_is_security_groups data source.

This rule validates `ibm_is_security_groups` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `FloatingIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_security_groups_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `security_groups` | no | — |
| `vpc_crn` | no | crn format |
| `vpc_id` | no | — |
| `vpc_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_security_groups_data_invalid_resource_group` | `resource_group` |
| `ibm_is_security_groups_data_invalid_vpc_crn` | `vpc_crn` |
| `ibm_is_security_groups_data_invalid_vpc_id` | `vpc_id` |
| `ibm_is_security_groups_data_invalid_vpc_name` | `vpc_name` |

## Example

### Valid

```hcl
data "ibm_is_security_groups" "test" {
}
```
