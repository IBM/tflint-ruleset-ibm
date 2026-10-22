# ibm_lbaas_server_instance_attachment

Schema validation for the ibm_lbaas_server_instance_attachment resource.

This rule validates `ibm_lbaas_server_instance_attachment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_lbaas_server_instance_attachment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `lbaas_id` | yes | — |
| `private_ip_address` | yes | — |
| `uuid` | no | — |
| `weight` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_lbaas_server_instance_attachment_invalid_lbaas_id` | `lbaas_id` |
| `ibm_lbaas_server_instance_attachment_invalid_private_ip_address` | `private_ip_address` |
| `ibm_lbaas_server_instance_attachment_invalid_weight` | `weight` |

## Example

### Valid

```hcl
resource "ibm_lbaas_server_instance_attachment" "test" {
  lbaas_id = "01234567-89ab-cdef-0123-456789abcdef"
  private_ip_address = "example"
}
```
