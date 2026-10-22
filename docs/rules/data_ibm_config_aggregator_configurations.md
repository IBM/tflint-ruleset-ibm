# ibm_config_aggregator_configurations_data

Schema validation for the ibm_config_aggregator_configurations data source.

This rule validates `ibm_config_aggregator_configurations` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `About` (`app-configuration`).

## Configuration

```hcl
rule "ibm_config_aggregator_configurations_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^[a-zA-Z0-9:-]+$` |
| `config_type` | no | pattern `^[a-zA-Z0-9.\:/-]+$` |
| `configs` | no | — |
| `location` | no | pattern `^$|^[a-z]+-[a-z]+$` |
| `prev` | no | — |
| `resource_crn` | no | crn format |
| `resource_group_id` | no | pattern `^[a-zA-Z0-9-]*$` |
| `service_name` | no | pattern `^[a-zA-Z0-9.\:/-]+$` |
| `service_tags` | no | pattern `^[a-zA-Z0-9:-]+$` |
| `sub_account` | no | pattern `^[a-zA-Z0-9]+$` |
| `user_tags` | no | pattern `^[a-zA-Z0-9:-]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_config_aggregator_configurations_data_invalid_access_tags` | `access_tags` |
| `ibm_config_aggregator_configurations_data_invalid_config_type` | `config_type` |
| `ibm_config_aggregator_configurations_data_invalid_location` | `location` |
| `ibm_config_aggregator_configurations_data_invalid_resource_crn` | `resource_crn` |
| `ibm_config_aggregator_configurations_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_config_aggregator_configurations_data_invalid_service_name` | `service_name` |
| `ibm_config_aggregator_configurations_data_invalid_service_tags` | `service_tags` |
| `ibm_config_aggregator_configurations_data_invalid_sub_account` | `sub_account` |
| `ibm_config_aggregator_configurations_data_invalid_user_tags` | `user_tags` |

## Example

### Valid

```hcl
data "ibm_config_aggregator_configurations" "test" {
}
```
