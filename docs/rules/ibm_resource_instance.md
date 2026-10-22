# ibm_resource_instance

Manages ibm_resource_instance. A name for the resource instance

This rule validates `ibm_resource_instance` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/service_instances/{instance_id}:patch` (`resource-controller-ibm-cloud-osb-api`).

## Configuration

```hcl
rule "ibm_resource_instance" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `allow_cleanup` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `dashboard_url` | no | — |
| `deleted_at` | no | — |
| `deleted_by` | no | — |
| `extensions` | no | — |
| `guid` | no | — |
| `last_operation` | no | — |
| `location` | yes | — |
| `locked` | no | — |
| `name` | yes | API: Name of the instance |
| `onetime_credentials` | no | — |
| `parameters` | no | API: Configuration options for the service instance or binding operation. Brokers should ensure that the client has provided valid configuration parameters and values for the operation. If this field is not present in the request message, then the broker must NOT change the parameters of the instance or binindg as a result of this request. |
| `parameters_json` | no | — |
| `plan` | yes | — |
| `plan_history` | no | — |
| `resource_aliases_url` | no | — |
| `resource_bindings_url` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_crn` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `resource_id` | no | — |
| `resource_keys_url` | no | — |
| `resource_name` | no | — |
| `resource_plan_id` | no | — |
| `resource_status` | no | — |
| `restored_at` | no | — |
| `restored_by` | no | — |
| `scheduled_reclaim_at` | no | — |
| `scheduled_reclaim_by` | no | — |
| `service` | yes | — |
| `service_endpoints` | no | one of: `public`, `private`, `public-and-private` |
| `state` | no | — |
| `status` | no | — |
| `sub_type` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `target_crn` | no | — |
| `type` | no | — |
| `update_at` | no | — |
| `update_by` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_resource_instance_invalid_location` | `location` |
| `ibm_resource_instance_invalid_name` | `name` |
| `ibm_resource_instance_invalid_parameters` | `parameters` |
| `ibm_resource_instance_invalid_parameters_json` | `parameters_json` |
| `ibm_resource_instance_invalid_plan` | `plan` |
| `ibm_resource_instance_invalid_resource_group_id` | `resource_group_id` |
| `ibm_resource_instance_invalid_service` | `service` |
| `ibm_resource_instance_invalid_service_endpoints` | `service_endpoints` |
| `ibm_resource_instance_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_resource_instance" "test" {
  location = "example"
  name = "example"
  plan = "example"
  service = "example"
}
```
