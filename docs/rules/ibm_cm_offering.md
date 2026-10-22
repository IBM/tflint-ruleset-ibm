# ibm_cm_offering

Manages ibm_cm_offering. The programmatic name of this offering.

This rule validates `ibm_cm_offering` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Offering` (`resource-catalog-private-catalog`).

## Configuration

```hcl
rule "ibm_cm_offering" {
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
| `deprecate` | no | — |
| `deprecate_pending` | no | — |
| `disclaimer` | no | API: A disclaimer for this offering |
| `hidden` | no | API: Determine if this offering should be displayed in the Consumption UI |
| `keywords` | no | API: List of keywords associated with offering, typically used to search for it |
| `kinds` | no | — |
| `label` | no | API: Display Name in the requested language |
| `label_i18n` | no | API: A map of translated strings, by language code |
| `long_description` | no | API: Long description in the requested language |
| `long_description_i18n` | no | API: A map of translated strings, by language code |
| `metadata` | no | API: Map of metadata values for this offering |
| `name` | no | API: The programmatic name of this offering |
| `offering_docs_url` | no | API: URL for an additional docs with this offering |
| `offering_icon_url` | no | API: URL for an icon associated with this offering |
| `offering_id` | no | API: Offering ID |
| `offering_identifier` | no | — |
| `offering_support_url` | no | API: [deprecated] - Use offering.support instead. URL to be displayed in the Consumption UI for getting support on this offering |
| `pc_managed` | no | — |
| `portal_approval_record` | no | API: The portal's approval record ID |
| `portal_ui_url` | no | API: The portal UI URL |
| `product_kind` | no | API: The product kind. Valid values are module, solution, or empty string |
| `public_original_crn` | no | crn format |
| `publish_approved` | no | — |
| `publish_public_crn` | no | crn format |
| `rating` | no | — |
| `repo_info` | no | — |
| `rev` | no | — |
| `share_enabled` | no | API: Denotes sharing including access list availability of an Offering is enabled |
| `share_with_access_list` | no | — |
| `share_with_all` | no | API: Denotes public availability of an Offering |
| `share_with_ibm` | no | API: Denotes IBM employee availability of an Offering - if share_enabled is true |
| `short_description` | no | API: Short description in the requested language |
| `short_description_i18n` | no | API: A map of translated strings, by language code |
| `support` | no | — |
| `tags` | no | API: List of tags associated with this catalog |
| `updated` | no | — |
| `url` | no | — |
| `features` (block) | no | nested block |
| `features.description` | no | API: Feature description |
| `features.description_i18n` | no | API: A map of translated strings, by language code |
| `features.title` | no | API: Heading |
| `features.title_i18n` | no | API: A map of translated strings, by language code |
| `image_pull_keys` (block) | no | nested block |
| `image_pull_keys.description` | no | API: Key description |
| `image_pull_keys.name` | no | API: Key name |
| `image_pull_keys.value` | no | API: Key value |
| `media` (block) | no | nested block |
| `media.api_url` | no | API: CM API specific URL of the specified media item |
| `media.caption` | no | API: Caption for this media item |
| `media.caption_i18n` | no | API: A map of translated strings, by language code |
| `media.thumbnail_url` | no | API: Thumbnail URL for this media item |
| `media.type` | no | API: Type of this media item |
| `media.url` | no | API: URL of the specified media item |
| `media.url_proxy` (block) | no | nested block |
| `media.url_proxy.sha` | no | API: SHA256 fingerprint of image |
| `media.url_proxy.url` | no | API: URL of the specified media item being proxied |
| `provider_info` (block) | no | nested block |
| `provider_info.id` | no | API: The id of this provider |
| `provider_info.name` | no | API: The name of this provider |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cm_offering_invalid_catalog_id` | `catalog_id` |
| `ibm_cm_offering_invalid_deprecate` | `deprecate` |
| `ibm_cm_offering_invalid_disclaimer` | `disclaimer` |
| `ibm_cm_offering_invalid_hidden` | `hidden` |
| `ibm_cm_offering_invalid_keywords` | `keywords` |
| `ibm_cm_offering_invalid_label` | `label` |
| `ibm_cm_offering_invalid_label_i18n` | `label_i18n` |
| `ibm_cm_offering_invalid_long_description` | `long_description` |
| `ibm_cm_offering_invalid_long_description_i18n` | `long_description_i18n` |
| `ibm_cm_offering_invalid_metadata` | `metadata` |
| `ibm_cm_offering_invalid_name` | `name` |
| `ibm_cm_offering_invalid_offering_docs_url` | `offering_docs_url` |
| `ibm_cm_offering_invalid_offering_icon_url` | `offering_icon_url` |
| `ibm_cm_offering_invalid_offering_id` | `offering_id` |
| `ibm_cm_offering_invalid_offering_support_url` | `offering_support_url` |
| `ibm_cm_offering_invalid_portal_approval_record` | `portal_approval_record` |
| `ibm_cm_offering_invalid_portal_ui_url` | `portal_ui_url` |
| `ibm_cm_offering_invalid_product_kind` | `product_kind` |
| `ibm_cm_offering_invalid_public_original_crn` | `public_original_crn` |
| `ibm_cm_offering_invalid_publish_public_crn` | `publish_public_crn` |
| `ibm_cm_offering_invalid_share_enabled` | `share_enabled` |
| `ibm_cm_offering_invalid_share_with_access_list` | `share_with_access_list` |
| `ibm_cm_offering_invalid_share_with_all` | `share_with_all` |
| `ibm_cm_offering_invalid_share_with_ibm` | `share_with_ibm` |
| `ibm_cm_offering_invalid_short_description` | `short_description` |
| `ibm_cm_offering_invalid_short_description_i18n` | `short_description_i18n` |
| `ibm_cm_offering_invalid_tags` | `tags` |
| `ibm_cm_offering_invalid_features__description` | `features.description` |
| `ibm_cm_offering_invalid_features__description_i18n` | `features.description_i18n` |
| `ibm_cm_offering_invalid_features__title` | `features.title` |
| `ibm_cm_offering_invalid_features__title_i18n` | `features.title_i18n` |
| `ibm_cm_offering_invalid_image_pull_keys__description` | `image_pull_keys.description` |
| `ibm_cm_offering_invalid_image_pull_keys__name` | `image_pull_keys.name` |
| `ibm_cm_offering_invalid_image_pull_keys__value` | `image_pull_keys.value` |
| `ibm_cm_offering_invalid_media__api_url` | `media.api_url` |
| `ibm_cm_offering_invalid_media__caption` | `media.caption` |
| `ibm_cm_offering_invalid_media__caption_i18n` | `media.caption_i18n` |
| `ibm_cm_offering_invalid_media__thumbnail_url` | `media.thumbnail_url` |
| `ibm_cm_offering_invalid_media__type` | `media.type` |
| `ibm_cm_offering_invalid_media__url` | `media.url` |
| `ibm_cm_offering_invalid_media__url_proxy__sha` | `media.url_proxy.sha` |
| `ibm_cm_offering_invalid_media__url_proxy__url` | `media.url_proxy.url` |
| `ibm_cm_offering_invalid_provider_info__id` | `provider_info.id` |
| `ibm_cm_offering_invalid_provider_info__name` | `provider_info.name` |

## Example

### Valid

```hcl
resource "ibm_cm_offering" "test" {
  catalog_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
