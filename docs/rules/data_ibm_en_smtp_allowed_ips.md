# ibm_en_smtp_allowed_ips_data

Schema validation for the ibm_en_smtp_allowed_ips data source.

This rule validates `ibm_en_smtp_allowed_ips` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `IBMCloudObjectStorageDestinationConfig` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_smtp_allowed_ips_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `en_smtp_allowed_ips_id` | yes | — |
| `instance_id` | yes | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]` |
| `subnets` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_smtp_allowed_ips_data_invalid_en_smtp_allowed_ips_id` | `en_smtp_allowed_ips_id` |
| `ibm_en_smtp_allowed_ips_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_en_smtp_allowed_ips" "test" {
  en_smtp_allowed_ips_id = "01234567-89ab-cdef-0123-456789abcdef"
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
