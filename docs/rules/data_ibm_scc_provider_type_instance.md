# ibm_scc_provider_type_instance_data

Reads ibm_scc_provider_type_instance. The name of the provider type instance.

This rule validates `ibm_scc_provider_type_instance` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_scc_provider_type_instance_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `attributes` | no | — |
| `created_at` | no | — |
| `instance_id` | yes | — |
| `name` | no | — |
| `provider_type_id` | yes | pattern `^[a-zA-Z0-9 ,\-_]+$` |
| `provider_type_instance_id` | yes | — |
| `provider_type_instance_item_id` | no | — |
| `type` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_provider_type_instance_data_invalid_instance_id` | `instance_id` |
| `ibm_scc_provider_type_instance_data_invalid_provider_type_id` | `provider_type_id` |
| `ibm_scc_provider_type_instance_data_invalid_provider_type_instance_id` | `provider_type_instance_id` |

## Example

### Valid

```hcl
data "ibm_scc_provider_type_instance" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  provider_type_id = "example-namexxxxxxxxxxxxxxxxxxxx"
  provider_type_instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
