# ibm_is_bare_metal_server_capacities_data

Schema validation for the ibm_is_bare_metal_server_capacities data source.

This rule validates `ibm_is_bare_metal_server_capacities` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerCapacity` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_bare_metal_server_capacities_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `capacities` | no | — |
| `profile` | no | API: The [profile](https://cloud.ibm.com/docs/vpc?topic=vpc-bare-metal-servers-profile) available in the `zone`. |
| `zone` | no | zone format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_bare_metal_server_capacities_data_invalid_profile` | `profile` |
| `ibm_is_bare_metal_server_capacities_data_invalid_zone` | `zone` |

## Example

### Valid

```hcl
data "ibm_is_bare_metal_server_capacities" "test" {
}
```
