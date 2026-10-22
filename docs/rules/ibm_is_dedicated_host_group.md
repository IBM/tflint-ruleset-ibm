# ibm_is_dedicated_host_group

Manages ibm_is_dedicated_host_group. The unique user-defined name for this dedicated host group.

This rule validates `ibm_is_dedicated_host_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DedicatedHostGroupPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_dedicated_host_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `class` | yes | — |
| `created_at` | no | — |
| `crn` | no | — |
| `dedicated_hosts` | no | — |
| `family` | yes | one of: `balanced`, `compute`, `memory`; pattern `^[a-z]+([-_][a-z]+){0,5}$` |
| `href` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_type` | no | — |
| `supported_instance_profiles` | no | — |
| `zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_dedicated_host_group_invalid_class` | `class` |
| `ibm_is_dedicated_host_group_invalid_family` | `family` |
| `ibm_is_dedicated_host_group_invalid_name` | `name` |
| `ibm_is_dedicated_host_group_invalid_resource_group` | `resource_group` |
| `ibm_is_dedicated_host_group_invalid_zone` | `zone` |

## Example

### Valid

```hcl
resource "ibm_is_dedicated_host_group" "test" {
  class = "example"
  family = "balanced"
  zone = "us-south-1"
}
```
