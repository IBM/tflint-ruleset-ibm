# ibm_is_security_group_target_data

Reads ibm_is_security_group_target. Security group target name

This rule validates `ibm_is_security_group_target` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_security_group_target_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `more_info` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `resource_type` | no | — |
| `security_group` | yes | pattern `^[-0-9a-z_]+$` |
| `target` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_security_group_target_data_invalid_name` | `name` |
| `ibm_is_security_group_target_data_invalid_security_group` | `security_group` |

## Example

### Valid

```hcl
data "ibm_is_security_group_target" "test" {
  name = "example-name"
  security_group = "example-name"
}
```
