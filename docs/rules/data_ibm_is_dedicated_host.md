# ibm_is_dedicated_host_data

Reads ibm_is_dedicated_host. The unique name of this dedicated host

This rule validates `ibm_is_dedicated_host` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `FloatingIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_dedicated_host_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `available_memory` | no | — |
| `available_vcpu` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `disks` | no | — |
| `host_group` | yes | — |
| `href` | no | — |
| `instance_placement_enabled` | no | — |
| `instances` | no | — |
| `lifecycle_state` | no | — |
| `memory` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `numa` | no | — |
| `profile` | no | — |
| `provisionable` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_type` | no | — |
| `socket_count` | no | — |
| `state` | no | — |
| `supported_instance_profiles` | no | — |
| `vcpu` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_dedicated_host_data_invalid_host_group` | `host_group` |
| `ibm_is_dedicated_host_data_invalid_name` | `name` |
| `ibm_is_dedicated_host_data_invalid_resource_group` | `resource_group` |

## Example

### Valid

```hcl
data "ibm_is_dedicated_host" "test" {
  host_group = "example"
  name = "example-name"
}
```
