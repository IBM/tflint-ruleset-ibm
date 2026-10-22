# ibm_en_integration_data

Schema validation for the ibm_en_integration data source.

This rule validates `ibm_en_integration` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_en_integration_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_guid` | yes | — |
| `integration_id` | yes | — |
| `metadata` | no | — |
| `type` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_integration_data_invalid_instance_guid` | `instance_guid` |
| `ibm_en_integration_data_invalid_integration_id` | `integration_id` |

## Example

### Valid

```hcl
data "ibm_en_integration" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  integration_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
