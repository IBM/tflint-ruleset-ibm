# ibm_config_aggregator_resource_collection_status_data

Schema validation for the ibm_config_aggregator_resource_collection_status data source.

This rule validates `ibm_config_aggregator_resource_collection_status` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_config_aggregator_resource_collection_status_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `last_config_refresh_time` | no | — |
| `status` | no | — |

## Example

### Valid

```hcl
data "ibm_config_aggregator_resource_collection_status" "test" {
}
```
