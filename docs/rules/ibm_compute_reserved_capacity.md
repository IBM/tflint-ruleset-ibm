# ibm_compute_reserved_capacity

Manages ibm_compute_reserved_capacity. Name

This rule validates `ibm_compute_reserved_capacity` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_reserved_capacity" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `datacenter` | yes | — |
| `flavor` | yes | — |
| `force_create` | no | — |
| `instances` | yes | — |
| `name` | yes | — |
| `pod` | yes | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_reserved_capacity_invalid_datacenter` | `datacenter` |
| `ibm_compute_reserved_capacity_invalid_flavor` | `flavor` |
| `ibm_compute_reserved_capacity_invalid_force_create` | `force_create` |
| `ibm_compute_reserved_capacity_invalid_instances` | `instances` |
| `ibm_compute_reserved_capacity_invalid_name` | `name` |
| `ibm_compute_reserved_capacity_invalid_pod` | `pod` |
| `ibm_compute_reserved_capacity_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_compute_reserved_capacity" "test" {
  datacenter = "example"
  flavor = "example"
  instances = 1
  name = "example"
  pod = "example"
}
```
