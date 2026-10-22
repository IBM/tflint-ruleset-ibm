# ibm_vmaas_vdc

Manages ibm_vmaas_vdc. A human readable ID for the virtual data center (VDC).

This rule validates `ibm_vmaas_vdc` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VDCPrototype` (`vmware-service`).

## Configuration

```hcl
rule "ibm_vmaas_vdc" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | pattern `^[A-Za-z0-9-,;=\.\*\s]{1,256}$` |
| `cpu` | no | range 0–2000 |
| `crn` | no | — |
| `deleted_at` | no | — |
| `edges` | no | — |
| `fast_provisioning_enabled` | no | API: Indicates whether to enable or not fast provisioning. |
| `ha` | no | — |
| `href` | no | — |
| `name` | yes | pattern `^[a-zA-Z][A-Za-z0-9_\-]{1,128}$` |
| `ordered_at` | no | — |
| `org_href` | no | — |
| `org_name` | no | — |
| `provisioned_at` | no | — |
| `ram` | no | range 0–40960 |
| `rhel_byol` | no | API: Indicates if the RHEL VMs will be using the license from IBM or the customer will use their own license (BYOL). |
| `status` | no | — |
| `status_reasons` | no | — |
| `type` | no | — |
| `windows_byol` | no | API: Indicates if the Microsoft Windows VMs will be using the license from IBM or the customer will use their own license (BYOL). |
| `director_site` (block) | yes | nested block |
| `director_site.id` | yes | pattern `^[A-Za-z0-9_-]{1,128}$` |
| `director_site.pvdc` (block) | yes | nested block |
| `director_site.pvdc.compute_ha_enabled` | no | API: Specifies whether compute HA is enabled for this VDC. |
| `director_site.pvdc.id` | yes | pattern `^[A-Za-z0-9_-]{1,128}$` |
| `director_site.pvdc.provider_type` (block) | no | nested block |
| `director_site.pvdc.provider_type.name` | yes | one of: `paygo`, `on_demand`, `reserved` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_vmaas_vdc_invalid_accept_language` | `accept_language` |
| `ibm_vmaas_vdc_invalid_cpu` | `cpu` |
| `ibm_vmaas_vdc_invalid_fast_provisioning_enabled` | `fast_provisioning_enabled` |
| `ibm_vmaas_vdc_invalid_name` | `name` |
| `ibm_vmaas_vdc_invalid_ram` | `ram` |
| `ibm_vmaas_vdc_invalid_rhel_byol` | `rhel_byol` |
| `ibm_vmaas_vdc_invalid_windows_byol` | `windows_byol` |
| `ibm_vmaas_vdc_invalid_director_site__id` | `director_site.id` |
| `ibm_vmaas_vdc_invalid_director_site__pvdc__compute_ha_enabled` | `director_site.pvdc.compute_ha_enabled` |
| `ibm_vmaas_vdc_invalid_director_site__pvdc__id` | `director_site.pvdc.id` |
| `ibm_vmaas_vdc_invalid_director_site__pvdc__provider_type__name` | `director_site.pvdc.provider_type.name` |

## Example

### Valid

```hcl
resource "ibm_vmaas_vdc" "test" {
  name = "example-name"
  director_site {
    id = "example-name"
    pvdc {
      id = "example-name"
    }
  }
}
```
