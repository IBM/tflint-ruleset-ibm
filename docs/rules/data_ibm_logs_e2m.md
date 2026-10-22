# ibm_logs_e2m_data

Reads ibm_logs_e2m. Name of the E2M.

This rule validates `ibm_logs_e2m` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_e2m_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `create_time` | no | — |
| `description` | no | — |
| `is_internal` | no | — |
| `logs_e2m_id` | yes | — |
| `logs_query` | no | — |
| `metric_fields` | no | — |
| `metric_labels` | no | — |
| `name` | no | — |
| `permutations` | no | — |
| `type` | no | — |
| `update_time` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_e2m_data_invalid_logs_e2m_id` | `logs_e2m_id` |

## Example

### Valid

```hcl
data "ibm_logs_e2m" "test" {
  logs_e2m_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
