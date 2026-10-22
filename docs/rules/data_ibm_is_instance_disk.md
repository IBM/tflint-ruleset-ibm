# ibm_is_instance_disk_data

Reads ibm_is_instance_disk. The user-defined name for this disk.

This rule validates `ibm_is_instance_disk` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerDNSPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_disk_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `disk` | yes | — |
| `href` | no | — |
| `instance` | yes | API: Identifies a DNS instance by a unique property. |
| `interface_type` | no | — |
| `name` | no | — |
| `resource_type` | no | — |
| `size` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_disk_data_invalid_disk` | `disk` |
| `ibm_is_instance_disk_data_invalid_instance` | `instance` |

## Example

### Valid

```hcl
data "ibm_is_instance_disk" "test" {
  disk = "example"
  instance = "example"
}
```
