# ibm_pha_api_key

Schema validation for the ibm_pha_api_key resource.

This rule validates `ibm_pha_api_key` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pha_api_key" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | pattern `^[a-zA-Z0-9\-_,;=.*]+$` |
| `api_key` | no | pattern `^[A-Za-z0-9._:-]+$` |
| `etag` | no | — |
| `id` | no | — |
| `if_none_match` | no | pattern `^[a-zA-Z0-9\-_,;=.*]+$` |
| `instance_id` | yes | pattern `^[a-zA-Z0-9-]+$` |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pha_api_key_invalid_accept_language` | `accept_language` |
| `ibm_pha_api_key_invalid_api_key` | `api_key` |
| `ibm_pha_api_key_invalid_if_none_match` | `if_none_match` |
| `ibm_pha_api_key_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
resource "ibm_pha_api_key" "test" {
  instance_id = "example-name"
}
```
