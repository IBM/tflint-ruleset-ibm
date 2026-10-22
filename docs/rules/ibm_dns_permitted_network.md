# ibm_dns_permitted_network

Schema validation for the ibm_dns_permitted_network resource.

This rule validates `ibm_dns_permitted_network` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `permitted_network_input` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_permitted_network" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_on` | no | — |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `modified_on` | no | — |
| `permitted_network_id` | no | — |
| `state` | no | — |
| `type` | no | one of: `vpc` |
| `vpc_crn` | yes | crn format |
| `zone_id` | yes | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_permitted_network_invalid_instance_id` | `instance_id` |
| `ibm_dns_permitted_network_invalid_type` | `type` |
| `ibm_dns_permitted_network_invalid_vpc_crn` | `vpc_crn` |
| `ibm_dns_permitted_network_invalid_zone_id` | `zone_id` |

## Example

### Valid

```hcl
resource "ibm_dns_permitted_network" "test" {
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  vpc_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  zone_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
}
```
