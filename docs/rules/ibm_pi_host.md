# ibm_pi_host

Schema validation for the ibm_pi_host resource.

This rule validates `ibm_pi_host` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_host" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `capacity` | no | — |
| `crn` | no | — |
| `display_name` | no | — |
| `host_group` | no | — |
| `host_id` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_host_group_id` | yes | — |
| `state` | no | — |
| `status` | no | — |
| `sys_type` | no | — |
| `user_tags` | no | — |
| `pi_host` (block) | yes | nested block |
| `pi_host.display_name` | yes | — |
| `pi_host.sys_type` | yes | — |
| `pi_host.user_tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_host_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_host_invalid_pi_host_group_id` | `pi_host_group_id` |
| `ibm_pi_host_invalid_pi_host__display_name` | `pi_host.display_name` |
| `ibm_pi_host_invalid_pi_host__sys_type` | `pi_host.sys_type` |
| `ibm_pi_host_invalid_pi_host__user_tags` | `pi_host.user_tags` |

## Example

### Valid

```hcl
resource "ibm_pi_host" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_host_group_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_host {
    display_name = "example"
    sys_type = "example"
  }
}
```
