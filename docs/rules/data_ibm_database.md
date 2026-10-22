# ibm_database_data

Reads ibm_database. Resource instance name for example, my Database instance

This rule validates `ibm_database` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_database_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `adminpassword` | no | — |
| `adminuser` | no | — |
| `allowlist` | no | — |
| `auto_scaling` | no | — |
| `configuration_schema` | no | — |
| `groups` | no | — |
| `guid` | no | — |
| `location` | no | — |
| `name` | yes | — |
| `plan` | no | — |
| `platform_options` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `service` | no | one of: `databases-for-etcd`, `databases-for-postgresql`, `databases-for-redis`, `databases-for-valkey`, `databases-for-valkey-cdp-dev`, `databases-for-elasticsearch`, `databases-for-mongodb`, `messages-for-rabbitmq`, `databases-for-mysql`, `databases-for-enterprisedb` |
| `status` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |
| `users` | no | — |
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_database_data_invalid_location` | `location` |
| `ibm_database_data_invalid_name` | `name` |
| `ibm_database_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_database_data_invalid_service` | `service` |
| `ibm_database_data_invalid_tags` | `tags` |

## Example

### Valid

```hcl
data "ibm_database" "test" {
  name = "example"
}
```
