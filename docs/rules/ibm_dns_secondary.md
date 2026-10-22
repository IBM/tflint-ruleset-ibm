# ibm_dns_secondary

Schema validation for the ibm_dns_secondary resource.

This rule validates `ibm_dns_secondary` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dns_secondary" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `master_ip_address` | yes | — |
| `status_id` | no | — |
| `status_text` | no | — |
| `tags` | no | — |
| `transfer_frequency` | yes | — |
| `zone_name` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_secondary_invalid_master_ip_address` | `master_ip_address` |
| `ibm_dns_secondary_invalid_tags` | `tags` |
| `ibm_dns_secondary_invalid_transfer_frequency` | `transfer_frequency` |
| `ibm_dns_secondary_invalid_zone_name` | `zone_name` |

## Example

### Valid

```hcl
resource "ibm_dns_secondary" "test" {
  master_ip_address = "example"
  transfer_frequency = 1
  zone_name = "us-south-1"
}
```
