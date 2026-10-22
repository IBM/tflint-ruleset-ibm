# ibm_hardware_firewall_shared

Schema validation for the ibm_hardware_firewall_shared resource.

This rule validates `ibm_hardware_firewall_shared` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_hardware_firewall_shared" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `billing_item_id` | no | — |
| `firewall_type` | yes | one of: `10MBPS_HARDWARE_FIREWALL`, `20MBPS_HARDWARE_FIREWALL`, `100MBPS_HARDWARE_FIREWALL`, `1000MBPS_HARDWARE_FIREWALL`, `200MBPS_HARDWARE_FIREWALL`, `2000MBPS_HARDWARE_FIREWALL` |
| `hardware_instance_id` | no | — |
| `virtual_instance_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_hardware_firewall_shared_invalid_firewall_type` | `firewall_type` |
| `ibm_hardware_firewall_shared_invalid_hardware_instance_id` | `hardware_instance_id` |
| `ibm_hardware_firewall_shared_invalid_virtual_instance_id` | `virtual_instance_id` |

## Example

### Valid

```hcl
resource "ibm_hardware_firewall_shared" "test" {
  firewall_type = "10MBPS_HARDWARE_FIREWALL"
}
```
