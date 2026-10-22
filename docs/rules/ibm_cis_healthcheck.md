# ibm_cis_healthcheck

Schema validation for the ibm_cis_healthcheck resource.

This rule validates `ibm_cis_healthcheck` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/{crn}/load_balancers/monitors:post` (`cis`).

## Configuration

```hcl
rule "ibm_cis_healthcheck" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allow_insecure` | no | API: Allow insecure |
| `cis_id` | yes | — |
| `create_on` | no | — |
| `description` | no | API: Login page monitor |
| `expected_body` | no | API: Expected body |
| `expected_codes` | no | API: Expected codes |
| `follow_redirects` | no | API: Follow redirects |
| `interval` | no | range 5–3600 |
| `method` | no | one of: `GET`, `HEAD` |
| `modified_on` | no | — |
| `monitor_id` | no | — |
| `path` | no | API: Path |
| `port` | no | range 1–65535 |
| `retries` | no | range 1–3 |
| `timeout` | no | range 1–10 |
| `type` | no | one of: `http`, `https`, `tcp` |
| `headers` (block) | no | nested block |
| `headers.header` | yes | API: Header |
| `headers.values` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_healthcheck_invalid_allow_insecure` | `allow_insecure` |
| `ibm_cis_healthcheck_invalid_cis_id` | `cis_id` |
| `ibm_cis_healthcheck_invalid_description` | `description` |
| `ibm_cis_healthcheck_invalid_expected_body` | `expected_body` |
| `ibm_cis_healthcheck_invalid_expected_codes` | `expected_codes` |
| `ibm_cis_healthcheck_invalid_follow_redirects` | `follow_redirects` |
| `ibm_cis_healthcheck_invalid_interval` | `interval` |
| `ibm_cis_healthcheck_invalid_method` | `method` |
| `ibm_cis_healthcheck_invalid_path` | `path` |
| `ibm_cis_healthcheck_invalid_port` | `port` |
| `ibm_cis_healthcheck_invalid_retries` | `retries` |
| `ibm_cis_healthcheck_invalid_timeout` | `timeout` |
| `ibm_cis_healthcheck_invalid_type` | `type` |
| `ibm_cis_healthcheck_invalid_headers__header` | `headers.header` |
| `ibm_cis_healthcheck_invalid_headers__values` | `headers.values` |

## Example

### Valid

```hcl
resource "ibm_cis_healthcheck" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
