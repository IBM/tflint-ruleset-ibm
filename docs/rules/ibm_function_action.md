# ibm_function_action

Manages ibm_function_action. Name of action.

This rule validates `ibm_function_action` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_function_action" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action_id` | no | — |
| `annotations` | no | — |
| `name` | yes | pattern `^[^/*][a-zA-Z0-9/_@.-]` |
| `namespace` | yes | — |
| `parameters` | no | — |
| `publish` | no | — |
| `target_endpoint_url` | no | — |
| `user_defined_annotations` | no | json format |
| `user_defined_parameters` | no | json format |
| `version` | no | — |
| `exec` (block) | yes | nested block |
| `exec.code` | no | — |
| `exec.code_path` | no | — |
| `exec.components` | no | — |
| `exec.image` | no | — |
| `exec.init` | no | — |
| `exec.kind` | yes | — |
| `exec.main` | no | — |
| `limits` (block) | no | nested block |
| `limits.log_size` | no | — |
| `limits.memory` | no | — |
| `limits.timeout` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_function_action_invalid_name` | `name` |
| `ibm_function_action_invalid_namespace` | `namespace` |
| `ibm_function_action_invalid_publish` | `publish` |
| `ibm_function_action_invalid_user_defined_annotations` | `user_defined_annotations` |
| `ibm_function_action_invalid_user_defined_parameters` | `user_defined_parameters` |
| `ibm_function_action_invalid_exec__code` | `exec.code` |
| `ibm_function_action_invalid_exec__code_path` | `exec.code_path` |
| `ibm_function_action_invalid_exec__components` | `exec.components` |
| `ibm_function_action_invalid_exec__image` | `exec.image` |
| `ibm_function_action_invalid_exec__init` | `exec.init` |
| `ibm_function_action_invalid_exec__kind` | `exec.kind` |
| `ibm_function_action_invalid_exec__main` | `exec.main` |
| `ibm_function_action_invalid_limits__log_size` | `limits.log_size` |
| `ibm_function_action_invalid_limits__memory` | `limits.memory` |
| `ibm_function_action_invalid_limits__timeout` | `limits.timeout` |

## Example

### Valid

```hcl
resource "ibm_function_action" "test" {
  name = "example-name"
  namespace = "example"
  exec {
    kind = "example"
  }
}
```
