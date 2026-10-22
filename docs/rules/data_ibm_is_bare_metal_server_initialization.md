# ibm_is_bare_metal_server_initialization_data

Schema validation for the ibm_is_bare_metal_server_initialization data source.

This rule validates `ibm_is_bare_metal_server_initialization` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ImageAllowedUsePatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_server_initialization_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bare_metal_server` | yes | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `default_trusted_profile` | no | — |
| `image` | no | — |
| `image_name` | no | — |
| `keys` | no | — |
| `passphrase` | no | — |
| `private_key` | no | — |
| `user_accounts` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_server_initialization_data_invalid_bare_metal_server` | `bare_metal_server` |
| `ibm_is_bare_metal_server_initialization_data_invalid_passphrase` | `passphrase` |
| `ibm_is_bare_metal_server_initialization_data_invalid_private_key` | `private_key` |

## Example

### Valid

```hcl
data "ibm_is_bare_metal_server_initialization" "test" {
  bare_metal_server = "example-name"
}
```
