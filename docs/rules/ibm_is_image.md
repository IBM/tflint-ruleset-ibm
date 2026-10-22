# ibm_is_image

Manages ibm_is_image. Image name

This rule validates `ibm_is_image` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Image` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_image" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `checksum` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `deprecate` | no | — |
| `deprecation_at` | no | pattern `^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[\+-]\d{2}:\d{2})?)$` |
| `encrypted_data_key` | no | pattern `^(?:[A-Za-z0-9+\/]{4})*(?:[A-Za-z0-9+\/]{2}==|[A-Za-z0-9+\/]{3}=|[A-Za-z0-9+\/]{4})$` |
| `encryption` | no | — |
| `encryption_key` | no | API: The key that will be used to encrypt volumes created from this image (unless an alternate `encryption_key` is specified at volume creation). This property will be present for images with an `encryption` type of `user_managed`. |
| `file` | no | — |
| `href` | no | pattern `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$` |
| `minimum_acceptable_status` | no | — |
| `minimum_provisioned_size` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `obsolescence_at` | no | pattern `^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[\+-]\d{2}:\d{2})?)$` |
| `obsolete` | no | — |
| `operating_system` | no | API: The operating system included in this image |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | API: The resource group for this image |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `size` | no | — |
| `source_volume` | no | API: The volume used to create this image (this may be [deleted](https://cloud.ibm.com/apidocs/vpc#deleted-resources)). If absent, this image was not created from a volume. |
| `status` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `user_data_format` | no | — |
| `visibility` | no | — |
| `zones` | no | — |
| `allowed_use` (block) | no | nested block |
| `allowed_use.api_version` | no | pattern `^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$` |
| `allowed_use.bare_metal_server` | no | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `allowed_use.instance` | no | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_image_invalid_access_tags` | `access_tags` |
| `ibm_is_image_invalid_deprecate` | `deprecate` |
| `ibm_is_image_invalid_deprecation_at` | `deprecation_at` |
| `ibm_is_image_invalid_encrypted_data_key` | `encrypted_data_key` |
| `ibm_is_image_invalid_encryption_key` | `encryption_key` |
| `ibm_is_image_invalid_href` | `href` |
| `ibm_is_image_invalid_minimum_acceptable_status` | `minimum_acceptable_status` |
| `ibm_is_image_invalid_name` | `name` |
| `ibm_is_image_invalid_obsolescence_at` | `obsolescence_at` |
| `ibm_is_image_invalid_obsolete` | `obsolete` |
| `ibm_is_image_invalid_operating_system` | `operating_system` |
| `ibm_is_image_invalid_resource_group` | `resource_group` |
| `ibm_is_image_invalid_source_volume` | `source_volume` |
| `ibm_is_image_invalid_tags` | `tags` |
| `ibm_is_image_invalid_allowed_use__api_version` | `allowed_use.api_version` |
| `ibm_is_image_invalid_allowed_use__bare_metal_server` | `allowed_use.bare_metal_server` |
| `ibm_is_image_invalid_allowed_use__instance` | `allowed_use.instance` |

## Example

### Valid

```hcl
resource "ibm_is_image" "test" {
  name = "example-name"
}
```
