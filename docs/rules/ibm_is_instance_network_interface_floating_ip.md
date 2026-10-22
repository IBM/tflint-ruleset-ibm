# ibm_is_instance_network_interface_floating_ip

Manages ibm_is_instance_network_interface_floating_ip. Name of the floating IP

This rule validates `ibm_is_instance_network_interface_floating_ip` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerDNSPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_network_interface_floating_ip" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `address` | no | — |
| `crn` | no | — |
| `floating_ip` | yes | — |
| `instance` | yes | API: Identifies a DNS instance by a unique property. |
| `name` | no | — |
| `network_interface` | yes | — |
| `status` | no | — |
| `target` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_network_interface_floating_ip_invalid_floating_ip` | `floating_ip` |
| `ibm_is_instance_network_interface_floating_ip_invalid_instance` | `instance` |
| `ibm_is_instance_network_interface_floating_ip_invalid_network_interface` | `network_interface` |

## Example

### Valid

```hcl
resource "ibm_is_instance_network_interface_floating_ip" "test" {
  floating_ip = "10.0.0.1"
  instance = "example"
  network_interface = "example"
}
```
