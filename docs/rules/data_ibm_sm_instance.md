# ibm_sm_instance_data

Reads ibm_sm_instance. The instance name.

This rule validates `ibm_sm_instance` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_sm_instance_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `encryption` | no | — |
| `endpoints` | no | — |
| `href` | no | — |
| `instance_crn` | no | — |
| `instance_id` | yes | — |
| `name` | no | — |
| `plan` | no | — |
| `vault_cluster` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_instance_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_sm_instance" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
