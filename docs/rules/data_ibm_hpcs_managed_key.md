# ibm_hpcs_managed_key_data

Schema validation for the ibm_hpcs_managed_key data source.

This rule validates `ibm_hpcs_managed_key` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Key` (`hs-crypto`).

## Configuration

```hcl
rule "ibm_hpcs_managed_key_data" {
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
| `description` | no | — |
| `expiration_date` | no | — |
| `href` | no | — |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `instances` | no | — |
| `key_id` | yes | — |
| `label` | no | — |
| `referenced_keystores` | no | — |
| `region` | yes | one of: `au-syd`, `in-che`, `jp-osa`, `jp-tok`, `kr-seo`, `eu-de`, `eu-gb`, `ca-tor`, `us-south`, `us-south-test`, `us-east`, `br-sao` |
| `size` | no | — |
| `state` | no | — |
| `tags` | no | — |
| `template` | no | — |
| `uko_vault` | yes | — |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `vault` | no | — |
| `verification_patterns` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_hpcs_managed_key_data_invalid_instance_id` | `instance_id` |
| `ibm_hpcs_managed_key_data_invalid_key_id` | `key_id` |
| `ibm_hpcs_managed_key_data_invalid_region` | `region` |
| `ibm_hpcs_managed_key_data_invalid_uko_vault` | `uko_vault` |

## Example

### Valid

```hcl
data "ibm_hpcs_managed_key" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  key_id = "01234567-89ab-cdef-0123-456789abcdef"
  region = "au-syd"
  uko_vault = "example"
}
```
