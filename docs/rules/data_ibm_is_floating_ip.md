# ibm_is_floating_ip_data

Reads ibm_is_floating_ip. Name of the floating IP

This rule validates `ibm_is_floating_ip` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_floating_ip_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `address` | no | — |
| `crn` | no | — |
| `name` | yes | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `status` | no | — |
| `tags` | no | — |
| `target` | no | — |
| `target_list` | no | — |
| `zone` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_floating_ip_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_floating_ip" "test" {
  name = "example-name"
}
```
