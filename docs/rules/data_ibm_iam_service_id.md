# ibm_iam_service_id_data

Reads ibm_iam_service_id. Name of the serviceID

This rule validates `ibm_iam_service_id` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `UpdateServiceIdGroupRequest` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_service_id_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | yes | API: Name of the service ID group. Unique in the account. |
| `service_ids` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_service_id_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_iam_service_id" "test" {
  name = "example"
}
```
