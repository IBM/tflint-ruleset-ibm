# ibm_satellite_endpoint_data

Schema validation for the ibm_satellite_endpoint data source.

This rule validates `ibm_satellite_endpoint` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_endpoint_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `certs` | no | — |
| `client_host` | no | — |
| `client_mutual_auth` | no | — |
| `client_port` | no | — |
| `client_protocol` | no | — |
| `connection_type` | no | — |
| `connector_port` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `display_name` | no | — |
| `endpoint_id` | yes | — |
| `last_change` | no | — |
| `location` | yes | — |
| `performance` | no | — |
| `reject_unauth` | no | — |
| `server_host` | no | — |
| `server_mutual_auth` | no | — |
| `server_port` | no | — |
| `server_protocol` | no | — |
| `service_name` | no | — |
| `sni` | no | — |
| `sources` | no | — |
| `status` | no | — |
| `timeout` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_endpoint_data_invalid_endpoint_id` | `endpoint_id` |
| `ibm_satellite_endpoint_data_invalid_location` | `location` |

## Example

### Valid

```hcl
data "ibm_satellite_endpoint" "test" {
  endpoint_id = "01234567-89ab-cdef-0123-456789abcdef"
  location = "example"
}
```
