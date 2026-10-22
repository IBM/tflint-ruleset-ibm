# ibm_pdr_last_operation_data

Schema validation for the ibm_pdr_last_operation data source.

This rule validates `ibm_pdr_last_operation` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pdr_last_operation_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | — |
| `crn` | no | — |
| `deployment_name` | no | — |
| `instance_id` | yes | — |
| `is_api_key_expired` | no | — |
| `last_updated_orchestrator_deployment_time` | no | — |
| `last_updated_standby_orchestrator_deployment_time` | no | — |
| `mfa_enabled` | no | — |
| `orch_ext_connectivity_status` | no | — |
| `orch_standby_node_addition_status` | no | — |
| `orchestrator_cluster_message` | no | — |
| `orchestrator_config_status` | no | — |
| `orchestrator_ha` | no | — |
| `plan_name` | no | — |
| `primary_description` | no | — |
| `primary_error_description` | no | — |
| `primary_ip_address` | no | — |
| `primary_orchestrator_status` | no | — |
| `recovery_location` | no | — |
| `resource_group` | no | — |
| `standby_description` | no | — |
| `standby_error_description` | no | — |
| `standby_ip_address` | no | — |
| `standby_status` | no | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pdr_last_operation_data_invalid_accept_language` | `accept_language` |
| `ibm_pdr_last_operation_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_pdr_last_operation" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
