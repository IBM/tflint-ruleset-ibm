# ibm_cis_mtls_app

Manages ibm_cis_mtls_app. App Name

This rule validates `ibm_cis_mtls_app` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `access_app_input` (`cis`).

## Configuration

```hcl
rule "ibm_cis_mtls_app" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `app_created_at` | no | — |
| `app_id` | no | — |
| `app_updated_at` | no | — |
| `cert_rule_val` | no | — |
| `cis_id` | yes | — |
| `common_rule_val` | no | — |
| `domain` | yes | API: The domain and path that Access blocks |
| `domain_id` | yes | — |
| `name` | yes | API: Application name |
| `pol_created_at` | no | — |
| `pol_updated_at` | no | — |
| `policy_decision` | no | — |
| `policy_id` | no | — |
| `policy_name` | no | — |
| `session_duration` | no | one of: `30m`, `6h`, `12h`, `24h`, `168h`, `730h` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_mtls_app_invalid_cert_rule_val` | `cert_rule_val` |
| `ibm_cis_mtls_app_invalid_cis_id` | `cis_id` |
| `ibm_cis_mtls_app_invalid_common_rule_val` | `common_rule_val` |
| `ibm_cis_mtls_app_invalid_domain` | `domain` |
| `ibm_cis_mtls_app_invalid_domain_id` | `domain_id` |
| `ibm_cis_mtls_app_invalid_name` | `name` |
| `ibm_cis_mtls_app_invalid_policy_decision` | `policy_decision` |
| `ibm_cis_mtls_app_invalid_policy_name` | `policy_name` |
| `ibm_cis_mtls_app_invalid_session_duration` | `session_duration` |

## Example

### Valid

```hcl
resource "ibm_cis_mtls_app" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain = "example"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  name = "example"
}
```
