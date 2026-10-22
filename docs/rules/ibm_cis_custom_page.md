# ibm_cis_custom_page

Schema validation for the ibm_cis_custom_page resource.

This rule validates `ibm_cis_custom_page` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `custom_page_request` (`cis`).

## Configuration

```hcl
rule "ibm_cis_custom_page" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `created_on` | no | — |
| `description` | no | — |
| `domain_id` | yes | — |
| `modified_on` | no | — |
| `page_id` | yes | one of: `basic_challenge`, `waf_challenge`, `waf_block`, `ratelimit_block`, `country_challenge`, `ip_block`, `under_attack`, `500_errors`, `1000_errors`, `always_online` |
| `preview_target` | no | — |
| `required_tokens` | no | — |
| `state` | no | — |
| `url` | yes | API: A URL that is associated with the Custom Page |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_custom_page_invalid_cis_id` | `cis_id` |
| `ibm_cis_custom_page_invalid_domain_id` | `domain_id` |
| `ibm_cis_custom_page_invalid_page_id` | `page_id` |
| `ibm_cis_custom_page_invalid_url` | `url` |

## Example

### Valid

```hcl
resource "ibm_cis_custom_page" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  page_id = "basic_challenge"
  url = "example"
}
```
