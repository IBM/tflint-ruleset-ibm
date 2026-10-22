# ibm_appid_apm

Schema validation for the ibm_appid_apm resource.

This rule validates `ibm_appid_apm` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `refreshTokenConfigParams` (`app-id-management`).

## Configuration

```hcl
rule "ibm_appid_apm" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `enabled` | yes | — |
| `prevent_password_with_username` | no | — |
| `tenant_id` | yes | — |
| `lockout_policy` (block) | yes | nested block |
| `lockout_policy.enabled` | yes | — |
| `lockout_policy.lockout_time_sec` | no | — |
| `lockout_policy.num_of_attempts` | no | — |
| `min_password_change_interval` (block) | yes | nested block |
| `min_password_change_interval.enabled` | yes | — |
| `min_password_change_interval.min_hours_to_change_password` | no | — |
| `password_expiration` (block) | yes | nested block |
| `password_expiration.days_to_expire` | no | — |
| `password_expiration.enabled` | yes | — |
| `password_reuse` (block) | yes | nested block |
| `password_reuse.enabled` | yes | — |
| `password_reuse.max_password_reuse` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_appid_apm_invalid_enabled` | `enabled` |
| `ibm_appid_apm_invalid_prevent_password_with_username` | `prevent_password_with_username` |
| `ibm_appid_apm_invalid_tenant_id` | `tenant_id` |
| `ibm_appid_apm_invalid_lockout_policy__enabled` | `lockout_policy.enabled` |
| `ibm_appid_apm_invalid_lockout_policy__lockout_time_sec` | `lockout_policy.lockout_time_sec` |
| `ibm_appid_apm_invalid_lockout_policy__num_of_attempts` | `lockout_policy.num_of_attempts` |
| `ibm_appid_apm_invalid_min_password_change_interval__enabled` | `min_password_change_interval.enabled` |
| `ibm_appid_apm_invalid_min_password_change_interval__min_hours_to_change_password` | `min_password_change_interval.min_hours_to_change_password` |
| `ibm_appid_apm_invalid_password_expiration__days_to_expire` | `password_expiration.days_to_expire` |
| `ibm_appid_apm_invalid_password_expiration__enabled` | `password_expiration.enabled` |
| `ibm_appid_apm_invalid_password_reuse__enabled` | `password_reuse.enabled` |
| `ibm_appid_apm_invalid_password_reuse__max_password_reuse` | `password_reuse.max_password_reuse` |

## Example

### Valid

```hcl
resource "ibm_appid_apm" "test" {
  enabled = true
  tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
  lockout_policy {
    enabled = true
  }
  min_password_change_interval {
    enabled = true
  }
  password_expiration {
    enabled = true
  }
  password_reuse {
    enabled = true
  }
}
```
