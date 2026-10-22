# ibm_scc_provider_types_data

Schema validation for the ibm_scc_provider_types data source.

This rule validates `ibm_scc_provider_types` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_scc_provider_types_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_id` | yes | — |
| `provider_types` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_provider_types_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_scc_provider_types" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
