# ibm_function_package_data

Reads ibm_function_package. Name of the package.

This rule validates `ibm_function_package` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_function_package_data" {
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
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_function_package_data_invalid_name` | `name` |
| `ibm_function_package_data_invalid_namespace` | `namespace` |

## Example

### Valid

```hcl
data "ibm_function_package" "test" {
  name = "A"
  namespace = "example"
}
```
