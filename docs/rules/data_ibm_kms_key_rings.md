# ibm_kms_key_rings_data

Schema validation for the ibm_kms_key_rings data source.

This rule validates `ibm_kms_key_rings` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `InstanceConnectionRead` (`hs-crypto`).

## Configuration

```hcl
rule "ibm_kms_key_rings_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `endpoint_type` | no | one of: `public`, `private` |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `key_rings` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_kms_key_rings_data_invalid_endpoint_type` | `endpoint_type` |
| `ibm_kms_key_rings_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_kms_key_rings" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
