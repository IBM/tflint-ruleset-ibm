# ibm_is_operating_system_data

Reads ibm_is_operating_system. The globally unique name for this operating system

This rule validates `ibm_is_operating_system` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_operating_system_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allow_user_image_creation` | no | — |
| `architecture` | no | — |
| `dedicated_host_only` | no | — |
| `display_name` | no | — |
| `family` | no | — |
| `href` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `user_data_format` | no | — |
| `vendor` | no | — |
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_operating_system_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_operating_system" "test" {
  name = "example-name"
}
```
