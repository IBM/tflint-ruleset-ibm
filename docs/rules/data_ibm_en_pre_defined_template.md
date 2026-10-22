# ibm_en_pre_defined_template_data

Reads ibm_en_pre_defined_template. Template name.

This rule validates `ibm_en_pre_defined_template` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_en_pre_defined_template_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | — |
| `instance_guid` | yes | — |
| `name` | no | — |
| `params` | no | — |
| `source` | no | — |
| `template_id` | yes | — |
| `type` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_pre_defined_template_data_invalid_instance_guid` | `instance_guid` |
| `ibm_en_pre_defined_template_data_invalid_template_id` | `template_id` |

## Example

### Valid

```hcl
data "ibm_en_pre_defined_template" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  template_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
