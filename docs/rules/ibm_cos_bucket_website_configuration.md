# ibm_cos_bucket_website_configuration

Schema validation for the ibm_cos_bucket_website_configuration resource.

This rule validates `ibm_cos_bucket_website_configuration` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cos_bucket_website_configuration" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bucket_crn` | yes | crn format |
| `bucket_location` | yes | — |
| `endpoint_type` | no | one of: `public`, `private`, `direct` |
| `website_endpoint` | no | — |
| `website_configuration` (block) | yes | nested block |
| `website_configuration.routing_rules` | no | json format |
| `website_configuration.error_document` (block) | no | nested block |
| `website_configuration.error_document.key` | yes | — |
| `website_configuration.index_document` (block) | no | nested block |
| `website_configuration.index_document.suffix` | yes | — |
| `website_configuration.redirect_all_requests_to` (block) | no | nested block |
| `website_configuration.redirect_all_requests_to.host_name` | yes | — |
| `website_configuration.redirect_all_requests_to.protocol` | no | — |
| `website_configuration.routing_rule` (block) | no | nested block |
| `website_configuration.routing_rule.condition` (block) | no | nested block |
| `website_configuration.routing_rule.condition.http_error_code_returned_equals` | no | — |
| `website_configuration.routing_rule.condition.key_prefix_equals` | no | — |
| `website_configuration.routing_rule.redirect` (block) | yes | nested block |
| `website_configuration.routing_rule.redirect.host_name` | no | — |
| `website_configuration.routing_rule.redirect.http_redirect_code` | no | — |
| `website_configuration.routing_rule.redirect.protocol` | no | — |
| `website_configuration.routing_rule.redirect.replace_key_prefix_with` | no | — |
| `website_configuration.routing_rule.redirect.replace_key_with` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cos_bucket_website_configuration_invalid_bucket_crn` | `bucket_crn` |
| `ibm_cos_bucket_website_configuration_invalid_bucket_location` | `bucket_location` |
| `ibm_cos_bucket_website_configuration_invalid_endpoint_type` | `endpoint_type` |
| `ibm_cos_bucket_website_configuration_invalid_website_configuration__routing_rules` | `website_configuration.routing_rules` |
| `ibm_cos_bucket_website_configuration_invalid_website_configuration__error_document__key` | `website_configuration.error_document.key` |
| `ibm_cos_bucket_website_configuration_invalid_website_configuration__index_document__suffix` | `website_configuration.index_document.suffix` |
| `ibm_cos_bucket_website_configuration_invalid_website_configuration__redirect_all_requests_to__host_name` | `website_configuration.redirect_all_requests_to.host_name` |
| `ibm_cos_bucket_website_configuration_invalid_website_configuration__redirect_all_requests_to__protocol` | `website_configuration.redirect_all_requests_to.protocol` |
| `ibm_cos_bucket_website_configuration_invalid_website_configuration__routing_rule__condition__http_error_code_returned_equals` | `website_configuration.routing_rule.condition.http_error_code_returned_equals` |
| `ibm_cos_bucket_website_configuration_invalid_website_configuration__routing_rule__condition__key_prefix_equals` | `website_configuration.routing_rule.condition.key_prefix_equals` |
| `ibm_cos_bucket_website_configuration_invalid_website_configuration__routing_rule__redirect__host_name` | `website_configuration.routing_rule.redirect.host_name` |
| `ibm_cos_bucket_website_configuration_invalid_website_configuration__routing_rule__redirect__http_redirect_code` | `website_configuration.routing_rule.redirect.http_redirect_code` |
| `ibm_cos_bucket_website_configuration_invalid_website_configuration__routing_rule__redirect__protocol` | `website_configuration.routing_rule.redirect.protocol` |
| `ibm_cos_bucket_website_configuration_invalid_website_configuration__routing_rule__redirect__replace_key_prefix_with` | `website_configuration.routing_rule.redirect.replace_key_prefix_with` |
| `ibm_cos_bucket_website_configuration_invalid_website_configuration__routing_rule__redirect__replace_key_with` | `website_configuration.routing_rule.redirect.replace_key_with` |

## Example

### Valid

```hcl
resource "ibm_cos_bucket_website_configuration" "test" {
  bucket_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  bucket_location = "example"
  website_configuration {
  }
}
```
