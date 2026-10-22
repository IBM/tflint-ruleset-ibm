# ibm_compute_autoscale_policy

Manages ibm_compute_autoscale_policy. Name

This rule validates `ibm_compute_autoscale_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_autoscale_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cooldown` | no | — |
| `name` | yes | — |
| `scale_amount` | yes | — |
| `scale_group_id` | yes | — |
| `scale_type` | yes | — |
| `tags` | no | — |
| `triggers` (block) | no | nested block |
| `triggers.date` | no | — |
| `triggers.schedule` | no | — |
| `triggers.type` | yes | — |
| `triggers.watches` (block) | no | nested block |
| `triggers.watches.metric` | yes | — |
| `triggers.watches.operator` | yes | — |
| `triggers.watches.period` | yes | — |
| `triggers.watches.value` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_autoscale_policy_invalid_cooldown` | `cooldown` |
| `ibm_compute_autoscale_policy_invalid_name` | `name` |
| `ibm_compute_autoscale_policy_invalid_scale_amount` | `scale_amount` |
| `ibm_compute_autoscale_policy_invalid_scale_group_id` | `scale_group_id` |
| `ibm_compute_autoscale_policy_invalid_scale_type` | `scale_type` |
| `ibm_compute_autoscale_policy_invalid_tags` | `tags` |
| `ibm_compute_autoscale_policy_invalid_triggers__date` | `triggers.date` |
| `ibm_compute_autoscale_policy_invalid_triggers__schedule` | `triggers.schedule` |
| `ibm_compute_autoscale_policy_invalid_triggers__type` | `triggers.type` |
| `ibm_compute_autoscale_policy_invalid_triggers__watches__metric` | `triggers.watches.metric` |
| `ibm_compute_autoscale_policy_invalid_triggers__watches__operator` | `triggers.watches.operator` |
| `ibm_compute_autoscale_policy_invalid_triggers__watches__period` | `triggers.watches.period` |
| `ibm_compute_autoscale_policy_invalid_triggers__watches__value` | `triggers.watches.value` |

## Example

### Valid

```hcl
resource "ibm_compute_autoscale_policy" "test" {
  name = "example"
  scale_amount = 1
  scale_group_id = 1
  scale_type = "example"
}
```
