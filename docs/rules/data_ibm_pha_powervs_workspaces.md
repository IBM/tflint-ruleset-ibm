# ibm_pha_powervs_workspaces_data

Schema validation for the ibm_pha_powervs_workspaces data source.

This rule validates `ibm_pha_powervs_workspaces` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pha_powervs_workspaces_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | — |
| `if_none_match` | no | — |
| `instance_id` | yes | — |
| `location_id` | yes | — |
| `workspaces` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pha_powervs_workspaces_data_invalid_accept_language` | `accept_language` |
| `ibm_pha_powervs_workspaces_data_invalid_if_none_match` | `if_none_match` |
| `ibm_pha_powervs_workspaces_data_invalid_instance_id` | `instance_id` |
| `ibm_pha_powervs_workspaces_data_invalid_location_id` | `location_id` |

## Example

### Valid

```hcl
data "ibm_pha_powervs_workspaces" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  location_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
