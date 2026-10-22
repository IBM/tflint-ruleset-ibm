# ibm_is_instance_software_attachments_data

Schema validation for the ibm_is_instance_software_attachments data source.

This rule validates `ibm_is_instance_software_attachments` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_instance_software_attachments_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_id` | yes | — |
| `software_attachments` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_software_attachments_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_is_instance_software_attachments" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
