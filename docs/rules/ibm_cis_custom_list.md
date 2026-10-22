# ibm_cis_custom_list

Manages ibm_cis_custom_list. Type of the List

This rule validates `ibm_cis_custom_list` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `create_custom_list_req` (`cis`).

## Configuration

```hcl
rule "ibm_cis_custom_list" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `description` | no | API: An informative summary of the list. |
| `kind` | yes | one of: `ip`, `asn`, `hostname` |
| `list_id` | no | — |
| `name` | yes | API: An informative name for the list. Use this name in rule expressions. |
| `num_items` | no | — |
| `num_referencing_filters` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_custom_list_invalid_cis_id` | `cis_id` |
| `ibm_cis_custom_list_invalid_description` | `description` |
| `ibm_cis_custom_list_invalid_kind` | `kind` |
| `ibm_cis_custom_list_invalid_name` | `name` |

## Example

### Valid

```hcl
resource "ibm_cis_custom_list" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  kind = "ip"
  name = "example"
}
```
