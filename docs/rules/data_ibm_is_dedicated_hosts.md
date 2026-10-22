# ibm_is_dedicated_hosts_data

Reads ibm_is_dedicated_hosts. The name of the dedicated host

This rule validates `ibm_is_dedicated_hosts` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DedicatedHostGroupPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_dedicated_hosts_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `dedicated_hosts` | no | — |
| `host_group` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `total_count` | no | — |
| `zone` | no | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_dedicated_hosts_data_invalid_host_group` | `host_group` |
| `ibm_is_dedicated_hosts_data_invalid_name` | `name` |
| `ibm_is_dedicated_hosts_data_invalid_resource_group` | `resource_group` |
| `ibm_is_dedicated_hosts_data_invalid_zone` | `zone` |

## Example

### Valid

```hcl
data "ibm_is_dedicated_hosts" "test" {
}
```
