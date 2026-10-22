# ibm_is_public_gateway_data

Reads ibm_is_public_gateway. Public gateway Name

This rule validates `ibm_is_public_gateway` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `FloatingIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_public_gateway_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `crn` | no | — |
| `floating_ip` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `status` | no | — |
| `tags` | no | — |
| `vpc` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_public_gateway_data_invalid_name` | `name` |
| `ibm_is_public_gateway_data_invalid_resource_group` | `resource_group` |

## Example

### Valid

```hcl
data "ibm_is_public_gateway" "test" {
  name = "example-name"
}
```
