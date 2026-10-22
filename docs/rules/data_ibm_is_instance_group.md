# ibm_is_instance_group_data

Reads ibm_is_instance_group. The user-defined name for this instance group

This rule validates `ibm_is_instance_group` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_group_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `application_port` | no | — |
| `crn` | no | — |
| `instance_template` | no | — |
| `load_balancer_pool` | no | — |
| `managers` | no | — |
| `membership_count` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `resource_group` | no | — |
| `status` | no | — |
| `subnets` | no | — |
| `vpc` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_group_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_instance_group" "test" {
  name = "example-name"
}
```
