# ibm_hpcs_key_template_data

Reads ibm_hpcs_key_template. Name of the key template.

This rule validates `ibm_hpcs_key_template` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Key` (`hs-crypto`).

## Configuration

```hcl
rule "ibm_hpcs_key_template_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `created_by` | no | — |
| `description` | no | — |
| `href` | no | — |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `key` | no | — |
| `keystores` | no | — |
| `name` | no | — |
| `region` | yes | one of: `au-syd`, `in-che`, `jp-osa`, `jp-tok`, `kr-seo`, `eu-de`, `eu-gb`, `ca-tor`, `us-south`, `us-south-test`, `us-east`, `br-sao` |
| `template_id` | yes | — |
| `uko_vault` | yes | — |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `vault` | no | — |
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_hpcs_key_template_data_invalid_instance_id` | `instance_id` |
| `ibm_hpcs_key_template_data_invalid_region` | `region` |
| `ibm_hpcs_key_template_data_invalid_template_id` | `template_id` |
| `ibm_hpcs_key_template_data_invalid_uko_vault` | `uko_vault` |

## Example

### Valid

```hcl
data "ibm_hpcs_key_template" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  region = "au-syd"
  template_id = "01234567-89ab-cdef-0123-456789abcdef"
  uko_vault = "example"
}
```
