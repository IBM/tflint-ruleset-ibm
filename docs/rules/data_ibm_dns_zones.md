# ibm_dns_zones_data

Schema validation for the ibm_dns_zones data source.

This rule validates `ibm_dns_zones` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `leaf-unique` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_zones_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `dns_zones` | no | — |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_zones_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_dns_zones" "test" {
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
}
```
