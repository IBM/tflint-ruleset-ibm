# ibm_function_trigger

Manages ibm_function_trigger. Name of Trigger.

This rule validates `ibm_function_trigger` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_function_trigger" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `annotations` | no | — |
| `name` | yes | pattern `\A([\w]|[\w][\w@ .-]*[\w@.-]+)\z` |
| `namespace` | yes | — |
| `parameters` | no | — |
| `publish` | no | — |
| `trigger_id` | no | — |
| `user_defined_annotations` | no | json format |
| `user_defined_parameters` | no | json format |
| `version` | no | — |
| `feed` (block) | no | nested block |
| `feed.name` | yes | — |
| `feed.parameters` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_function_trigger_invalid_name` | `name` |
| `ibm_function_trigger_invalid_namespace` | `namespace` |
| `ibm_function_trigger_invalid_user_defined_annotations` | `user_defined_annotations` |
| `ibm_function_trigger_invalid_user_defined_parameters` | `user_defined_parameters` |
| `ibm_function_trigger_invalid_feed__name` | `feed.name` |
| `ibm_function_trigger_invalid_feed__parameters` | `feed.parameters` |

## Example

### Valid

```hcl
resource "ibm_function_trigger" "test" {
  name = "A"
  namespace = "example"
}
```
