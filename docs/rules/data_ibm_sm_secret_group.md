# ibm_sm_secret_group_data

Reads ibm_sm_secret_group. The name of your secret group.

This rule validates `ibm_sm_secret_group` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Secret` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_secret_group_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `description` | no | — |
| `name` | no | — |
| `secret_group_id` | yes | pattern `^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|default)$` |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_secret_group_data_invalid_secret_group_id` | `secret_group_id` |

## Example

### Valid

```hcl
variable "secret_group_id" {
  sensitive = true
}

data "ibm_sm_secret_group" "test" {
  secret_group_id = var.secret_group_id
}
```
