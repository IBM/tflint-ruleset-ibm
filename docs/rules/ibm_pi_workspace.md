# ibm_pi_workspace

Schema validation for the ibm_pi_workspace resource.

This rule validates `ibm_pi_workspace` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_workspace" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `pi_datacenter` | yes | — |
| `pi_name` | yes | — |
| `pi_parameters` | no | — |
| `pi_plan` | no | one of: `private`, `public` |
| `pi_resource_group_id` | yes | — |
| `pi_user_tags` | no | — |
| `pi_workspace_details` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_workspace_invalid_pi_datacenter` | `pi_datacenter` |
| `ibm_pi_workspace_invalid_pi_name` | `pi_name` |
| `ibm_pi_workspace_invalid_pi_parameters` | `pi_parameters` |
| `ibm_pi_workspace_invalid_pi_plan` | `pi_plan` |
| `ibm_pi_workspace_invalid_pi_resource_group_id` | `pi_resource_group_id` |
| `ibm_pi_workspace_invalid_pi_user_tags` | `pi_user_tags` |

## Example

### Valid

```hcl
resource "ibm_pi_workspace" "test" {
  pi_datacenter = "example"
  pi_name = "example"
  pi_resource_group_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
