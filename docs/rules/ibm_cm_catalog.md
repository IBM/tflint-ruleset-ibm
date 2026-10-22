# ibm_cm_catalog

Schema validation for the ibm_cm_catalog resource.

This rule validates `ibm_cm_catalog` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Catalog` (`resource-catalog-private-catalog`).

## Configuration

```hcl
rule "ibm_cm_catalog" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `catalog_banner_url` | no | API: URL for a banner image for this catalog |
| `catalog_icon_url` | no | API: URL for an icon associated with this catalog |
| `created` | no | — |
| `crn` | no | — |
| `disabled` | no | API: Denotes whether a catalog is disabled |
| `kind` | no | API: Kind of catalog. Supported kinds are offering and vpe |
| `label` | no | API: Display Name in the requested language |
| `label_i18n` | no | API: A map of translated strings, by language code |
| `metadata` | no | API: Catalog specific metadata |
| `offerings_url` | no | — |
| `owning_account` | no | — |
| `resource_group_id` | no | API: Resource group id the catalog is owned by |
| `rev` | no | — |
| `short_description` | no | API: Description in the requested language |
| `short_description_i18n` | no | API: A map of translated strings, by language code |
| `tags` | no | API: List of tags associated with this catalog |
| `updated` | no | — |
| `url` | no | — |
| `catalog_filters` (block) | no | nested block |
| `catalog_filters.include_all` | no | API: -&gt; true - Include all of the public catalog when filtering. Further settings will specifically exclude some offerings. false - Exclude all of the public catalog when filtering. Further settings will specifically include some offerings. |
| `catalog_filters.category_filters` (block) | no | nested block |
| `catalog_filters.category_filters.category_name` | yes | — |
| `catalog_filters.category_filters.include` | no | API: Offering filter terms |
| `catalog_filters.category_filters.filter` (block) | no | nested block |
| `catalog_filters.category_filters.filter.filter_terms` | no | API: List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded. |
| `catalog_filters.id_filters` (block) | no | nested block |
| `catalog_filters.id_filters.exclude` (block) | no | nested block |
| `catalog_filters.id_filters.exclude.filter_terms` | no | API: List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded. |
| `catalog_filters.id_filters.include` (block) | no | nested block |
| `catalog_filters.id_filters.include.filter_terms` | no | API: List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded. |
| `features` (block) | no | nested block |
| `features.description` | no | API: Feature description |
| `features.description_i18n` | no | API: A map of translated strings, by language code |
| `features.title` | no | API: Heading |
| `features.title_i18n` | no | API: A map of translated strings, by language code |
| `target_account_contexts` (block) | no | nested block |
| `target_account_contexts.api_key` | no | API: API Key |
| `target_account_contexts.label` | no | API: Display name |
| `target_account_contexts.name` | no | API: Unique identifier/name |
| `target_account_contexts.project_id` | no | API: Project ID |
| `target_account_contexts.trusted_profile` (block) | no | nested block |
| `target_account_contexts.trusted_profile.target_service_id` | no | API: Target service ID |
| `target_account_contexts.trusted_profile.trusted_profile_id` | no | API: Trusted profile ID |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cm_catalog_invalid_catalog_banner_url` | `catalog_banner_url` |
| `ibm_cm_catalog_invalid_catalog_icon_url` | `catalog_icon_url` |
| `ibm_cm_catalog_invalid_disabled` | `disabled` |
| `ibm_cm_catalog_invalid_kind` | `kind` |
| `ibm_cm_catalog_invalid_label` | `label` |
| `ibm_cm_catalog_invalid_label_i18n` | `label_i18n` |
| `ibm_cm_catalog_invalid_metadata` | `metadata` |
| `ibm_cm_catalog_invalid_resource_group_id` | `resource_group_id` |
| `ibm_cm_catalog_invalid_short_description` | `short_description` |
| `ibm_cm_catalog_invalid_short_description_i18n` | `short_description_i18n` |
| `ibm_cm_catalog_invalid_tags` | `tags` |
| `ibm_cm_catalog_invalid_catalog_filters__include_all` | `catalog_filters.include_all` |
| `ibm_cm_catalog_invalid_catalog_filters__category_filters__category_name` | `catalog_filters.category_filters.category_name` |
| `ibm_cm_catalog_invalid_catalog_filters__category_filters__include` | `catalog_filters.category_filters.include` |
| `ibm_cm_catalog_invalid_catalog_filters__category_filters__filter__filter_terms` | `catalog_filters.category_filters.filter.filter_terms` |
| `ibm_cm_catalog_invalid_catalog_filters__id_filters__exclude__filter_terms` | `catalog_filters.id_filters.exclude.filter_terms` |
| `ibm_cm_catalog_invalid_catalog_filters__id_filters__include__filter_terms` | `catalog_filters.id_filters.include.filter_terms` |
| `ibm_cm_catalog_invalid_features__description` | `features.description` |
| `ibm_cm_catalog_invalid_features__description_i18n` | `features.description_i18n` |
| `ibm_cm_catalog_invalid_features__title` | `features.title` |
| `ibm_cm_catalog_invalid_features__title_i18n` | `features.title_i18n` |
| `ibm_cm_catalog_invalid_target_account_contexts__api_key` | `target_account_contexts.api_key` |
| `ibm_cm_catalog_invalid_target_account_contexts__label` | `target_account_contexts.label` |
| `ibm_cm_catalog_invalid_target_account_contexts__name` | `target_account_contexts.name` |
| `ibm_cm_catalog_invalid_target_account_contexts__project_id` | `target_account_contexts.project_id` |
| `ibm_cm_catalog_invalid_target_account_contexts__trusted_profile__target_service_id` | `target_account_contexts.trusted_profile.target_service_id` |
| `ibm_cm_catalog_invalid_target_account_contexts__trusted_profile__trusted_profile_id` | `target_account_contexts.trusted_profile.trusted_profile_id` |

## Example

### Valid

```hcl
resource "ibm_cm_catalog" "test" {
}
```
