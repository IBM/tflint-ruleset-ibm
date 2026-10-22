# ibm_cbr_rule_data

Schema validation for the ibm_cbr_rule data source.

This rule validates `ibm_cbr_rule` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cbr_rule_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `contexts` | no | — |
| `created_at` | no | — |
| `created_by_id` | no | — |
| `crn` | no | — |
| `description` | no | — |
| `enforcement_mode` | no | — |
| `href` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `operations` | no | — |
| `resources` | no | — |
| `rule_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cbr_rule_data_invalid_rule_id` | `rule_id` |

## Example

### Valid

```hcl
data "ibm_cbr_rule" "test" {
  rule_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
