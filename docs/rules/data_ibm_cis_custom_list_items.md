# ibm_cis_custom_list_items_data

Schema validation for the ibm_cis_custom_list_items data source.

This rule validates `ibm_cis_custom_list_items` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cis_custom_list_items_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `item_id` | no | — |
| `items` | no | — |
| `list_id` | yes | — |
| `search` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_custom_list_items_data_invalid_cis_id` | `cis_id` |
| `ibm_cis_custom_list_items_data_invalid_item_id` | `item_id` |
| `ibm_cis_custom_list_items_data_invalid_list_id` | `list_id` |
| `ibm_cis_custom_list_items_data_invalid_search` | `search` |

## Example

### Valid

```hcl
data "ibm_cis_custom_list_items" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  list_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
