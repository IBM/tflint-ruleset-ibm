# ibm_sm_service_credentials_secret

Manages ibm_sm_service_credentials_secret. A human-readable name to assign to your secret.

This rule validates `ibm_sm_service_credentials_secret` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CustomCredentialsSecretPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_service_credentials_secret" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `created_by` | no | — |
| `credentials` | no | — |
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
| `secret_group_id` | no | pattern `^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|default)$` |
| `secret_id` | no | — |
| `secret_type` | no | — |
| `state` | no | — |
| `state_description` | no | — |
| `ttl` | no | pattern `^[0-9]+[s,m,h,d]{0,1}$` |
| `updated_at` | no | — |
| `version_custom_metadata` | no | API: The secret version metadata that a user can customize. |
| `versions_total` | no | — |
| `rotation` (block) | no | nested block |
| `rotation.auto_rotate` | no | API: This field indicates whether Secrets Manager rotates your secret automatically. The default is `false`. If `auto_rotate` is set to `true` the service rotates your secret based on the defined interval. |
| `rotation.interval` | no | range 1–1000 |
| `rotation.unit` | no | one of: `day`, `month`, `hour` |
| `source_service` (block) | yes | nested block |
| `source_service.parameters` | no | API: The parameters that are passed to custom credentials engine. Allowed types are string, integer, and boolean. |
| `source_service.instance` (block) | yes | nested block |
| `source_service.instance.crn` | yes | crn format |
| `source_service.role` (block) | no | nested block |
| `source_service.role.crn` | no | crn format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_service_credentials_secret_invalid_custom_metadata` | `custom_metadata` |
| `ibm_sm_service_credentials_secret_invalid_description` | `description` |
| `ibm_sm_service_credentials_secret_invalid_labels` | `labels` |
| `ibm_sm_service_credentials_secret_invalid_name` | `name` |
| `ibm_sm_service_credentials_secret_invalid_secret_group_id` | `secret_group_id` |
| `ibm_sm_service_credentials_secret_invalid_ttl` | `ttl` |
| `ibm_sm_service_credentials_secret_invalid_version_custom_metadata` | `version_custom_metadata` |
| `ibm_sm_service_credentials_secret_invalid_rotation__auto_rotate` | `rotation.auto_rotate` |
| `ibm_sm_service_credentials_secret_invalid_rotation__interval` | `rotation.interval` |
| `ibm_sm_service_credentials_secret_invalid_rotation__unit` | `rotation.unit` |
| `ibm_sm_service_credentials_secret_invalid_source_service__parameters` | `source_service.parameters` |
| `ibm_sm_service_credentials_secret_invalid_source_service__instance__crn` | `source_service.instance.crn` |
| `ibm_sm_service_credentials_secret_invalid_source_service__role__crn` | `source_service.role.crn` |

## Example

### Valid

```hcl
resource "ibm_sm_service_credentials_secret" "test" {
  name = "example-name"
  source_service {
    instance {
      crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
    }
  }
}
```
