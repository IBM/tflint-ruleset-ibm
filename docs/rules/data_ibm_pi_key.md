# ibm_pi_key_data

Reads ibm_pi_key. Name of SSH key.

This rule validates `ibm_pi_key` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_key_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `creation_date` | no | — |
| `description` | no | — |
| `name` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_key_name` | no | — |
| `pi_ssh_key_id` | no | — |
| `primary_workspace` | no | — |
| `ssh_key` | no | — |
| `ssh_key_id` | no | — |
| `visibility` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_key_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_key_data_invalid_pi_key_name` | `pi_key_name` |
| `ibm_pi_key_data_invalid_pi_ssh_key_id` | `pi_ssh_key_id` |

## Example

### Valid

```hcl
data "ibm_pi_key" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
