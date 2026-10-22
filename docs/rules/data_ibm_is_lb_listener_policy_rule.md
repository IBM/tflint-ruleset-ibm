# ibm_is_lb_listener_policy_rule_data

Schema validation for the ibm_is_lb_listener_policy_rule data source.

This rule validates `ibm_is_lb_listener_policy_rule` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerReservationAffinityPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_lb_listener_policy_rule_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `condition` | no | — |
| `created_at` | no | — |
| `field` | no | — |
| `href` | no | — |
| `lb` | yes | — |
| `listener` | yes | — |
| `policy` | yes | one of: `automatic`, `disabled`, `manual`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `provisioning_status` | no | — |
| `rule` | yes | — |
| `type` | no | — |
| `value` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_listener_policy_rule_data_invalid_lb` | `lb` |
| `ibm_is_lb_listener_policy_rule_data_invalid_listener` | `listener` |
| `ibm_is_lb_listener_policy_rule_data_invalid_policy` | `policy` |
| `ibm_is_lb_listener_policy_rule_data_invalid_rule` | `rule` |

## Example

### Valid

```hcl
data "ibm_is_lb_listener_policy_rule" "test" {
  lb = "example"
  listener = "example"
  policy = "automatic"
  rule = "example"
}
```
