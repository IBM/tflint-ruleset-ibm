# ibm_is_instance_reinitialize

Schema validation for the ibm_is_instance_reinitialize resource.

This rule validates `ibm_is_instance_reinitialize` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Instance` (`vpc-metadata-latest`).

## Configuration

```hcl
rule "ibm_is_instance_reinitialize" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `image` | no | API: The image the virtual server instance was provisioned from |
| `instance_id` | yes | — |
| `keys` | no | — |
| `status` | no | — |
| `user_data` | no | API: The [user data](https://cloud.ibm.com/docs/vpc?topic=vpc-user-data) to make available when setting up the bare metal server. |
| `boot_volume_attachment` (block) | no | nested block |
| `boot_volume_attachment.delete_volume_on_instance_delete` | no | — |
| `boot_volume_attachment.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `boot_volume_attachment.volume` (block) | no | nested block |
| `boot_volume_attachment.volume.bandwidth` | no | API: The total bandwidth (in megabits per second) shared across the instance network attachments or instance network interfaces and storage volumes of the virtual server instance |
| `boot_volume_attachment.volume.capacity` | no | — |
| `boot_volume_attachment.volume.id` | no | pattern `^[-0-9a-z_]+$` |
| `boot_volume_attachment.volume.iops` | no | range 100–9007199254740991 |
| `boot_volume_attachment.volume.name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `boot_volume_attachment.volume.resource_group` | no | API: The resource group for this instance |
| `boot_volume_attachment.volume.user_tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `boot_volume_attachment.volume.allowed_use` (block) | no | nested block |
| `boot_volume_attachment.volume.allowed_use.api_version` | yes | pattern `^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$` |
| `boot_volume_attachment.volume.allowed_use.bare_metal_server` | no | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `boot_volume_attachment.volume.allowed_use.instance` | no | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `boot_volume_attachment.volume.encryption_key` (block) | no | nested block |
| `boot_volume_attachment.volume.encryption_key.crn` | yes | crn format |
| `boot_volume_attachment.volume.profile` (block) | no | nested block |
| `boot_volume_attachment.volume.profile.name` | yes | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `boot_volume_attachment.volume.source_snapshot` (block) | no | nested block |
| `boot_volume_attachment.volume.source_snapshot.id` | yes | pattern `^[-0-9a-z_]+$` |
| `default_trusted_profile` (block) | no | nested block |
| `default_trusted_profile.auto_link` | yes | — |
| `default_trusted_profile.target` (block) | yes | nested block |
| `default_trusted_profile.target.crn` | no | crn format |
| `default_trusted_profile.target.id` | no | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_reinitialize_invalid_image` | `image` |
| `ibm_is_instance_reinitialize_invalid_instance_id` | `instance_id` |
| `ibm_is_instance_reinitialize_invalid_keys` | `keys` |
| `ibm_is_instance_reinitialize_invalid_user_data` | `user_data` |
| `ibm_is_instance_reinitialize_invalid_boot_volume_attachment__delete_volume_on_instance_delete` | `boot_volume_attachment.delete_volume_on_instance_delete` |
| `ibm_is_instance_reinitialize_invalid_boot_volume_attachment__name` | `boot_volume_attachment.name` |
| `ibm_is_instance_reinitialize_invalid_boot_volume_attachment__volume__bandwidth` | `boot_volume_attachment.volume.bandwidth` |
| `ibm_is_instance_reinitialize_invalid_boot_volume_attachment__volume__capacity` | `boot_volume_attachment.volume.capacity` |
| `ibm_is_instance_reinitialize_invalid_boot_volume_attachment__volume__id` | `boot_volume_attachment.volume.id` |
| `ibm_is_instance_reinitialize_invalid_boot_volume_attachment__volume__iops` | `boot_volume_attachment.volume.iops` |
| `ibm_is_instance_reinitialize_invalid_boot_volume_attachment__volume__name` | `boot_volume_attachment.volume.name` |
| `ibm_is_instance_reinitialize_invalid_boot_volume_attachment__volume__resource_group` | `boot_volume_attachment.volume.resource_group` |
| `ibm_is_instance_reinitialize_invalid_boot_volume_attachment__volume__user_tags` | `boot_volume_attachment.volume.user_tags` |
| `ibm_is_instance_reinitialize_invalid_boot_volume_attachment__volume__allowed_use__api_version` | `boot_volume_attachment.volume.allowed_use.api_version` |
| `ibm_is_instance_reinitialize_invalid_boot_volume_attachment__volume__allowed_use__bare_metal_server` | `boot_volume_attachment.volume.allowed_use.bare_metal_server` |
| `ibm_is_instance_reinitialize_invalid_boot_volume_attachment__volume__allowed_use__instance` | `boot_volume_attachment.volume.allowed_use.instance` |
| `ibm_is_instance_reinitialize_invalid_boot_volume_attachment__volume__encryption_key__crn` | `boot_volume_attachment.volume.encryption_key.crn` |
| `ibm_is_instance_reinitialize_invalid_boot_volume_attachment__volume__profile__name` | `boot_volume_attachment.volume.profile.name` |
| `ibm_is_instance_reinitialize_invalid_boot_volume_attachment__volume__source_snapshot__id` | `boot_volume_attachment.volume.source_snapshot.id` |
| `ibm_is_instance_reinitialize_invalid_default_trusted_profile__auto_link` | `default_trusted_profile.auto_link` |
| `ibm_is_instance_reinitialize_invalid_default_trusted_profile__target__crn` | `default_trusted_profile.target.crn` |
| `ibm_is_instance_reinitialize_invalid_default_trusted_profile__target__id` | `default_trusted_profile.target.id` |

## Example

### Valid

```hcl
resource "ibm_is_instance_reinitialize" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
