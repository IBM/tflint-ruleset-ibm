# ibm_sm_arbitrary_secret

Manages ibm_sm_arbitrary_secret. A human-readable name to assign to your secret.

This rule validates `ibm_sm_arbitrary_secret` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ArbitrarySecretPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_arbitrary_secret" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `custom_metadata` | no | API: The secret metadata that a user can customize. |
| `description` | no | API: An extended description of your secret. To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group. |
| `downloaded` | no | — |
| `expiration_date` | no | API: The date when the secret material expires. The date format follows the `RFC 3339` format. Supported secret types: Arbitrary, username_password |
| `labels` | no | API: Label that you can use to search for secrets in your instance. Label can be between 2-64 characters, including spaces. To protect your privacy, do not use personal data, such as your name or location, as a label for your secret. |
| `locks_total` | no | — |
| `name` | yes | pattern `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$` |
| `payload` | yes | API: The secret data that is assigned to an `ARBITRARY` secret. |
| `retrieved_at` | no | — |
| `secret_group_id` | no | pattern `^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|default)$` |
| `secret_id` | no | — |
| `secret_type` | no | — |
| `state` | no | — |
| `state_description` | no | — |
| `updated_at` | no | — |
| `version_custom_metadata` | no | API: The secret version metadata that a user can customize. |
| `versions_total` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_arbitrary_secret_invalid_custom_metadata` | `custom_metadata` |
| `ibm_sm_arbitrary_secret_invalid_description` | `description` |
| `ibm_sm_arbitrary_secret_invalid_expiration_date` | `expiration_date` |
| `ibm_sm_arbitrary_secret_invalid_labels` | `labels` |
| `ibm_sm_arbitrary_secret_invalid_name` | `name` |
| `ibm_sm_arbitrary_secret_invalid_payload` | `payload` |
| `ibm_sm_arbitrary_secret_invalid_secret_group_id` | `secret_group_id` |
| `ibm_sm_arbitrary_secret_invalid_version_custom_metadata` | `version_custom_metadata` |

## Example

### Valid

```hcl
variable "payload" {
  sensitive = true
}

resource "ibm_sm_arbitrary_secret" "test" {
  name = "example-name"
  payload = var.payload
}
```
