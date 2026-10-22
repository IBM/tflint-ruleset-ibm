# ibm_cis_logpush_job

Manages ibm_cis_logpush_job. Logpush Job Name

This rule validates `ibm_cis_logpush_job` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `logpush_job_cos_req` (`cis`).

## Configuration

```hcl
rule "ibm_cis_logpush_job" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `cos` | no | API: Information to identify the COS bucket where the data will be pushed. |
| `dataset` | yes | one of: `http_requests`, `range_events`, `firewall_events`, `dns_logs` |
| `destination_conf` | no | — |
| `domain_id` | yes | — |
| `enabled` | no | API: Indicates whether the logpush job is enabled or not. |
| `error_message` | no | — |
| `frequency` | no | one of: `high`, `low` |
| `job_id` | no | — |
| `last_complete` | no | — |
| `last_error` | no | — |
| `logdna` | no | — |
| `logpull_options` | no | API: Configuration string |
| `name` | no | API: Logpush job name |
| `ownership_challenge` | no | API: Ownership challenge token to prove destination ownership. |
| `ibmcl` (block) | no | nested block |
| `ibmcl.api_key` | no | — |
| `ibmcl.instance_id` | yes | — |
| `ibmcl.region` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_logpush_job_invalid_cis_id` | `cis_id` |
| `ibm_cis_logpush_job_invalid_cos` | `cos` |
| `ibm_cis_logpush_job_invalid_dataset` | `dataset` |
| `ibm_cis_logpush_job_invalid_destination_conf` | `destination_conf` |
| `ibm_cis_logpush_job_invalid_domain_id` | `domain_id` |
| `ibm_cis_logpush_job_invalid_enabled` | `enabled` |
| `ibm_cis_logpush_job_invalid_frequency` | `frequency` |
| `ibm_cis_logpush_job_invalid_logdna` | `logdna` |
| `ibm_cis_logpush_job_invalid_logpull_options` | `logpull_options` |
| `ibm_cis_logpush_job_invalid_name` | `name` |
| `ibm_cis_logpush_job_invalid_ownership_challenge` | `ownership_challenge` |
| `ibm_cis_logpush_job_invalid_ibmcl__api_key` | `ibmcl.api_key` |
| `ibm_cis_logpush_job_invalid_ibmcl__instance_id` | `ibmcl.instance_id` |
| `ibm_cis_logpush_job_invalid_ibmcl__region` | `ibmcl.region` |

## Example

### Valid

```hcl
resource "ibm_cis_logpush_job" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  dataset = "http_requests"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
