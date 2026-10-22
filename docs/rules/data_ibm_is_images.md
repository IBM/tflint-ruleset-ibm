# ibm_is_images_data

Reads ibm_is_images. The name of the image

This rule validates `ibm_is_images` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `FloatingIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_images_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `catalog_managed` | no | — |
| `images` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `remote_account_id` | no | — |
| `resource_group` | no | API: The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) will be used. |
| `status` | no | one of: `available`, `deleting`, `deprecated`, `failed`, `obsolete`, `pending`, `tentative`, `unusable` |
| `user_data_format` | no | one of: `cloud_init`, `esxi_kickstart`, `ipxe`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `visibility` | no | one of: `private`, `public`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_images_data_invalid_catalog_managed` | `catalog_managed` |
| `ibm_is_images_data_invalid_name` | `name` |
| `ibm_is_images_data_invalid_remote_account_id` | `remote_account_id` |
| `ibm_is_images_data_invalid_resource_group` | `resource_group` |
| `ibm_is_images_data_invalid_status` | `status` |
| `ibm_is_images_data_invalid_user_data_format` | `user_data_format` |
| `ibm_is_images_data_invalid_visibility` | `visibility` |

## Example

### Valid

```hcl
data "ibm_is_images" "test" {
}
```
