# ibm_is_subnet_reserved_ip_data

Reads ibm_is_subnet_reserved_ip. The user-defined or system-provided name for this reserved IP.

This rule validates `ibm_is_subnet_reserved_ip` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ClusterNetworkInterfacePrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_subnet_reserved_ip_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `address` | no | — |
| `auto_delete` | no | — |
| `created_at` | no | — |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | — |
| `owner` | no | — |
| `reserved_ip` | yes | — |
| `resource_type` | no | — |
| `subnet` | yes | API: Identifies a cluster network subnet by a unique property. |
| `target` | no | — |
| `target_crn` | no | — |
| `target_reference` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_subnet_reserved_ip_data_invalid_reserved_ip` | `reserved_ip` |
| `ibm_is_subnet_reserved_ip_data_invalid_subnet` | `subnet` |

## Example

### Valid

```hcl
data "ibm_is_subnet_reserved_ip" "test" {
  reserved_ip = "10.0.0.1"
  subnet = "example"
}
```
