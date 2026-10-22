# ibm_is_zones_data

Schema validation for the ibm_is_zones data source.

This rule validates `ibm_is_zones` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BackupPolicyPlanRemote` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_zones_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `region` | yes | API: If present, this property indicates that the referenced resource is remote to this region, and identifies the native region. |
| `status` | no | — |
| `zone_info` | no | — |
| `zones` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_zones_data_invalid_region` | `region` |
| `ibm_is_zones_data_invalid_status` | `status` |

## Example

### Valid

```hcl
data "ibm_is_zones" "test" {
  region = "example"
}
```
