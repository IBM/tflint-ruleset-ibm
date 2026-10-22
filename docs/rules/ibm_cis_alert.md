# ibm_cis_alert

Manages ibm_cis_alert. Policy name

This rule validates `ibm_cis_alert` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `create_alert_policy_input` (`cis`).

## Configuration

```hcl
rule "ibm_cis_alert" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alert_type` | yes | one of: `dos_attack_l7`, `g6_pool_toggle_alert`, `clickhouse_alert_fw_anomaly`, `clickhouse_alert_fw_ent_anomaly`, `universal_ssl_event_type`, `dedicated_ssl_certificate_event_type`, `health_check_status_notification`, `web_analytics_metrics_update`, `maintenance_event_notification` |
| `cis_id` | yes | — |
| `conditions` | no | API: Conditions depending on the alert type. HTTP DDOS Attack Alerter does not have any conditions. The Load Balancing Pool Enablement Alerter takes conditions that describe for all pools whether the pool is being enabled, disabled, or both. This field is not required when creating a new alert. |
| `description` | no | API: Policy description |
| `enabled` | yes | API: Is the alert policy active |
| `filters` | no | API: Optional filters depending for the alert type |
| `name` | yes | API: Policy name |
| `policy_id` | no | — |
| `mechanisms` (block) | yes | nested block |
| `mechanisms.email` | no | — |
| `mechanisms.webhooks` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_alert_invalid_alert_type` | `alert_type` |
| `ibm_cis_alert_invalid_cis_id` | `cis_id` |
| `ibm_cis_alert_invalid_conditions` | `conditions` |
| `ibm_cis_alert_invalid_description` | `description` |
| `ibm_cis_alert_invalid_enabled` | `enabled` |
| `ibm_cis_alert_invalid_filters` | `filters` |
| `ibm_cis_alert_invalid_name` | `name` |
| `ibm_cis_alert_invalid_mechanisms__email` | `mechanisms.email` |
| `ibm_cis_alert_invalid_mechanisms__webhooks` | `mechanisms.webhooks` |

## Example

### Valid

```hcl
resource "ibm_cis_alert" "test" {
  alert_type = "dos_attack_l7"
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  enabled = true
  name = "example"
  mechanisms {
  }
}
```
