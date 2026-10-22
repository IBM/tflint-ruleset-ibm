# ibm_hpcs_managed_key

Schema validation for the ibm_hpcs_managed_key resource.

This rule validates `ibm_hpcs_managed_key` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `template` (`uko`).

## Configuration

```hcl
rule "ibm_hpcs_managed_key" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `activation_date` | no | — |
| `algorithm` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `description` | no | API: Description of the key template |
| `etag` | no | — |
| `expiration_date` | no | — |
| `href` | no | — |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `instances` | no | — |
| `key_id` | no | — |
| `label` | yes | pattern `^[A-Za-z0-9._ \/-]+$` |
| `referenced_keystores` | no | — |
| `region` | yes | one of: `au-syd`, `in-che`, `jp-osa`, `jp-tok`, `kr-seo`, `eu-de`, `eu-gb`, `ca-tor`, `us-south`, `us-south-test`, `us-east`, `br-sao` |
| `size` | no | — |
| `state` | no | one of: `archived`, `unarchived` |
| `template` | no | — |
| `template_name` | yes | pattern `^[A-Za-z][A-Za-z0-9-]+$` |
| `uko_vault` | yes | — |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `verification_patterns` | no | — |
| `tags` (block) | no | nested block |
| `tags.name` | yes | pattern `^[A-Za-z0-9#@!$%'_-][A-Za-z0-9#@!$% '_-]*$` |
| `tags.value` | yes | — |
| `vault` (block) | yes | nested block |
| `vault.id` | yes | pattern `^[-0-9a-z]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_hpcs_managed_key_invalid_description` | `description` |
| `ibm_hpcs_managed_key_invalid_instance_id` | `instance_id` |
| `ibm_hpcs_managed_key_invalid_label` | `label` |
| `ibm_hpcs_managed_key_invalid_region` | `region` |
| `ibm_hpcs_managed_key_invalid_state` | `state` |
| `ibm_hpcs_managed_key_invalid_template_name` | `template_name` |
| `ibm_hpcs_managed_key_invalid_uko_vault` | `uko_vault` |
| `ibm_hpcs_managed_key_invalid_tags__name` | `tags.name` |
| `ibm_hpcs_managed_key_invalid_tags__value` | `tags.value` |
| `ibm_hpcs_managed_key_invalid_vault__id` | `vault.id` |

## Example

### Valid

```hcl
resource "ibm_hpcs_managed_key" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  label = "example-name"
  region = "au-syd"
  template_name = "example-name"
  uko_vault = "example"
  vault {
    id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  }
}
```
