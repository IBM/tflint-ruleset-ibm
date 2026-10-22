# ibm_is_instance_software_attachment

Manages ibm_is_instance_software_attachment. The name for this instance software attachment.

This rule validates `ibm_is_instance_software_attachment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_software_attachment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `catalog_offering` | no | — |
| `created_at` | no | — |
| `entitlement` | no | — |
| `href` | no | — |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `instance_software_attachment_id` | yes | pattern `^[-0-9a-z_]+$` |
| `lifecycle_reasons` | no | — |
| `lifecycle_state` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `offering_instance` | no | — |
| `resource_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_software_attachment_invalid_instance_id` | `instance_id` |
| `ibm_is_instance_software_attachment_invalid_instance_software_attachment_id` | `instance_software_attachment_id` |
| `ibm_is_instance_software_attachment_invalid_name` | `name` |

## Example

### Valid

```hcl
resource "ibm_is_instance_software_attachment" "test" {
  instance_id = "example-name"
  instance_software_attachment_id = "example-name"
}
```
