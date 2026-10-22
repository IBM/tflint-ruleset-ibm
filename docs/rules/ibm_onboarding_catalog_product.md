# ibm_onboarding_catalog_product

Manages ibm_onboarding_catalog_product. The programmatic name of this product.

This rule validates `ibm_onboarding_catalog_product` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/products/{product_id}/catalog:patch` (`partner-center-sell`).

## Configuration

```hcl
rule "ibm_onboarding_catalog_product" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `active` | yes | — |
| `catalog_product_id` | no | — |
| `disabled` | yes | — |
| `env` | no | pattern `^[a-z_.-]+$` |
| `geo_tags` | no | — |
| `group` | no | — |
| `kind` | yes | one of: `composite`, `iaas`, `platform_service`, `service` |
| `name` | yes | API: The name of the product. |
| `object_id` | no | — |
| `pricing_tags` | no | — |
| `product_id` | yes | pattern `^[a-zA-Z0-9]{32}:o:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `tags` | yes | pattern `^[a-z0-9]+([ _:-]?[a-z0-9]+){0,10}$` |
| `url` | no | — |
| `images` (block) | no | nested block |
| `images.image` | no | — |
| `metadata` (block) | no | nested block |
| `metadata.rc_compatible` | no | — |
| `metadata.other` (block) | no | nested block |
| `metadata.other.product_code` | no | — |
| `metadata.other.product_code_type` | no | — |
| `metadata.other.composite` (block) | no | nested block |
| `metadata.other.composite.composite_kind` | no | — |
| `metadata.other.composite.composite_tag` | no | — |
| `metadata.other.composite.children` (block) | no | nested block |
| `metadata.other.composite.children.id` | no | — |
| `metadata.other.composite.children.kind` | no | — |
| `metadata.other.composite.children.name` | no | API: The name of the product. |
| `metadata.other.pc` (block) | no | nested block |
| `metadata.other.pc.support` (block) | no | nested block |
| `metadata.other.pc.support.languages` | no | — |
| `metadata.other.pc.support.locations` | no | one of: `af`, `ax`, `al`, `dz`, `as`, `ad`, `ao`, `ai`, `aq`, `ag`, `ar`, `am`, `aw`, `au`, `at`, `az`, `bs`, `bh`, `bd`, `bb`, `by`, `be`, `bz`, `bj`, `bm`, `bt`, `bo`, `bq`, `ba`, `bw`, `bv`, `br`, `io`, `bn`, `bg`, `bf`, `bi`, `kh`, `cm`, `ca`, `cv`, `ky`, `cf`, `td`, `cl`, `cn`, `cx`, `cc`, `co`, `km`, `cg`, `cd`, `ck`, `cr`, `ci`, `hr`, `cw`, `cy`, `cz`, `dk`, `dj`, `dm`, `do`, `ec`, `eg`, `sv`, `gq`, `er`, `ee`, `et`, `fk`, `fo`, `fj`, `fi`, `fr`, `gf`, `pf`, `tf`, `ga`, `gm`, `ge`, `de`, `gh`, `gi`, `gr`, `gl`, `gd`, `gp`, `gu`, `gt`, `gg`, `gn`, `gw`, `gy`, `ht`, `hm`, `va`, `hn`, `hk`, `hu`, `is`, `in`, `id`, `iq`, `ie`, `im`, `il`, `it`, `jm`, `jp`, `je`, `jo`, `kz`, `ke`, `ki`, `kr`, `kw`, `kg`, `la`, `lv`, `lb`, `ls`, `lr`, `ly`, `li`, `lt`, `lu`, `mo`, `mk`, `mg`, `mw`, `my`, `mv`, `ml`, `mt`, `mh`, `mq`, `mr`, `mu`, `yt`, `mx`, `fm`, `md`, `mc`, `mn`, `me`, `ms`, `ma`, `mz`, `mm`, `na`, `nr`, `np`, `nl`, `an`, `nc`, `nz`, `ni`, `ne`, `ng`, `nu`, `nf`, `mp`, `no`, `om`, `pk`, `pw`, `ps`, `pa`, `pg`, `py`, `pe`, `ph`, `pn`, `pl`, `pt`, `pr`, `qa`, `re`, `ro`, `ru`, `rw`, `bl`, `sh`, `kn`, `lc`, `mf`, `pm`, `vc`, `ws`, `sm`, `st`, `sa`, `sn`, `rs`, `sc`, `sl`, `sg`, `sx`, `sk`, `si`, `sb`, `so`, `za`, `gs`, `ss`, `es`, `lka`, `sd`, `sr`, `sj`, `sz`, `se`, `ch`, `tw`, `tj`, `tz`, `th`, `tl`, `tg`, `tk`, `to`, `tt`, `tn`, `tr`, `tm`, `tc`, `tv`, `ug`, `ua`, `ae`, `gb`, `us`, `um`, `uy`, `uz`, `vu`, `ve`, `vn`, `vg`, `vi`, `wf`, `eh`, `ye`, `zr`, `zm`, `zw` |
| `metadata.other.pc.support.process` | no | API: Provide details about how and when customers can get in touch with your support team. |
| `metadata.other.pc.support.process_i18n` | no | — |
| `metadata.other.pc.support.status_url` | no | — |
| `metadata.other.pc.support.support_type` | no | one of: `third_party`, `community` |
| `metadata.other.pc.support.url` | no | uri format |
| `metadata.other.pc.support.support_details` (block) | no | nested block |
| `metadata.other.pc.support.support_escalation` (block) | no | nested block |
| `metadata.service` (block) | no | nested block |
| `metadata.service.async_provisioning_supported` | no | — |
| `metadata.service.async_unprovisioning_supported` | no | — |
| `metadata.service.custom_create_page_hybrid_enabled` | no | — |
| `metadata.service.iam_compatible` | no | — |
| `metadata.service.rc_provisionable` | no | — |
| `metadata.service.service_key_supported` | no | — |
| `metadata.service.unique_api_key` | no | — |
| `metadata.service.parameters` (block) | no | nested block |
| `metadata.service.parameters.description` | no | API: The description of the product. |
| `metadata.service.parameters.displayname` | no | — |
| `metadata.service.parameters.hidden` | no | — |
| `metadata.service.parameters.invalidmessage` | no | — |
| `metadata.service.parameters.layout` | no | — |
| `metadata.service.parameters.name` | no | API: The name of the product. |
| `metadata.service.parameters.options_url` | no | — |
| `metadata.service.parameters.pattern` | no | — |
| `metadata.service.parameters.placeholder` | no | — |
| `metadata.service.parameters.readonly` | no | — |
| `metadata.service.parameters.required` | no | — |
| `metadata.service.parameters.type` | no | one of: `image`, `youtube`, `video_mp_4`, `video_webm` |
| `metadata.service.parameters.validation_url` | no | — |
| `metadata.service.parameters.value` | no | API: The amount per unit of time. |
| `metadata.service.parameters.associations` (block) | no | nested block |
| `metadata.service.parameters.associations.location` (block) | no | nested block |
| `metadata.service.parameters.associations.parameters` (block) | no | nested block |
| `metadata.service.parameters.associations.plan` (block) | no | nested block |
| `metadata.service.parameters.i18n` (block) | no | nested block |
| `metadata.service.parameters.i18n.de` (block) | no | nested block |
| `metadata.service.parameters.i18n.en` (block) | no | nested block |
| `metadata.service.parameters.i18n.es` (block) | no | nested block |
| `metadata.service.parameters.i18n.fr` (block) | no | nested block |
| `metadata.service.parameters.i18n.it` (block) | no | nested block |
| `metadata.service.parameters.i18n.ja` (block) | no | nested block |
| `metadata.service.parameters.i18n.ko` (block) | no | nested block |
| `metadata.service.parameters.i18n.pt_br` (block) | no | nested block |
| `metadata.service.parameters.i18n.zh_cn` (block) | no | nested block |
| `metadata.service.parameters.i18n.zh_tw` (block) | no | nested block |
| `metadata.service.parameters.options` (block) | no | nested block |
| `metadata.service.parameters.options.displayname` | no | — |
| `metadata.service.parameters.options.value` | no | API: The amount per unit of time. |
| `metadata.service.parameters.options.i18n` (block) | no | nested block |
| `metadata.ui` (block) | no | nested block |
| `metadata.ui.accessible_during_provision` | no | — |
| `metadata.ui.embeddable_dashboard` | no | — |
| `metadata.ui.hidden` | no | — |
| `metadata.ui.primary_offering_id` | no | — |
| `metadata.ui.side_by_side_index` | no | — |
| `metadata.ui.strings` (block) | no | nested block |
| `metadata.ui.strings.en` (block) | no | nested block |
| `metadata.ui.strings.en.bullets` (block) | no | nested block |
| `metadata.ui.strings.en.media` (block) | no | nested block |
| `metadata.ui.strings.en.navigation_items` (block) | no | nested block |
| `metadata.ui.urls` (block) | no | nested block |
| `metadata.ui.urls.apidocs_url` | no | — |
| `metadata.ui.urls.catalog_details_url` | no | — |
| `metadata.ui.urls.custom_create_page_url` | no | — |
| `metadata.ui.urls.dashboard` | no | — |
| `metadata.ui.urls.doc_url` | no | — |
| `metadata.ui.urls.instructions_url` | no | — |
| `metadata.ui.urls.terms_url` | no | — |
| `object_provider` (block) | yes | nested block |
| `object_provider.email` | no | pattern `^\S+@\S+$` |
| `object_provider.name` | no | API: The name of the product. |
| `overview_ui` (block) | no | nested block |
| `overview_ui.en` (block) | no | nested block |
| `overview_ui.en.description` | no | API: The description of the product. |
| `overview_ui.en.display_name` | no | — |
| `overview_ui.en.long_description` | no | API: The description about the details of the product. You can use markdown syntax to provide this description. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_onboarding_catalog_product_invalid_active` | `active` |
| `ibm_onboarding_catalog_product_invalid_disabled` | `disabled` |
| `ibm_onboarding_catalog_product_invalid_env` | `env` |
| `ibm_onboarding_catalog_product_invalid_kind` | `kind` |
| `ibm_onboarding_catalog_product_invalid_name` | `name` |
| `ibm_onboarding_catalog_product_invalid_object_id` | `object_id` |
| `ibm_onboarding_catalog_product_invalid_product_id` | `product_id` |
| `ibm_onboarding_catalog_product_invalid_tags` | `tags` |
| `ibm_onboarding_catalog_product_invalid_images__image` | `images.image` |
| `ibm_onboarding_catalog_product_invalid_metadata__rc_compatible` | `metadata.rc_compatible` |
| `ibm_onboarding_catalog_product_invalid_metadata__other__product_code` | `metadata.other.product_code` |
| `ibm_onboarding_catalog_product_invalid_metadata__other__product_code_type` | `metadata.other.product_code_type` |
| `ibm_onboarding_catalog_product_invalid_metadata__other__composite__composite_kind` | `metadata.other.composite.composite_kind` |
| `ibm_onboarding_catalog_product_invalid_metadata__other__composite__composite_tag` | `metadata.other.composite.composite_tag` |
| `ibm_onboarding_catalog_product_invalid_metadata__other__composite__children__id` | `metadata.other.composite.children.id` |
| `ibm_onboarding_catalog_product_invalid_metadata__other__composite__children__kind` | `metadata.other.composite.children.kind` |
| `ibm_onboarding_catalog_product_invalid_metadata__other__composite__children__name` | `metadata.other.composite.children.name` |
| `ibm_onboarding_catalog_product_invalid_metadata__other__pc__support__languages` | `metadata.other.pc.support.languages` |
| `ibm_onboarding_catalog_product_invalid_metadata__other__pc__support__locations` | `metadata.other.pc.support.locations` |
| `ibm_onboarding_catalog_product_invalid_metadata__other__pc__support__process` | `metadata.other.pc.support.process` |
| `ibm_onboarding_catalog_product_invalid_metadata__other__pc__support__process_i18n` | `metadata.other.pc.support.process_i18n` |
| `ibm_onboarding_catalog_product_invalid_metadata__other__pc__support__status_url` | `metadata.other.pc.support.status_url` |
| `ibm_onboarding_catalog_product_invalid_metadata__other__pc__support__support_type` | `metadata.other.pc.support.support_type` |
| `ibm_onboarding_catalog_product_invalid_metadata__other__pc__support__url` | `metadata.other.pc.support.url` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__async_provisioning_supported` | `metadata.service.async_provisioning_supported` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__async_unprovisioning_supported` | `metadata.service.async_unprovisioning_supported` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__custom_create_page_hybrid_enabled` | `metadata.service.custom_create_page_hybrid_enabled` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__iam_compatible` | `metadata.service.iam_compatible` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__rc_provisionable` | `metadata.service.rc_provisionable` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__service_key_supported` | `metadata.service.service_key_supported` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__unique_api_key` | `metadata.service.unique_api_key` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__description` | `metadata.service.parameters.description` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__displayname` | `metadata.service.parameters.displayname` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__hidden` | `metadata.service.parameters.hidden` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__invalidmessage` | `metadata.service.parameters.invalidmessage` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__layout` | `metadata.service.parameters.layout` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__name` | `metadata.service.parameters.name` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__options_url` | `metadata.service.parameters.options_url` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__pattern` | `metadata.service.parameters.pattern` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__placeholder` | `metadata.service.parameters.placeholder` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__readonly` | `metadata.service.parameters.readonly` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__required` | `metadata.service.parameters.required` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__type` | `metadata.service.parameters.type` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__validation_url` | `metadata.service.parameters.validation_url` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__value` | `metadata.service.parameters.value` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__options__displayname` | `metadata.service.parameters.options.displayname` |
| `ibm_onboarding_catalog_product_invalid_metadata__service__parameters__options__value` | `metadata.service.parameters.options.value` |
| `ibm_onboarding_catalog_product_invalid_metadata__ui__accessible_during_provision` | `metadata.ui.accessible_during_provision` |
| `ibm_onboarding_catalog_product_invalid_metadata__ui__embeddable_dashboard` | `metadata.ui.embeddable_dashboard` |
| `ibm_onboarding_catalog_product_invalid_metadata__ui__hidden` | `metadata.ui.hidden` |
| `ibm_onboarding_catalog_product_invalid_metadata__ui__primary_offering_id` | `metadata.ui.primary_offering_id` |
| `ibm_onboarding_catalog_product_invalid_metadata__ui__side_by_side_index` | `metadata.ui.side_by_side_index` |
| `ibm_onboarding_catalog_product_invalid_metadata__ui__urls__apidocs_url` | `metadata.ui.urls.apidocs_url` |
| `ibm_onboarding_catalog_product_invalid_metadata__ui__urls__catalog_details_url` | `metadata.ui.urls.catalog_details_url` |
| `ibm_onboarding_catalog_product_invalid_metadata__ui__urls__custom_create_page_url` | `metadata.ui.urls.custom_create_page_url` |
| `ibm_onboarding_catalog_product_invalid_metadata__ui__urls__dashboard` | `metadata.ui.urls.dashboard` |
| `ibm_onboarding_catalog_product_invalid_metadata__ui__urls__doc_url` | `metadata.ui.urls.doc_url` |
| `ibm_onboarding_catalog_product_invalid_metadata__ui__urls__instructions_url` | `metadata.ui.urls.instructions_url` |
| `ibm_onboarding_catalog_product_invalid_metadata__ui__urls__terms_url` | `metadata.ui.urls.terms_url` |
| `ibm_onboarding_catalog_product_invalid_object_provider__email` | `object_provider.email` |
| `ibm_onboarding_catalog_product_invalid_object_provider__name` | `object_provider.name` |
| `ibm_onboarding_catalog_product_invalid_overview_ui__en__description` | `overview_ui.en.description` |
| `ibm_onboarding_catalog_product_invalid_overview_ui__en__display_name` | `overview_ui.en.display_name` |
| `ibm_onboarding_catalog_product_invalid_overview_ui__en__long_description` | `overview_ui.en.long_description` |

## Example

### Valid

```hcl
resource "ibm_onboarding_catalog_product" "test" {
  active = true
  disabled = true
  kind = "composite"
  name = "aaaaaaa"
  product_id = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa:o:01234567-89ab-cdef-0123-456789abcdef"
  tags = "example-name"
  object_provider {
  }
}
```
