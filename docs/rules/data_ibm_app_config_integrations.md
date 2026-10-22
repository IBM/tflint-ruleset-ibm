# ibm_app_config_integrations_data

Schema validation for the ibm_app_config_integrations data source.

This rule validates `ibm_app_config_integrations` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Config` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_integrations_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `first` | no | — |
| `guid` | yes | — |
| `integrations` | no | — |
| `last` | no | — |
| `limit` | no | — |
| `next` | no | — |
| `offset` | no | API: The number of records that are skipped in a list. |
| `previous` | no | — |
| `total_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_integrations_data_invalid_guid` | `guid` |
| `ibm_app_config_integrations_data_invalid_limit` | `limit` |
| `ibm_app_config_integrations_data_invalid_offset` | `offset` |

## Example

### Valid

```hcl
data "ibm_app_config_integrations" "test" {
  guid = "example"
}
```
