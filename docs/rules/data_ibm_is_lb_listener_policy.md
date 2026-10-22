# ibm_is_lb_listener_policy_data

Reads ibm_is_lb_listener_policy. The user-defined name for this policy.

This rule validates `ibm_is_lb_listener_policy` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerListenerHTTPSRedirectPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_lb_listener_policy_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | no | — |
| `created_at` | no | — |
| `href` | no | — |
| `lb` | yes | — |
| `listener` | yes | API: Identifies a load balancer listener by a unique property. |
| `name` | no | — |
| `policy_id` | yes | — |
| `priority` | no | — |
| `provisioning_status` | no | — |
| `rules` | no | — |
| `target` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_listener_policy_data_invalid_lb` | `lb` |
| `ibm_is_lb_listener_policy_data_invalid_listener` | `listener` |
| `ibm_is_lb_listener_policy_data_invalid_policy_id` | `policy_id` |

## Example

### Valid

```hcl
data "ibm_is_lb_listener_policy" "test" {
  lb = "example"
  listener = "example"
  policy_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
