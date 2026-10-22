# ibm_enterprise

Manages ibm_enterprise. The name of the enterprise.

This rule validates `ibm_enterprise` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/enterprises/{enterprise_id}:patch` (`enterprise-apis-enterprise`).

## Configuration

```hcl
rule "ibm_enterprise" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `domain` | no | API: The new domain of the enterprise. This field has a limit of 60 characters. |
| `enterprise_account_id` | no | — |
| `name` | yes | API: The new name of the enterprise. This field must have 3 - 60 characters. |
| `primary_contact_email` | no | — |
| `primary_contact_iam_id` | yes | API: The IAM ID of the user to be the new primary contact for the enterprise. |
| `source_account_id` | yes | — |
| `state` | no | — |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `url` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_enterprise_invalid_domain` | `domain` |
| `ibm_enterprise_invalid_name` | `name` |
| `ibm_enterprise_invalid_primary_contact_iam_id` | `primary_contact_iam_id` |
| `ibm_enterprise_invalid_source_account_id` | `source_account_id` |

## Example

### Valid

```hcl
resource "ibm_enterprise" "test" {
  name = "example"
  primary_contact_iam_id = "01234567-89ab-cdef-0123-456789abcdef"
  source_account_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
