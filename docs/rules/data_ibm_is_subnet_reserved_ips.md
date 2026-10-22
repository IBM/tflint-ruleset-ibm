# ibm_is_subnet_reserved_ips_data

Schema validation for the ibm_is_subnet_reserved_ips data source.

This rule validates `ibm_is_subnet_reserved_ips` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ClusterNetworkInterfacePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_subnet_reserved_ips_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `reserved_ips` | no | — |
| `subnet` | yes | API: Identifies a cluster network subnet by a unique property. |
| `total_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_subnet_reserved_ips_data_invalid_subnet` | `subnet` |

## Example

### Valid

```hcl
data "ibm_is_subnet_reserved_ips" "test" {
  subnet = "example"
}
```
