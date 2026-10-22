# ibm_is_instance_volume_attachment

Manages ibm_is_instance_volume_attachment. The user-defined name for this volume attachment.

This rule validates `ibm_is_instance_volume_attachment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Volume` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_volume_attachment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bandwidth` | no | range 100–8192 |
| `capacity` | no | range 10–64000 |
| `delete_volume_on_attachment_delete` | no | — |
| `delete_volume_on_instance_delete` | no | API: Indicates whether deleting the instance will also delete the attached volume |
| `device` | no | — |
| `encryption_key` | no | API: The root key used to wrap the data encryption key for the volume. This property will be present for volumes with an `encryption` type of `user_managed`. |
| `href` | no | — |
| `instance` | yes | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `iops` | no | API: The maximum I/O operations per second (IOPS) for this volume. |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `profile` | no | one of: `general-purpose`, `5iops-tier`, `10iops-tier`, `custom`, `sdp` |
| `snapshot` | no | — |
| `snapshot_crn` | no | crn format |
| `status` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `type` | no | — |
| `version` | no | — |
| `volume` | no | — |
| `volume_attachment_id` | no | — |
| `volume_crn` | no | — |
| `volume_deleted` | no | — |
| `volume_href` | no | — |
| `volume_name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `allowed_use` (block) | no | nested block |
| `allowed_use.api_version` | no | pattern `^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$` |
| `allowed_use.bare_metal_server` | no | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `allowed_use.instance` | no | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_volume_attachment_invalid_bandwidth` | `bandwidth` |
| `ibm_is_instance_volume_attachment_invalid_capacity` | `capacity` |
| `ibm_is_instance_volume_attachment_invalid_delete_volume_on_attachment_delete` | `delete_volume_on_attachment_delete` |
| `ibm_is_instance_volume_attachment_invalid_delete_volume_on_instance_delete` | `delete_volume_on_instance_delete` |
| `ibm_is_instance_volume_attachment_invalid_encryption_key` | `encryption_key` |
| `ibm_is_instance_volume_attachment_invalid_instance` | `instance` |
| `ibm_is_instance_volume_attachment_invalid_iops` | `iops` |
| `ibm_is_instance_volume_attachment_invalid_name` | `name` |
| `ibm_is_instance_volume_attachment_invalid_profile` | `profile` |
| `ibm_is_instance_volume_attachment_invalid_snapshot` | `snapshot` |
| `ibm_is_instance_volume_attachment_invalid_snapshot_crn` | `snapshot_crn` |
| `ibm_is_instance_volume_attachment_invalid_tags` | `tags` |
| `ibm_is_instance_volume_attachment_invalid_volume` | `volume` |
| `ibm_is_instance_volume_attachment_invalid_volume_name` | `volume_name` |
| `ibm_is_instance_volume_attachment_invalid_allowed_use__api_version` | `allowed_use.api_version` |
| `ibm_is_instance_volume_attachment_invalid_allowed_use__bare_metal_server` | `allowed_use.bare_metal_server` |
| `ibm_is_instance_volume_attachment_invalid_allowed_use__instance` | `allowed_use.instance` |

## Example

### Valid

```hcl
resource "ibm_is_instance_volume_attachment" "test" {
  instance = "example-name"
}
```
