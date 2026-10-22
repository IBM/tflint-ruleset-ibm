# ibm_function_namespace_data

Reads ibm_function_namespace. Name of namespace.

This rule validates `ibm_function_namespace` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_function_namespace_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `description` | no | — |
| `location` | no | — |
| `name` | yes | — |
| `resource_group_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_function_namespace_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_function_namespace" "test" {
  name = "example"
}
```
