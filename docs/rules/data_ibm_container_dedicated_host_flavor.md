# ibm_container_dedicated_host_flavor_data

Schema validation for the ibm_container_dedicated_host_flavor data source.

This rule validates `ibm_container_dedicated_host_flavor` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CreateDedicatedHost contains the fields in requests to create dedicated host instances.` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_dedicated_host_flavor_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `deprecated` | no | — |
| `flavor_class` | no | — |
| `host_flavor_id` | yes | — |
| `instance_storage` | no | — |
| `max_memory` | no | — |
| `max_vcpus` | no | — |
| `region` | no | — |
| `zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_dedicated_host_flavor_data_invalid_host_flavor_id` | `host_flavor_id` |
| `ibm_container_dedicated_host_flavor_data_invalid_zone` | `zone` |

## Example

### Valid

```hcl
data "ibm_container_dedicated_host_flavor" "test" {
  host_flavor_id = "01234567-89ab-cdef-0123-456789abcdef"
  zone = "us-south-1"
}
```
