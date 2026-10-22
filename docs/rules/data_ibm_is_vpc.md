# ibm_is_vpc_data

Schema validation for the ibm_is_vpc data source.

This rule validates `ibm_is_vpc` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `BareMetalServerDiskPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_vpc_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_tags` | no | — |
| `classic_access` | no | — |
| `crn` | no | — |
| `cse_source_addresses` | no | — |
| `default_address_prefixes` | no | — |
| `default_network_acl` | no | — |
| `default_network_acl_crn` | no | — |
| `default_network_acl_name` | no | — |
| `default_routing_table` | no | — |
| `default_routing_table_name` | no | — |
| `default_security_group` | no | — |
| `default_security_group_crn` | no | — |
| `default_security_group_name` | no | — |
| `dns` | no | — |
| `health_reasons` | no | — |
| `health_state` | no | — |
| `identifier` | no | — |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$` |
| `public_address_ranges` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `security_group` | no | — |
| `status` | no | — |
| `subnets` | no | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_vpc_data_invalid_identifier` | `identifier` |
| `ibm_is_vpc_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_is_vpc" "test" {
}
```
