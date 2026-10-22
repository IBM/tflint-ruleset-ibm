# ibm_sds_volume_mapping

Schema validation for the ibm_sds_volume_mapping resource.

This rule validates `ibm_sds_volume_mapping` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Namespace` (`block-storage`).

## Configuration

```hcl
rule "ibm_sds_volume_mapping" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `gateways` | no | — |
| `host` | no | — |
| `host_id` | yes | pattern `^\S+$` |
| `href` | no | — |
| `namespace` | no | — |
| `sds_endpoint` | no | — |
| `status` | no | — |
| `storage_identifier` | no | — |
| `subsystem_nqn` | no | — |
| `volume_mapping_id` | no | — |
| `volume` (block) | yes | nested block |
| `volume.id` | yes | range 1–32 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sds_volume_mapping_invalid_host_id` | `host_id` |
| `ibm_sds_volume_mapping_invalid_sds_endpoint` | `sds_endpoint` |
| `ibm_sds_volume_mapping_invalid_volume__id` | `volume.id` |

## Example

### Valid

```hcl
resource "ibm_sds_volume_mapping" "test" {
  host_id = "S"
  volume {
    id = 1
  }
}
```
