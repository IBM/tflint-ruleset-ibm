# ibm_is_volumes_data

Schema validation for the ibm_is_volumes data source.

This rule validates `ibm_is_volumes` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `leaf-unique` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_volumes_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `attachment_state` | no | one of: `attached`, `unattached`, `unusable`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `encryption` | no | — |
| `operating_system_architecture` | no | — |
| `operating_system_family` | no | — |
| `volume_name` | no | — |
| `volumes` | no | — |
| `zone_name` | no | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_volumes_data_invalid_attachment_state` | `attachment_state` |
| `ibm_is_volumes_data_invalid_encryption` | `encryption` |
| `ibm_is_volumes_data_invalid_operating_system_architecture` | `operating_system_architecture` |
| `ibm_is_volumes_data_invalid_operating_system_family` | `operating_system_family` |
| `ibm_is_volumes_data_invalid_volume_name` | `volume_name` |
| `ibm_is_volumes_data_invalid_zone_name` | `zone_name` |

## Example

### Valid

```hcl
data "ibm_is_volumes" "test" {
}
```
