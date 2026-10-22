# ibm_is_dedicated_host_disk_data

Reads ibm_is_dedicated_host_disk. The user-defined or system-provided name for this disk.

This rule validates `ibm_is_dedicated_host_disk` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_dedicated_host_disk_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `available` | no | — |
| `created_at` | no | — |
| `dedicated_host` | yes | — |
| `disk` | yes | — |
| `href` | no | — |
| `instance_disks` | no | — |
| `interface_type` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | — |
| `provisionable` | no | — |
| `resource_type` | no | — |
| `size` | no | — |
| `supported_instance_interface_types` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_dedicated_host_disk_data_invalid_dedicated_host` | `dedicated_host` |
| `ibm_is_dedicated_host_disk_data_invalid_disk` | `disk` |

## Example

### Valid

```hcl
data "ibm_is_dedicated_host_disk" "test" {
  dedicated_host = "example"
  disk = "example"
}
```
