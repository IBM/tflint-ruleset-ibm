# ibm_is_image_deprecate

Manages ibm_is_image_deprecate. Image name

This rule validates `ibm_is_image_deprecate` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `InstanceByImageAndCatalogOffering` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_image_deprecate" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `checksum` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `deprecation_at` | no | — |
| `encrypted_data_key` | no | — |
| `encryption` | no | — |
| `encryption_key` | no | — |
| `file` | no | — |
| `href` | no | — |
| `image` | yes | API: Identifies an image by a unique property. |
| `name` | no | — |
| `obsolescence_at` | no | — |
| `operating_system` | no | — |
| `resource_group` | no | — |
| `size` | no | — |
| `source_volume` | no | — |
| `status` | no | — |
| `tags` | no | — |
| `visibility` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_image_deprecate_invalid_image` | `image` |

## Example

### Valid

```hcl
resource "ibm_is_image_deprecate" "test" {
  image = "example"
}
```
