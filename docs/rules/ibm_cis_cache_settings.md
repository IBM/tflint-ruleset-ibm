# ibm_cis_cache_settings

Schema validation for the ibm_cis_cache_settings resource.

This rule validates `ibm_cis_cache_settings` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cis_cache_settings" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `browser_expiration` | no | — |
| `caching_level` | no | one of: `basic`, `simplified`, `aggressive` |
| `cis_id` | yes | — |
| `development_mode` | no | one of: `on`, `off` |
| `domain_id` | yes | — |
| `purge_all` | no | — |
| `purge_by_hosts` | no | — |
| `purge_by_tags` | no | — |
| `purge_by_urls` | no | — |
| `query_string_sort` | no | one of: `on`, `off` |
| `serve_stale_content` | no | one of: `on`, `off` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_cache_settings_invalid_browser_expiration` | `browser_expiration` |
| `ibm_cis_cache_settings_invalid_caching_level` | `caching_level` |
| `ibm_cis_cache_settings_invalid_cis_id` | `cis_id` |
| `ibm_cis_cache_settings_invalid_development_mode` | `development_mode` |
| `ibm_cis_cache_settings_invalid_domain_id` | `domain_id` |
| `ibm_cis_cache_settings_invalid_purge_all` | `purge_all` |
| `ibm_cis_cache_settings_invalid_purge_by_hosts` | `purge_by_hosts` |
| `ibm_cis_cache_settings_invalid_purge_by_tags` | `purge_by_tags` |
| `ibm_cis_cache_settings_invalid_purge_by_urls` | `purge_by_urls` |
| `ibm_cis_cache_settings_invalid_query_string_sort` | `query_string_sort` |
| `ibm_cis_cache_settings_invalid_serve_stale_content` | `serve_stale_content` |

## Example

### Valid

```hcl
resource "ibm_cis_cache_settings" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
