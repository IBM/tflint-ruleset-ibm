# ibm_sds_volume

Manages ibm_sds_volume. Unique name of the host.

This rule validates `ibm_sds_volume` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Volume` (`block-storage`).

## Configuration

```hcl
rule "ibm_sds_volume" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bandwidth` | no | — |
| `capacity` | yes | range 1–32000 |
| `created_at` | no | — |
| `href` | no | — |
| `iops` | no | — |
| `name` | no | pattern `^[a-z][a-z0-9]*(-[a-z0-9]+)*$` |
| `resource_type` | no | — |
| `sds_endpoint` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `volume_mappings` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sds_volume_invalid_capacity` | `capacity` |
| `ibm_sds_volume_invalid_name` | `name` |
| `ibm_sds_volume_invalid_sds_endpoint` | `sds_endpoint` |

## Example

### Valid

```hcl
resource "ibm_sds_volume" "test" {
  capacity = 1
}
```
