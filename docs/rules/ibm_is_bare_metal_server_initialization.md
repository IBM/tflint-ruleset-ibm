# ibm_is_bare_metal_server_initialization

Schema validation for the ibm_is_bare_metal_server_initialization resource.

This rule validates `ibm_is_bare_metal_server_initialization` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerInitializationPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_server_initialization" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bare_metal_server` | yes | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `image` | yes | API: Identifies an image by a unique property. |
| `keys` | yes | API: Identifies a key by a unique property. |
| `user_data` | no | API: The [user data](https://cloud.ibm.com/docs/vpc?topic=vpc-user-data) to make available when setting up the bare metal server. |
| `default_trusted_profile` (block) | no | nested block |
| `default_trusted_profile.auto_link` | no | API: If set to `true`, the system will create a link to the specified `target` trusted profile during server creation. Regardless of whether a link is created by the system or manually using the IAM Identity service, it will be automatically deleted when the server is deleted. |
| `default_trusted_profile.target` (block) | no | nested block |
| `default_trusted_profile.target.crn` | no | crn format |
| `default_trusted_profile.target.id` | no | pattern `^Profile-[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_server_initialization_invalid_bare_metal_server` | `bare_metal_server` |
| `ibm_is_bare_metal_server_initialization_invalid_image` | `image` |
| `ibm_is_bare_metal_server_initialization_invalid_keys` | `keys` |
| `ibm_is_bare_metal_server_initialization_invalid_user_data` | `user_data` |
| `ibm_is_bare_metal_server_initialization_invalid_default_trusted_profile__auto_link` | `default_trusted_profile.auto_link` |
| `ibm_is_bare_metal_server_initialization_invalid_default_trusted_profile__target__crn` | `default_trusted_profile.target.crn` |
| `ibm_is_bare_metal_server_initialization_invalid_default_trusted_profile__target__id` | `default_trusted_profile.target.id` |

## Example

### Valid

```hcl
resource "ibm_is_bare_metal_server_initialization" "test" {
  bare_metal_server = "example-name"
  image = "example"
  keys = "example"
}
```
