# ibm_is_lb_listener_policy_rule

Schema validation for the ibm_is_lb_listener_policy_rule resource.

This rule validates `ibm_is_lb_listener_policy_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerListenerPolicyRulePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_lb_listener_policy_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `condition` | yes | one of: `contains`, `equals`, `matches_regex`, `starts_with`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `field` | no | pattern `^[a-zA-Z0-9-_]+$` |
| `lb` | yes | — |
| `listener` | yes | — |
| `policy` | yes | one of: `automatic`, `disabled`, `manual`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `provisioning_status` | no | — |
| `related_crn` | no | — |
| `rule` | no | — |
| `type` | yes | one of: `header`, `hostname`, `path`, `body`, `query`, `sni_hostname`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `value` | yes | pattern `^[ -~]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_listener_policy_rule_invalid_condition` | `condition` |
| `ibm_is_lb_listener_policy_rule_invalid_field` | `field` |
| `ibm_is_lb_listener_policy_rule_invalid_lb` | `lb` |
| `ibm_is_lb_listener_policy_rule_invalid_listener` | `listener` |
| `ibm_is_lb_listener_policy_rule_invalid_policy` | `policy` |
| `ibm_is_lb_listener_policy_rule_invalid_type` | `type` |
| `ibm_is_lb_listener_policy_rule_invalid_value` | `value` |

## Example

### Valid

```hcl
resource "ibm_is_lb_listener_policy_rule" "test" {
  condition = "contains"
  lb = "example"
  listener = "example"
  policy = "automatic"
  type = "header"
  value = "example-name"
}
```
