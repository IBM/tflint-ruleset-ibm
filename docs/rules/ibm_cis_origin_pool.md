# ibm_cis_origin_pool

Manages ibm_cis_origin_pool. name

This rule validates `ibm_cis_origin_pool` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `load_balancer_pool_req` (`cis`).

## Configuration

```hcl
rule "ibm_cis_origin_pool" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `check_regions` | yes | API: Geographic region code |
| `cis_id` | yes | — |
| `created_on` | no | — |
| `description` | no | API: Desc |
| `enabled` | yes | API: Enabled/disabled |
| `health` | no | — |
| `healthy` | no | — |
| `minimum_origins` | no | API: The minimum number of origins that must be healthy for this pool to serve traffic |
| `modified_on` | no | — |
| `monitor` | no | API: Monitor |
| `name` | yes | API: Name |
| `notification_email` | no | API: Notification email |
| `pool_id` | no | — |
| `origins` (block) | yes | nested block |
| `origins.address` | yes | API: Address |
| `origins.enabled` | yes | API: Enabled/disabled |
| `origins.name` | yes | API: Name |
| `origins.weight` | no | API: Weight |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_origin_pool_invalid_check_regions` | `check_regions` |
| `ibm_cis_origin_pool_invalid_cis_id` | `cis_id` |
| `ibm_cis_origin_pool_invalid_description` | `description` |
| `ibm_cis_origin_pool_invalid_enabled` | `enabled` |
| `ibm_cis_origin_pool_invalid_minimum_origins` | `minimum_origins` |
| `ibm_cis_origin_pool_invalid_monitor` | `monitor` |
| `ibm_cis_origin_pool_invalid_name` | `name` |
| `ibm_cis_origin_pool_invalid_notification_email` | `notification_email` |
| `ibm_cis_origin_pool_invalid_origins__address` | `origins.address` |
| `ibm_cis_origin_pool_invalid_origins__enabled` | `origins.enabled` |
| `ibm_cis_origin_pool_invalid_origins__name` | `origins.name` |
| `ibm_cis_origin_pool_invalid_origins__weight` | `origins.weight` |

## Example

### Valid

```hcl
resource "ibm_cis_origin_pool" "test" {
  check_regions = "example"
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  enabled = true
  name = "example"
  origins {
    address = "10.0.0.1"
    enabled = true
    name = "example"
  }
}
```
