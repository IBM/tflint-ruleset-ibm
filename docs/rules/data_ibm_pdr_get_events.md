# ibm_pdr_get_events_data

Schema validation for the ibm_pdr_get_events data source.

This rule validates `ibm_pdr_get_events` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `leaf-unique` (`dr-automation-powervs`).

## Configuration

```hcl
rule "ibm_pdr_get_events_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | — |
| `event` | no | — |
| `events` | no | — |
| `from_time` | no | — |
| `instance_id` | yes | — |
| `time` | no | pattern `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[+\-]\d{2}:\d{2})$` |
| `to_time` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pdr_get_events_data_invalid_accept_language` | `accept_language` |
| `ibm_pdr_get_events_data_invalid_from_time` | `from_time` |
| `ibm_pdr_get_events_data_invalid_instance_id` | `instance_id` |
| `ibm_pdr_get_events_data_invalid_time` | `time` |
| `ibm_pdr_get_events_data_invalid_to_time` | `to_time` |

## Example

### Valid

```hcl
data "ibm_pdr_get_events" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
