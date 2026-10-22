# ibm_scc_scope_data

Reads ibm_scc_scope. The scope name.

This rule validates `ibm_scc_scope` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ScopePropertyExclusionItem` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_scope_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `attachment_count` | no | — |
| `created_by` | no | — |
| `created_on` | no | — |
| `description` | no | — |
| `environment` | no | — |
| `exclusions` | no | — |
| `instance_id` | yes | pattern `^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}$` |
| `name` | no | — |
| `properties` | no | — |
| `scope_id` | yes | pattern `^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}$` |
| `updated_by` | no | — |
| `updated_on` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_scope_data_invalid_instance_id` | `instance_id` |
| `ibm_scc_scope_data_invalid_scope_id` | `scope_id` |

## Example

### Valid

```hcl
data "ibm_scc_scope" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  scope_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
