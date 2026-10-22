# ibm_cm_preset_data

Schema validation for the ibm_cm_preset data source.

This rule validates `ibm_cm_preset` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ProviderInfo` (`resource-catalog-private-catalog`).

## Configuration

```hcl
rule "ibm_cm_preset_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `id` | yes | API: The id of this provider |
| `preset` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cm_preset_data_invalid_id` | `id` |

## Example

### Valid

```hcl
data "ibm_cm_preset" "test" {
  id = "example"
}
```
