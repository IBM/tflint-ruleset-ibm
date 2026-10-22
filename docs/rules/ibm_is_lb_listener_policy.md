# ibm_is_lb_listener_policy

Manages ibm_is_lb_listener_policy. Policy name

This rule validates `ibm_is_lb_listener_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerListenerPolicyTargetPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_lb_listener_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | yes | one of: `forward`, `forward_to_pool`, `forward_to_listener`, `redirect`, `reject`, `https_redirect` |
| `lb` | yes | — |
| `listener` | yes | API: Identifies a load balancer listener by a unique property. |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `policy_id` | no | — |
| `priority` | yes | — |
| `provisioning_status` | no | — |
| `related_crn` | no | — |
| `target_http_status_code` | no | — |
| `target_https_redirect_listener` | no | — |
| `target_https_redirect_status_code` | no | — |
| `target_https_redirect_uri` | no | — |
| `target_id` | no | — |
| `target_url` | no | — |
| `rules` (block) | no | nested block |
| `rules.condition` | yes | one of: `contains`, `equals`, `matches_regex`, `starts_with`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `rules.field` | no | — |
| `rules.type` | yes | — |
| `rules.value` | yes | — |
| `target` (block) | no | nested block |
| `target.http_status_code` | no | one of: `301`, `302`, `303`, `307`, `308` |
| `target.id` | no | pattern `^[-0-9a-z_]+$` |
| `target.uri` | no | pattern `^\/(([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})+(\/([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})*)*)?(\?([a-zA-Z0-9-._~!$&'()*+,;=:@\/?]|%[a-fA-F0-9]{2})*)?$` |
| `target.url` | no | pattern `^(http(s)?|\{protocol\}):\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$` |
| `target.listener` (block) | no | nested block |
| `target.listener.id` | yes | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_listener_policy_invalid_action` | `action` |
| `ibm_is_lb_listener_policy_invalid_lb` | `lb` |
| `ibm_is_lb_listener_policy_invalid_listener` | `listener` |
| `ibm_is_lb_listener_policy_invalid_name` | `name` |
| `ibm_is_lb_listener_policy_invalid_priority` | `priority` |
| `ibm_is_lb_listener_policy_invalid_target_http_status_code` | `target_http_status_code` |
| `ibm_is_lb_listener_policy_invalid_target_https_redirect_listener` | `target_https_redirect_listener` |
| `ibm_is_lb_listener_policy_invalid_target_https_redirect_status_code` | `target_https_redirect_status_code` |
| `ibm_is_lb_listener_policy_invalid_target_https_redirect_uri` | `target_https_redirect_uri` |
| `ibm_is_lb_listener_policy_invalid_target_id` | `target_id` |
| `ibm_is_lb_listener_policy_invalid_target_url` | `target_url` |
| `ibm_is_lb_listener_policy_invalid_rules__condition` | `rules.condition` |
| `ibm_is_lb_listener_policy_invalid_rules__field` | `rules.field` |
| `ibm_is_lb_listener_policy_invalid_rules__type` | `rules.type` |
| `ibm_is_lb_listener_policy_invalid_rules__value` | `rules.value` |
| `ibm_is_lb_listener_policy_invalid_target__http_status_code` | `target.http_status_code` |
| `ibm_is_lb_listener_policy_invalid_target__id` | `target.id` |
| `ibm_is_lb_listener_policy_invalid_target__uri` | `target.uri` |
| `ibm_is_lb_listener_policy_invalid_target__url` | `target.url` |
| `ibm_is_lb_listener_policy_invalid_target__listener__id` | `target.listener.id` |

## Example

### Valid

```hcl
resource "ibm_is_lb_listener_policy" "test" {
  action = "forward"
  lb = "example"
  listener = "example"
  priority = 1
}
```
