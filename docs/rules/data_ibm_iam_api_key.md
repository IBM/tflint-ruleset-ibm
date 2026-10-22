# ibm_iam_api_key_data

Reads ibm_iam_api_key. Name of the API key.

This rule validates `ibm_iam_api_key` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_iam_api_key_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `apikey_id` | yes | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `description` | no | — |
| `entity_tag` | no | — |
| `expires_at` | no | — |
| `iam_id` | no | — |
| `locked` | no | — |
| `modified_at` | no | — |
| `name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_api_key_data_invalid_apikey_id` | `apikey_id` |

## Example

### Valid

```hcl
variable "apikey_id" {
  sensitive = true
}

data "ibm_iam_api_key" "test" {
  apikey_id = var.apikey_id
}
```
