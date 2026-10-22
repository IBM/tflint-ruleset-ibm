# ibm_is_image_data

Reads ibm_is_image. Image name

This rule validates `ibm_is_image` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_image_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `allowed_use` | no | — |
| `architecture` | no | — |
| `catalog_offering` | no | — |
| `checksum` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `deprecation_at` | no | — |
| `encryption` | no | — |
| `encryption_key` | no | — |
| `identifier` | no | — |
| `minimum_provisioned_size` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `obsolescence_at` | no | — |
| `operating_system` | no | — |
| `os` | no | — |
| `remote` | no | — |
| `resource_group` | no | — |
| `source_volume` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `user_data_format` | no | — |
| `visibility` | no | one of: `public`, `private`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `zones` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_image_data_invalid_identifier` | `identifier` |
| `ibm_is_image_data_invalid_name` | `name` |
| `ibm_is_image_data_invalid_visibility` | `visibility` |

## Example

### Valid

```hcl
data "ibm_is_image" "test" {
}
```
