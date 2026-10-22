# ibm_is_instance_group

Manages ibm_is_instance_group. The user-defined name for this instance group

This rule validates `ibm_is_instance_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `InstanceGroupPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `application_port` | no | range 1–65535 |
| `crn` | no | — |
| `instance_count` | no | range 0–1000 |
| `instance_template` | yes | API: The instance template to use when creating new instances. The specified template must not have `default_trusted_profile.auto_link` set to `true`. |
| `instances` | no | — |
| `load_balancer` | no | API: The load balancer associated with the specified load balancer pool. Required if `load_balancer_pool` is specified. The load balancer must have `instance_groups_supported` set to `true`. |
| `load_balancer_pool` | no | API: If specified, this instance group will manage the load balancer pool. A pool member will be created for each instance created by this group. The specified load balancer pool must not be used by another instance group in the VPC. If specified, `load_balancer` and `application_port` must also be specified. |
| `managers` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `status` | no | — |
| `subnets` | yes | API: Identifies a subnet by a unique property. |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `vpc` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_group_invalid_access_tags` | `access_tags` |
| `ibm_is_instance_group_invalid_application_port` | `application_port` |
| `ibm_is_instance_group_invalid_instance_count` | `instance_count` |
| `ibm_is_instance_group_invalid_instance_template` | `instance_template` |
| `ibm_is_instance_group_invalid_load_balancer` | `load_balancer` |
| `ibm_is_instance_group_invalid_load_balancer_pool` | `load_balancer_pool` |
| `ibm_is_instance_group_invalid_name` | `name` |
| `ibm_is_instance_group_invalid_resource_group` | `resource_group` |
| `ibm_is_instance_group_invalid_subnets` | `subnets` |
| `ibm_is_instance_group_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_is_instance_group" "test" {
  instance_template = "example"
  name = "example-name"
  subnets = "example"
}
```
