# ibm_appid_cloud_directory_template_data

Schema validation for the ibm_appid_cloud_directory_template data source.

This rule validates `ibm_appid_cloud_directory_template` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_appid_cloud_directory_template_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `base64_encoded_html_body` | no | — |
| `html_body` | no | — |
| `language` | no | — |
| `plain_text_body` | no | — |
| `subject` | no | — |
| `template_name` | yes | — |
| `tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_cloud_directory_template_data_invalid_language` | `language` |
| `ibm_appid_cloud_directory_template_data_invalid_template_name` | `template_name` |
| `ibm_appid_cloud_directory_template_data_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
data "ibm_appid_cloud_directory_template" "test" {
  template_name = "example"
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
