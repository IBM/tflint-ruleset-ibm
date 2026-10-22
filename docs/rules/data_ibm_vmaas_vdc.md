# ibm_vmaas_vdc_data

Reads ibm_vmaas_vdc. A human readable ID for the virtual data center (VDC).

This rule validates `ibm_vmaas_vdc` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_vmaas_vdc_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | pattern `^[A-Za-z0-9-,;=\.\*\s]{1,256}$` |
| `cpu` | no | — |
| `crn` | no | — |
| `deleted_at` | no | — |
| `director_site` | no | — |
| `edges` | no | — |
| `fast_provisioning_enabled` | no | — |
| `ha` | no | — |
| `href` | no | — |
| `name` | no | — |
| `ordered_at` | no | — |
| `org_href` | no | — |
| `org_name` | no | — |
| `provisioned_at` | no | — |
| `ram` | no | — |
| `rhel_byol` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `type` | no | — |
| `vmaas_vdc_id` | yes | — |
| `windows_byol` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_vmaas_vdc_data_invalid_accept_language` | `accept_language` |
| `ibm_vmaas_vdc_data_invalid_vmaas_vdc_id` | `vmaas_vdc_id` |

## Example

### Valid

```hcl
data "ibm_vmaas_vdc" "test" {
  vmaas_vdc_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
