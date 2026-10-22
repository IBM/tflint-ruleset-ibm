# ibm_dns_glb_pool

Manages ibm_dns_glb_pool. The unique identifier of a service instance.

This rule validates `ibm_dns_glb_pool` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `pool_input` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_glb_pool" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_on` | no | — |
| `description` | no | API: Descriptive text of the load balancer pool. |
| `enabled` | no | API: Whether the load balancer pool is enabled. |
| `health` | no | — |
| `healthcheck_region` | no | one of: ``, `us-south`, `us-east`, `eu-gb`, `eu-du`, `au-syd`, `jp-tok` |
| `healthcheck_subnets` | no | API: Subnet CRN. |
| `healthy_origins_threshold` | no | API: The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins falls below this number, the pool will be marked unhealthy and we will failover to the next available pool. |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `modified_on` | no | — |
| `monitor` | no | API: The ID of the load balancer monitor to be associated to this pool. |
| `name` | yes | API: Name of the load balancer pool. |
| `notification_channel` | no | API: The notification channel. |
| `pool_id` | no | — |
| `origins` (block) | yes | nested block |
| `origins.address` | yes | API: The address of the origin server. It can be a hostname or an IP address. |
| `origins.description` | no | API: Description of the origin server. |
| `origins.enabled` | yes | API: Whether the origin server is enabled. |
| `origins.name` | yes | API: The name of the origin server. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_glb_pool_invalid_description` | `description` |
| `ibm_dns_glb_pool_invalid_enabled` | `enabled` |
| `ibm_dns_glb_pool_invalid_healthcheck_region` | `healthcheck_region` |
| `ibm_dns_glb_pool_invalid_healthcheck_subnets` | `healthcheck_subnets` |
| `ibm_dns_glb_pool_invalid_healthy_origins_threshold` | `healthy_origins_threshold` |
| `ibm_dns_glb_pool_invalid_instance_id` | `instance_id` |
| `ibm_dns_glb_pool_invalid_monitor` | `monitor` |
| `ibm_dns_glb_pool_invalid_name` | `name` |
| `ibm_dns_glb_pool_invalid_notification_channel` | `notification_channel` |
| `ibm_dns_glb_pool_invalid_origins__address` | `origins.address` |
| `ibm_dns_glb_pool_invalid_origins__description` | `origins.description` |
| `ibm_dns_glb_pool_invalid_origins__enabled` | `origins.enabled` |
| `ibm_dns_glb_pool_invalid_origins__name` | `origins.name` |

## Example

### Valid

```hcl
resource "ibm_dns_glb_pool" "test" {
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  name = "example"
  origins {
    address = "10.0.0.1"
    enabled = true
    name = "example"
  }
}
```
