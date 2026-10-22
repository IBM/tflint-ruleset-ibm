# ibm_is_dedicated_host_disks_data

Schema validation for the ibm_is_dedicated_host_disks data source.

This rule validates `ibm_is_dedicated_host_disks` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_dedicated_host_disks_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `dedicated_host` | yes | — |
| `disks` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_dedicated_host_disks_data_invalid_dedicated_host` | `dedicated_host` |

## Example

### Valid

```hcl
data "ibm_is_dedicated_host_disks" "test" {
  dedicated_host = "example"
}
```
