# ibm_mqcloud_virtual_private_endpoint_gateway_data

Reads ibm_mqcloud_virtual_private_endpoint_gateway. The name of the virtual private endpoint gateway, created by the user.

This rule validates `ibm_mqcloud_virtual_private_endpoint_gateway` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_mqcloud_virtual_private_endpoint_gateway_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `href` | no | — |
| `name` | no | — |
| `service_instance_guid` | yes | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `status` | no | — |
| `target_crn` | no | — |
| `trusted_profile` | no | pattern `^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\.\/]*$|^crn:\[\.\.\.\]$` |
| `virtual_private_endpoint_gateway_guid` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_mqcloud_virtual_private_endpoint_gateway_data_invalid_service_instance_guid` | `service_instance_guid` |
| `ibm_mqcloud_virtual_private_endpoint_gateway_data_invalid_trusted_profile` | `trusted_profile` |
| `ibm_mqcloud_virtual_private_endpoint_gateway_data_invalid_virtual_private_endpoint_gateway_guid` | `virtual_private_endpoint_gateway_guid` |

## Example

### Valid

```hcl
data "ibm_mqcloud_virtual_private_endpoint_gateway" "test" {
  service_instance_guid = "01234567-89ab-cdef-0123-456789abcdef"
  virtual_private_endpoint_gateway_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
