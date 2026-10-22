# ibm_cm_offering_data

Reads ibm_cm_offering. The programmatic name of this offering.

This rule validates `ibm_cm_offering` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Offering` (`resource-catalog-private-catalog`).

## Configuration

```hcl
rule "ibm_cm_offering_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `badges` | no | — |
| `catalog_id` | yes | API: The id of the catalog containing this offering |
| `catalog_name` | no | — |
| `created` | no | — |
| `crn` | no | — |
| `deprecate_pending` | no | — |
| `disclaimer` | no | — |
| `features` | no | — |
| `hidden` | no | — |
| `image_pull_keys` | no | — |
| `keywords` | no | — |
| `kinds` | no | — |
| `label` | no | — |
| `label_i18n` | no | — |
| `long_description` | no | — |
| `long_description_i18n` | no | — |
| `media` | no | — |
| `metadata` | no | — |
| `name` | no | — |
| `offering_docs_url` | no | — |
| `offering_icon_url` | no | — |
| `offering_id` | yes | API: Offering ID |
| `offering_identifier` | no | — |
| `offering_support_url` | no | — |
| `pc_managed` | no | — |
| `portal_approval_record` | no | — |
| `portal_ui_url` | no | — |
| `product_kind` | no | — |
| `provider_info` | no | — |
| `public_original_crn` | no | — |
| `publish_approved` | no | — |
| `publish_public_crn` | no | — |
| `rating` | no | — |
| `repo_info` | no | — |
| `rev` | no | — |
| `share_enabled` | no | — |
| `share_with_access_list` | no | — |
| `share_with_all` | no | — |
| `share_with_ibm` | no | — |
| `short_description` | no | — |
| `short_description_i18n` | no | — |
| `support` | no | — |
| `tags` | no | — |
| `updated` | no | — |
| `url` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cm_offering_data_invalid_catalog_id` | `catalog_id` |
| `ibm_cm_offering_data_invalid_offering_id` | `offering_id` |

## Example

### Valid

```hcl
data "ibm_cm_offering" "test" {
  catalog_id = "01234567-89ab-cdef-0123-456789abcdef"
  offering_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
