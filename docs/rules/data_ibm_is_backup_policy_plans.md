# ibm_is_backup_policy_plans_data

Reads ibm_is_backup_policy_plans. The unique user-defined name for this backup policy plan.

This rule validates `ibm_is_backup_policy_plans` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_backup_policy_plans_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `backup_policy_id` | yes | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `plans` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_backup_policy_plans_data_invalid_backup_policy_id` | `backup_policy_id` |
| `ibm_is_backup_policy_plans_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_backup_policy_plans" "test" {
  backup_policy_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
