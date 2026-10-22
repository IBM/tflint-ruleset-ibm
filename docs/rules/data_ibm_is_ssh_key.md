# ibm_is_ssh_key_data

Reads ibm_is_ssh_key. The name of the ssh key

This rule validates `ibm_is_ssh_key` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `FloatingIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_ssh_key_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `fingerprint` | no | — |
| `href` | no | — |
| `id` | no | pattern `^[0-9a-f]{32}$` |
| `length` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `public_key` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `tags` | no | — |
| `type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_ssh_key_data_invalid_id` | `id` |
| `ibm_is_ssh_key_data_invalid_name` | `name` |
| `ibm_is_ssh_key_data_invalid_resource_group` | `resource_group` |

## Example

### Valid

```hcl
data "ibm_is_ssh_key" "test" {
}
```
