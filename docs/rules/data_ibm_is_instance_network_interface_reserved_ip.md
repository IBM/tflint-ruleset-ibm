# ibm_is_instance_network_interface_reserved_ip_data

Reads ibm_is_instance_network_interface_reserved_ip. The user-defined or system-provided name for this reserved IP.

This rule validates `ibm_is_instance_network_interface_reserved_ip` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerDNSPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_network_interface_reserved_ip_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `address` | no | — |
| `auto_delete` | no | — |
| `created_at` | no | — |
| `href` | no | — |
| `instance` | yes | API: Identifies a DNS instance by a unique property. |
| `name` | no | — |
| `network_interface` | yes | — |
| `owner` | no | — |
| `reserved_ip` | yes | — |
| `resource_type` | no | — |
| `target` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_network_interface_reserved_ip_data_invalid_instance` | `instance` |
| `ibm_is_instance_network_interface_reserved_ip_data_invalid_network_interface` | `network_interface` |
| `ibm_is_instance_network_interface_reserved_ip_data_invalid_reserved_ip` | `reserved_ip` |

## Example

### Valid

```hcl
data "ibm_is_instance_network_interface_reserved_ip" "test" {
  instance = "example"
  network_interface = "example"
  reserved_ip = "10.0.0.1"
}
```
