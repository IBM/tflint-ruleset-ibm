# ibm_dns_custom_resolver_forwarding_rule

Schema validation for the ibm_dns_custom_resolver_forwarding_rule resource.

This rule validates `ibm_dns_custom_resolver_forwarding_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `forwarding_rule_input` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_custom_resolver_forwarding_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | API: Descriptive text of the forwarding rule. |
| `forward_to` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `match` | no | API: The matching zone or hostname. |
| `resolver_id` | yes | — |
| `rule_id` | no | — |
| `type` | no | one of: `hostname`, `zone`, `Default` |
| `views` (block) | no | nested block |
| `views.description` | no | API: Descriptive text of the forwarding rule. |
| `views.expression` | yes | API: The custom resolver will only apply the rule to the incoming DNS queries if the expression evaluates to true. The expression follows [Common Expression Language](https://github.com/google/cel-spec/blob/master/doc/langdef.md), but does not support CEL build-in functions and macros. Currently, the expression only supports the following custom functions, variables and operators: - Functions: - `ipInR |
| `views.forward_to` | yes | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `views.name` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_custom_resolver_forwarding_rule_invalid_description` | `description` |
| `ibm_dns_custom_resolver_forwarding_rule_invalid_forward_to` | `forward_to` |
| `ibm_dns_custom_resolver_forwarding_rule_invalid_instance_id` | `instance_id` |
| `ibm_dns_custom_resolver_forwarding_rule_invalid_match` | `match` |
| `ibm_dns_custom_resolver_forwarding_rule_invalid_resolver_id` | `resolver_id` |
| `ibm_dns_custom_resolver_forwarding_rule_invalid_type` | `type` |
| `ibm_dns_custom_resolver_forwarding_rule_invalid_views__description` | `views.description` |
| `ibm_dns_custom_resolver_forwarding_rule_invalid_views__expression` | `views.expression` |
| `ibm_dns_custom_resolver_forwarding_rule_invalid_views__forward_to` | `views.forward_to` |
| `ibm_dns_custom_resolver_forwarding_rule_invalid_views__name` | `views.name` |

## Example

### Valid

```hcl
resource "ibm_dns_custom_resolver_forwarding_rule" "test" {
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  resolver_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
