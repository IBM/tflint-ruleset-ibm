# ibm_iam_role_template_data

Schema validation for the ibm_iam_role_template data source.

This rule validates `ibm_iam_role_template` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Role` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_role_template_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `role_templates` | no | — |

## Example

### Valid

```hcl
data "ibm_iam_role_template" "test" {
}
```
