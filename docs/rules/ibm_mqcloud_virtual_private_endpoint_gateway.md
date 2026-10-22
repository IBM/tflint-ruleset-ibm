# ibm_mqcloud_virtual_private_endpoint_gateway

Manages ibm_mqcloud_virtual_private_endpoint_gateway. The name of the virtual private endpoint gateway, created by the user.

This rule validates `ibm_mqcloud_virtual_private_endpoint_gateway` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `VirtualPrivateEndpointGatewayDetailsPrototype` (`mq-on-cloud`).

## Configuration

```hcl
rule "ibm_mqcloud_virtual_private_endpoint_gateway" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `href` | no | — |
| `name` | yes | pattern `^[a-z]|[a-z][-a-z0-9]*[a-z0-9]$` |
| `service_instance_guid` | yes | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `status` | no | — |
| `target_crn` | yes | pattern `^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\.\/]*$|^crn:\[\.\.\.\]$` |
| `trusted_profile` | no | pattern `^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\.\/]*$|^crn:\[\.\.\.\]$` |
| `virtual_private_endpoint_gateway_guid` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_mqcloud_virtual_private_endpoint_gateway_invalid_name` | `name` |
| `ibm_mqcloud_virtual_private_endpoint_gateway_invalid_service_instance_guid` | `service_instance_guid` |
| `ibm_mqcloud_virtual_private_endpoint_gateway_invalid_target_crn` | `target_crn` |
| `ibm_mqcloud_virtual_private_endpoint_gateway_invalid_trusted_profile` | `trusted_profile` |

## Example

### Valid

```hcl
resource "ibm_mqcloud_virtual_private_endpoint_gateway" "test" {
  name = "example-name"
  service_instance_guid = "01234567-89ab-cdef-0123-456789abcdef"
  target_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
