# ibm_is_public_address_range

Manages ibm_is_public_address_range. The name for this public address range.

This rule validates `ibm_is_public_address_range` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PublicGatewayFloatingIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_public_address_range" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | pattern `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$` |
| `cidr` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `href` | no | — |
| `ipv4_address_count` | yes | min 1 |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_type` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `resource_group` (block) | no | nested block |
| `resource_group.id` | no | pattern `^[0-9a-f]{32}$` |
| `target` (block) | no | nested block |
| `target.vpc` (block) | yes | nested block |
| `target.vpc.crn` | no | crn format |
| `target.vpc.href` | no | pattern `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$` |
| `target.vpc.id` | no | pattern `^[-0-9a-z_]+$` |
| `target.zone` (block) | yes | nested block |
| `target.zone.href` | no | pattern `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$` |
| `target.zone.name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_public_address_range_invalid_access_tags` | `access_tags` |
| `ibm_is_public_address_range_invalid_ipv4_address_count` | `ipv4_address_count` |
| `ibm_is_public_address_range_invalid_name` | `name` |
| `ibm_is_public_address_range_invalid_tags` | `tags` |
| `ibm_is_public_address_range_invalid_resource_group__id` | `resource_group.id` |
| `ibm_is_public_address_range_invalid_target__vpc__crn` | `target.vpc.crn` |
| `ibm_is_public_address_range_invalid_target__vpc__href` | `target.vpc.href` |
| `ibm_is_public_address_range_invalid_target__vpc__id` | `target.vpc.id` |
| `ibm_is_public_address_range_invalid_target__zone__href` | `target.zone.href` |
| `ibm_is_public_address_range_invalid_target__zone__name` | `target.zone.name` |

## Example

### Valid

```hcl
resource "ibm_is_public_address_range" "test" {
  ipv4_address_count = 1
}
```
