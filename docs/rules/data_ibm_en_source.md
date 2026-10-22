# ibm_en_source_data

Reads ibm_en_source. Source name.

This rule validates `ibm_en_source` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_en_source_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | — |
| `enabled` | no | — |
| `instance_guid` | yes | — |
| `name` | no | — |
| `source_id` | yes | — |
| `store_notifications` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_source_data_invalid_instance_guid` | `instance_guid` |
| `ibm_en_source_data_invalid_source_id` | `source_id` |

## Example

### Valid

```hcl
data "ibm_en_source" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  source_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
