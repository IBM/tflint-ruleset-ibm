# ibm_code_engine_binding

Schema validation for the ibm_code_engine_binding resource.

This rule validates `ibm_code_engine_binding` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Binding` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_binding" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `binding_id` | no | — |
| `href` | no | — |
| `prefix` | yes | pattern `^[A-Z]([_A-Z0-9]*[A-Z0-9])*$` |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `resource_type` | no | — |
| `secret_name` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `status` | no | — |
| `component` (block) | yes | nested block |
| `component.name` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?$` |
| `component.resource_type` | yes | API: The type of the referenced resource. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_binding_invalid_prefix` | `prefix` |
| `ibm_code_engine_binding_invalid_project_id` | `project_id` |
| `ibm_code_engine_binding_invalid_secret_name` | `secret_name` |
| `ibm_code_engine_binding_invalid_component__name` | `component.name` |
| `ibm_code_engine_binding_invalid_component__resource_type` | `component.resource_type` |

## Example

### Valid

```hcl
variable "secret_name" {
  sensitive = true
}

resource "ibm_code_engine_binding" "test" {
  prefix = "PREFIX"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
  secret_name = var.secret_name
  component {
    name = "example-name"
    resource_type = "aaaaaaa"
  }
}
```
