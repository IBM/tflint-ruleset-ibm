# ibm_schematics_policy_data

Reads ibm_schematics_policy. Name of Schematics customization policy.

This rule validates `ibm_schematics_policy` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:patch:/v2/settings/policies/{policy_id}` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_policy_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `description` | no | — |
| `id` | no | — |
| `kind` | no | — |
| `location` | no | — |
| `name` | no | — |
| `parameter` | no | — |
| `policy_id` | yes | — |
| `resource_group` | no | — |
| `scoped_resources` | no | — |
| `state` | no | — |
| `tags` | no | — |
| `target` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_policy_data_invalid_policy_id` | `policy_id` |

## Example

### Valid

```hcl
data "ibm_schematics_policy" "test" {
  policy_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
