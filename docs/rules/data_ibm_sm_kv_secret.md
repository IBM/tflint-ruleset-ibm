# ibm_sm_kv_secret_data

Reads ibm_sm_kv_secret. The human-readable name of your secret.

This rule validates `ibm_sm_kv_secret` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SecretGroupPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_kv_secret_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `custom_metadata` | no | — |
| `data` | no | — |
| `description` | no | — |
| `downloaded` | no | — |
| `labels` | no | — |
| `locks_total` | no | — |
| `name` | no | pattern `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$` |
| `retrieved_at` | no | — |
| `secret_group_id` | no | — |
| `secret_group_name` | no | — |
| `secret_id` | no | pattern `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}` |
| `secret_type` | no | — |
| `state` | no | — |
| `state_description` | no | — |
| `updated_at` | no | — |
| `versions_total` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_kv_secret_data_invalid_name` | `name` |
| `ibm_sm_kv_secret_data_invalid_secret_group_name` | `secret_group_name` |
| `ibm_sm_kv_secret_data_invalid_secret_id` | `secret_id` |

## Example

### Valid

```hcl
data "ibm_sm_kv_secret" "test" {
}
```
