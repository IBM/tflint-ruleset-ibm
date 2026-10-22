# ibm_is_vpc_routing_tables_data

Schema validation for the ibm_is_vpc_routing_tables data source.

This rule validates `ibm_is_vpc_routing_tables` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ShareMountTargetByAccessControlModeVPC` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpc_routing_tables_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `is_default` | no | — |
| `routing_tables` | no | — |
| `vpc` | yes | API: Identifies a VPC by a unique property. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpc_routing_tables_data_invalid_is_default` | `is_default` |
| `ibm_is_vpc_routing_tables_data_invalid_vpc` | `vpc` |

## Example

### Valid

```hcl
data "ibm_is_vpc_routing_tables" "test" {
  vpc = "example"
}
```
