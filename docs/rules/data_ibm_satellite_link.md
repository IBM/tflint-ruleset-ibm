# ibm_satellite_link_data

Schema validation for the ibm_satellite_link data source.

This rule validates `ibm_satellite_link` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_link_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `crn` | no | — |
| `description` | no | — |
| `last_change` | no | — |
| `location` | yes | — |
| `performance` | no | — |
| `satellite_link_host` | no | — |
| `status` | no | — |
| `ws_endpoint` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_link_data_invalid_location` | `location` |

## Example

### Valid

```hcl
data "ibm_satellite_link" "test" {
  location = "example"
}
```
