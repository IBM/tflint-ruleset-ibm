# ibm_dns_glbs_data

Schema validation for the ibm_dns_glbs data source.

This rule validates `ibm_dns_glbs` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `access_request` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_glbs_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `dns_glbs` | no | — |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `zone_id` | yes | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_glbs_data_invalid_instance_id` | `instance_id` |
| `ibm_dns_glbs_data_invalid_zone_id` | `zone_id` |

## Example

### Valid

```hcl
data "ibm_dns_glbs" "test" {
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  zone_id = "example-name"
}
```
