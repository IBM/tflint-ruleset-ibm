# ibm_schematics_policies_data

Schema validation for the ibm_schematics_policies data source.

This rule validates `ibm_schematics_policies` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/settings/policies/{policy_id}:patch` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_policies_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `limit` | no | — |
| `offset` | no | — |
| `policies` | no | — |
| `policy_kind` | no | one of: `agent_assignment_policy` |
| `total_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_policies_data_invalid_policy_kind` | `policy_kind` |

## Example

### Valid

```hcl
data "ibm_schematics_policies" "test" {
}
```
