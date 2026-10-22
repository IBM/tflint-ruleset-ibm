# ibm_is_instance_group_membership

Manages ibm_is_instance_group_membership. The user-defined name for this instance group membership.

This rule validates `ibm_is_instance_group_membership` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_group_membership" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action_delete` | no | — |
| `delete_instance_on_membership_delete` | no | — |
| `instance` | no | — |
| `instance_group` | yes | pattern `^[-0-9a-z_]+$` |
| `instance_group_membership` | yes | pattern `^[-0-9a-z_]+$` |
| `instance_template` | no | — |
| `load_balancer_pool_member` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_group_membership_invalid_action_delete` | `action_delete` |
| `ibm_is_instance_group_membership_invalid_instance_group` | `instance_group` |
| `ibm_is_instance_group_membership_invalid_instance_group_membership` | `instance_group_membership` |
| `ibm_is_instance_group_membership_invalid_name` | `name` |

## Example

### Valid

```hcl
resource "ibm_is_instance_group_membership" "test" {
  instance_group = "example-name"
  instance_group_membership = "example-name"
}
```
