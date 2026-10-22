# ibm_logs_stream

Manages ibm_logs_stream. The name of the Event stream.

This rule validates `ibm_logs_stream` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Stream` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_stream" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `compression_type` | no | one of: `gzip`, `unspecified` |
| `created_at` | no | — |
| `dpxl_expression` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `is_active` | no | API: Whether the event stream is active. |
| `name` | yes | pattern `^[a-zA-ZÀ-ÖØ-öø-ÿĀ-ſΑ-ωА-я一-龥ぁ-ゔァ-ヴー々〆〤0-9_\.,\-"{}()\[\]=!:#\/$|' ]+$` |
| `streams_id` | no | — |
| `updated_at` | no | — |
| `ibm_event_streams` (block) | no | nested block |
| `ibm_event_streams.brokers` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `ibm_event_streams.topic` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_stream_invalid_compression_type` | `compression_type` |
| `ibm_logs_stream_invalid_dpxl_expression` | `dpxl_expression` |
| `ibm_logs_stream_invalid_is_active` | `is_active` |
| `ibm_logs_stream_invalid_name` | `name` |
| `ibm_logs_stream_invalid_ibm_event_streams__brokers` | `ibm_event_streams.brokers` |
| `ibm_logs_stream_invalid_ibm_event_streams__topic` | `ibm_event_streams.topic` |

## Example

### Valid

```hcl
resource "ibm_logs_stream" "test" {
  dpxl_expression = "example-name"
  name = "example-name"
}
```
