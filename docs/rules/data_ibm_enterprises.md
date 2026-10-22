# ibm_enterprises_data

Reads ibm_enterprises. The name of the enterprise.

This rule validates `ibm_enterprises` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/enterprises/{enterprise_id}:patch` (`enterprise-apis-enterprise`).

## Configuration

```hcl
rule "ibm_enterprises_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `enterprises` | no | — |
| `name` | no | API: The new name of the enterprise. This field must have 3 - 60 characters. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_enterprises_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_enterprises" "test" {
}
```
