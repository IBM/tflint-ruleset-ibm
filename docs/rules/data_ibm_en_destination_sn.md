# ibm_en_destination_sn_data

Reads ibm_en_destination_sn. Destination name.

This rule validates `ibm_en_destination_sn` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Destination` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_destination_sn_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `collect_failed_events` | no | — |
| `config` | no | — |
| `description` | no | — |
| `destination_id` | yes | pattern `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}` |
| `instance_guid` | yes | — |
| `name` | no | — |
| `subscription_count` | no | — |
| `subscription_names` | no | — |
| `type` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_destination_sn_data_invalid_destination_id` | `destination_id` |
| `ibm_en_destination_sn_data_invalid_instance_guid` | `instance_guid` |

## Example

### Valid

```hcl
data "ibm_en_destination_sn" "test" {
  destination_id = "01234567-89ab-cdef-0123-456789abcdef"
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
