# ibm_cis_domain_settings

Schema validation for the ibm_cis_domain_settings resource.

This rule validates `ibm_cis_domain_settings` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/{crn}/zones/{zone_identifier}/ssl/universal/settings:patch` (`cis`).

## Configuration

```hcl
rule "ibm_cis_domain_settings" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `always_use_https` | no | one of: `on`, `off` |
| `automatic_https_rewrites` | no | one of: `on`, `off` |
| `brotli` | no | one of: `on`, `off` |
| `browser_check` | no | one of: `on`, `off` |
| `certificate_status` | no | — |
| `challenge_ttl` | no | — |
| `cipher` | no | one of: `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-CHACHA20-POLY1305`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-CHACHA20-POLY1305`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES128-SHA`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES128-SHA`, `AES128-GCM-SHA256`, `AES128-SHA256`, `AES128-SHA`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES256-SHA384`, `ECDHE-RSA-AES256-SHA`, `AES256-GCM-SHA384`, `AES256-SHA256`, `AES256-SHA`, `DES-CBC3-SHA`, `AEAD-AES128-GCM-SHA256`, `AEAD-AES256-GCM-SHA384`, `AEAD-CHACHA20-POLY1305-SHA256` |
| `cis_id` | yes | — |
| `cname_flattening` | no | one of: `flatten_at_root`, `flatten_all`, `flatten_none` |
| `dnssec` | no | one of: `active`, `disabled` |
| `domain_id` | yes | — |
| `email_obfuscation` | no | one of: `on`, `off` |
| `hotlink_protection` | no | one of: `on`, `off` |
| `http2` | no | one of: `on`, `off` |
| `http3` | no | one of: `on`, `off` |
| `image_load_optimization` | no | one of: `on`, `off` |
| `image_size_optimization` | no | one of: `lossless`, `off`, `lossy` |
| `ip_geolocation` | no | one of: `on`, `off` |
| `ipv6` | no | one of: `on`, `off` |
| `log_retention` | no | — |
| `max_upload` | no | — |
| `min_tls_version` | no | one of: `1.1`, `1.2`, `1.3`, `1.4` |
| `opportunistic_encryption` | no | one of: `on`, `off` |
| `opportunistic_onion` | no | one of: `on`, `off` |
| `origin_error_page_pass_thru` | no | one of: `on`, `off` |
| `origin_max_http_version` | no | one of: `1`, `2` |
| `origin_post_quantum_encryption` | no | one of: `off`, `preferred`, `supported` |
| `prefetch_preload` | no | one of: `on`, `off` |
| `proxy_read_timeout` | no | range 1–6000 |
| `pseudo_ipv4` | no | one of: `overwrite_header`, `off`, `add_header` |
| `replace_insecure_js` | no | one of: `on`, `off` |
| `response_buffering` | no | one of: `on`, `off` |
| `script_load_optimization` | no | one of: `on`, `off` |
| `security_level` | no | one of: `essentially_off`, `low`, `medium`, `high`, `under_attack` |
| `server_side_exclude` | no | one of: `on`, `off` |
| `ssl` | no | one of: `off`, `flexible`, `full`, `strict`, `origin_pull` |
| `tls_client_auth` | no | one of: `on`, `off` |
| `true_client_ip_header` | no | one of: `on`, `off` |
| `waf` | no | one of: `on`, `off` |
| `websockets` | no | one of: `on`, `off` |
| `minify` (block) | no | nested block |
| `minify.css` | yes | one of: `on`, `off` |
| `minify.html` | yes | one of: `on`, `off` |
| `minify.js` | yes | one of: `on`, `off` |
| `mobile_redirect` (block) | no | nested block |
| `mobile_redirect.mobile_subdomain` | no | — |
| `mobile_redirect.status` | yes | — |
| `mobile_redirect.strip_uri` | no | — |
| `security_header` (block) | no | nested block |
| `security_header.enabled` | yes | API: Enabled |
| `security_header.include_subdomains` | yes | — |
| `security_header.max_age` | yes | — |
| `security_header.nosniff` | yes | — |
| `security_header.preload` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_domain_settings_invalid_always_use_https` | `always_use_https` |
| `ibm_cis_domain_settings_invalid_automatic_https_rewrites` | `automatic_https_rewrites` |
| `ibm_cis_domain_settings_invalid_brotli` | `brotli` |
| `ibm_cis_domain_settings_invalid_browser_check` | `browser_check` |
| `ibm_cis_domain_settings_invalid_challenge_ttl` | `challenge_ttl` |
| `ibm_cis_domain_settings_invalid_cipher` | `cipher` |
| `ibm_cis_domain_settings_invalid_cis_id` | `cis_id` |
| `ibm_cis_domain_settings_invalid_cname_flattening` | `cname_flattening` |
| `ibm_cis_domain_settings_invalid_dnssec` | `dnssec` |
| `ibm_cis_domain_settings_invalid_domain_id` | `domain_id` |
| `ibm_cis_domain_settings_invalid_email_obfuscation` | `email_obfuscation` |
| `ibm_cis_domain_settings_invalid_hotlink_protection` | `hotlink_protection` |
| `ibm_cis_domain_settings_invalid_http2` | `http2` |
| `ibm_cis_domain_settings_invalid_http3` | `http3` |
| `ibm_cis_domain_settings_invalid_image_load_optimization` | `image_load_optimization` |
| `ibm_cis_domain_settings_invalid_image_size_optimization` | `image_size_optimization` |
| `ibm_cis_domain_settings_invalid_ip_geolocation` | `ip_geolocation` |
| `ibm_cis_domain_settings_invalid_ipv6` | `ipv6` |
| `ibm_cis_domain_settings_invalid_log_retention` | `log_retention` |
| `ibm_cis_domain_settings_invalid_max_upload` | `max_upload` |
| `ibm_cis_domain_settings_invalid_min_tls_version` | `min_tls_version` |
| `ibm_cis_domain_settings_invalid_opportunistic_encryption` | `opportunistic_encryption` |
| `ibm_cis_domain_settings_invalid_opportunistic_onion` | `opportunistic_onion` |
| `ibm_cis_domain_settings_invalid_origin_error_page_pass_thru` | `origin_error_page_pass_thru` |
| `ibm_cis_domain_settings_invalid_origin_max_http_version` | `origin_max_http_version` |
| `ibm_cis_domain_settings_invalid_origin_post_quantum_encryption` | `origin_post_quantum_encryption` |
| `ibm_cis_domain_settings_invalid_prefetch_preload` | `prefetch_preload` |
| `ibm_cis_domain_settings_invalid_proxy_read_timeout` | `proxy_read_timeout` |
| `ibm_cis_domain_settings_invalid_pseudo_ipv4` | `pseudo_ipv4` |
| `ibm_cis_domain_settings_invalid_replace_insecure_js` | `replace_insecure_js` |
| `ibm_cis_domain_settings_invalid_response_buffering` | `response_buffering` |
| `ibm_cis_domain_settings_invalid_script_load_optimization` | `script_load_optimization` |
| `ibm_cis_domain_settings_invalid_security_level` | `security_level` |
| `ibm_cis_domain_settings_invalid_server_side_exclude` | `server_side_exclude` |
| `ibm_cis_domain_settings_invalid_ssl` | `ssl` |
| `ibm_cis_domain_settings_invalid_tls_client_auth` | `tls_client_auth` |
| `ibm_cis_domain_settings_invalid_true_client_ip_header` | `true_client_ip_header` |
| `ibm_cis_domain_settings_invalid_waf` | `waf` |
| `ibm_cis_domain_settings_invalid_websockets` | `websockets` |
| `ibm_cis_domain_settings_invalid_minify__css` | `minify.css` |
| `ibm_cis_domain_settings_invalid_minify__html` | `minify.html` |
| `ibm_cis_domain_settings_invalid_minify__js` | `minify.js` |
| `ibm_cis_domain_settings_invalid_mobile_redirect__mobile_subdomain` | `mobile_redirect.mobile_subdomain` |
| `ibm_cis_domain_settings_invalid_mobile_redirect__status` | `mobile_redirect.status` |
| `ibm_cis_domain_settings_invalid_mobile_redirect__strip_uri` | `mobile_redirect.strip_uri` |
| `ibm_cis_domain_settings_invalid_security_header__enabled` | `security_header.enabled` |
| `ibm_cis_domain_settings_invalid_security_header__include_subdomains` | `security_header.include_subdomains` |
| `ibm_cis_domain_settings_invalid_security_header__max_age` | `security_header.max_age` |
| `ibm_cis_domain_settings_invalid_security_header__nosniff` | `security_header.nosniff` |
| `ibm_cis_domain_settings_invalid_security_header__preload` | `security_header.preload` |

## Example

### Valid

```hcl
resource "ibm_cis_domain_settings" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
