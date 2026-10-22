# ibm_is_volume_job_data

Reads ibm_is_volume_job. The name for this volume job.

This rule validates `ibm_is_volume_job` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_volume_job_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `auto_delete` | no | — |
| `completed_at` | no | — |
| `created_at` | no | — |
| `estimated_completion_at` | no | — |
| `href` | no | — |
| `job_type` | no | — |
| `name` | no | — |
| `parameters` | no | — |
| `resource_type` | no | — |
| `started_at` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `volume_id` | yes | pattern `^[-0-9a-z_]+$` |
| `volume_job_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_volume_job_data_invalid_volume_id` | `volume_id` |
| `ibm_is_volume_job_data_invalid_volume_job_id` | `volume_job_id` |

## Example

### Valid

```hcl
data "ibm_is_volume_job" "test" {
  volume_id = "example-name"
  volume_job_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
