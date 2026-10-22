# ibm_scc_provider_type_instance

Manages ibm_scc_provider_type_instance. The name of the provider type instance.

This rule validates `ibm_scc_provider_type_instance` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/instances/{instance_id}/v3/provider_types/{provider_type_id}/provider_type_instances/{provider_type_instance_id}:patch` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_provider_type_instance" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `attributes` | no | API: The attributes for connecting to the provider type instance. |
| `created_at` | no | — |
| `instance_id` | yes | — |
| `name` | yes | pattern `[A-Za-z0-9]+` |
| `provider_type_id` | yes | pattern `^[a-zA-Z0-9 ,\-_]+$` |
| `provider_type_instance_id` | no | — |
| `type` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_provider_type_instance_invalid_attributes` | `attributes` |
| `ibm_scc_provider_type_instance_invalid_instance_id` | `instance_id` |
| `ibm_scc_provider_type_instance_invalid_name` | `name` |
| `ibm_scc_provider_type_instance_invalid_provider_type_id` | `provider_type_id` |

## Example

### Valid

```hcl
resource "ibm_scc_provider_type_instance" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  name = "example-name"
  provider_type_id = "example-namexxxxxxxxxxxxxxxxxxxx"
}
```
