# ibm_logs_enrichment

Schema validation for the ibm_logs_enrichment resource.

This rule validates `ibm_logs_enrichment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Enrichment` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_enrichment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `enrichment_id` | no | — |
| `field_name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `enrichment_type` (block) | yes | nested block |
| `enrichment_type.custom_enrichment` (block) | no | nested block |
| `enrichment_type.custom_enrichment.id` | yes | range 0–4294967295 |
| `enrichment_type.geo_ip` (block) | no | nested block |
| `enrichment_type.suspicious_ip` (block) | no | nested block |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_enrichment_invalid_field_name` | `field_name` |
| `ibm_logs_enrichment_invalid_enrichment_type__custom_enrichment__id` | `enrichment_type.custom_enrichment.id` |

## Example

### Valid

```hcl
resource "ibm_logs_enrichment" "test" {
  field_name = "example-name"
  enrichment_type {
  }
}
```
