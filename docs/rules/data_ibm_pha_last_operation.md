# ibm_pha_last_operation_data

Schema validation for the ibm_pha_last_operation data source.

This rule validates `ibm_pha_last_operation` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pha_last_operation_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | — |
| `deployment_name` | no | — |
| `if_none_match` | no | — |
| `instance_id` | yes | — |
| `provision_id` | no | — |
| `resource_group` | no | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pha_last_operation_data_invalid_accept_language` | `accept_language` |
| `ibm_pha_last_operation_data_invalid_if_none_match` | `if_none_match` |
| `ibm_pha_last_operation_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_pha_last_operation" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
