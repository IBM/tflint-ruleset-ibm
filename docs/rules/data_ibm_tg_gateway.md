# ibm_tg_gateway_data

Reads ibm_tg_gateway. The Transit Gateway identifier

This rule validates `ibm_tg_gateway` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `RedundancyGroupPatch` (`transit-gateway`).

## Configuration

```hcl
rule "ibm_tg_gateway_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `connections` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `global` | no | — |
| `gre_enhanced_route_propagation` | no | — |
| `location` | no | — |
| `name` | yes | pattern `^([a-zA-Z]|[a-zA-Z][-_a-zA-Z0-9]*[a-zA-Z0-9])$` |
| `resource_group` | no | — |
| `status` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_tg_gateway_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_tg_gateway" "test" {
  name = "example-name"
}
```
