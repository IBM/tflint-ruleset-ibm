# ibm_pi_datacenter_data

Schema validation for the ibm_pi_datacenter data source.

This rule validates `ibm_pi_datacenter` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_datacenter_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `capability_details` | no | — |
| `pi_cloud_instance_id` | no | — |
| `pi_datacenter_capabilities` | no | — |
| `pi_datacenter_href` | no | — |
| `pi_datacenter_location` | no | — |
| `pi_datacenter_status` | no | — |
| `pi_datacenter_type` | no | — |
| `pi_datacenter_zone` | no | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_datacenter_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |
| `ibm_pi_datacenter_data_invalid_pi_datacenter_zone` | `pi_datacenter_zone` |

## Example

### Valid

```hcl
data "ibm_pi_datacenter" "test" {
}
```
