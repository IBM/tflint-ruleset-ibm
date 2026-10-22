# ibm_atracker_targets_data

Reads ibm_atracker_targets. The name of the target resource.

This rule validates `ibm_atracker_targets` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `TargetPrototype` (`atracker`).

## Configuration

```hcl
rule "ibm_atracker_targets_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | no | pattern `^[a-zA-Z0-9 -._:]+$` |
| `region` | no | pattern `^[a-zA-Z0-9 -._:]+$` |
| `targets` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_atracker_targets_data_invalid_name` | `name` |
| `ibm_atracker_targets_data_invalid_region` | `region` |

## Example

### Valid

```hcl
data "ibm_atracker_targets" "test" {
}
```
