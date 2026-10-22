# ibm_is_snapshot_clones_data

Schema validation for the ibm_is_snapshot_clones data source.

This rule validates `ibm_is_snapshot_clones` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_snapshot_clones_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `clones` | no | — |
| `snapshot` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_snapshot_clones_data_invalid_snapshot` | `snapshot` |

## Example

### Valid

```hcl
data "ibm_is_snapshot_clones" "test" {
  snapshot = "example"
}
```
