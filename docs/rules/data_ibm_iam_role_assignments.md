# ibm_iam_role_assignments_data

Schema validation for the ibm_iam_role_assignments data source.

This rule validates `ibm_iam_role_assignments` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/role_assignments/{assignment_id}:patch` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_role_assignments_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `assignments` | no | — |
| `template_id` | no | pattern `^policyTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `template_version` | no | pattern `^[0-9]*$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_role_assignments_data_invalid_template_id` | `template_id` |
| `ibm_iam_role_assignments_data_invalid_template_version` | `template_version` |

## Example

### Valid

```hcl
data "ibm_iam_role_assignments" "test" {
}
```
