# ibm_is_flow_logs_data

Reads ibm_is_flow_logs. The name of the flow log

This rule validates `ibm_is_flow_logs` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `NetworkACLPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_flow_logs_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `flow_log_collectors` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `target` | no | — |
| `target_resource_type` | no | — |
| `vpc` | no | API: Identifies a VPC by a unique property. |
| `vpc_crn` | no | crn format |
| `vpc_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_flow_logs_data_invalid_name` | `name` |
| `ibm_is_flow_logs_data_invalid_resource_group` | `resource_group` |
| `ibm_is_flow_logs_data_invalid_target` | `target` |
| `ibm_is_flow_logs_data_invalid_target_resource_type` | `target_resource_type` |
| `ibm_is_flow_logs_data_invalid_vpc` | `vpc` |
| `ibm_is_flow_logs_data_invalid_vpc_crn` | `vpc_crn` |
| `ibm_is_flow_logs_data_invalid_vpc_name` | `vpc_name` |

## Example

### Valid

```hcl
data "ibm_is_flow_logs" "test" {
}
```
