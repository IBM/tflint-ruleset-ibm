# ibm_is_source_share_data

Reads ibm_is_source_share. Name of the share.

This rule validates `ibm_is_source_share` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_source_share_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `href` | no | — |
| `name` | no | — |
| `resource_type` | no | — |
| `share_replica` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_source_share_data_invalid_share_replica` | `share_replica` |

## Example

### Valid

```hcl
data "ibm_is_source_share" "test" {
  share_replica = "example"
}
```
