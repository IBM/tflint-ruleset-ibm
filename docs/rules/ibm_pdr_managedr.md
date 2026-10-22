# ibm_pdr_managedr

Schema validation for the ibm_pdr_managedr resource.

This rule validates `ibm_pdr_managedr` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ServiceInstanceManageDRRequest` (`dr-automation-powervs`).

## Configuration

```hcl
rule "ibm_pdr_managedr" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | — |
| `accepts_incomplete` | no | — |
| `action` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `api_key` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `client_id` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `client_secret` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `dashboard_url` | no | — |
| `etag` | no | — |
| `guid` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `id` | no | — |
| `instance_id` | yes | — |
| `location_id` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `machine_type` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `managed_apikey` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `orchestrator_ha` | no | API: Indicates whether the orchestrator High Availability (HA) is enabled for the service instance. |
| `orchestrator_location_type` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `orchestrator_name` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `orchestrator_network_ids` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `orchestrator_password` | no | API: The password that you can use to access your orchestrator. |
| `orchestrator_workspace_id` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `proxy_ip` | no | pattern `^[a-zA-Z0-9.\-_:]+$` |
| `region_id` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `resource_instance` | no | API: The uniquie identifier of the associated IBM Cloud resource instance. |
| `secret` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `secret_group` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `ssh_key_name` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `stand_by_redeploy` | no | one of: `false`, `true` |
| `standby_machine_type` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `standby_orchestrator_name` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `standby_orchestrator_network_ids` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `standby_orchestrator_workspace_id` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `standby_ssh_key_name` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `standby_tier` | no | pattern `^[a-zA-Z0-9\-_]+$` |
| `tenant_name` | no | pattern `^[a-zA-Z0-9.\-]+$` |
| `tier` | no | pattern `^[a-zA-Z0-9\-_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pdr_managedr_invalid_accept_language` | `accept_language` |
| `ibm_pdr_managedr_invalid_accepts_incomplete` | `accepts_incomplete` |
| `ibm_pdr_managedr_invalid_action` | `action` |
| `ibm_pdr_managedr_invalid_api_key` | `api_key` |
| `ibm_pdr_managedr_invalid_client_id` | `client_id` |
| `ibm_pdr_managedr_invalid_client_secret` | `client_secret` |
| `ibm_pdr_managedr_invalid_guid` | `guid` |
| `ibm_pdr_managedr_invalid_instance_id` | `instance_id` |
| `ibm_pdr_managedr_invalid_location_id` | `location_id` |
| `ibm_pdr_managedr_invalid_machine_type` | `machine_type` |
| `ibm_pdr_managedr_invalid_managed_apikey` | `managed_apikey` |
| `ibm_pdr_managedr_invalid_orchestrator_ha` | `orchestrator_ha` |
| `ibm_pdr_managedr_invalid_orchestrator_location_type` | `orchestrator_location_type` |
| `ibm_pdr_managedr_invalid_orchestrator_name` | `orchestrator_name` |
| `ibm_pdr_managedr_invalid_orchestrator_network_ids` | `orchestrator_network_ids` |
| `ibm_pdr_managedr_invalid_orchestrator_password` | `orchestrator_password` |
| `ibm_pdr_managedr_invalid_orchestrator_workspace_id` | `orchestrator_workspace_id` |
| `ibm_pdr_managedr_invalid_proxy_ip` | `proxy_ip` |
| `ibm_pdr_managedr_invalid_region_id` | `region_id` |
| `ibm_pdr_managedr_invalid_resource_instance` | `resource_instance` |
| `ibm_pdr_managedr_invalid_secret` | `secret` |
| `ibm_pdr_managedr_invalid_secret_group` | `secret_group` |
| `ibm_pdr_managedr_invalid_ssh_key_name` | `ssh_key_name` |
| `ibm_pdr_managedr_invalid_stand_by_redeploy` | `stand_by_redeploy` |
| `ibm_pdr_managedr_invalid_standby_machine_type` | `standby_machine_type` |
| `ibm_pdr_managedr_invalid_standby_orchestrator_name` | `standby_orchestrator_name` |
| `ibm_pdr_managedr_invalid_standby_orchestrator_network_ids` | `standby_orchestrator_network_ids` |
| `ibm_pdr_managedr_invalid_standby_orchestrator_workspace_id` | `standby_orchestrator_workspace_id` |
| `ibm_pdr_managedr_invalid_standby_ssh_key_name` | `standby_ssh_key_name` |
| `ibm_pdr_managedr_invalid_standby_tier` | `standby_tier` |
| `ibm_pdr_managedr_invalid_tenant_name` | `tenant_name` |
| `ibm_pdr_managedr_invalid_tier` | `tier` |

## Example

### Valid

```hcl
resource "ibm_pdr_managedr" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
