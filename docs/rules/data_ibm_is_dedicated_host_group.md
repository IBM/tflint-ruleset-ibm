# ibm_is_dedicated_host_group_data

Reads ibm_is_dedicated_host_group. The unique user-defined name for this dedicated host.

This rule validates `ibm_is_dedicated_host_group` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_dedicated_host_group_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `class` | no | — |
| `created_at` | no | — |
| `crn` | no | — |
| `dedicated_hosts` | no | — |
| `family` | no | — |
| `href` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | — |
| `resource_type` | no | — |
| `supported_instance_profiles` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_dedicated_host_group_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_dedicated_host_group" "test" {
  name = "example-name"
}
```
