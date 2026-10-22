# ibm_ipsec_vpn

Schema validation for the ibm_ipsec_vpn resource.

This rule validates `ibm_ipsec_vpn` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_ipsec_vpn" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `customer_peer_ip` | no | — |
| `datacenter` | yes | — |
| `internal_peer_ip_address` | no | — |
| `internal_subnet_id` | no | — |
| `name` | no | — |
| `preshared_key` | no | — |
| `remote_subnet_id` | no | — |
| `service_subnet_id` | no | — |
| `address_translation` (block) | no | nested block |
| `address_translation.internal_ip_adress` | yes | — |
| `address_translation.notes` | no | — |
| `address_translation.remote_ip_adress` | yes | — |
| `phase_one` (block) | no | nested block |
| `phase_one.authentication` | no | — |
| `phase_one.diffie_hellman_group` | no | — |
| `phase_one.encryption` | no | — |
| `phase_one.keylife` | no | — |
| `phase_two` (block) | no | nested block |
| `phase_two.authentication` | no | — |
| `phase_two.diffie_hellman_group` | no | — |
| `phase_two.encryption` | no | — |
| `phase_two.keylife` | no | — |
| `remote_subnet` (block) | no | nested block |
| `remote_subnet.account_id` | no | — |
| `remote_subnet.remote_ip_adress` | yes | — |
| `remote_subnet.remote_ip_cidr` | yes | cidr format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_ipsec_vpn_invalid_customer_peer_ip` | `customer_peer_ip` |
| `ibm_ipsec_vpn_invalid_datacenter` | `datacenter` |
| `ibm_ipsec_vpn_invalid_internal_subnet_id` | `internal_subnet_id` |
| `ibm_ipsec_vpn_invalid_preshared_key` | `preshared_key` |
| `ibm_ipsec_vpn_invalid_remote_subnet_id` | `remote_subnet_id` |
| `ibm_ipsec_vpn_invalid_service_subnet_id` | `service_subnet_id` |
| `ibm_ipsec_vpn_invalid_address_translation__internal_ip_adress` | `address_translation.internal_ip_adress` |
| `ibm_ipsec_vpn_invalid_address_translation__notes` | `address_translation.notes` |
| `ibm_ipsec_vpn_invalid_address_translation__remote_ip_adress` | `address_translation.remote_ip_adress` |
| `ibm_ipsec_vpn_invalid_phase_one__authentication` | `phase_one.authentication` |
| `ibm_ipsec_vpn_invalid_phase_one__diffie_hellman_group` | `phase_one.diffie_hellman_group` |
| `ibm_ipsec_vpn_invalid_phase_one__encryption` | `phase_one.encryption` |
| `ibm_ipsec_vpn_invalid_phase_one__keylife` | `phase_one.keylife` |
| `ibm_ipsec_vpn_invalid_phase_two__authentication` | `phase_two.authentication` |
| `ibm_ipsec_vpn_invalid_phase_two__diffie_hellman_group` | `phase_two.diffie_hellman_group` |
| `ibm_ipsec_vpn_invalid_phase_two__encryption` | `phase_two.encryption` |
| `ibm_ipsec_vpn_invalid_phase_two__keylife` | `phase_two.keylife` |
| `ibm_ipsec_vpn_invalid_remote_subnet__account_id` | `remote_subnet.account_id` |
| `ibm_ipsec_vpn_invalid_remote_subnet__remote_ip_adress` | `remote_subnet.remote_ip_adress` |
| `ibm_ipsec_vpn_invalid_remote_subnet__remote_ip_cidr` | `remote_subnet.remote_ip_cidr` |

## Example

### Valid

```hcl
resource "ibm_ipsec_vpn" "test" {
  datacenter = "example"
}
```
