# ibm_function_namespace

Manages ibm_function_namespace. Name of namespace.

This rule validates `ibm_function_namespace` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_function_namespace" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | — |
| `location` | no | — |
| `name` | yes | — |
| `resource_group_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_function_namespace_invalid_description` | `description` |
| `ibm_function_namespace_invalid_name` | `name` |
| `ibm_function_namespace_invalid_resource_group_id` | `resource_group_id` |

## Example

### Valid

```hcl
resource "ibm_function_namespace" "test" {
  name = "example"
  resource_group_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
