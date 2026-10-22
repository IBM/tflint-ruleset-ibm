# ibm_is_vpc_dns_resolution_binding

Manages ibm_is_vpc_dns_resolution_binding. The name for this DNS resolution binding.

This rule validates `ibm_is_vpc_dns_resolution_binding` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `FloatingIPTargetPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpc_dns_resolution_binding" {
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
| `lifecycle_state` | no | — |
| `name` | no | — |
| `resource_type` | no | — |
| `vpc_id` | yes | — |
| `vpc` (block) | yes | nested block |
| `vpc.crn` | no | crn format |
| `vpc.href` | no | API: The URL for this virtual network interface |
| `vpc.id` | no | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpc_dns_resolution_binding_invalid_name` | `name` |
| `ibm_is_vpc_dns_resolution_binding_invalid_vpc_id` | `vpc_id` |
| `ibm_is_vpc_dns_resolution_binding_invalid_vpc__crn` | `vpc.crn` |
| `ibm_is_vpc_dns_resolution_binding_invalid_vpc__href` | `vpc.href` |
| `ibm_is_vpc_dns_resolution_binding_invalid_vpc__id` | `vpc.id` |

## Example

### Valid

```hcl
resource "ibm_is_vpc_dns_resolution_binding" "test" {
  vpc_id = "01234567-89ab-cdef-0123-456789abcdef"
  vpc {
  }
}
```
