# ibm_is_instance_network_attachments_data

Schema validation for the ibm_is_instance_network_attachments data source.

This rule validates `ibm_is_instance_network_attachments` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerDNSPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_network_attachments_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance` | yes | API: Identifies a DNS instance by a unique property. |
| `network_attachments` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_network_attachments_data_invalid_instance` | `instance` |

## Example

### Valid

```hcl
data "ibm_is_instance_network_attachments" "test" {
  instance = "example"
}
```
