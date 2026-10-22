# ibm_dns_glb_monitors_data

Schema validation for the ibm_dns_glb_monitors data source.

This rule validates `ibm_dns_glb_monitors` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `list_monitors` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_glb_monitors_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `dns_glb_monitors` | no | — |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_glb_monitors_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_dns_glb_monitors" "test" {
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
}
```
