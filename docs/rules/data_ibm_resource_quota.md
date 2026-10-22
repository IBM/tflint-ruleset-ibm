# ibm_resource_quota_data

Reads ibm_resource_quota. Resource quota name, for example Trial Quota

This rule validates `ibm_resource_quota` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/resource_keys/{id}:patch` (`resource-controller-resource-controller`).

## Configuration

```hcl
rule "ibm_resource_quota_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `max_app_instance_memory` | no | — |
| `max_apps` | no | — |
| `max_instances_per_app` | no | — |
| `max_service_instances` | no | — |
| `name` | yes | pattern `^([\p{L}\p{N}._: -]{1,256})$` |
| `total_app_memory` | no | — |
| `type` | no | — |
| `vsi_limit` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_resource_quota_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_resource_quota" "test" {
  name = "example-name"
}
```
