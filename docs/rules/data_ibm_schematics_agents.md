# ibm_schematics_agents_data

Reads ibm_schematics_agents. The name of the agent (must be unique, for an account).

This rule validates `ibm_schematics_agents` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CommandsInfo` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_agents_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `agents` | no | — |
| `limit` | no | — |
| `name` | no | API: Name of the command |
| `offset` | no | — |
| `total_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_agents_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_schematics_agents" "test" {
}
```
