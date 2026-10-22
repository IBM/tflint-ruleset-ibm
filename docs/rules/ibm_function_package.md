# ibm_function_package

Manages ibm_function_package. Name of package.

This rule validates `ibm_function_package` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_function_package" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `annotations` | no | — |
| `bind_package_name` | no | — |
| `name` | yes | pattern `\A([\w]|[\w][\w@ .-]*[\w@.-]+)\z` |
| `namespace` | yes | — |
| `package_id` | no | — |
| `parameters` | no | — |
| `publish` | no | — |
| `user_defined_annotations` | no | json format |
| `user_defined_parameters` | no | — |
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_function_package_invalid_bind_package_name` | `bind_package_name` |
| `ibm_function_package_invalid_name` | `name` |
| `ibm_function_package_invalid_namespace` | `namespace` |
| `ibm_function_package_invalid_publish` | `publish` |
| `ibm_function_package_invalid_user_defined_annotations` | `user_defined_annotations` |
| `ibm_function_package_invalid_user_defined_parameters` | `user_defined_parameters` |

## Example

### Valid

```hcl
resource "ibm_function_package" "test" {
  name = "A"
  namespace = "example"
}
```
