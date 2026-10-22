# ibm_kms_keys_data

Schema validation for the ibm_kms_keys data source.

This rule validates `ibm_kms_keys` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/api/v2/keys/{id}:patch` (`key-protect`).

## Configuration

```hcl
rule "ibm_kms_keys_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `alias` | no | — |
| `endpoint_type` | no | one of: `public`, `private` |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `key_id` | no | — |
| `key_name` | no | — |
| `keys` | no | — |
| `limit` | no | range 1–1000 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_kms_keys_data_invalid_alias` | `alias` |
| `ibm_kms_keys_data_invalid_endpoint_type` | `endpoint_type` |
| `ibm_kms_keys_data_invalid_instance_id` | `instance_id` |
| `ibm_kms_keys_data_invalid_key_id` | `key_id` |
| `ibm_kms_keys_data_invalid_key_name` | `key_name` |
| `ibm_kms_keys_data_invalid_limit` | `limit` |

## Example

### Valid

```hcl
data "ibm_kms_keys" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
