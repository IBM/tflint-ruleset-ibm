# ibm_is_vpc_dns_resolution_binding_data

Reads ibm_is_vpc_dns_resolution_binding. The name for this DNS resolution binding.

This rule validates `ibm_is_vpc_dns_resolution_binding` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_vpc_dns_resolution_binding_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `endpoint_gateways` | no | — |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `href` | no | — |
| `identifier` | yes | — |
| `lifecycle_state` | no | — |
| `name` | no | — |
| `resource_type` | no | — |
| `vpc` | no | — |
| `vpc_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpc_dns_resolution_binding_data_invalid_identifier` | `identifier` |
| `ibm_is_vpc_dns_resolution_binding_data_invalid_vpc_id` | `vpc_id` |

## Example

### Valid

```hcl
data "ibm_is_vpc_dns_resolution_binding" "test" {
  identifier = "example"
  vpc_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
