# ibm_network_interface_sg_attachment

Schema validation for the ibm_network_interface_sg_attachment resource.

This rule validates `ibm_network_interface_sg_attachment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_network_interface_sg_attachment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `network_interface_id` | yes | — |
| `security_group_id` | yes | — |
| `soft_reboot` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_network_interface_sg_attachment_invalid_network_interface_id` | `network_interface_id` |
| `ibm_network_interface_sg_attachment_invalid_security_group_id` | `security_group_id` |
| `ibm_network_interface_sg_attachment_invalid_soft_reboot` | `soft_reboot` |

## Example

### Valid

```hcl
resource "ibm_network_interface_sg_attachment" "test" {
  network_interface_id = 1
  security_group_id = 1
}
```
