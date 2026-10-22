# ibm_container_nlb_dns

Schema validation for the ibm_container_nlb_dns resource.

This rule validates `ibm_container_nlb_dns` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/nlb-dns/deleteSecret:post` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_nlb_dns" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster` | yes | API: The cluster value. |
| `nlb_dns_type` | no | — |
| `nlb_host` | yes | — |
| `nlb_ips` | yes | — |
| `nlb_monitor_state` | no | — |
| `nlb_ssl_secret_name` | no | — |
| `nlb_ssl_secret_status` | no | — |
| `nlb_type` | no | — |
| `resource_group_id` | no | — |
| `secret_namespace` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_nlb_dns_invalid_cluster` | `cluster` |
| `ibm_container_nlb_dns_invalid_nlb_host` | `nlb_host` |
| `ibm_container_nlb_dns_invalid_nlb_ips` | `nlb_ips` |
| `ibm_container_nlb_dns_invalid_resource_group_id` | `resource_group_id` |

## Example

### Valid

```hcl
resource "ibm_container_nlb_dns" "test" {
  cluster = "aaaaaaa"
  nlb_host = "example"
  nlb_ips = "example"
}
```
