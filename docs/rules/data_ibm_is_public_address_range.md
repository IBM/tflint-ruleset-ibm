# ibm_is_public_address_range_data

Reads ibm_is_public_address_range. The unique user-defined name for this public-address-range

This rule validates `ibm_is_public_address_range` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_public_address_range_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `cidr` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `href` | no | — |
| `identifier` | no | — |
| `ipv4_address_count` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `tags` | no | — |
| `target` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_public_address_range_data_invalid_identifier` | `identifier` |
| `ibm_is_public_address_range_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_public_address_range" "test" {
}
```
