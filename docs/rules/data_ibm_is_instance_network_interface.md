# ibm_is_instance_network_interface_data

Reads ibm_is_instance_network_interface. The user-defined name for this network interface.

This rule validates `ibm_is_instance_network_interface` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_instance_network_interface_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allow_ip_spoofing` | no | — |
| `created_at` | no | — |
| `floating_ips` | no | — |
| `href` | no | — |
| `instance_name` | yes | — |
| `name` | no | — |
| `network_interface_name` | yes | — |
| `port_speed` | no | — |
| `primary_ip` | no | — |
| `primary_ipv4_address` | no | — |
| `resource_type` | no | — |
| `security_groups` | no | — |
| `status` | no | — |
| `subnet` | no | — |
| `type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_network_interface_data_invalid_instance_name` | `instance_name` |
| `ibm_is_instance_network_interface_data_invalid_network_interface_name` | `network_interface_name` |

## Example

### Valid

```hcl
data "ibm_is_instance_network_interface" "test" {
  instance_name = "example"
  network_interface_name = "example"
}
```
