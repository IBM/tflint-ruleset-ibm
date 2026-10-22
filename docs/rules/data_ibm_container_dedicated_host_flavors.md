# ibm_container_dedicated_host_flavors_data

Schema validation for the ibm_container_dedicated_host_flavors data source.

This rule validates `ibm_container_dedicated_host_flavors` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CreateDedicatedHost contains the fields in requests to create dedicated host instances.` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_dedicated_host_flavors_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `host_flavors` | no | — |
| `zone` | yes | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_dedicated_host_flavors_data_invalid_zone` | `zone` |

## Example

### Valid

```hcl
data "ibm_container_dedicated_host_flavors" "test" {
  zone = "us-south-1"
}
```
