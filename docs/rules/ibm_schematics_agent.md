# ibm_schematics_agent

Manages ibm_schematics_agent. The name of the agent (must be unique, for an account).

This rule validates `ibm_schematics_agent` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VariableMetadata` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_agent" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `agent_crn` | no | — |
| `agent_kpi` | no | — |
| `agent_location` | yes | API: The location where agent is deployed in the user environment. |
| `created_at` | no | — |
| `creation_by` | no | — |
| `description` | no | API: The description of the meta data. |
| `name` | yes | API: Name of the Service Data |
| `recent_deploy_job` | no | — |
| `recent_health_job` | no | — |
| `recent_prs_job` | no | — |
| `resource_group` | yes | API: The ID of the resource group where you want to provision the workspace. |
| `run_destroy_resources` | no | — |
| `schematics_location` | yes | one of: `us-south`, `us-east`, `eu-gb`, `eu-de` |
| `system_state` | no | — |
| `tags` | no | API: A list of tags that are associated with the workspace. |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `version` | yes | API: A version of the terraform template. |
| `agent_infrastructure` (block) | yes | nested block |
| `agent_infrastructure.cluster_id` | no | API: The ID of the cluster where you want to provision the resources of all IBM Cloud catalog templates that are included in the catalog offering. |
| `agent_infrastructure.cluster_resource_group` | no | API: The resource group of the cluster (is it required?) |
| `agent_infrastructure.cos_bucket_name` | no | API: The COS bucket name used to store the logs. |
| `agent_infrastructure.cos_bucket_region` | no | API: The COS bucket region. |
| `agent_infrastructure.cos_instance_name` | no | API: The COS instance name to store the agent logs. |
| `agent_infrastructure.infra_type` | no | one of: `ibm_kubernetes`, `ibm_openshift`, `ibm_satellite` |
| `agent_inputs` (block) | no | nested block |
| `agent_inputs.name` | no | API: Name of the Service Data |
| `agent_inputs.use_default` | no | — |
| `agent_inputs.value` | no | API: Values of service data |
| `agent_inputs.metadata` (block) | no | nested block |
| `agent_inputs.metadata.aliases` | no | API: The list of aliases for the variable name. |
| `agent_inputs.metadata.cloud_data_type` | no | API: Cloud data type of the variable. eg. resource_group_id, region, vpc_id. |
| `agent_inputs.metadata.default_value` | no | API: Default value for the variable only if the override value is not specified. |
| `agent_inputs.metadata.description` | no | API: The description of the meta data. |
| `agent_inputs.metadata.group_by` | no | API: The display name of the group this variable belongs to. |
| `agent_inputs.metadata.hidden` | no | API: If **true**, the variable is not displayed on UI or Command line. |
| `agent_inputs.metadata.immutable` | no | API: Is the variable readonly ? |
| `agent_inputs.metadata.link_status` | no | one of: `normal`, `broken` |
| `agent_inputs.metadata.matches` | no | API: The regex for the variable value. |
| `agent_inputs.metadata.max_length` | no | range 0–10000 |
| `agent_inputs.metadata.max_value` | no | range 0–10000 |
| `agent_inputs.metadata.min_length` | no | range 0–10000 |
| `agent_inputs.metadata.min_value` | no | range 0–10000 |
| `agent_inputs.metadata.options` | no | API: The list of possible values for this variable. If type is **integer** or **date**, then the array of string is converted to array of integers or date during the runtime. |
| `agent_inputs.metadata.position` | no | range 0–10000 |
| `agent_inputs.metadata.required` | no | API: If the variable required? |
| `agent_inputs.metadata.secure` | no | API: Is the variable secure or sensitive ? |
| `agent_inputs.metadata.source` | no | API: The source of this meta-data. |
| `agent_inputs.metadata.type` | no | one of: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link` |
| `agent_metadata` (block) | no | nested block |
| `agent_metadata.name` | no | API: Name of the Service Data |
| `agent_metadata.value` | no | API: Values of service data |
| `user_state` (block) | no | nested block |
| `user_state.state` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_agent_invalid_agent_location` | `agent_location` |
| `ibm_schematics_agent_invalid_description` | `description` |
| `ibm_schematics_agent_invalid_name` | `name` |
| `ibm_schematics_agent_invalid_resource_group` | `resource_group` |
| `ibm_schematics_agent_invalid_run_destroy_resources` | `run_destroy_resources` |
| `ibm_schematics_agent_invalid_schematics_location` | `schematics_location` |
| `ibm_schematics_agent_invalid_tags` | `tags` |
| `ibm_schematics_agent_invalid_version` | `version` |
| `ibm_schematics_agent_invalid_agent_infrastructure__cluster_id` | `agent_infrastructure.cluster_id` |
| `ibm_schematics_agent_invalid_agent_infrastructure__cluster_resource_group` | `agent_infrastructure.cluster_resource_group` |
| `ibm_schematics_agent_invalid_agent_infrastructure__cos_bucket_name` | `agent_infrastructure.cos_bucket_name` |
| `ibm_schematics_agent_invalid_agent_infrastructure__cos_bucket_region` | `agent_infrastructure.cos_bucket_region` |
| `ibm_schematics_agent_invalid_agent_infrastructure__cos_instance_name` | `agent_infrastructure.cos_instance_name` |
| `ibm_schematics_agent_invalid_agent_infrastructure__infra_type` | `agent_infrastructure.infra_type` |
| `ibm_schematics_agent_invalid_agent_inputs__name` | `agent_inputs.name` |
| `ibm_schematics_agent_invalid_agent_inputs__use_default` | `agent_inputs.use_default` |
| `ibm_schematics_agent_invalid_agent_inputs__value` | `agent_inputs.value` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__aliases` | `agent_inputs.metadata.aliases` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__cloud_data_type` | `agent_inputs.metadata.cloud_data_type` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__default_value` | `agent_inputs.metadata.default_value` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__description` | `agent_inputs.metadata.description` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__group_by` | `agent_inputs.metadata.group_by` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__hidden` | `agent_inputs.metadata.hidden` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__immutable` | `agent_inputs.metadata.immutable` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__link_status` | `agent_inputs.metadata.link_status` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__matches` | `agent_inputs.metadata.matches` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__max_length` | `agent_inputs.metadata.max_length` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__max_value` | `agent_inputs.metadata.max_value` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__min_length` | `agent_inputs.metadata.min_length` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__min_value` | `agent_inputs.metadata.min_value` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__options` | `agent_inputs.metadata.options` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__position` | `agent_inputs.metadata.position` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__required` | `agent_inputs.metadata.required` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__secure` | `agent_inputs.metadata.secure` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__source` | `agent_inputs.metadata.source` |
| `ibm_schematics_agent_invalid_agent_inputs__metadata__type` | `agent_inputs.metadata.type` |
| `ibm_schematics_agent_invalid_agent_metadata__name` | `agent_metadata.name` |
| `ibm_schematics_agent_invalid_agent_metadata__value` | `agent_metadata.value` |
| `ibm_schematics_agent_invalid_user_state__state` | `user_state.state` |

## Example

### Valid

```hcl
resource "ibm_schematics_agent" "test" {
  agent_location = "aaaaaaa"
  name = "aaaaaaa"
  resource_group = "aaaaaaa"
  schematics_location = "us-south"
  version = "aaaaaaa"
  agent_infrastructure {
  }
}
```
