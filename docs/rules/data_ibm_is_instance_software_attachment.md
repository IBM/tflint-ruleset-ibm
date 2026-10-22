# ibm_is_instance_software_attachment_data

Reads ibm_is_instance_software_attachment. The name for this instance software attachment.

This rule validates `ibm_is_instance_software_attachment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_instance_software_attachment_data" {
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
| `name` | no | — |
| `offering_instance` | no | — |
| `resource_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_software_attachment_data_invalid_instance_id` | `instance_id` |
| `ibm_is_instance_software_attachment_data_invalid_instance_software_attachment_id` | `instance_software_attachment_id` |

## Example

### Valid

```hcl
data "ibm_is_instance_software_attachment" "test" {
  instance_id = "example-name"
  instance_software_attachment_id = "example-name"
}
```
