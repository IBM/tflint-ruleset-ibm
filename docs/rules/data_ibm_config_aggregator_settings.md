# ibm_config_aggregator_settings_data

Schema validation for the ibm_config_aggregator_settings data source.

This rule validates `ibm_config_aggregator_settings` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_config_aggregator_settings_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `additional_scope` | no | — |
| `last_updated` | no | — |
| `regions` | no | — |
| `resource_collection_enabled` | no | — |
| `trusted_profile_id` | no | — |

## Example

### Valid

```hcl
data "ibm_config_aggregator_settings" "test" {
}
```
