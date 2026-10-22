# ibm_scc_scope_collection_data

Schema validation for the ibm_scc_scope_collection data source.

This rule validates `ibm_scc_scope_collection` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Scope` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_scope_collection_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_id` | yes | pattern `^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-4[0-9A-Fa-f]{3}-[89ABab][0-9A-Fa-f]{3}-[0-9A-Fa-f]{12}$|^$` |
| `scopes` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_scope_collection_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_scc_scope_collection" "test" {
  instance_id = "01234567-89ab-4cde-8123-456789abcdef"
}
```
