# ibm_is_share_replica_operations

Schema validation for the ibm_is_share_replica_operations resource.

This rule validates `ibm_is_share_replica_operations` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ShareFailoverOptions` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_share_replica_operations" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `fallback_policy` | no | one of: `fail`, `split`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `share_replica` | yes | — |
| `split_share` | no | — |
| `timeout` | no | range 300–3600 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_share_replica_operations_invalid_fallback_policy` | `fallback_policy` |
| `ibm_is_share_replica_operations_invalid_share_replica` | `share_replica` |
| `ibm_is_share_replica_operations_invalid_split_share` | `split_share` |
| `ibm_is_share_replica_operations_invalid_timeout` | `timeout` |

## Example

### Valid

```hcl
resource "ibm_is_share_replica_operations" "test" {
  share_replica = "example"
}
```
