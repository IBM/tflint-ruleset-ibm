# ibm_iam_serviceid_group

Manages ibm_iam_serviceid_group. Name of the service ID group.

This rule validates `ibm_iam_serviceid_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CreateServiceIdGroupRequest` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_serviceid_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | yes | API: ID of the account the service ID group belongs to. |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `description` | no | API: Description of the service ID group. |
| `entity_tag` | no | — |
| `modified_at` | no | — |
| `name` | yes | API: Name of the service ID group. Unique in the account. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_serviceid_group_invalid_account_id` | `account_id` |
| `ibm_iam_serviceid_group_invalid_description` | `description` |
| `ibm_iam_serviceid_group_invalid_entity_tag` | `entity_tag` |
| `ibm_iam_serviceid_group_invalid_name` | `name` |

## Example

### Valid

```hcl
resource "ibm_iam_serviceid_group" "test" {
  account_id = "01234567-89ab-cdef-0123-456789abcdef"
  name = "example"
}
```
