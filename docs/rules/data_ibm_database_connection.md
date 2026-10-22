# ibm_database_connection_data

Schema validation for the ibm_database_connection data source.

This rule validates `ibm_database_connection` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:post:/deployments/{id}/users/{user_type}/{user_id}/connections/{endpoint_type}` (`cloud-databases-api-cloud-databases-api-v5`).

## Configuration

```hcl
rule "ibm_database_connection_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `amqps` | no | — |
| `analytics` | no | — |
| `bi_connector` | no | — |
| `certificate_root` | no | — |
| `cli` | no | — |
| `deployment_id` | yes | — |
| `emp` | no | — |
| `endpoint_type` | yes | one of: `public`, `private`, `public-and-private` |
| `grpc` | no | — |
| `https` | no | — |
| `mongodb` | no | — |
| `mqtts` | no | — |
| `mysql` | no | — |
| `ops_manager` | no | — |
| `postgres` | no | — |
| `rediss` | no | — |
| `secure` | no | — |
| `stomp_ssl` | no | — |
| `user_id` | yes | — |
| `user_type` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_database_connection_data_invalid_certificate_root` | `certificate_root` |
| `ibm_database_connection_data_invalid_deployment_id` | `deployment_id` |
| `ibm_database_connection_data_invalid_endpoint_type` | `endpoint_type` |
| `ibm_database_connection_data_invalid_user_id` | `user_id` |
| `ibm_database_connection_data_invalid_user_type` | `user_type` |

## Example

### Valid

```hcl
data "ibm_database_connection" "test" {
  deployment_id = "01234567-89ab-cdef-0123-456789abcdef"
  endpoint_type = "public"
  user_id = "01234567-89ab-cdef-0123-456789abcdef"
  user_type = "example"
}
```
