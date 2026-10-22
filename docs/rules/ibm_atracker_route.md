# ibm_atracker_route

Manages ibm_atracker_route. The name of the route.

This rule validates `ibm_atracker_route` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `RulePrototype` (`atracker`).

## Configuration

```hcl
rule "ibm_atracker_route" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `api_version` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `managed_by` | no | one of: `account`, `enterprise` |
| `message` | no | — |
| `name` | yes | pattern `^[a-zA-Z0-9 -._:]+$` |
| `updated_at` | no | — |
| `version` | no | — |
| `rules` (block) | yes | nested block |
| `rules.locations` | yes | pattern `^[a-zA-Z0-9 -._:]+$` |
| `rules.target_ids` | yes | pattern `^[a-zA-Z0-9 -._:]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_atracker_route_invalid_managed_by` | `managed_by` |
| `ibm_atracker_route_invalid_name` | `name` |
| `ibm_atracker_route_invalid_rules__locations` | `rules.locations` |
| `ibm_atracker_route_invalid_rules__target_ids` | `rules.target_ids` |

## Example

### Valid

```hcl
resource "ibm_atracker_route" "test" {
  name = "example-name"
  rules {
    locations = "example-name"
    target_ids = "example-name"
  }
}
```
