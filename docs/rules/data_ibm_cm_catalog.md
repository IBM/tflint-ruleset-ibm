# ibm_cm_catalog_data

Schema validation for the ibm_cm_catalog data source.

This rule validates `ibm_cm_catalog` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Catalog` (`resource-catalog-private-catalog`).

## Configuration

```hcl
rule "ibm_cm_catalog_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `catalog_banner_url` | no | — |
| `catalog_filters` | no | — |
| `catalog_icon_url` | no | — |
| `catalog_identifier` | no | — |
| `created` | no | — |
| `crn` | no | — |
| `disabled` | no | — |
| `features` | no | — |
| `id` | no | — |
| `kind` | no | — |
| `label` | no | API: Display Name in the requested language |
| `label_i18n` | no | — |
| `metadata` | no | — |
| `offerings_url` | no | — |
| `owning_account` | no | — |
| `resource_group_id` | no | — |
| `rev` | no | — |
| `short_description` | no | — |
| `short_description_i18n` | no | — |
| `tags` | no | — |
| `target_account_contexts` | no | — |
| `updated` | no | — |
| `url` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cm_catalog_data_invalid_catalog_identifier` | `catalog_identifier` |
| `ibm_cm_catalog_data_invalid_label` | `label` |

## Example

### Valid

```hcl
data "ibm_cm_catalog" "test" {
}
```
