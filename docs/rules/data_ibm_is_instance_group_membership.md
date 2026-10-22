# ibm_is_instance_group_membership_data

Reads ibm_is_instance_group_membership. The user-defined name for this instance group membership.

This rule validates `ibm_is_instance_group_membership` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_group_membership_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `delete_instance_on_membership_delete` | no | — |
| `instance` | no | — |
| `instance_group` | yes | pattern `^[-0-9a-z_]+$` |
| `instance_group_membership` | no | — |
| `instance_template` | no | — |
| `load_balancer_pool_member` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_group_membership_data_invalid_instance_group` | `instance_group` |
| `ibm_is_instance_group_membership_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_instance_group_membership" "test" {
  instance_group = "example-name"
  name = "example-name"
}
```
