# ibm_dns_custom_resolver_secondary_zone

Schema validation for the ibm_dns_custom_resolver_secondary_zone resource.

This rule validates `ibm_dns_custom_resolver_secondary_zone` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `secondary_zone_input` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_custom_resolver_secondary_zone" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_on` | no | — |
| `description` | no | API: Descriptive text of the secondary zone |
| `enabled` | yes | API: Enable/Disable the secondary zone |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `modified_on` | no | — |
| `resolver_id` | yes | — |
| `secondary_zone_id` | no | — |
| `transfer_from` | yes | API: The address of DNS server where the secondary zone data should be transferred from |
| `zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_custom_resolver_secondary_zone_invalid_description` | `description` |
| `ibm_dns_custom_resolver_secondary_zone_invalid_enabled` | `enabled` |
| `ibm_dns_custom_resolver_secondary_zone_invalid_instance_id` | `instance_id` |
| `ibm_dns_custom_resolver_secondary_zone_invalid_resolver_id` | `resolver_id` |
| `ibm_dns_custom_resolver_secondary_zone_invalid_transfer_from` | `transfer_from` |
| `ibm_dns_custom_resolver_secondary_zone_invalid_zone` | `zone` |

## Example

### Valid

```hcl
resource "ibm_dns_custom_resolver_secondary_zone" "test" {
  enabled = true
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  resolver_id = "01234567-89ab-cdef-0123-456789abcdef"
  transfer_from = "aaaaaaa"
  zone = "us-south-1"
}
```
