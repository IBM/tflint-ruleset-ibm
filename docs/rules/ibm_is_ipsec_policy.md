# ibm_is_ipsec_policy

Manages ibm_is_ipsec_policy. IPSEC name

This rule validates `ibm_is_ipsec_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `IPsecPolicyPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_ipsec_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `authentication_algorithm` | no | one of: `md5`, `sha1`, `sha256`, `sha512`, `sha384`, `disabled`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `authentication_algorithms` | no | one of: `disabled`, `sha256`, `sha384`, `sha512` |
| `encapsulation_mode` | no | — |
| `encryption_algorithm` | no | one of: `triple_des`, `aes128`, `aes256`, `aes128gcm16`, `aes192gcm16`, `aes256gcm16`, `aes192`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `encryption_algorithms` | no | one of: `aes128`, `aes128gcm16`, `aes192`, `aes192gcm16`, `aes256`, `aes256gcm16` |
| `key_lifetime` | no | range 1800–86400 |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `pfs` | no | one of: `disabled`, `group_2`, `group_5`, `group_14`, `group_19`, `group_15`, `group_16`, `group_17`, `group_18`, `group_20`, `group_21`, `group_22`, `group_23`, `group_24`, `group_31`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `pfs_groups` | no | one of: `disabled`, `group_14`, `group_15`, `group_16`, `group_17`, `group_18`, `group_19`, `group_20`, `group_21`, `group_22`, `group_23`, `group_24`, `group_31` |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `transform_protocol` | no | — |
| `vpn_connections` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_ipsec_policy_invalid_authentication_algorithm` | `authentication_algorithm` |
| `ibm_is_ipsec_policy_invalid_authentication_algorithms` | `authentication_algorithms` |
| `ibm_is_ipsec_policy_invalid_encryption_algorithm` | `encryption_algorithm` |
| `ibm_is_ipsec_policy_invalid_encryption_algorithms` | `encryption_algorithms` |
| `ibm_is_ipsec_policy_invalid_key_lifetime` | `key_lifetime` |
| `ibm_is_ipsec_policy_invalid_name` | `name` |
| `ibm_is_ipsec_policy_invalid_pfs` | `pfs` |
| `ibm_is_ipsec_policy_invalid_pfs_groups` | `pfs_groups` |
| `ibm_is_ipsec_policy_invalid_resource_group` | `resource_group` |

## Example

### Valid

```hcl
resource "ibm_is_ipsec_policy" "test" {
  name = "example-name"
}
```
