# ibm_is_snapshot_clone_data

Schema validation for the ibm_is_snapshot_clone data source.

This rule validates `ibm_is_snapshot_clone` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SnapshotClonePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_snapshot_clone_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `available` | no | — |
| `created_at` | no | — |
| `snapshot` | yes | — |
| `zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_snapshot_clone_data_invalid_snapshot` | `snapshot` |
| `ibm_is_snapshot_clone_data_invalid_zone` | `zone` |

## Example

### Valid

```hcl
data "ibm_is_snapshot_clone" "test" {
  snapshot = "example"
  zone = "us-south-1"
}
```
