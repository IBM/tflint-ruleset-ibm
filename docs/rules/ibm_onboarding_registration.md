# ibm_onboarding_registration

Schema validation for the ibm_onboarding_registration resource.

This rule validates `ibm_onboarding_registration` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `EscalationContactsUpdate` (`partner-center-sell`).

## Configuration

```hcl
rule "ibm_onboarding_registration" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | yes | pattern `^[a-zA-Z0-9]+$` |
| `company_name` | yes | — |
| `created_at` | no | — |
| `default_private_catalog_id` | no | pattern `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}` |
| `provider_access_group` | no | pattern `^AccessGroupId-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `updated_at` | no | — |
| `primary_contact` (block) | yes | nested block |
| `primary_contact.email` | yes | pattern `^\S+@\S+$` |
| `primary_contact.name` | yes | pattern `^\b([A-ZÀ-ÿ][-,a-z. ']+[ ]*)+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_onboarding_registration_invalid_account_id` | `account_id` |
| `ibm_onboarding_registration_invalid_company_name` | `company_name` |
| `ibm_onboarding_registration_invalid_default_private_catalog_id` | `default_private_catalog_id` |
| `ibm_onboarding_registration_invalid_provider_access_group` | `provider_access_group` |
| `ibm_onboarding_registration_invalid_primary_contact__email` | `primary_contact.email` |
| `ibm_onboarding_registration_invalid_primary_contact__name` | `primary_contact.name` |

## Example

### Valid

```hcl
resource "ibm_onboarding_registration" "test" {
  account_id = "example"
  company_name = "aaaaaaa"
  primary_contact {
    email = "user@example.com"
    name = "Example"
  }
}
```
