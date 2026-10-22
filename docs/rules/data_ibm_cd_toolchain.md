# ibm_cd_toolchain_data

Reads ibm_cd_toolchain. Toolchain name.

This rule validates `ibm_cd_toolchain` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:patch:/toolchains/{toolchain_id}` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `description` | no | — |
| `href` | no | — |
| `location` | no | — |
| `name` | no | — |
| `resource_group_id` | no | — |
| `tags` | no | — |
| `toolchain_id` | yes | pattern `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$` |
| `ui_href` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_data_invalid_toolchain_id` | `toolchain_id` |

## Example

### Valid

```hcl
data "ibm_cd_toolchain" "test" {
  toolchain_id = "01234567-89ab-4cde-8123-456789abcdef"
}
```
