# ibm_is_share_accessor_binding_data

Schema validation for the ibm_is_share_accessor_binding data source.

This rule validates `ibm_is_share_accessor_binding` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_share_accessor_binding_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accessor` | no | — |
| `accessor_binding` | yes | — |
| `created_at` | no | — |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `resource_type` | no | — |
| `share` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_share_accessor_binding_data_invalid_accessor_binding` | `accessor_binding` |
| `ibm_is_share_accessor_binding_data_invalid_share` | `share` |

## Example

### Valid

```hcl
data "ibm_is_share_accessor_binding" "test" {
  accessor_binding = "example"
  share = "example"
}
```
