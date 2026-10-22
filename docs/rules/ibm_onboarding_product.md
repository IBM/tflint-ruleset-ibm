# ibm_onboarding_product

Schema validation for the ibm_onboarding_product resource.

This rule validates `ibm_onboarding_product` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `EscalationContactsUpdate` (`partner-center-sell`).

## Configuration

```hcl
rule "ibm_onboarding_product" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `approver_resource_id` | no | — |
| `eccn_number` | no | — |
| `ero_class` | no | — |
| `global_catalog_offering_id` | no | — |
| `iam_registration_id` | no | — |
| `oss_id` | no | — |
| `private_catalog_id` | no | — |
| `private_catalog_offering_id` | no | — |
| `staging_global_catalog_offering_id` | no | — |
| `tax_assessment` | no | — |
| `type` | yes | one of: `professional_service`, `service`, `software` |
| `unspsc` | no | — |
| `primary_contact` (block) | yes | nested block |
| `primary_contact.email` | yes | pattern `^\S+@\S+$` |
| `primary_contact.name` | yes | pattern `^\b([A-ZÀ-ÿ][-,a-z. ']+[ ]*)+$` |
| `support` (block) | no | nested block |
| `support.escalation_contacts` (block) | no | nested block |
| `support.escalation_contacts.email` | no | pattern `^\S+@\S+$` |
| `support.escalation_contacts.name` | no | pattern `^\b([A-ZÀ-ÿ][-,a-z. ']+[ ]*)+$` |
| `support.escalation_contacts.role` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_onboarding_product_invalid_eccn_number` | `eccn_number` |
| `ibm_onboarding_product_invalid_ero_class` | `ero_class` |
| `ibm_onboarding_product_invalid_tax_assessment` | `tax_assessment` |
| `ibm_onboarding_product_invalid_type` | `type` |
| `ibm_onboarding_product_invalid_unspsc` | `unspsc` |
| `ibm_onboarding_product_invalid_primary_contact__email` | `primary_contact.email` |
| `ibm_onboarding_product_invalid_primary_contact__name` | `primary_contact.name` |
| `ibm_onboarding_product_invalid_support__escalation_contacts__email` | `support.escalation_contacts.email` |
| `ibm_onboarding_product_invalid_support__escalation_contacts__name` | `support.escalation_contacts.name` |
| `ibm_onboarding_product_invalid_support__escalation_contacts__role` | `support.escalation_contacts.role` |

## Example

### Valid

```hcl
resource "ibm_onboarding_product" "test" {
  type = "professional_service"
  primary_contact {
    email = "user@example.com"
    name = "Example"
  }
}
```
