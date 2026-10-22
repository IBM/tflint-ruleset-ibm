# ibm_dns_resource_records_data

Schema validation for the ibm_dns_resource_records data source.

This rule validates `ibm_dns_resource_records` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `list_resource_records` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_resource_records_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `dns_resource_records` | no | — |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `zone_id` | yes | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_resource_records_data_invalid_instance_id` | `instance_id` |
| `ibm_dns_resource_records_data_invalid_zone_id` | `zone_id` |

## Example

### Valid

```hcl
data "ibm_dns_resource_records" "test" {
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  zone_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
}
```
