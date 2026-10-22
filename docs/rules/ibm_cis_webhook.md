# ibm_cis_webhook

Manages ibm_cis_webhook. Webhook Name

This rule validates `ibm_cis_webhook` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `alert_webhook_input` (`cis`).

## Configuration

```hcl
rule "ibm_cis_webhook" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `name` | yes | API: Webhook Name |
| `secret` | no | API: The optional secret or API key needed to use the webhook |
| `type` | no | — |
| `url` | no | API: Webhook url |
| `webhook_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_webhook_invalid_cis_id` | `cis_id` |
| `ibm_cis_webhook_invalid_name` | `name` |
| `ibm_cis_webhook_invalid_secret` | `secret` |
| `ibm_cis_webhook_invalid_url` | `url` |

## Example

### Valid

```hcl
resource "ibm_cis_webhook" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  name = "example"
}
```
