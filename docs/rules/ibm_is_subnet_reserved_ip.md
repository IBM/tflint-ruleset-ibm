# ibm_is_subnet_reserved_ip

Manages ibm_is_subnet_reserved_ip. The user-defined or system-provided name for this reserved IP.

This rule validates `ibm_is_subnet_reserved_ip` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ReservedIPPrototype` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_subnet_reserved_ip" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `address` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$` |
| `auto_delete` | no | API: Indicates whether this reserved IP member will be automatically deleted when either `target` is deleted, or the reserved IP is unbound. Must be `false` if the reserved IP is unbound. |
| `created_at` | no | — |
| `href` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `owner` | no | — |
| `reserved_ip` | no | — |
| `resource_type` | no | — |
| `subnet` | yes | — |
| `target` | no | API: The target to bind this reserved IP to. The target must be in the same VPC. The following targets are supported: - An endpoint gateway not already bound to a reserved IP in the subnet's zone. - A virtual network interface. If unspecified, the reserved IP will be created unbound. |
| `target_crn` | no | crn format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_subnet_reserved_ip_invalid_address` | `address` |
| `ibm_is_subnet_reserved_ip_invalid_auto_delete` | `auto_delete` |
| `ibm_is_subnet_reserved_ip_invalid_name` | `name` |
| `ibm_is_subnet_reserved_ip_invalid_subnet` | `subnet` |
| `ibm_is_subnet_reserved_ip_invalid_target` | `target` |
| `ibm_is_subnet_reserved_ip_invalid_target_crn` | `target_crn` |

## Example

### Valid

```hcl
resource "ibm_is_subnet_reserved_ip" "test" {
  subnet = "example"
}
```
