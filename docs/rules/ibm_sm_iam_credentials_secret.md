# ibm_sm_iam_credentials_secret

Manages ibm_sm_iam_credentials_secret. A human-readable name to assign to your secret.

This rule validates `ibm_sm_iam_credentials_secret` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `IAMCredentialsSecretPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_iam_credentials_secret" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_groups` | no | pattern `^AccessGroupId-[a-z0-9-]+[a-z0-9]$` |
| `account_id` | no | pattern `^[a-z0-9]*$` |
| `api_key` | no | — |
| `api_key_id` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `custom_metadata` | no | API: The secret metadata that a user can customize. |
| `description` | no | API: An extended description of your secret. To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group. |
| `downloaded` | no | — |
| `expiration_date` | no | — |
| `labels` | no | API: Label that you can use to search for secrets in your instance. Label can be between 2-64 characters, including spaces. To protect your privacy, do not use personal data, such as your name or location, as a label for your secret. |
| `locks_total` | no | — |
| `name` | yes | pattern `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$` |
| `next_rotation_date` | no | — |
| `retrieved_at` | no | — |
| `reuse_api_key` | no | API: (IAM credentials) This parameter indicates whether to reuse the service ID and API key for future read operations. If it is set to `true`, the service reuses the current credentials. If it is set to `false`, a new service ID and API key are generated each time that the secret is read or accessed. |
| `secret_group_id` | no | pattern `^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|default)$` |
| `secret_id` | no | — |
| `secret_type` | no | — |
| `service_id` | no | pattern `^ServiceId-[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `service_id_is_static` | no | — |
| `state` | no | — |
| `state_description` | no | — |
| `ttl` | yes | pattern `^[0-9]+[s,m,h,d]{0,1}$` |
| `updated_at` | no | — |
| `version_custom_metadata` | no | API: The secret version metadata that a user can customize. |
| `versions_total` | no | — |
| `rotation` (block) | no | nested block |
| `rotation.auto_rotate` | no | API: This field indicates whether Secrets Manager rotates your secret automatically. The default is `false`. If `auto_rotate` is set to `true` the service rotates your secret based on the defined interval. |
| `rotation.interval` | no | range 1–1000 |
| `rotation.unit` | no | one of: `day`, `month`, `hour` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_iam_credentials_secret_invalid_access_groups` | `access_groups` |
| `ibm_sm_iam_credentials_secret_invalid_account_id` | `account_id` |
| `ibm_sm_iam_credentials_secret_invalid_custom_metadata` | `custom_metadata` |
| `ibm_sm_iam_credentials_secret_invalid_description` | `description` |
| `ibm_sm_iam_credentials_secret_invalid_labels` | `labels` |
| `ibm_sm_iam_credentials_secret_invalid_name` | `name` |
| `ibm_sm_iam_credentials_secret_invalid_reuse_api_key` | `reuse_api_key` |
| `ibm_sm_iam_credentials_secret_invalid_secret_group_id` | `secret_group_id` |
| `ibm_sm_iam_credentials_secret_invalid_service_id` | `service_id` |
| `ibm_sm_iam_credentials_secret_invalid_ttl` | `ttl` |
| `ibm_sm_iam_credentials_secret_invalid_version_custom_metadata` | `version_custom_metadata` |
| `ibm_sm_iam_credentials_secret_invalid_rotation__auto_rotate` | `rotation.auto_rotate` |
| `ibm_sm_iam_credentials_secret_invalid_rotation__interval` | `rotation.interval` |
| `ibm_sm_iam_credentials_secret_invalid_rotation__unit` | `rotation.unit` |

## Example

### Valid

```hcl
resource "ibm_sm_iam_credentials_secret" "test" {
  name = "example-name"
  ttl = "1"
}
```
