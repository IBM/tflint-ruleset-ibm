# ibm_iam_api_key

Manages ibm_iam_api_key. Name of the API key.

This rule validates `ibm_iam_api_key` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ApiKeyInsideCreateServiceIdRequest` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_api_key" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `apikey` | no | API: You can optionally passthrough the API key value for this API key. If passed, a minimum length validation of 32 characters for that apiKey value is done, i.e. the value can contain any characters and can even be non-URL safe, but the minimum length requirement must be met. If omitted, the API key management will create an URL safe opaque API key value. The value of the API key is checked for uniqu |
| `apikey_id` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `description` | no | API: The optional description of the API key. The 'description' property is only available if a description was provided during a create of an API key. |
| `entity_lock` | no | — |
| `entity_tag` | no | — |
| `expires_at` | no | API: Date and time when the API key becomes invalid, ISO 8601 datetime in the format 'yyyy-MM-ddTHH:mm+0000'. **WARNING** An API key will be permanently and irrevocably deleted when both the expires_at and modified_at timestamps are more than ninety (90) days in the past, regardless of the key’s locked status or any other state. |
| `file` | no | — |
| `iam_id` | no | — |
| `locked` | no | — |
| `modified_at` | no | — |
| `name` | yes | API: Name of the API key. The name is not checked for uniqueness. Therefore multiple names with the same value can exist. Access is done via the UUID of the API key. |
| `store_value` | no | API: Send true or false to set whether the API key value is retrievable in the future by using the Get details of an API key request. If you create an API key for a user, you must specify `false` or omit the value. We don't allow storing of API keys for users. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_api_key_invalid_apikey` | `apikey` |
| `ibm_iam_api_key_invalid_description` | `description` |
| `ibm_iam_api_key_invalid_entity_lock` | `entity_lock` |
| `ibm_iam_api_key_invalid_expires_at` | `expires_at` |
| `ibm_iam_api_key_invalid_file` | `file` |
| `ibm_iam_api_key_invalid_name` | `name` |
| `ibm_iam_api_key_invalid_store_value` | `store_value` |

## Example

### Valid

```hcl
resource "ibm_iam_api_key" "test" {
  name = "example"
}
```
