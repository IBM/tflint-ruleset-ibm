# ibm_is_ssh_key

Manages ibm_is_ssh_key. SSH Key name

This rule validates `ibm_is_ssh_key` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `KeyPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_ssh_key" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `fingerprint` | no | — |
| `href` | no | — |
| `length` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `public_key` | yes | API: The public SSH key to use, in OpenSSH format (consisting of three space-separated fields: the algorithm name, base64-encoded key value, and a comment). The algorithm and comment fields may be omitted, as only the key field is used. The key field must not match another key in the region. Keys of type `rsa` must be 2048 or 4096 bits in length (4096 is recommended). Keys of type `ed25519` must be 256 |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `type` | no | one of: `ed25519`, `rsa`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_ssh_key_invalid_access_tags` | `access_tags` |
| `ibm_is_ssh_key_invalid_name` | `name` |
| `ibm_is_ssh_key_invalid_public_key` | `public_key` |
| `ibm_is_ssh_key_invalid_resource_group` | `resource_group` |
| `ibm_is_ssh_key_invalid_tags` | `tags` |
| `ibm_is_ssh_key_invalid_type` | `type` |

## Example

### Valid

```hcl
resource "ibm_is_ssh_key" "test" {
  name = "example-name"
  public_key = "aaaaaaa"
}
```
