# ibm_onboarding_resource_broker

Manages ibm_onboarding_resource_broker. The name of the broker.

This rule validates `ibm_onboarding_resource_broker` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `EscalationContactsUpdate` (`partner-center-sell`).

## Configuration

```hcl
rule "ibm_onboarding_resource_broker" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `allow_context_updates` | no | — |
| `auth_password` | no | — |
| `auth_scheme` | yes | one of: `bearer`, `bearer-crn` |
| `auth_username` | no | one of: `apikey` |
| `broker_url` | yes | — |
| `catalog_type` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `deleted_at` | no | — |
| `deleted_by` | no | — |
| `env` | no | pattern `^[a-z_.-]+$` |
| `guid` | no | — |
| `name` | yes | pattern `^\b([A-ZÀ-ÿ][-,a-z. ']+[ ]*)+$` |
| `region` | no | — |
| `resource_group_crn` | no | crn format |
| `state` | no | one of: `active`, `removed` |
| `type` | yes | one of: `provision_behind`, `provision_through` |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `url` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_onboarding_resource_broker_invalid_allow_context_updates` | `allow_context_updates` |
| `ibm_onboarding_resource_broker_invalid_auth_password` | `auth_password` |
| `ibm_onboarding_resource_broker_invalid_auth_scheme` | `auth_scheme` |
| `ibm_onboarding_resource_broker_invalid_auth_username` | `auth_username` |
| `ibm_onboarding_resource_broker_invalid_broker_url` | `broker_url` |
| `ibm_onboarding_resource_broker_invalid_catalog_type` | `catalog_type` |
| `ibm_onboarding_resource_broker_invalid_env` | `env` |
| `ibm_onboarding_resource_broker_invalid_name` | `name` |
| `ibm_onboarding_resource_broker_invalid_region` | `region` |
| `ibm_onboarding_resource_broker_invalid_resource_group_crn` | `resource_group_crn` |
| `ibm_onboarding_resource_broker_invalid_state` | `state` |
| `ibm_onboarding_resource_broker_invalid_type` | `type` |

## Example

### Valid

```hcl
resource "ibm_onboarding_resource_broker" "test" {
  auth_scheme = "bearer"
  broker_url = "example"
  name = "Example"
  type = "provision_behind"
}
```
