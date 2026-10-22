# ibm_schematics_agent_data

Reads ibm_schematics_agent. The name of the agent (must be unique, for an account).

This rule validates `ibm_schematics_agent` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/v2/agents/{agent_id}` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_agent_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `agent_crn` | no | — |
| `agent_id` | yes | — |
| `agent_infrastructure` | no | — |
| `agent_inputs` | no | — |
| `agent_kpi` | no | — |
| `agent_location` | no | — |
| `agent_metadata` | no | — |
| `created_at` | no | — |
| `creation_by` | no | — |
| `description` | no | — |
| `id` | no | — |
| `name` | no | — |
| `recent_deploy_job` | no | — |
| `recent_health_job` | no | — |
| `recent_prs_job` | no | — |
| `resource_group` | no | — |
| `schematics_location` | no | — |
| `system_state` | no | — |
| `tags` | no | — |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `user_state` | no | — |
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_agent_data_invalid_agent_id` | `agent_id` |

## Example

### Valid

```hcl
data "ibm_schematics_agent" "test" {
  agent_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
