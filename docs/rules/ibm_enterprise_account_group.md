# ibm_enterprise_account_group

Manages ibm_enterprise_account_group. The name of the account group.

This rule validates `ibm_enterprise_account_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CreateAccountGroupRequest` (`enterprise-apis-enterprise`).

## Configuration

```hcl
rule "ibm_enterprise_account_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `enterprise_account_id` | no | — |
| `enterprise_id` | no | — |
| `enterprise_path` | no | — |
| `name` | yes | API: The name of the account group. This field must have 3 - 60 characters. |
| `parent` | yes | API: The CRN of the parent under which the account group will be created. The parent can be an existing account group or the enterprise itself. |
| `primary_contact_email` | no | — |
| `primary_contact_iam_id` | yes | API: The IAM ID of the primary contact for this account group, such as `IBMid-0123ABC`. The IAM ID must already exist. |
| `state` | no | — |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `url` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_enterprise_account_group_invalid_name` | `name` |
| `ibm_enterprise_account_group_invalid_parent` | `parent` |
| `ibm_enterprise_account_group_invalid_primary_contact_iam_id` | `primary_contact_iam_id` |

## Example

### Valid

```hcl
resource "ibm_enterprise_account_group" "test" {
  name = "example"
  parent = "example"
  primary_contact_iam_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
