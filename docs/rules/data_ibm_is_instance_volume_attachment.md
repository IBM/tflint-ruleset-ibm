# ibm_is_instance_volume_attachment_data

Reads ibm_is_instance_volume_attachment. The user-defined name for this volume attachment.

This rule validates `ibm_is_instance_volume_attachment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_volume_attachment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bandwidth` | no | — |
| `delete_volume_on_instance_delete` | no | — |
| `device` | no | — |
| `href` | no | — |
| `instance` | yes | pattern `^([a-zA-Z_][a-zA-Z0-9_]*|[-+*/%]|&&|\|\||!|==|!=|<|<=|>|>=|~|\bin\b|\(|\)|\[|\]|,|\.|"|'|"|'|\s+|\d+)+$` |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `status` | no | — |
| `type` | no | — |
| `volume_attachment_id` | no | — |
| `volume_reference` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_volume_attachment_data_invalid_instance` | `instance` |
| `ibm_is_instance_volume_attachment_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_instance_volume_attachment" "test" {
  instance = "example-name"
  name = "example-name"
}
```
