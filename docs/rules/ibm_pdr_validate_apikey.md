# ibm_pdr_validate_apikey

Schema validation for the ibm_pdr_validate_apikey resource.

This rule validates `ibm_pdr_validate_apikey` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `UpdateKeyRequest` (`dr-automation-powervs`).

## Configuration

```hcl
rule "ibm_pdr_validate_apikey" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | — |
| `api_key` | yes | pattern `^[a-zA-Z0-9\-_]+$` |
| `description` | no | — |
| `etag` | no | — |
| `id` | no | — |
| `instance_id` | yes | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pdr_validate_apikey_invalid_accept_language` | `accept_language` |
| `ibm_pdr_validate_apikey_invalid_api_key` | `api_key` |
| `ibm_pdr_validate_apikey_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
variable "api_key" {
  sensitive = true
}

resource "ibm_pdr_validate_apikey" "test" {
  api_key = var.api_key
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
