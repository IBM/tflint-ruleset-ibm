# ibm_satellite_endpoint

Schema validation for the ibm_satellite_endpoint resource.

This rule validates `ibm_satellite_endpoint` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_endpoint" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `client_host` | no | — |
| `client_mutual_auth` | no | — |
| `client_port` | no | — |
| `client_protocol` | yes | one of: `http`, `http-tunnel`, `https`, `tcp`, `tls`, `udp` |
| `connection_type` | yes | one of: `cloud`, `location` |
| `connector_port` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `display_name` | yes | — |
| `endpoint_id` | no | — |
| `last_change` | no | — |
| `location` | yes | — |
| `performance` | no | — |
| `reject_unauth` | no | — |
| `server_host` | yes | — |
| `server_mutual_auth` | no | — |
| `server_port` | yes | range 1–65535; port format |
| `server_protocol` | no | one of: `tcp`, `tls`, `udp` |
| `service_name` | no | — |
| `sni` | no | — |
| `sources` | no | — |
| `status` | no | — |
| `timeout` | no | range 1–180 |
| `certs` (block) | no | nested block |
| `certs.client` (block) | no | nested block |
| `certs.client.cert` (block) | no | nested block |
| `certs.client.cert.file_contents` | no | — |
| `certs.client.cert.filename` | no | — |
| `certs.connector` (block) | no | nested block |
| `certs.connector.cert` (block) | no | nested block |
| `certs.connector.cert.file_contents` | no | — |
| `certs.connector.cert.filename` | no | — |
| `certs.connector.key` (block) | no | nested block |
| `certs.connector.key.file_contents` | no | — |
| `certs.connector.key.filename` | no | — |
| `certs.server` (block) | no | nested block |
| `certs.server.cert` (block) | no | nested block |
| `certs.server.cert.file_contents` | no | — |
| `certs.server.cert.filename` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_endpoint_invalid_client_mutual_auth` | `client_mutual_auth` |
| `ibm_satellite_endpoint_invalid_client_protocol` | `client_protocol` |
| `ibm_satellite_endpoint_invalid_connection_type` | `connection_type` |
| `ibm_satellite_endpoint_invalid_created_by` | `created_by` |
| `ibm_satellite_endpoint_invalid_display_name` | `display_name` |
| `ibm_satellite_endpoint_invalid_location` | `location` |
| `ibm_satellite_endpoint_invalid_reject_unauth` | `reject_unauth` |
| `ibm_satellite_endpoint_invalid_server_host` | `server_host` |
| `ibm_satellite_endpoint_invalid_server_mutual_auth` | `server_mutual_auth` |
| `ibm_satellite_endpoint_invalid_server_port` | `server_port` |
| `ibm_satellite_endpoint_invalid_server_protocol` | `server_protocol` |
| `ibm_satellite_endpoint_invalid_sni` | `sni` |
| `ibm_satellite_endpoint_invalid_timeout` | `timeout` |
| `ibm_satellite_endpoint_invalid_certs__client__cert__file_contents` | `certs.client.cert.file_contents` |
| `ibm_satellite_endpoint_invalid_certs__client__cert__filename` | `certs.client.cert.filename` |
| `ibm_satellite_endpoint_invalid_certs__connector__cert__file_contents` | `certs.connector.cert.file_contents` |
| `ibm_satellite_endpoint_invalid_certs__connector__cert__filename` | `certs.connector.cert.filename` |
| `ibm_satellite_endpoint_invalid_certs__connector__key__file_contents` | `certs.connector.key.file_contents` |
| `ibm_satellite_endpoint_invalid_certs__connector__key__filename` | `certs.connector.key.filename` |
| `ibm_satellite_endpoint_invalid_certs__server__cert__file_contents` | `certs.server.cert.file_contents` |
| `ibm_satellite_endpoint_invalid_certs__server__cert__filename` | `certs.server.cert.filename` |

## Example

### Valid

```hcl
resource "ibm_satellite_endpoint" "test" {
  client_protocol = "http"
  connection_type = "cloud"
  display_name = "example"
  location = "example"
  server_host = "example"
  server_port = 80
}
```
