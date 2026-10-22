# ibm_org_quota_data

Reads ibm_org_quota. Org quota name, for example qIBM

This rule validates `ibm_org_quota` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_org_quota_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `app_instance_limit` | no | — |
| `app_tasks_limit` | no | — |
| `instance_memory_limit` | no | — |
| `memory_limit` | no | — |
| `name` | yes | — |
| `non_basic_services_allowed` | no | — |
| `total_private_domains` | no | — |
| `total_reserved_route_ports` | no | — |
| `total_routes` | no | — |
| `total_service_keys` | no | — |
| `total_services` | no | — |
| `trial_db_allowed` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_org_quota_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_org_quota" "test" {
  name = "example"
}
```
