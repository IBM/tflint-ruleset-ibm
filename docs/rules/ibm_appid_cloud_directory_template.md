# ibm_appid_cloud_directory_template

Schema validation for the ibm_appid_cloud_directory_template resource.

This rule validates `ibm_appid_cloud_directory_template` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `getTemplate` (`app-id-management`).

## Configuration

```hcl
rule "ibm_appid_cloud_directory_template" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `base64_encoded_html_body` | no | — |
| `html_body` | no | API: The HTML body of the email. |
| `language` | no | — |
| `plain_text_body` | no | API: The plain text body of the email. |
| `subject` | yes | API: The subject of the email. |
| `template_name` | yes | — |
| `tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_cloud_directory_template_invalid_html_body` | `html_body` |
| `ibm_appid_cloud_directory_template_invalid_language` | `language` |
| `ibm_appid_cloud_directory_template_invalid_plain_text_body` | `plain_text_body` |
| `ibm_appid_cloud_directory_template_invalid_subject` | `subject` |
| `ibm_appid_cloud_directory_template_invalid_template_name` | `template_name` |
| `ibm_appid_cloud_directory_template_invalid_tenant_id` | `tenant_id` |

## Example

### Valid

```hcl
resource "ibm_appid_cloud_directory_template" "test" {
  subject = "example"
  template_name = "example"
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
