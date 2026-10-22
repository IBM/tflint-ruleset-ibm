# ibm_cis_custom_list_items

Schema validation for the ibm_cis_custom_list_items resource.

This rule validates `ibm_cis_custom_list_items` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `list_item` (`cis`).

## Configuration

```hcl
rule "ibm_cis_custom_list_items" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `list_id` | yes | — |
| `items` (block) | yes | nested block |
| `items.asn` | no | API: An autonomous system number. |
| `items.comment` | no | API: An informative summary of the list item. |
| `items.hostname` | no | API: Valid characters for hostnames are ASCII(7) letters from a to z, the digits from 0 to 9, wildcards (*), and the hyphen (-). |
| `items.ip` | no | API: An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a maximum of /64. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_custom_list_items_invalid_cis_id` | `cis_id` |
| `ibm_cis_custom_list_items_invalid_list_id` | `list_id` |
| `ibm_cis_custom_list_items_invalid_items__asn` | `items.asn` |
| `ibm_cis_custom_list_items_invalid_items__comment` | `items.comment` |
| `ibm_cis_custom_list_items_invalid_items__hostname` | `items.hostname` |
| `ibm_cis_custom_list_items_invalid_items__ip` | `items.ip` |

## Example

### Valid

```hcl
resource "ibm_cis_custom_list_items" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  list_id = "01234567-89ab-cdef-0123-456789abcdef"
  items {
  }
}
```
