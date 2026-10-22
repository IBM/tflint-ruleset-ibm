# ibm_is_share_profile_data

Reads ibm_is_share_profile. The file share profile name.

This rule validates `ibm_is_share_profile` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_share_profile_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allowed_access_protocols` | no | — |
| `allowed_transit_encryption_modes` | no | — |
| `availability_modes` | no | — |
| `bandwidth` | no | — |
| `capacity` | no | — |
| `family` | no | — |
| `href` | no | — |
| `iops` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_share_profile_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_share_profile" "test" {
  name = "example-name"
}
```
