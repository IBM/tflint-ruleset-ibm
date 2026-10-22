# ibm_db2_connection_info_data

Schema validation for the ibm_db2_connection_info data source.

This rule validates `ibm_db2_connection_info` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SuccessTakeoverDRResponse` (`db2-on-cloud-db2-on-cloud-v4`).

## Configuration

```hcl
rule "ibm_db2_connection_info_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `deployment_id` | yes | — |
| `private` | no | — |
| `public` | no | — |
| `x_deployment_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_db2_connection_info_data_invalid_deployment_id` | `deployment_id` |
| `ibm_db2_connection_info_data_invalid_x_deployment_id` | `x_deployment_id` |

## Example

### Valid

```hcl
data "ibm_db2_connection_info" "test" {
  deployment_id = "01234567-89ab-cdef-0123-456789abcdef"
  x_deployment_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
