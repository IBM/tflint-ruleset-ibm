# ibm_hpcs

Manages ibm_hpcs. A name for the HPCS instance

This rule validates `ibm_hpcs` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v4/templates/{id}:patch` (`uko`).

## Configuration

```hcl
rule "ibm_hpcs" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `dashboard_url` | no | — |
| `deleted_at` | no | — |
| `deleted_by` | no | — |
| `extensions` | no | — |
| `failover_units` | no | — |
| `guid` | no | — |
| `hsm_info` | no | — |
| `location` | yes | pattern `^[A-Za-z0-9][A-Za-z0-9 ._-]*$` |
| `name` | yes | pattern `^[A-Za-z0-9#@!$%'_-][A-Za-z0-9#@!$% '_-]*$` |
| `plan` | yes | — |
| `resource_aliases_url` | no | — |
| `resource_bindings_url` | no | — |
| `resource_group_id` | no | — |
| `resource_keys_url` | no | — |
| `restored_at` | no | — |
| `restored_by` | no | — |
| `revocation_threshold` | yes | — |
| `scheduled_reclaim_at` | no | — |
| `scheduled_reclaim_by` | no | — |
| `service` | no | — |
| `service_endpoints` | no | one of: `public-and-private`, `private-only` |
| `signature_server_url` | no | — |
| `signature_threshold` | yes | — |
| `state` | no | — |
| `status` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `units` | yes | — |
| `update_at` | no | — |
| `update_by` | no | — |
| `admins` (block) | yes | nested block |
| `admins.key` | yes | API: Updated key related properties |
| `admins.name` | yes | pattern `^[A-Za-z0-9#@!$%'_-][A-Za-z0-9#@!$% '_-]*$` |
| `admins.token` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_hpcs_invalid_failover_units` | `failover_units` |
| `ibm_hpcs_invalid_location` | `location` |
| `ibm_hpcs_invalid_name` | `name` |
| `ibm_hpcs_invalid_plan` | `plan` |
| `ibm_hpcs_invalid_resource_group_id` | `resource_group_id` |
| `ibm_hpcs_invalid_revocation_threshold` | `revocation_threshold` |
| `ibm_hpcs_invalid_service` | `service` |
| `ibm_hpcs_invalid_service_endpoints` | `service_endpoints` |
| `ibm_hpcs_invalid_signature_server_url` | `signature_server_url` |
| `ibm_hpcs_invalid_signature_threshold` | `signature_threshold` |
| `ibm_hpcs_invalid_tags` | `tags` |
| `ibm_hpcs_invalid_units` | `units` |
| `ibm_hpcs_invalid_admins__key` | `admins.key` |
| `ibm_hpcs_invalid_admins__name` | `admins.name` |
| `ibm_hpcs_invalid_admins__token` | `admins.token` |

## Example

### Valid

```hcl
variable "token" {
  sensitive = true
}

resource "ibm_hpcs" "test" {
  location = "example-name"
  name = "example-name"
  plan = "example"
  revocation_threshold = 1
  signature_threshold = 1
  units = 1
  admins {
    key = "example"
    name = "example-name"
    token = var.token
  }
}
```
