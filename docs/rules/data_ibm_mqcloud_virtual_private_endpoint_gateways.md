# ibm_mqcloud_virtual_private_endpoint_gateways_data

Reads ibm_mqcloud_virtual_private_endpoint_gateways. The name of the virtual private endpoint gateway, created by the user.

This rule validates `ibm_mqcloud_virtual_private_endpoint_gateways` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ApplicationDetailsPrototype` (`mq-on-cloud`).

## Configuration

```hcl
rule "ibm_mqcloud_virtual_private_endpoint_gateways_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | no | pattern `^[a-z][-a-z0-9]*$` |
| `service_instance_guid` | yes | — |
| `trusted_profile` | no | — |
| `virtual_private_endpoint_gateways` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_mqcloud_virtual_private_endpoint_gateways_data_invalid_name` | `name` |
| `ibm_mqcloud_virtual_private_endpoint_gateways_data_invalid_service_instance_guid` | `service_instance_guid` |
| `ibm_mqcloud_virtual_private_endpoint_gateways_data_invalid_trusted_profile` | `trusted_profile` |

## Example

### Valid

```hcl
data "ibm_mqcloud_virtual_private_endpoint_gateways" "test" {
  service_instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
