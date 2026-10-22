# ibm_is_instance_network_attachment_data

Reads ibm_is_instance_network_attachment. The name for this instance network attachment.

This rule validates `ibm_is_instance_network_attachment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerDNSPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_network_attachment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `href` | no | — |
| `instance` | yes | pattern `^[-0-9a-z_]+$` |
| `lifecycle_state` | no | — |
| `name` | no | — |
| `network_attachment` | yes | — |
| `port_speed` | no | — |
| `primary_ip` | no | — |
| `resource_type` | no | — |
| `subnet` | no | — |
| `type` | no | — |
| `virtual_network_interface` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_network_attachment_data_invalid_instance` | `instance` |
| `ibm_is_instance_network_attachment_data_invalid_network_attachment` | `network_attachment` |

## Example

### Valid

```hcl
data "ibm_is_instance_network_attachment" "test" {
  instance = "example-name"
  network_attachment = "example"
}
```
