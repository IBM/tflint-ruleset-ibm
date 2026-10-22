# ibm_iam_service_id

Manages ibm_iam_service_id. Name of the Service Id.

This rule validates `ibm_iam_service_id` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `UpdateServiceIdGroupRequest` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_service_id" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `description` | no | API: Description of the service ID group. |
| `iam_id` | no | — |
| `locked` | no | — |
| `modified_at` | no | — |
| `name` | yes | API: Name of the service ID group. Unique in the account. |
| `tags` | no | — |
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_service_id_invalid_description` | `description` |
| `ibm_iam_service_id_invalid_name` | `name` |
| `ibm_iam_service_id_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_iam_service_id" "test" {
  name = "example"
}
```
