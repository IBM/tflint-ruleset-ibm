# ibm_is_instance_network_interfaces_data

Schema validation for the ibm_is_instance_network_interfaces data source.

This rule validates `ibm_is_instance_network_interfaces` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_instance_network_interfaces_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_name` | yes | — |
| `network_interfaces` | no | — |
| `total_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_network_interfaces_data_invalid_instance_name` | `instance_name` |

## Example

### Valid

```hcl
data "ibm_is_instance_network_interfaces" "test" {
  instance_name = "example"
}
```
