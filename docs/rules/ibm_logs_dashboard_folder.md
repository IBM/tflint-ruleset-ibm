# ibm_logs_dashboard_folder

Manages ibm_logs_dashboard_folder. The dashboard folder name, required.

This rule validates `ibm_logs_dashboard_folder` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DashboardFolder` (`logs-service-api`).

## Configuration

```hcl
rule "ibm_logs_dashboard_folder" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `dashboard_folder_id` | no | — |
| `name` | yes | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |
| `parent_id` | no | pattern `^[\p{L}\p{N}\p{P}\p{Z}\p{S}\p{M}]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_logs_dashboard_folder_invalid_name` | `name` |
| `ibm_logs_dashboard_folder_invalid_parent_id` | `parent_id` |

## Example

### Valid

```hcl
resource "ibm_logs_dashboard_folder" "test" {
  name = "example-name"
}
```
