# ibm_cbr_zone_data

Reads ibm_cbr_zone. The name of the zone.

This rule validates `ibm_cbr_zone` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cbr_zone_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `address_count` | no | — |
| `addresses` | no | — |
| `created_at` | no | — |
| `created_by_id` | no | — |
| `crn` | no | — |
| `description` | no | — |
| `excluded` | no | — |
| `excluded_count` | no | — |
| `href` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `name` | no | — |
| `zone_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cbr_zone_data_invalid_zone_id` | `zone_id` |

## Example

### Valid

```hcl
data "ibm_cbr_zone" "test" {
  zone_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
