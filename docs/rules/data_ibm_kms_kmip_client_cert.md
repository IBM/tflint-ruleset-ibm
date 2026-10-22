# ibm_kms_kmip_client_cert_data

Reads ibm_kms_kmip_client_cert. The name of the KMIP Client Certificate to be fetched

This rule validates `ibm_kms_kmip_client_cert` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `InstanceConnectionRead` (`hs-crypto`).

## Configuration

```hcl
rule "ibm_kms_kmip_client_cert_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `adapter_id` | no | — |
| `adapter_name` | no | — |
| `cert_id` | no | — |
| `certificate` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `endpoint_type` | no | one of: `public`, `private` |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_kms_kmip_client_cert_data_invalid_adapter_id` | `adapter_id` |
| `ibm_kms_kmip_client_cert_data_invalid_adapter_name` | `adapter_name` |
| `ibm_kms_kmip_client_cert_data_invalid_endpoint_type` | `endpoint_type` |
| `ibm_kms_kmip_client_cert_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_kms_kmip_client_cert" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
