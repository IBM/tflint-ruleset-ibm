# ibm_pi_host_group

Manages ibm_pi_host_group. Name of the host group.

This rule validates `ibm_pi_host_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Host` (`power-cloud`).

## Configuration

```hcl
rule "ibm_pi_host_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `creation_date` | no | — |
| `host_group_id` | no | — |
| `hosts` | no | — |
| `name` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_name` | yes | — |
| `pi_remove` | no | — |
| `primary` | no | — |
| `secondaries` | no | — |
| `pi_hosts` (block) | yes | nested block |
| `pi_hosts.display_name` | yes | — |
| `pi_hosts.sys_type` | yes | — |
| `pi_hosts.user_tags` | no | — |
| `pi_secondaries` (block) | no | nested block |
| `pi_secondaries.name` | no | API: Name of the host group |
| `pi_secondaries.workspace` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_host_group_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_host_group_invalid_pi_name` | `pi_name` |
| `ibm_pi_host_group_invalid_pi_remove` | `pi_remove` |
| `ibm_pi_host_group_invalid_pi_hosts__display_name` | `pi_hosts.display_name` |
| `ibm_pi_host_group_invalid_pi_hosts__sys_type` | `pi_hosts.sys_type` |
| `ibm_pi_host_group_invalid_pi_hosts__user_tags` | `pi_hosts.user_tags` |
| `ibm_pi_host_group_invalid_pi_secondaries__name` | `pi_secondaries.name` |
| `ibm_pi_host_group_invalid_pi_secondaries__workspace` | `pi_secondaries.workspace` |

## Example

### Valid

```hcl
resource "ibm_pi_host_group" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_name = "example"
  pi_hosts {
    display_name = "example"
    sys_type = "example"
  }
}
```
