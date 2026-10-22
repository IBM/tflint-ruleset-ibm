# ibm_logs_view_folder_data

Reads ibm_logs_view_folder. Folder name.

This rule validates `ibm_logs_view_folder` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_view_folder_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `logs_view_folder_id` | yes | — |
| `name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_view_folder_data_invalid_logs_view_folder_id` | `logs_view_folder_id` |

## Example

### Valid

```hcl
data "ibm_logs_view_folder" "test" {
  logs_view_folder_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
