# ibm_pi_vpn_connection

Schema validation for the ibm_pi_vpn_connection resource.

This rule validates `ibm_pi_vpn_connection` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_vpn_connection" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `connection_id` | no | — |
| `connection_status` | no | — |
| `dead_peer_detections` | no | — |
| `gateway_address` | no | — |
| `local_gateway_address` | no | — |

## Example

### Valid

```hcl
resource "ibm_pi_vpn_connection" "test" {
}
```
