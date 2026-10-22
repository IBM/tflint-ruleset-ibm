# ibm_tg_location_data

Reads ibm_tg_location. Name of the Location.

This rule validates `ibm_tg_location` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `RedundancyGroupPatch` (`transit-gateway`).

## Configuration

```hcl
rule "ibm_tg_location_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `billing_location` | no | — |
| `local_connection_locations` | no | — |
| `name` | yes | pattern `^([a-zA-Z]|[a-zA-Z][-_a-zA-Z0-9]*[a-zA-Z0-9])$` |
| `type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_tg_location_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_tg_location" "test" {
  name = "example-name"
}
```
