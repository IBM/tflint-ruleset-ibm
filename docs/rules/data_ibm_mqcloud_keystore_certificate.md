# ibm_mqcloud_keystore_certificate_data

Schema validation for the ibm_mqcloud_keystore_certificate data source.

This rule validates `ibm_mqcloud_keystore_certificate` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/{service_instance_guid}/queue_managers/{queue_manager_id}/certificates/trust_store:post` (`mq-on-cloud`).

## Configuration

```hcl
rule "ibm_mqcloud_keystore_certificate_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `key_store` | no | — |
| `label` | no | pattern `^[a-zA-Z0-9_.]*$` |
| `queue_manager_id` | yes | pattern `^[0-9a-fA-F]{32}$` |
| `service_instance_guid` | yes | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `total_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_mqcloud_keystore_certificate_data_invalid_label` | `label` |
| `ibm_mqcloud_keystore_certificate_data_invalid_queue_manager_id` | `queue_manager_id` |
| `ibm_mqcloud_keystore_certificate_data_invalid_service_instance_guid` | `service_instance_guid` |

## Example

### Valid

```hcl
data "ibm_mqcloud_keystore_certificate" "test" {
  queue_manager_id = "0123456789abcdef0123456789abcdef"
  service_instance_guid = "01234567-89ab-cdef-0123-456789abcdef"
}
```
