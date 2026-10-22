# ibm_pi_key

Manages ibm_pi_key. User defined name for the SSH key.

This rule validates `ibm_pi_key` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_key" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `creation_date` | no | — |
| `key` | no | — |
| `name` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_description` | no | — |
| `pi_key_name` | yes | — |
| `pi_ssh_key` | yes | — |
| `pi_visibility` | no | one of: `account`, `workspace` |
| `primary_workspace` | no | — |
| `ssh_key_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_key_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_key_invalid_pi_description` | `pi_description` |
| `ibm_pi_key_invalid_pi_key_name` | `pi_key_name` |
| `ibm_pi_key_invalid_pi_ssh_key` | `pi_ssh_key` |
| `ibm_pi_key_invalid_pi_visibility` | `pi_visibility` |

## Example

### Valid

```hcl
resource "ibm_pi_key" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_key_name = "example"
  pi_ssh_key = "example"
}
```
