# ibm_code_engine_domain_mapping

Manages ibm_code_engine_domain_mapping. The name of the domain mapping.

This rule validates `ibm_code_engine_domain_mapping` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DomainMappingPrototype` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_domain_mapping" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cname_target` | no | — |
| `created_at` | no | — |
| `domain_mapping_id` | no | — |
| `entity_tag` | no | — |
| `etag` | no | — |
| `href` | no | — |
| `name` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)+$` |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `region` | no | — |
| `resource_type` | no | — |
| `status` | no | — |
| `status_details` | no | — |
| `tls_secret` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$` |
| `user_managed` | no | — |
| `visibility` | no | — |
| `component` (block) | yes | nested block |
| `component.name` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?$` |
| `component.resource_type` | yes | API: The type of the referenced resource. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_domain_mapping_invalid_name` | `name` |
| `ibm_code_engine_domain_mapping_invalid_project_id` | `project_id` |
| `ibm_code_engine_domain_mapping_invalid_tls_secret` | `tls_secret` |
| `ibm_code_engine_domain_mapping_invalid_component__name` | `component.name` |
| `ibm_code_engine_domain_mapping_invalid_component__resource_type` | `component.resource_type` |

## Example

### Valid

```hcl
variable "tls_secret" {
  sensitive = true
}

resource "ibm_code_engine_domain_mapping" "test" {
  name = "10.0.0.1"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
  tls_secret = var.tls_secret
  component {
    name = "example-name"
    resource_type = "aaaaaaa"
  }
}
```
