# ibm_hpcs_data

Reads ibm_hpcs. Resource instance name for example, myobjectstorage

This rule validates `ibm_hpcs` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v4/vaults/{id}:patch` (`uko`).

## Configuration

```hcl
rule "ibm_hpcs_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `extensions` | no | — |
| `failover_units` | no | — |
| `guid` | no | — |
| `hsm_info` | no | — |
| `location` | no | pattern `^[A-Za-z0-9][A-Za-z0-9 ._-]*$` |
| `name` | yes | pattern `^[A-Za-z0-9#@!$%'_-][A-Za-z0-9#@!$% '_-]*$` |
| `plan` | no | — |
| `resource_group_id` | no | — |
| `service` | no | — |
| `service_endpoints` | no | — |
| `status` | no | — |
| `units` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_hpcs_data_invalid_location` | `location` |
| `ibm_hpcs_data_invalid_name` | `name` |
| `ibm_hpcs_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_hpcs_data_invalid_service` | `service` |

## Example

### Valid

```hcl
data "ibm_hpcs" "test" {
  name = "example-name"
}
```
