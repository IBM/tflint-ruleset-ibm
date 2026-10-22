# ibm_iam_idp_data

Reads ibm_iam_idp. Speaking name of the Identity Provider.

This rule validates `ibm_iam_idp` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ConsumersResponse` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_idp_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `active` | no | — |
| `created_at` | no | — |
| `entity_tag` | no | — |
| `idp_id` | yes | pattern `^[A-Za-z0-9]+$` |
| `modified_at` | no | — |
| `name` | no | — |
| `share_scope` | no | — |
| `type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_idp_data_invalid_idp_id` | `idp_id` |

## Example

### Valid

```hcl
data "ibm_iam_idp" "test" {
  idp_id = "example"
}
```
