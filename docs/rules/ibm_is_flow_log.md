# ibm_is_flow_log

Manages ibm_is_flow_log. Flow Log Collector name

This rule validates `ibm_is_flow_log` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `FlowLogCollectorPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_flow_log" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `active` | no | API: Indicates whether this collector will be active upon creation. |
| `auto_delete` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `storage_bucket` | yes | API: Identifies a Cloud Object Storage bucket by a unique property. |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `target` | yes | API: The target this collector will collect flow logs for. If the target is an instance, subnet, or VPC, flow logs will not be collected for any instance network attachments, virtual network interfaces or instance network interfaces within the target that are themselves the target of a more specific flow log collector. For this request to succeed, the target must: - not be a virtual network interface t |
| `vpc` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_flow_log_invalid_access_tags` | `access_tags` |
| `ibm_is_flow_log_invalid_active` | `active` |
| `ibm_is_flow_log_invalid_name` | `name` |
| `ibm_is_flow_log_invalid_resource_group` | `resource_group` |
| `ibm_is_flow_log_invalid_storage_bucket` | `storage_bucket` |
| `ibm_is_flow_log_invalid_tags` | `tags` |
| `ibm_is_flow_log_invalid_target` | `target` |

## Example

### Valid

```hcl
resource "ibm_is_flow_log" "test" {
  name = "example-name"
  storage_bucket = "example"
  target = "example"
}
```
