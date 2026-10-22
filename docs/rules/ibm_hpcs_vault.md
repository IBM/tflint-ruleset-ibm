# ibm_hpcs_vault

Manages ibm_hpcs_vault. A human-readable name to assign to your vault.

This rule validates `ibm_hpcs_vault` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v4/vaults/{id}:patch` (`uko`).

## Configuration

```hcl
rule "ibm_hpcs_vault" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `created_by` | no | — |
| `description` | no | API: Updated description of the vault |
| `etag` | no | — |
| `href` | no | — |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `name` | yes | pattern `^[A-Za-z0-9#@!$%'_-][A-Za-z0-9#@!$% '_-]*$` |
| `region` | yes | one of: `au-syd`, `in-che`, `jp-osa`, `jp-tok`, `kr-seo`, `eu-de`, `eu-gb`, `ca-tor`, `us-south`, `us-south-test`, `us-east`, `br-sao` |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `vault_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_hpcs_vault_invalid_description` | `description` |
| `ibm_hpcs_vault_invalid_instance_id` | `instance_id` |
| `ibm_hpcs_vault_invalid_name` | `name` |
| `ibm_hpcs_vault_invalid_region` | `region` |

## Example

### Valid

```hcl
resource "ibm_hpcs_vault" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  name = "example-name"
  region = "au-syd"
}
```
