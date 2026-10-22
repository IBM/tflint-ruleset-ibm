# ibm_scc_scope

Manages ibm_scc_scope. The scope name.

This rule validates `ibm_scc_scope` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ScopePrototype` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_scope" {
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
| `description` | no | pattern `^[a-zA-Z0-9_,'\s\-\.]*$` |
| `environment` | yes | pattern `^[a-zA-Z0-9_,'\s\-\.]*$` |
| `instance_id` | yes | pattern `^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}$` |
| `name` | yes | pattern `^[a-zA-Z0-9_,'\s\-\.]*$` |
| `properties` | yes | API: Attribute that details what kind of type of scope |
| `scope_id` | no | — |
| `updated_by` | no | — |
| `updated_on` | no | — |
| `exclusions` (block) | no | nested block |
| `exclusions.scope_id` | no | — |
| `exclusions.scope_type` | no | one of: `enterprise`, `enterprise.account_group`, `enterprise.account`, `account`, `account.resource_group` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_scope_invalid_description` | `description` |
| `ibm_scc_scope_invalid_environment` | `environment` |
| `ibm_scc_scope_invalid_instance_id` | `instance_id` |
| `ibm_scc_scope_invalid_name` | `name` |
| `ibm_scc_scope_invalid_properties` | `properties` |
| `ibm_scc_scope_invalid_exclusions__scope_id` | `exclusions.scope_id` |
| `ibm_scc_scope_invalid_exclusions__scope_type` | `exclusions.scope_type` |

## Example

### Valid

```hcl
resource "ibm_scc_scope" "test" {
  environment = "example-name"
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  name = "example-name"
  properties = "example"
}
```
