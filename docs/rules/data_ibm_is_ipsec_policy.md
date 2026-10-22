# ibm_is_ipsec_policy_data

Reads ibm_is_ipsec_policy. The IPsec policy name.

This rule validates `ibm_is_ipsec_policy` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_ipsec_policy_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `authentication_algorithm` | no | — |
| `authentication_algorithms` | no | — |
| `connections` | no | — |
| `created_at` | no | — |
| `encapsulation_mode` | no | — |
| `encryption_algorithm` | no | — |
| `encryption_algorithms` | no | — |
| `href` | no | — |
| `ipsec_policy` | no | — |
| `key_lifetime` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `pfs` | no | — |
| `pfs_groups` | no | — |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `transform_protocol` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_ipsec_policy_data_invalid_ipsec_policy` | `ipsec_policy` |
| `ibm_is_ipsec_policy_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_ipsec_policy" "test" {
}
```
