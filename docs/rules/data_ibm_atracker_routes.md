# ibm_atracker_routes_data

Reads ibm_atracker_routes. The name of the route.

This rule validates `ibm_atracker_routes` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `RoutePrototype` (`atracker`).

## Configuration

```hcl
rule "ibm_atracker_routes_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | no | pattern `^[a-zA-Z0-9 -._:]+$` |
| `routes` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_atracker_routes_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_atracker_routes" "test" {
}
```
