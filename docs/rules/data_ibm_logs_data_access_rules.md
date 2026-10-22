# ibm_logs_data_access_rules_data

Schema validation for the ibm_logs_data_access_rules data source.

This rule validates `ibm_logs_data_access_rules` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_data_access_rules_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `data_access_rules` | no | — |
| `logs_data_access_rules_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_data_access_rules_data_invalid_logs_data_access_rules_id` | `logs_data_access_rules_id` |

## Example

### Valid

```hcl
data "ibm_logs_data_access_rules" "test" {
}
```
