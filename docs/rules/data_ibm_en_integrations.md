# ibm_en_integrations_data

Schema validation for the ibm_en_integrations data source.

This rule validates `ibm_en_integrations` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_en_integrations_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_guid` | yes | — |
| `integrations` | no | — |
| `search_key` | no | — |
| `total_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_integrations_data_invalid_instance_guid` | `instance_guid` |
| `ibm_en_integrations_data_invalid_search_key` | `search_key` |

## Example

### Valid

```hcl
data "ibm_en_integrations" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
