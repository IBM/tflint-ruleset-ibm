# ibm_logs_view_folder

Manages ibm_logs_view_folder. Folder name.

This rule validates `ibm_logs_view_folder` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ViewFolderPrototype` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_view_folder" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `view_folder_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_view_folder_invalid_name` | `name` |

## Example

### Valid

```hcl
resource "ibm_logs_view_folder" "test" {
  name = "example-name"
}
```
