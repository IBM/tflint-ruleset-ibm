# ibm_pi_sap_profile_data

Schema validation for the ibm_pi_sap_profile data source.

This rule validates `ibm_pi_sap_profile` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_sap_profile_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `asaps` | no | — |
| `certified` | no | — |
| `cores` | no | — |
| `default_system` | no | — |
| `deprecated` | no | — |
| `full_system_profile` | no | — |
| `memory` | no | — |
| `pi_cloud_instance_id` | yes | — |
| `pi_sap_profile_id` | yes | — |
| `saps` | no | — |
| `supported_systems` | no | — |
| `type` | no | — |
| `vpmem_volume` | no | — |
| `workload_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_sap_profile_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_sap_profile_data_invalid_pi_sap_profile_id` | `pi_sap_profile_id` |

## Example

### Valid

```hcl
data "ibm_pi_sap_profile" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  pi_sap_profile_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
