# ibm_is_volume_job

Manages ibm_is_volume_job. The name for this volume job.

This rule validates `ibm_is_volume_job` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Volume` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_volume_job" {
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
| `job_type` | yes | one of: `migrate` |
| `limit` | no | range 1–100 |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_type` | no | — |
| `start` | no | pattern `^[ -~]+$` |
| `started_at` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `volume_id` | yes | pattern `^[-0-9a-z_]+$` |
| `volume_job_id` | no | — |
| `parameters` (block) | no | nested block |
| `parameters.bandwidth` | no | range 100–8192 |
| `parameters.iops` | no | API: The maximum I/O operations per second (IOPS) for this volume. |
| `parameters.profile` (block) | yes | nested block |
| `parameters.profile.href` | no | API: The URL for this volume |
| `parameters.profile.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_volume_job_invalid_job_type` | `job_type` |
| `ibm_is_volume_job_invalid_limit` | `limit` |
| `ibm_is_volume_job_invalid_name` | `name` |
| `ibm_is_volume_job_invalid_start` | `start` |
| `ibm_is_volume_job_invalid_volume_id` | `volume_id` |
| `ibm_is_volume_job_invalid_parameters__bandwidth` | `parameters.bandwidth` |
| `ibm_is_volume_job_invalid_parameters__iops` | `parameters.iops` |
| `ibm_is_volume_job_invalid_parameters__profile__href` | `parameters.profile.href` |
| `ibm_is_volume_job_invalid_parameters__profile__name` | `parameters.profile.name` |

## Example

### Valid

```hcl
resource "ibm_is_volume_job" "test" {
  job_type = "migrate"
  volume_id = "example-name"
}
```
