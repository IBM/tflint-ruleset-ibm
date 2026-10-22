# ibm_is_dedicated_host

Manages ibm_is_dedicated_host. The unique user-defined name for this dedicated host.

This rule validates `ibm_is_dedicated_host` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DedicatedHostPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_dedicated_host" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `available_memory` | no | — |
| `available_vcpu` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `disks` | no | — |
| `host_group` | yes | — |
| `href` | no | — |
| `instance_placement_enabled` | no | API: If set to true, instances can be placed on this dedicated host |
| `instances` | no | — |
| `lifecycle_state` | no | — |
| `memory` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `numa` | no | — |
| `profile` | yes | API: Identifies a dedicated host profile by a unique property |
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
| `ibm_is_dedicated_host_invalid_access_tags` | `access_tags` |
| `ibm_is_dedicated_host_invalid_host_group` | `host_group` |
| `ibm_is_dedicated_host_invalid_instance_placement_enabled` | `instance_placement_enabled` |
| `ibm_is_dedicated_host_invalid_name` | `name` |
| `ibm_is_dedicated_host_invalid_profile` | `profile` |
| `ibm_is_dedicated_host_invalid_resource_group` | `resource_group` |

## Example

### Valid

```hcl
resource "ibm_is_dedicated_host" "test" {
  host_group = "example"
  profile = "example"
}
```
