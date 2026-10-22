# ibm_is_volume_data

Reads ibm_is_volume. Volume name

This rule validates `ibm_is_volume` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DedicatedHostGroupPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_volume_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `active` | no | — |
| `adjustable_capacity_states` | no | — |
| `adjustable_iops_states` | no | — |
| `allowed_use` | no | — |
| `attachment_state` | no | — |
| `bandwidth` | no | — |
| `busy` | no | — |
| `capacity` | no | — |
| `catalog_offering` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `encryption_key` | no | — |
| `encryption_type` | no | — |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `identifier` | no | — |
| `iops` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `operating_system` | no | — |
| `profile` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `source_snapshot` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `storage_generation` | no | — |
| `tags` | no | — |
| `zone` | no | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_volume_data_invalid_identifier` | `identifier` |
| `ibm_is_volume_data_invalid_name` | `name` |
| `ibm_is_volume_data_invalid_zone` | `zone` |

## Example

### Valid

```hcl
data "ibm_is_volume" "test" {
}
```
