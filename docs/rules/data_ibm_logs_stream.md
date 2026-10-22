# ibm_logs_stream_data

Reads ibm_logs_stream. The name of the Event stream.

This rule validates `ibm_logs_stream` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_stream_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `compression_type` | no | — |
| `created_at` | no | — |
| `dpxl_expression` | no | — |
| `ibm_event_streams` | no | — |
| `is_active` | no | — |
| `logs_streams_id` | yes | — |
| `name` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_stream_data_invalid_logs_streams_id` | `logs_streams_id` |

## Example

### Valid

```hcl
data "ibm_logs_stream" "test" {
  logs_streams_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
