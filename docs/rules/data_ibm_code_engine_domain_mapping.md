# ibm_code_engine_domain_mapping_data

Reads ibm_code_engine_domain_mapping. The name of your domain mapping.

This rule validates `ibm_code_engine_domain_mapping` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:patch:/projects/{project_id}/domain_mappings/{name}` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_domain_mapping_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cname_target` | no | — |
| `component` | no | — |
| `created_at` | no | — |
| `domain_mapping_id` | no | — |
| `entity_tag` | no | — |
| `href` | no | — |
| `name` | yes | pattern `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)+$` |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `region` | no | — |
| `resource_type` | no | — |
| `status` | no | — |
| `status_details` | no | — |
| `tls_secret` | no | — |
| `user_managed` | no | — |
| `visibility` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_domain_mapping_data_invalid_name` | `name` |
| `ibm_code_engine_domain_mapping_data_invalid_project_id` | `project_id` |

## Example

### Valid

```hcl
data "ibm_code_engine_domain_mapping" "test" {
  name = "10.0.0.1"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
