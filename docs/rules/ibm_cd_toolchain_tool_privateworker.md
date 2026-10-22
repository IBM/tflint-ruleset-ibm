# ibm_cd_toolchain_tool_privateworker

Manages ibm_cd_toolchain_tool_privateworker. Name of the tool.

This rule validates `ibm_cd_toolchain_tool_privateworker` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Toolchain` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain_tool_privateworker" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `href` | no | — |
| `name` | no | pattern `^([^\x00-\x7F]|[a-zA-Z0-9-._ ])+$` |
| `referent` | no | — |
| `resource_group_id` | no | — |
| `state` | no | — |
| `tool_id` | no | — |
| `toolchain_crn` | no | — |
| `toolchain_id` | yes | pattern `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$` |
| `updated_at` | no | — |
| `parameters` (block) | yes | nested block |
| `parameters.name` | yes | pattern `^([^\x00-\x7F]|[a-zA-Z0-9-._ ])+$` |
| `parameters.worker_queue_credentials` | yes | — |
| `parameters.worker_queue_identifier` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_tool_privateworker_invalid_name` | `name` |
| `ibm_cd_toolchain_tool_privateworker_invalid_toolchain_id` | `toolchain_id` |
| `ibm_cd_toolchain_tool_privateworker_invalid_parameters__name` | `parameters.name` |
| `ibm_cd_toolchain_tool_privateworker_invalid_parameters__worker_queue_credentials` | `parameters.worker_queue_credentials` |
| `ibm_cd_toolchain_tool_privateworker_invalid_parameters__worker_queue_identifier` | `parameters.worker_queue_identifier` |

## Example

### Valid

```hcl
variable "worker_queue_credentials" {
  sensitive = true
}

resource "ibm_cd_toolchain_tool_privateworker" "test" {
  toolchain_id = "01234567-89ab-4cde-8123-456789abcdef"
  parameters {
    name = "example-name"
    worker_queue_credentials = var.worker_queue_credentials
  }
}
```
