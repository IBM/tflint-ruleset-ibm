# ibm_cis_logpush_jobs_data

Schema validation for the ibm_cis_logpush_jobs data source.

This rule validates `ibm_cis_logpush_jobs` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/v2/{crn}/zones/{zone_id}/logpush/jobs/{job_id}` (`cis`).

## Configuration

```hcl
rule "ibm_cis_logpush_jobs_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `job_id` | no | — |
| `logpush_job_pack` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_logpush_jobs_data_invalid_cis_id` | `cis_id` |
| `ibm_cis_logpush_jobs_data_invalid_domain_id` | `domain_id` |
| `ibm_cis_logpush_jobs_data_invalid_job_id` | `job_id` |

## Example

### Valid

```hcl
data "ibm_cis_logpush_jobs" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
