# ibm_dns_custom_resolver_forwarding_rules_data

Schema validation for the ibm_dns_custom_resolver_forwarding_rules data source.

This rule validates `ibm_dns_custom_resolver_forwarding_rules` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `custom_resolver` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_custom_resolver_forwarding_rules_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `resolver_id` | yes | — |
| `rules` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_custom_resolver_forwarding_rules_data_invalid_instance_id` | `instance_id` |
| `ibm_dns_custom_resolver_forwarding_rules_data_invalid_resolver_id` | `resolver_id` |

## Example

### Valid

```hcl
data "ibm_dns_custom_resolver_forwarding_rules" "test" {
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  resolver_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
