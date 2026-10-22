# ibm_iam_role_actions_data

Schema validation for the ibm_iam_role_actions data source.

This rule validates `ibm_iam_role_actions` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `IdentityTypesPatch` (`iam-policy-management`).

## Configuration

```hcl
rule "ibm_iam_role_actions_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `actions` | no | — |
| `manager` | no | — |
| `reader` | no | — |
| `reader_plus` | no | — |
| `service` | yes | API: The core set of properties associated with an identity type. |
| `writer` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_role_actions_data_invalid_service` | `service` |

## Example

### Valid

```hcl
data "ibm_iam_role_actions" "test" {
  service = "example"
}
```
