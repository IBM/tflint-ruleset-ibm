# ibm_dns_secondary_data

Schema validation for the ibm_dns_secondary data source.

This rule validates `ibm_dns_secondary` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dns_secondary_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `master_ip_address` | no | — |
| `status_id` | no | — |
| `status_text` | no | — |
| `transfer_frequency` | no | — |
| `zone_name` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_secondary_data_invalid_zone_name` | `zone_name` |

## Example

### Valid

```hcl
data "ibm_dns_secondary" "test" {
  zone_name = "us-south-1"
}
```
