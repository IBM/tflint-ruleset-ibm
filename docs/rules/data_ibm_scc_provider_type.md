# ibm_scc_provider_type_data

Reads ibm_scc_provider_type. The name of the provider type.

This rule validates `ibm_scc_provider_type` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_scc_provider_type_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `attributes` | no | — |
| `created_at` | no | — |
| `data_type` | no | — |
| `description` | no | — |
| `icon` | no | — |
| `id` | no | — |
| `instance_id` | yes | — |
| `instance_limit` | no | — |
| `label` | no | — |
| `mode` | no | — |
| `name` | no | — |
| `provider_type_id` | yes | — |
| `s2s_enabled` | no | — |
| `type` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_provider_type_data_invalid_instance_id` | `instance_id` |
| `ibm_scc_provider_type_data_invalid_provider_type_id` | `provider_type_id` |

## Example

### Valid

```hcl
data "ibm_scc_provider_type" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  provider_type_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
