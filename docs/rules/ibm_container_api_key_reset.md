# ibm_container_api_key_reset

Schema validation for the ibm_container_api_key_reset resource.

This rule validates `ibm_container_api_key_reset` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `IAAS` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_api_key_reset" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `region` | yes | API: The name of region in the infrastructure provider that is used for the Satellite location, such as the AWS 'us-east-2' region. |
| `reset_api_key` | no | — |
| `resource_group_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_api_key_reset_invalid_region` | `region` |
| `ibm_container_api_key_reset_invalid_reset_api_key` | `reset_api_key` |
| `ibm_container_api_key_reset_invalid_resource_group_id` | `resource_group_id` |

## Example

### Valid

```hcl
resource "ibm_container_api_key_reset" "test" {
  region = "aaaaaaa"
}
```
