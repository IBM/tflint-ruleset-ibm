# ibm_code_engine_secret

Manages ibm_code_engine_secret. The name of the secret.

This rule validates `ibm_code_engine_secret` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Secret` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_secret" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `data` | no | API: Data container that allows to specify config parameters and their values as a key-value map. Each key field must consist of alphanumeric characters, `-`, `_` or `.` and must not exceed a max length of 253 characters. Each value field can consists of any character and must not exceed a max length of 1048576 characters. |
| `entity_tag` | no | — |
| `etag` | no | — |
| `format` | yes | one of: `basic_auth`, `generic`, `hmac_auth`, `other`, `registry`, `service_access`, `service_operator`, `ssh_auth`, `tls`; pattern `^(generic|ssh_auth|basic_auth|hmac_auth|tls|service_access|registry|service_operator|other)$` |
| `generated_by` | no | — |
| `href` | no | — |
| `name` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `region` | no | — |
| `resource_type` | no | — |
| `secret_id` | no | — |
| `service_access` (block) | no | nested block |
| `service_access.resource_key` (block) | no | nested block |
| `service_access.resource_key.id` | no | pattern `^[a-z0-9][\-a-z0-9]*[a-z0-9]$` |
| `service_access.role` (block) | no | nested block |
| `service_access.role.crn` | no | crn format |
| `service_access.service_instance` (block) | yes | nested block |
| `service_access.service_instance.id` | no | pattern `^[a-z0-9][\-a-z0-9]*[a-z0-9]$` |
| `service_access.serviceid` (block) | no | nested block |
| `service_access.serviceid.id` | no | pattern `^ServiceId-[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `service_operator` (block) | no | nested block |
| `service_operator.resource_group_ids` | yes | pattern `^[a-z0-9]*$` |
| `service_operator.serviceid` (block) | yes | nested block |
| `service_operator.serviceid.id` | no | pattern `^ServiceId-[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_secret_invalid_data` | `data` |
| `ibm_code_engine_secret_invalid_format` | `format` |
| `ibm_code_engine_secret_invalid_name` | `name` |
| `ibm_code_engine_secret_invalid_project_id` | `project_id` |
| `ibm_code_engine_secret_invalid_service_access__resource_key__id` | `service_access.resource_key.id` |
| `ibm_code_engine_secret_invalid_service_access__role__crn` | `service_access.role.crn` |
| `ibm_code_engine_secret_invalid_service_access__service_instance__id` | `service_access.service_instance.id` |
| `ibm_code_engine_secret_invalid_service_access__serviceid__id` | `service_access.serviceid.id` |
| `ibm_code_engine_secret_invalid_service_operator__resource_group_ids` | `service_operator.resource_group_ids` |
| `ibm_code_engine_secret_invalid_service_operator__serviceid__id` | `service_operator.serviceid.id` |

## Example

### Valid

```hcl
resource "ibm_code_engine_secret" "test" {
  format = "basic_auth"
  name = "example-name"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
