# ibm_schematics_agent_deploy

Schema validation for the ibm_schematics_agent_deploy resource.

This rule validates `ibm_schematics_agent_deploy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/v2/agents/{agent_id}` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_agent_deploy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `agent_id` | yes | — |
| `agent_version` | no | — |
| `force` | no | — |
| `is_redeployed` | no | — |
| `job_id` | no | — |
| `log_url` | no | — |
| `status_code` | no | — |
| `status_message` | no | — |
| `updated_at` | no | — |
| `updated_by` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_agent_deploy_invalid_agent_id` | `agent_id` |
| `ibm_schematics_agent_deploy_invalid_force` | `force` |

## Example

### Valid

```hcl
resource "ibm_schematics_agent_deploy" "test" {
  agent_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
