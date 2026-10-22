# ibm_hpcs_vault_data

Reads ibm_hpcs_vault. Name of the vault.

This rule validates `ibm_hpcs_vault` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `vault` (`uko`).

## Configuration

```hcl
rule "ibm_hpcs_vault_data" {
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
| `name` | no | — |
| `region` | yes | one of: `au-syd`, `in-che`, `jp-osa`, `jp-tok`, `kr-seo`, `eu-de`, `eu-gb`, `ca-tor`, `us-south`, `us-south-test`, `us-east`, `br-sao` |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `vault_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_hpcs_vault_data_invalid_instance_id` | `instance_id` |
| `ibm_hpcs_vault_data_invalid_region` | `region` |
| `ibm_hpcs_vault_data_invalid_vault_id` | `vault_id` |

## Example

### Valid

```hcl
data "ibm_hpcs_vault" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  region = "au-syd"
  vault_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
