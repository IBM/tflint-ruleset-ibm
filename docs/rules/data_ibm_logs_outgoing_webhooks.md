# ibm_logs_outgoing_webhooks_data

Schema validation for the ibm_logs_outgoing_webhooks data source.

This rule validates `ibm_logs_outgoing_webhooks` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_logs_outgoing_webhooks_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `outgoing_webhooks` | no | — |
| `type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_outgoing_webhooks_data_invalid_type` | `type` |

## Example

### Valid

```hcl
data "ibm_logs_outgoing_webhooks" "test" {
}
```
