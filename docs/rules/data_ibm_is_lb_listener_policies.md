# ibm_is_lb_listener_policies_data

Schema validation for the ibm_is_lb_listener_policies data source.

This rule validates `ibm_is_lb_listener_policies` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerListenerHTTPSRedirectPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_lb_listener_policies_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `lb` | yes | — |
| `listener` | yes | API: Identifies a load balancer listener by a unique property. |
| `policies` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_lb_listener_policies_data_invalid_lb` | `lb` |
| `ibm_is_lb_listener_policies_data_invalid_listener` | `listener` |

## Example

### Valid

```hcl
data "ibm_is_lb_listener_policies" "test" {
  lb = "example"
  listener = "example"
}
```
