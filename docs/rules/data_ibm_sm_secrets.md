# ibm_sm_secrets_data

Schema validation for the ibm_sm_secrets data source.

This rule validates `ibm_sm_secrets` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_sm_secrets_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `groups` | no | — |
| `match_all_labels` | no | — |
| `search` | no | — |
| `secret_types` | no | — |
| `secrets` | no | — |
| `sort` | no | — |
| `total_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_secrets_data_invalid_groups` | `groups` |
| `ibm_sm_secrets_data_invalid_match_all_labels` | `match_all_labels` |
| `ibm_sm_secrets_data_invalid_search` | `search` |
| `ibm_sm_secrets_data_invalid_secret_types` | `secret_types` |
| `ibm_sm_secrets_data_invalid_sort` | `sort` |

## Example

### Valid

```hcl
data "ibm_sm_secrets" "test" {
}
```
