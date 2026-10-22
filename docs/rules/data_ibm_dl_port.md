# ibm_dl_port_data

Schema validation for the ibm_dl_port data source.

This rule validates `ibm_dl_port` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_dl_port_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `direct_link_count` | no | — |
| `label` | no | — |
| `location_display_name` | no | — |
| `location_name` | no | — |
| `port_id` | yes | — |
| `provider_name` | no | — |
| `supported_link_speeds` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dl_port_data_invalid_port_id` | `port_id` |

## Example

### Valid

```hcl
data "ibm_dl_port" "test" {
  port_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
