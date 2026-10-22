# ibm_is_instance_disk_management

Schema validation for the ibm_is_instance_disk_management resource.

This rule validates `ibm_is_instance_disk_management` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `InstanceClusterNetworkAttachmentBefore` (`vpc-metadata-latest`).

## Configuration

```hcl
rule "ibm_is_instance_disk_management" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance` | yes | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `disks` (block) | yes | nested block |
| `disks.id` | yes | pattern `^[-0-9a-z_]+$` |
| `disks.name` | yes | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_disk_management_invalid_instance` | `instance` |
| `ibm_is_instance_disk_management_invalid_disks__id` | `disks.id` |
| `ibm_is_instance_disk_management_invalid_disks__name` | `disks.name` |

## Example

### Valid

```hcl
resource "ibm_is_instance_disk_management" "test" {
  instance = "example-name"
  disks {
    id = "example-name"
    name = "-01234567-89ab-cdef-0123-456789abcdef"
  }
}
```
