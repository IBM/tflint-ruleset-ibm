# ibm_lbaas_data

Schema validation for the ibm_lbaas data source.

This rule validates `ibm_lbaas` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_lbaas_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `active_connections` | no | — |
| `datacenter` | no | — |
| `description` | no | — |
| `health_monitors` | no | — |
| `name` | yes | — |
| `protocols` | no | — |
| `server_instances` | no | — |
| `server_instances_down` | no | — |
| `server_instances_up` | no | — |
| `ssl_ciphers` | no | — |
| `status` | no | — |
| `type` | no | — |
| `use_system_public_ip_pool` | no | — |
| `vip` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_lbaas_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_lbaas" "test" {
  name = "example"
}
```
