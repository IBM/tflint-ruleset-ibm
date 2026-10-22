# ibm_is_instance_disks_data

Schema validation for the ibm_is_instance_disks data source.

This rule validates `ibm_is_instance_disks` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerDNSPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_disks_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `disks` | no | — |
| `instance` | yes | API: Identifies a DNS instance by a unique property. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_disks_data_invalid_instance` | `instance` |

## Example

### Valid

```hcl
data "ibm_is_instance_disks" "test" {
  instance = "example"
}
```
