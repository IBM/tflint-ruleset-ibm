# ibm_is_ike_policy

Manages ibm_is_ike_policy. IKE name

This rule validates `ibm_is_ike_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `IKEPolicyPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_ike_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `authentication_algorithm` | no | one of: `md5`, `sha1`, `sha256`, `sha512`, `sha384`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `authentication_algorithms` | no | one of: `sha256`, `sha384`, `sha512` |
| `dh_group` | no | one of: `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `22`, `23`, `24`, `31` |
| `dh_groups` | no | one of: `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `22`, `23`, `24`, `31` |
| `encryption_algorithm` | no | one of: `triple_des`, `aes128`, `aes192`, `aes256`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `encryption_algorithms` | no | one of: `aes128`, `aes192`, `aes256` |
| `href` | no | — |
| `ike_version` | no | one of: `1`, `2` |
| `key_lifetime` | no | range 1800–86400 |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `negotiation_mode` | no | — |
| `resource_controller_url` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `vpn_connections` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_ike_policy_invalid_authentication_algorithm` | `authentication_algorithm` |
| `ibm_is_ike_policy_invalid_authentication_algorithms` | `authentication_algorithms` |
| `ibm_is_ike_policy_invalid_dh_group` | `dh_group` |
| `ibm_is_ike_policy_invalid_dh_groups` | `dh_groups` |
| `ibm_is_ike_policy_invalid_encryption_algorithm` | `encryption_algorithm` |
| `ibm_is_ike_policy_invalid_encryption_algorithms` | `encryption_algorithms` |
| `ibm_is_ike_policy_invalid_ike_version` | `ike_version` |
| `ibm_is_ike_policy_invalid_key_lifetime` | `key_lifetime` |
| `ibm_is_ike_policy_invalid_name` | `name` |
| `ibm_is_ike_policy_invalid_resource_group` | `resource_group` |

## Example

### Valid

```hcl
resource "ibm_is_ike_policy" "test" {
  name = "example-name"
}
```
