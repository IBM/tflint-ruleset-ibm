# ibm_onboarding_catalog_plan

Manages ibm_onboarding_catalog_plan. The programmatic name of this plan.

This rule validates `ibm_onboarding_catalog_plan` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/products/{product_id}/catalog:patch` (`partner-center-sell`).

## Configuration

```hcl
rule "ibm_onboarding_catalog_plan" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `active` | yes | — |
| `catalog_plan_id` | no | — |
| `catalog_product_id` | yes | — |
| `disabled` | yes | — |
| `env` | no | pattern `^[a-z_.-]+$` |
| `geo_tags` | no | — |
| `kind` | yes | one of: `plan` |
| `name` | yes | API: The name of the product. |
| `object_id` | no | — |
| `pricing_tags` | no | — |
| `product_id` | yes | pattern `^[a-zA-Z0-9]{32}:o:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `tags` | no | pattern `^[a-z0-9]+([ _:-]?[a-z0-9]+){0,10}$` |
| `url` | no | — |
| `metadata` (block) | no | nested block |
| `metadata.rc_compatible` | no | — |
| `metadata.other` (block) | no | nested block |
| `metadata.other.resource_controller` (block) | no | nested block |
| `metadata.other.resource_controller.subscription_provider_id` | no | — |
| `metadata.plan` (block) | no | nested block |
| `metadata.plan.allow_internal_users` | no | — |
| `metadata.plan.provision_type` | no | — |
| `metadata.plan.reservable` | no | — |
| `metadata.pricing` (block) | no | nested block |
| `metadata.pricing.origin` | no | — |
| `metadata.pricing.sales_avenue` | no | — |
| `metadata.pricing.type` | no | one of: `image`, `youtube`, `video_mp_4`, `video_webm` |
| `metadata.service` (block) | no | nested block |
| `metadata.service.bindable` | no | — |
| `metadata.service.iam_compatible` | no | — |
| `metadata.service.plan_updateable` | no | — |
| `metadata.service.rc_provisionable` | no | — |
| `metadata.service.service_key_supported` | no | — |
| `metadata.ui` (block) | no | nested block |
| `metadata.ui.hidden` | no | — |
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
| `ibm_onboarding_catalog_plan_invalid_active` | `active` |
| `ibm_onboarding_catalog_plan_invalid_catalog_product_id` | `catalog_product_id` |
| `ibm_onboarding_catalog_plan_invalid_disabled` | `disabled` |
| `ibm_onboarding_catalog_plan_invalid_env` | `env` |
| `ibm_onboarding_catalog_plan_invalid_kind` | `kind` |
| `ibm_onboarding_catalog_plan_invalid_name` | `name` |
| `ibm_onboarding_catalog_plan_invalid_object_id` | `object_id` |
| `ibm_onboarding_catalog_plan_invalid_pricing_tags` | `pricing_tags` |
| `ibm_onboarding_catalog_plan_invalid_product_id` | `product_id` |
| `ibm_onboarding_catalog_plan_invalid_tags` | `tags` |
| `ibm_onboarding_catalog_plan_invalid_metadata__rc_compatible` | `metadata.rc_compatible` |
| `ibm_onboarding_catalog_plan_invalid_metadata__other__resource_controller__subscription_provider_id` | `metadata.other.resource_controller.subscription_provider_id` |
| `ibm_onboarding_catalog_plan_invalid_metadata__plan__allow_internal_users` | `metadata.plan.allow_internal_users` |
| `ibm_onboarding_catalog_plan_invalid_metadata__plan__provision_type` | `metadata.plan.provision_type` |
| `ibm_onboarding_catalog_plan_invalid_metadata__plan__reservable` | `metadata.plan.reservable` |
| `ibm_onboarding_catalog_plan_invalid_metadata__pricing__origin` | `metadata.pricing.origin` |
| `ibm_onboarding_catalog_plan_invalid_metadata__pricing__sales_avenue` | `metadata.pricing.sales_avenue` |
| `ibm_onboarding_catalog_plan_invalid_metadata__pricing__type` | `metadata.pricing.type` |
| `ibm_onboarding_catalog_plan_invalid_metadata__service__bindable` | `metadata.service.bindable` |
| `ibm_onboarding_catalog_plan_invalid_metadata__service__iam_compatible` | `metadata.service.iam_compatible` |
| `ibm_onboarding_catalog_plan_invalid_metadata__service__plan_updateable` | `metadata.service.plan_updateable` |
| `ibm_onboarding_catalog_plan_invalid_metadata__service__rc_provisionable` | `metadata.service.rc_provisionable` |
| `ibm_onboarding_catalog_plan_invalid_metadata__service__service_key_supported` | `metadata.service.service_key_supported` |
| `ibm_onboarding_catalog_plan_invalid_metadata__ui__hidden` | `metadata.ui.hidden` |
| `ibm_onboarding_catalog_plan_invalid_metadata__ui__side_by_side_index` | `metadata.ui.side_by_side_index` |
| `ibm_onboarding_catalog_plan_invalid_metadata__ui__urls__apidocs_url` | `metadata.ui.urls.apidocs_url` |
| `ibm_onboarding_catalog_plan_invalid_metadata__ui__urls__catalog_details_url` | `metadata.ui.urls.catalog_details_url` |
| `ibm_onboarding_catalog_plan_invalid_metadata__ui__urls__custom_create_page_url` | `metadata.ui.urls.custom_create_page_url` |
| `ibm_onboarding_catalog_plan_invalid_metadata__ui__urls__dashboard` | `metadata.ui.urls.dashboard` |
| `ibm_onboarding_catalog_plan_invalid_metadata__ui__urls__doc_url` | `metadata.ui.urls.doc_url` |
| `ibm_onboarding_catalog_plan_invalid_metadata__ui__urls__instructions_url` | `metadata.ui.urls.instructions_url` |
| `ibm_onboarding_catalog_plan_invalid_metadata__ui__urls__terms_url` | `metadata.ui.urls.terms_url` |
| `ibm_onboarding_catalog_plan_invalid_object_provider__email` | `object_provider.email` |
| `ibm_onboarding_catalog_plan_invalid_object_provider__name` | `object_provider.name` |
| `ibm_onboarding_catalog_plan_invalid_overview_ui__en__description` | `overview_ui.en.description` |
| `ibm_onboarding_catalog_plan_invalid_overview_ui__en__display_name` | `overview_ui.en.display_name` |
| `ibm_onboarding_catalog_plan_invalid_overview_ui__en__long_description` | `overview_ui.en.long_description` |

## Example

### Valid

```hcl
resource "ibm_onboarding_catalog_plan" "test" {
  active = true
  catalog_product_id = "01234567-89ab-cdef-0123-456789abcdef"
  disabled = true
  kind = "plan"
  name = "aaaaaaa"
  product_id = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa:o:01234567-89ab-cdef-0123-456789abcdef"
  object_provider {
  }
}
```
