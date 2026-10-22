# ibm_resource_instance_data

Reads ibm_resource_instance. Resource instance name for example, myobjectstorage

This rule validates `ibm_resource_instance` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/service_instances/{instance_id}:patch` (`resource-controller-ibm-cloud-osb-api`).

## Configuration

```hcl
rule "ibm_resource_instance_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `extensions` | no | — |
| `guid` | no | — |
| `identifier` | no | — |
| `location` | no | — |
| `name` | no | API: Name of the instance |
| `onetime_credentials` | no | — |
| `parameters_json` | no | — |
| `plan` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `service` | no | — |
| `status` | no | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_resource_instance_data_invalid_identifier` | `identifier` |
| `ibm_resource_instance_data_invalid_location` | `location` |
| `ibm_resource_instance_data_invalid_name` | `name` |
| `ibm_resource_instance_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_resource_instance_data_invalid_service` | `service` |

## Example

### Valid

```hcl
data "ibm_resource_instance" "test" {
}
```
