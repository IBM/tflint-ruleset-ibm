# ibm_pha_deployment

Schema validation for the ibm_pha_deployment resource.

This rule validates `ibm_pha_deployment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pha_deployment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | pattern `^[a-zA-Z0-9\-_,;=.*]+$` |
| `api_key` | no | pattern `^[a-zA-Z0-9._:-]+$` |
| `apikey` | no | — |
| `cloud_account_id` | no | — |
| `cluster_type` | no | pattern `^[a-zA-Z0-9._:-]+$` |
| `configure_type` | no | pattern `^[a-zA-Z0-9._:-]+$` |
| `connectivity_type` | no | — |
| `creation_time` | no | — |
| `custom_network` | no | — |
| `deprovision_time` | no | — |
| `etag` | no | — |
| `guid` | no | — |
| `id` | no | — |
| `if_none_match` | no | pattern `^[a-zA-Z0-9\-_,;=.*]+$` |
| `instance_id` | yes | pattern `^[a-zA-Z0-9-]+$` |
| `is_duplicate` | no | — |
| `location_id` | yes | pattern `^[a-zA-Z0-9._:-]+$` |
| `plan_id` | no | — |
| `plan_name` | no | — |
| `powerha_cluster_name` | no | — |
| `powerha_cluster_type` | no | — |
| `powerha_level` | no | — |
| `primary_cluster_nodes` | no | — |
| `primary_cluster_nodes_details` | no | — |
| `primary_region_name` | no | — |
| `primary_workspace` | yes | pattern `^[a-zA-Z0-9._:-]+$` |
| `primary_workspace_name` | no | — |
| `provision_end_time` | no | — |
| `provision_start_time` | no | — |
| `provision_status` | no | — |
| `region_id` | no | — |
| `resource_group` | no | — |
| `resource_group_crn` | no | — |
| `resource_instance` | no | — |
| `secondary_cluster_nodes` | no | — |
| `secondary_location` | no | pattern `^[a-zA-Z0-9._:-]+$` |
| `secondary_workspace` | no | pattern `^[a-zA-Z0-9._:-]+$` |
| `service_description` | no | — |
| `service_id` | no | — |
| `service_name` | no | — |
| `standby_cluster_nodes` | no | — |
| `standby_region_name` | no | — |
| `standby_workspace_name` | no | — |
| `user_tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pha_deployment_invalid_accept_language` | `accept_language` |
| `ibm_pha_deployment_invalid_api_key` | `api_key` |
| `ibm_pha_deployment_invalid_cluster_type` | `cluster_type` |
| `ibm_pha_deployment_invalid_configure_type` | `configure_type` |
| `ibm_pha_deployment_invalid_if_none_match` | `if_none_match` |
| `ibm_pha_deployment_invalid_instance_id` | `instance_id` |
| `ibm_pha_deployment_invalid_location_id` | `location_id` |
| `ibm_pha_deployment_invalid_primary_cluster_nodes` | `primary_cluster_nodes` |
| `ibm_pha_deployment_invalid_primary_workspace` | `primary_workspace` |
| `ibm_pha_deployment_invalid_secondary_location` | `secondary_location` |
| `ibm_pha_deployment_invalid_secondary_workspace` | `secondary_workspace` |
| `ibm_pha_deployment_invalid_standby_cluster_nodes` | `standby_cluster_nodes` |

## Example

### Valid

```hcl
resource "ibm_pha_deployment" "test" {
  instance_id = "example-name"
  location_id = "example-name"
  primary_workspace = "example-name"
}
```
