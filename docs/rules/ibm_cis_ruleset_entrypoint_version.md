# ibm_cis_ruleset_entrypoint_version

Schema validation for the ibm_cis_ruleset_entrypoint_version resource.

This rule validates `ibm_cis_ruleset_entrypoint_version` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ruleset_create` (`cis`).

## Configuration

```hcl
rule "ibm_cis_ruleset_entrypoint_version" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | no | — |
| `phase` | yes | one of: `ddos_l4`, `ddos_l7`, `http_config_settings`, `http_custom_errors`, `http_log_custom_fields`, `http_ratelimit`, `http_request_cache_settings`, `http_request_dynamic_redirect`, `http_request_firewall_custom`, `http_request_firewall_managed`, `http_request_late_transform`, `http_request_origin`, `http_request_redirect`, `http_request_sanitize`, `http_request_sbfm`, `http_request_select_configuration`, `http_request_transform`, `http_response_compression`, `http_response_firewall_managed`, `http_response_headers_transform` |
| `rulesets` | no | API: List of ruleset ids to apply action to. Use "current" to apply to the current ruleset. |
| `rulesets_response` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_ruleset_entrypoint_version_invalid_cis_id` | `cis_id` |
| `ibm_cis_ruleset_entrypoint_version_invalid_domain_id` | `domain_id` |
| `ibm_cis_ruleset_entrypoint_version_invalid_phase` | `phase` |
| `ibm_cis_ruleset_entrypoint_version_invalid_rulesets` | `rulesets` |

## Example

### Valid

```hcl
resource "ibm_cis_ruleset_entrypoint_version" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  phase = "ddos_l4"
}
```
