# ibm_is_backup_policies_data

Reads ibm_is_backup_policies. Filters the collection to resources with the exact specified name

This rule validates `ibm_is_backup_policies` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `FloatingIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_backup_policies_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `backup_policies` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `tag` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_backup_policies_data_invalid_name` | `name` |
| `ibm_is_backup_policies_data_invalid_resource_group` | `resource_group` |
| `ibm_is_backup_policies_data_invalid_tag` | `tag` |

## Example

### Valid

```hcl
data "ibm_is_backup_policies" "test" {
}
```
