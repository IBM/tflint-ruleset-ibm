# ibm_is_share_snapshot

Manages ibm_is_share_snapshot. The name for this share snapshot.

This rule validates `ibm_is_share_snapshot` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_share_snapshot" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `backup_policy_plan` | no | — |
| `captured_at` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `fingerprint` | no | — |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `minimum_size` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `share` | yes | pattern `^[-0-9a-z_]+$` |
| `share_snapshot` | no | — |
| `status` | no | — |
| `status_reasons` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_share_snapshot_invalid_access_tags` | `access_tags` |
| `ibm_is_share_snapshot_invalid_name` | `name` |
| `ibm_is_share_snapshot_invalid_share` | `share` |
| `ibm_is_share_snapshot_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_is_share_snapshot" "test" {
  share = "example-name"
}
```
