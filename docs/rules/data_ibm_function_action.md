# ibm_function_action_data

Reads ibm_function_action. Name of action.

This rule validates `ibm_function_action` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_function_action_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action_id` | no | — |
| `annotations` | no | — |
| `exec` | no | — |
| `limits` | no | — |
| `name` | yes | pattern `^[^/*][a-zA-Z0-9/_@.-]` |
| `namespace` | yes | — |
| `parameters` | no | — |
| `publish` | no | — |
| `target_endpoint_url` | no | — |
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_function_action_data_invalid_name` | `name` |
| `ibm_function_action_data_invalid_namespace` | `namespace` |

## Example

### Valid

```hcl
data "ibm_function_action" "test" {
  name = "example-name"
  namespace = "example"
}
```
