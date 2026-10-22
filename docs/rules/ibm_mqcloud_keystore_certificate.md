# ibm_mqcloud_keystore_certificate

Schema validation for the ibm_mqcloud_keystore_certificate resource.

This rule validates `ibm_mqcloud_keystore_certificate` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/{service_instance_guid}/queue_managers/{queue_manager_id}/certificates/trust_store:post` (`mq-on-cloud`).

## Configuration

```hcl
rule "ibm_mqcloud_keystore_certificate" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `certificate_file` | yes | API: The filename and path of the certificate to be uploaded. |
| `certificate_id` | no | — |
| `certificate_type` | no | — |
| `dns_names` | no | — |
| `dns_names_total_count` | no | — |
| `expiry` | no | — |
| `fingerprint_sha256` | no | — |
| `href` | no | — |
| `is_default` | no | — |
| `issued` | no | — |
| `issuer_cn` | no | — |
| `issuer_dn` | no | — |
| `label` | yes | pattern `^[a-zA-Z0-9_.]*$` |
| `queue_manager_id` | yes | pattern `^[0-9a-fA-F]{32}$` |
| `service_instance_guid` | yes | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `subject_cn` | no | — |
| `subject_dn` | no | — |
| `config` (block) | no | nested block |
| `config.ams` (block) | no | nested block |
| `config.ams.channels` (block) | no | nested block |
| `config.ams.channels.name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_mqcloud_keystore_certificate_invalid_certificate_file` | `certificate_file` |
| `ibm_mqcloud_keystore_certificate_invalid_label` | `label` |
| `ibm_mqcloud_keystore_certificate_invalid_queue_manager_id` | `queue_manager_id` |
| `ibm_mqcloud_keystore_certificate_invalid_service_instance_guid` | `service_instance_guid` |
| `ibm_mqcloud_keystore_certificate_invalid_config__ams__channels__name` | `config.ams.channels.name` |

## Example

### Valid

```hcl
resource "ibm_mqcloud_keystore_certificate" "test" {
  certificate_file = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  label = "example"
  queue_manager_id = "0123456789abcdef0123456789abcdef"
  service_instance_guid = "01234567-89ab-cdef-0123-456789abcdef"
}
```
