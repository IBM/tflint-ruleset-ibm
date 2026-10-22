# ibm_sm_configurations_data

Schema validation for the ibm_sm_configurations data source.

This rule validates `ibm_sm_configurations` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_sm_configurations_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `configurations` | no | — |
| `groups` | no | — |
| `search` | no | — |
| `secret_types` | no | — |
| `sort` | no | — |
| `total_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_configurations_data_invalid_groups` | `groups` |
| `ibm_sm_configurations_data_invalid_search` | `search` |
| `ibm_sm_configurations_data_invalid_secret_types` | `secret_types` |
| `ibm_sm_configurations_data_invalid_sort` | `sort` |

## Example

### Valid

```hcl
data "ibm_sm_configurations" "test" {
}
```
