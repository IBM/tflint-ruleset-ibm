# ibm_pdr_event_data

Schema validation for the ibm_pdr_event data source.

This rule validates `ibm_pdr_event` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Event` (`dr-automation-powervs`).

## Configuration

```hcl
rule "ibm_pdr_event_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | — |
| `action` | no | — |
| `api_source` | no | — |
| `event_id` | yes | pattern `^[a-zA-Z0-9\-_]+$` |
| `instance_id` | yes | — |
| `level` | no | — |
| `message` | no | — |
| `message_data` | no | — |
| `metadata` | no | — |
| `resource` | no | — |
| `time` | no | — |
| `timestamp` | no | — |
| `user` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pdr_event_data_invalid_accept_language` | `accept_language` |
| `ibm_pdr_event_data_invalid_event_id` | `event_id` |
| `ibm_pdr_event_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_pdr_event" "test" {
  event_id = "example-name"
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
