# ibm_is_dedicated_host_groups_data

Reads ibm_is_dedicated_host_groups. The name of the dedicated host group

This rule validates `ibm_is_dedicated_host_groups` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DedicatedHostGroupPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_dedicated_host_groups_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `host_groups` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `total_count` | no | — |
| `zone` | no | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_dedicated_host_groups_data_invalid_name` | `name` |
| `ibm_is_dedicated_host_groups_data_invalid_resource_group` | `resource_group` |
| `ibm_is_dedicated_host_groups_data_invalid_zone` | `zone` |

## Example

### Valid

```hcl
data "ibm_is_dedicated_host_groups" "test" {
}
```
