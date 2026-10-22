# ibm_pi_sap_profiles_data

Schema validation for the ibm_pi_sap_profiles data source.

This rule validates `ibm_pi_sap_profiles` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_sap_profiles_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `pi_cloud_instance_id` | yes | — |
| `pi_family_filter` | no | one of: `balanced`, `compute`, `memory`, `sap-rise`, `sap-rise-app`, `small`, `ultra-memory` |
| `pi_prefix_filter` | no | one of: `bh1`, `bh2`, `ch1`, `ch2`, `mh1`, `mh2`, `umh`, `ush1`, `sh2`, `sr2`, `sr3` |
| `profiles` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_sap_profiles_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_sap_profiles_data_invalid_pi_family_filter` | `pi_family_filter` |
| `ibm_pi_sap_profiles_data_invalid_pi_prefix_filter` | `pi_prefix_filter` |

## Example

### Valid

```hcl
data "ibm_pi_sap_profiles" "test" {
  pi_cloud_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
