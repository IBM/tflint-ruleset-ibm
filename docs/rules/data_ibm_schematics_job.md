# ibm_schematics_job_data

Reads ibm_schematics_job. Job name, uniquely derived from the related Workspace or Action.

This rule validates `ibm_schematics_job` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/v2/jobs/{job_id}` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_job_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bastion` | no | — |
| `command_name` | no | — |
| `command_object` | no | — |
| `command_object_id` | no | — |
| `command_options` | no | — |
| `command_parameter` | no | — |
| `data` | no | — |
| `description` | no | — |
| `duration` | no | — |
| `end_at` | no | — |
| `id` | no | — |
| `job_env_settings` | no | — |
| `job_id` | yes | — |
| `job_inputs` | no | — |
| `location` | no | — |
| `log_store_url` | no | — |
| `log_summary` | no | — |
| `name` | no | — |
| `resource_group` | no | — |
| `results_url` | no | — |
| `start_at` | no | — |
| `state_store_url` | no | — |
| `status` | no | — |
| `submitted_at` | no | — |
| `submitted_by` | no | — |
| `tags` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_job_data_invalid_job_id` | `job_id` |
| `ibm_schematics_job_data_invalid_location` | `location` |

## Example

### Valid

```hcl
data "ibm_schematics_job" "test" {
  job_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
