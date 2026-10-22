# ibm_is_dedicated_host_profile_data

Reads ibm_is_dedicated_host_profile. The globally unique name for this virtual server instance profile.

This rule validates `ibm_is_dedicated_host_profile` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_dedicated_host_profile_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `class` | no | — |
| `disks` | no | — |
| `family` | no | — |
| `href` | no | — |
| `memory` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `socket_count` | no | — |
| `status` | no | — |
| `supported_instance_profiles` | no | — |
| `vcpu_architecture` | no | — |
| `vcpu_count` | no | — |
| `vcpu_manufacturer` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_dedicated_host_profile_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_dedicated_host_profile" "test" {
}
```
