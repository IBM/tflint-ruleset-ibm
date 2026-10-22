# ibm_sm_private_certificate_metadata_data

Reads ibm_sm_private_certificate_metadata. The human-readable name of your secret.

This rule validates `ibm_sm_private_certificate_metadata` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SecretLocks` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_private_certificate_metadata_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alt_names` | no | — |
| `certificate_authority` | no | — |
| `certificate_template` | no | — |
| `common_name` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `custom_metadata` | no | — |
| `description` | no | — |
| `downloaded` | no | — |
| `expiration_date` | no | — |
| `issuer` | no | — |
| `key_algorithm` | no | — |
| `labels` | no | — |
| `locks_total` | no | — |
| `name` | no | — |
| `next_rotation_date` | no | — |
| `retrieved_at` | no | — |
| `revocation_time_rfc3339` | no | — |
| `revocation_time_seconds` | no | — |
| `rotation` | no | — |
| `secret_group_id` | no | — |
| `secret_id` | yes | pattern `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}` |
| `secret_type` | no | — |
| `serial_number` | no | — |
| `signing_algorithm` | no | — |
| `state` | no | — |
| `state_description` | no | — |
| `updated_at` | no | — |
| `validity` | no | — |
| `versions_total` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_private_certificate_metadata_data_invalid_secret_id` | `secret_id` |

## Example

### Valid

```hcl
variable "secret_id" {
  sensitive = true
}

data "ibm_sm_private_certificate_metadata" "test" {
  secret_id = var.secret_id
}
```
