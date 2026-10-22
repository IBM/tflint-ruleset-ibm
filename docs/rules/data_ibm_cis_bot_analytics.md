# ibm_cis_bot_analytics_data

Schema validation for the ibm_cis_bot_analytics data source.

This rule validates `ibm_cis_bot_analytics` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ratelimit_analytics_resp` (`cis`).

## Configuration

```hcl
rule "ibm_cis_bot_analytics_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `result` | no | — |
| `since` | yes | API: The beginning time of analytics. |
| `type` | yes | — |
| `until` | yes | API: The end time of analytics. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_bot_analytics_data_invalid_cis_id` | `cis_id` |
| `ibm_cis_bot_analytics_data_invalid_domain_id` | `domain_id` |
| `ibm_cis_bot_analytics_data_invalid_since` | `since` |
| `ibm_cis_bot_analytics_data_invalid_type` | `type` |
| `ibm_cis_bot_analytics_data_invalid_until` | `until` |

## Example

### Valid

```hcl
data "ibm_cis_bot_analytics" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  since = "example"
  type = "example"
  until = "example"
}
```
