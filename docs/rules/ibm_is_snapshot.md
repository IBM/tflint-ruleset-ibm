# ibm_is_snapshot

Manages ibm_is_snapshot. Snapshot name

This rule validates `ibm_is_snapshot` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Snapshot` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_snapshot" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `backup_policy_plan` | no | — |
| `bootable` | no | — |
| `catalog_offering` | no | — |
| `clones` | no | API: The clones for this snapshot |
| `copies` | no | — |
| `crn` | no | — |
| `encryption` | no | — |
| `encryption_key` | no | API: The root key used to wrap the data encryption key for the source volume. This property will be present for volumes with an `encryption` type of `user_managed`. |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `minimum_capacity` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `operating_system` | no | — |
| `resource_group` | no | API: The resource group for this snapshot |
| `resource_type` | no | — |
| `service_tags` | no | — |
| `size` | no | — |
| `snapshot_consistency_group` | no | — |
| `source_image` | no | — |
| `source_snapshot` | no | — |
| `source_snapshot_crn` | no | crn format |
| `source_volume` | no | API: The source volume this snapshot was created from (may be [deleted](https://cloud.ibm.com/apidocs/vpc#deleted-resources)). |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `allowed_use` (block) | no | nested block |
| `allowed_use.api_version` | no | pattern `^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$` |
| `allowed_use.bare_metal_server` | no | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `allowed_use.instance` | no | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_snapshot_invalid_access_tags` | `access_tags` |
| `ibm_is_snapshot_invalid_clones` | `clones` |
| `ibm_is_snapshot_invalid_encryption_key` | `encryption_key` |
| `ibm_is_snapshot_invalid_name` | `name` |
| `ibm_is_snapshot_invalid_resource_group` | `resource_group` |
| `ibm_is_snapshot_invalid_source_snapshot_crn` | `source_snapshot_crn` |
| `ibm_is_snapshot_invalid_source_volume` | `source_volume` |
| `ibm_is_snapshot_invalid_tags` | `tags` |
| `ibm_is_snapshot_invalid_allowed_use__api_version` | `allowed_use.api_version` |
| `ibm_is_snapshot_invalid_allowed_use__bare_metal_server` | `allowed_use.bare_metal_server` |
| `ibm_is_snapshot_invalid_allowed_use__instance` | `allowed_use.instance` |

## Example

### Valid

```hcl
resource "ibm_is_snapshot" "test" {
}
```
