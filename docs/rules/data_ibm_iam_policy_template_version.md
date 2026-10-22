# ibm_iam_policy_template_version_data

Reads ibm_iam_policy_template_version. name of template.

This rule validates `ibm_iam_policy_template_version` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `V2PolicySubjectAttribute` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_policy_template_version_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `committed` | no | — |
| `description` | no | — |
| `id` | no | — |
| `name` | no | — |
| `policy` | no | — |
| `policy_template_id` | yes | pattern `^policyTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `version` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_policy_template_version_data_invalid_policy_template_id` | `policy_template_id` |
| `ibm_iam_policy_template_version_data_invalid_version` | `version` |

## Example

### Valid

```hcl
data "ibm_iam_policy_template_version" "test" {
  policy_template_id = "policyTemplate-01234567-89ab-cdef-0123-456789abcdef"
  version = "example"
}
```
