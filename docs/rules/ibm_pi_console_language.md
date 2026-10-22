# ibm_pi_console_language

Schema validation for the ibm_pi_console_language resource.

This rule validates `ibm_pi_console_language` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_console_language" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `pi_cloud_instance_id` | yes | — |
| `pi_instance_name` | yes | — |
| `pi_language_code` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_console_language_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_console_language_invalid_pi_instance_name` | `pi_instance_name` |
| `ibm_pi_console_language_invalid_pi_language_code` | `pi_language_code` |

## Example

### Valid

```hcl
resource "ibm_pi_console_language" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_instance_name = "example"
  pi_language_code = "example"
}
```
