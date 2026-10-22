# ibm_cis_ruleset_entrypoint_versions_data

Schema validation for the ibm_cis_ruleset_entrypoint_versions data source.

This rule validates `ibm_cis_ruleset_entrypoint_versions` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ruleset_create` (`cis`).

## Configuration

```hcl
rule "ibm_cis_ruleset_entrypoint_versions_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | no | — |
| `list_all` | no | — |
| `phase` | yes | one of: `ddos_l4`, `ddos_l7`, `http_config_settings`, `http_custom_errors`, `http_log_custom_fields`, `http_ratelimit`, `http_request_cache_settings`, `http_request_dynamic_redirect`, `http_request_firewall_custom`, `http_request_firewall_managed`, `http_request_late_transform`, `http_request_origin`, `http_request_redirect`, `http_request_sanitize`, `http_request_sbfm`, `http_request_select_configuration`, `http_request_transform`, `http_response_compression`, `http_response_firewall_managed`, `http_response_headers_transform` |
| `ruleset_id` | no | — |
| `rulesets` | no | — |
| `version` | no | API: The version of the ruleset. Use "latest" to get the latest version. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_ruleset_entrypoint_versions_data_invalid_cis_id` | `cis_id` |
| `ibm_cis_ruleset_entrypoint_versions_data_invalid_domain_id` | `domain_id` |
| `ibm_cis_ruleset_entrypoint_versions_data_invalid_list_all` | `list_all` |
| `ibm_cis_ruleset_entrypoint_versions_data_invalid_phase` | `phase` |
| `ibm_cis_ruleset_entrypoint_versions_data_invalid_ruleset_id` | `ruleset_id` |
| `ibm_cis_ruleset_entrypoint_versions_data_invalid_version` | `version` |

## Example

### Valid

```hcl
data "ibm_cis_ruleset_entrypoint_versions" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  phase = "ddos_l4"
}
```
