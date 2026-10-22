# ibm_cis_rate_limit

Schema validation for the ibm_cis_rate_limit resource.

This rule validates `ibm_cis_rate_limit` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/{crn}/zones/{zone_identifier}/rate_limits/{rate_limit_identifier}:put` (`cis`).

## Configuration

```hcl
rule "ibm_cis_rate_limit" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `description` | no | API: A note that you can use to describe the reason for a rate limit. |
| `disabled` | no | API: Whether this ratelimit is currently disabled. |
| `domain_id` | yes | — |
| `period` | yes | range 1–86400 |
| `rule_id` | no | — |
| `threshold` | yes | range 1–1000000 |
| `action` (block) | yes | nested block |
| `action.mode` | yes | one of: `simulate`, `ban`, `challenge`, `js_challenge` |
| `action.timeout` | no | range 10–86400 |
| `action.response` (block) | no | nested block |
| `action.response.body` | yes | API: The body to return, the content here should conform to the content_type. |
| `action.response.content_type` | yes | one of: `text/plain`, `text/xml`, `application/json` |
| `bypass` (block) | no | nested block |
| `bypass.name` | no | one of: `url` |
| `bypass.value` | no | API: The URL to bypass. |
| `correlate` (block) | no | nested block |
| `correlate.by` | no | one of: `nat` |
| `match` (block) | no | nested block |
| `match.request` (block) | no | nested block |
| `match.request.methods` | no | one of: `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `HEAD`, `_ALL_` |
| `match.request.schemes` | no | one of: `HTTP`, `HTTPS`, `_ALL_` |
| `match.request.url` | no | API: The URL pattern to match comprised of the host and path, i.e. example.org/path. Wildcard are expanded to match applicable traffic, query strings are not matched. Use * for all traffic to your zone. |
| `match.response` (block) | no | nested block |
| `match.response.origin_traffic` | no | API: Deprecated, please use response headers instead and also provide "origin_traffic:false" to avoid legacy behaviour interacting with the response.headers property. |
| `match.response.status` | no | range 100–999 |
| `match.response.headers` (block) | no | nested block |
| `match.response.headers.name` | no | API: The name of the response header to match |
| `match.response.headers.op` | no | one of: `eq`, `ne` |
| `match.response.headers.value` | no | one of: `HIT` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_rate_limit_invalid_cis_id` | `cis_id` |
| `ibm_cis_rate_limit_invalid_description` | `description` |
| `ibm_cis_rate_limit_invalid_disabled` | `disabled` |
| `ibm_cis_rate_limit_invalid_domain_id` | `domain_id` |
| `ibm_cis_rate_limit_invalid_period` | `period` |
| `ibm_cis_rate_limit_invalid_threshold` | `threshold` |
| `ibm_cis_rate_limit_invalid_action__mode` | `action.mode` |
| `ibm_cis_rate_limit_invalid_action__timeout` | `action.timeout` |
| `ibm_cis_rate_limit_invalid_action__response__body` | `action.response.body` |
| `ibm_cis_rate_limit_invalid_action__response__content_type` | `action.response.content_type` |
| `ibm_cis_rate_limit_invalid_bypass__name` | `bypass.name` |
| `ibm_cis_rate_limit_invalid_bypass__value` | `bypass.value` |
| `ibm_cis_rate_limit_invalid_correlate__by` | `correlate.by` |
| `ibm_cis_rate_limit_invalid_match__request__methods` | `match.request.methods` |
| `ibm_cis_rate_limit_invalid_match__request__schemes` | `match.request.schemes` |
| `ibm_cis_rate_limit_invalid_match__request__url` | `match.request.url` |
| `ibm_cis_rate_limit_invalid_match__response__origin_traffic` | `match.response.origin_traffic` |
| `ibm_cis_rate_limit_invalid_match__response__status` | `match.response.status` |
| `ibm_cis_rate_limit_invalid_match__response__headers__name` | `match.response.headers.name` |
| `ibm_cis_rate_limit_invalid_match__response__headers__op` | `match.response.headers.op` |
| `ibm_cis_rate_limit_invalid_match__response__headers__value` | `match.response.headers.value` |

## Example

### Valid

```hcl
resource "ibm_cis_rate_limit" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  period = 1
  threshold = 1
  action {
    mode = "simulate"
  }
}
```
