# ibm_kms_instance_policies_data

Schema validation for the ibm_kms_instance_policies data source.

This rule validates `ibm_kms_instance_policies` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `InstanceConnectionRead` (`hs-crypto`).

## Configuration

```hcl
rule "ibm_kms_instance_policies_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `dual_auth_delete` | no | — |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `key_create_import_access` | no | — |
| `metrics` | no | — |
| `policy_type` | no | — |
| `rotation` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_kms_instance_policies_data_invalid_instance_id` | `instance_id` |
| `ibm_kms_instance_policies_data_invalid_policy_type` | `policy_type` |

## Example

### Valid

```hcl
data "ibm_kms_instance_policies" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
