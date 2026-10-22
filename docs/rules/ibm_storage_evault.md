# ibm_storage_evault

Schema validation for the ibm_storage_evault resource.

This rule validates `ibm_storage_evault` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_storage_evault" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `capacity` | yes | — |
| `datacenter` | yes | — |
| `hardware_instance_id` | no | — |
| `password` | no | — |
| `service_resource_name` | no | — |
| `tags` | no | — |
| `username` | no | — |
| `virtual_instance_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_storage_evault_invalid_capacity` | `capacity` |
| `ibm_storage_evault_invalid_datacenter` | `datacenter` |
| `ibm_storage_evault_invalid_hardware_instance_id` | `hardware_instance_id` |
| `ibm_storage_evault_invalid_tags` | `tags` |
| `ibm_storage_evault_invalid_virtual_instance_id` | `virtual_instance_id` |

## Example

### Valid

```hcl
resource "ibm_storage_evault" "test" {
  capacity = 1
  datacenter = "example"
}
```
