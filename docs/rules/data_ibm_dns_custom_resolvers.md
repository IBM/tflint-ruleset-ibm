# ibm_dns_custom_resolvers_data

Schema validation for the ibm_dns_custom_resolvers data source.

This rule validates `ibm_dns_custom_resolvers` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/instances/{instance_id}/custom_resolvers/{resolver_id}:patch` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_custom_resolvers_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `custom_resolvers` | no | — |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_custom_resolvers_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_dns_custom_resolvers" "test" {
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
}
```
