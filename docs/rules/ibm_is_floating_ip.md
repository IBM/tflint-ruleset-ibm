# ibm_is_floating_ip

Manages ibm_is_floating_ip. Name of the floating IP

This rule validates `ibm_is_floating_ip` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DedicatedHostGroupPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_floating_ip" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `address` | no | — |
| `crn` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `status` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `target` | no | — |
| `target_list` | no | — |
| `zone` | no | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_floating_ip_invalid_access_tags` | `access_tags` |
| `ibm_is_floating_ip_invalid_name` | `name` |
| `ibm_is_floating_ip_invalid_resource_group` | `resource_group` |
| `ibm_is_floating_ip_invalid_tags` | `tags` |
| `ibm_is_floating_ip_invalid_target` | `target` |
| `ibm_is_floating_ip_invalid_zone` | `zone` |

## Example

### Valid

```hcl
resource "ibm_is_floating_ip" "test" {
  name = "example-name"
}
```
