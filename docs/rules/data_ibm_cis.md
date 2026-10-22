# ibm_cis_data

Reads ibm_cis. Resource instance name for example, my cis instance

This rule validates `ibm_cis` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/{crn}/zones:post` (`cis`).

## Configuration

```hcl
rule "ibm_cis_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `guid` | no | — |
| `location` | no | — |
| `name` | yes | API: Name |
| `plan` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `service` | no | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_data_invalid_name` | `name` |
| `ibm_cis_data_invalid_resource_group_id` | `resource_group_id` |

## Example

### Valid

```hcl
data "ibm_cis" "test" {
  name = "example"
}
```
