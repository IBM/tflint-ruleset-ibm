# ibm_en_pre_defined_templates_data

Schema validation for the ibm_en_pre_defined_templates data source.

This rule validates `ibm_en_pre_defined_templates` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PredefinedTemplate` (`event-notifications`).

## Configuration

```hcl
rule "ibm_en_pre_defined_templates_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_guid` | yes | — |
| `search_key` | no | — |
| `source` | yes | API: The type of source |
| `templates` | no | — |
| `total_count` | no | — |
| `type` | yes | API: The type of template |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_pre_defined_templates_data_invalid_instance_guid` | `instance_guid` |
| `ibm_en_pre_defined_templates_data_invalid_search_key` | `search_key` |
| `ibm_en_pre_defined_templates_data_invalid_source` | `source` |
| `ibm_en_pre_defined_templates_data_invalid_type` | `type` |

## Example

### Valid

```hcl
data "ibm_en_pre_defined_templates" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  source = "aaaaaaa"
  type = "aaaaaaa"
}
```
