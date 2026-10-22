# ibm_onboarding_catalog_deployment

Manages ibm_onboarding_catalog_deployment. The programmatic name of this deployment.

This rule validates `ibm_onboarding_catalog_deployment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/products/{product_id}/catalog:patch` (`partner-center-sell`).

## Configuration

```hcl
rule "ibm_onboarding_catalog_deployment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `active` | yes | — |
| `catalog_deployment_id` | no | — |
| `catalog_plan_id` | yes | — |
| `catalog_product_id` | yes | — |
| `disabled` | yes | — |
| `env` | no | pattern `^[a-z_.-]+$` |
| `geo_tags` | no | — |
| `kind` | yes | one of: `deployment` |
| `name` | yes | API: The name of the product. |
| `object_id` | no | — |
| `product_id` | yes | pattern `^[a-zA-Z0-9]{32}:o:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `tags` | no | pattern `^[a-z0-9]+([ _:-]?[a-z0-9]+){0,10}$` |
| `url` | no | — |
| `metadata` (block) | no | nested block |
| `metadata.rc_compatible` | no | — |
| `metadata.deployment` (block) | no | nested block |
| `metadata.deployment.location` | no | — |
| `metadata.deployment.location_url` | no | — |
| `metadata.deployment.target_crn` | no | crn format |
| `metadata.deployment.broker` (block) | no | nested block |
| `metadata.deployment.broker.guid` | no | — |
| `metadata.deployment.broker.name` | no | API: The name of the product. |
| `metadata.other` (block) | no | nested block |
| `metadata.other.location_proxied_by` | no | — |
| `metadata.service` (block) | no | nested block |
| `metadata.service.iam_compatible` | no | — |
| `metadata.service.plan_updateable` | no | — |
| `metadata.service.rc_provisionable` | no | — |
| `metadata.service.service_key_supported` | no | — |
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
| `ibm_onboarding_catalog_deployment_invalid_active` | `active` |
| `ibm_onboarding_catalog_deployment_invalid_catalog_plan_id` | `catalog_plan_id` |
| `ibm_onboarding_catalog_deployment_invalid_catalog_product_id` | `catalog_product_id` |
| `ibm_onboarding_catalog_deployment_invalid_disabled` | `disabled` |
| `ibm_onboarding_catalog_deployment_invalid_env` | `env` |
| `ibm_onboarding_catalog_deployment_invalid_kind` | `kind` |
| `ibm_onboarding_catalog_deployment_invalid_name` | `name` |
| `ibm_onboarding_catalog_deployment_invalid_object_id` | `object_id` |
| `ibm_onboarding_catalog_deployment_invalid_product_id` | `product_id` |
| `ibm_onboarding_catalog_deployment_invalid_tags` | `tags` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__rc_compatible` | `metadata.rc_compatible` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__deployment__location` | `metadata.deployment.location` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__deployment__location_url` | `metadata.deployment.location_url` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__deployment__target_crn` | `metadata.deployment.target_crn` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__deployment__broker__guid` | `metadata.deployment.broker.guid` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__deployment__broker__name` | `metadata.deployment.broker.name` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__other__location_proxied_by` | `metadata.other.location_proxied_by` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__iam_compatible` | `metadata.service.iam_compatible` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__plan_updateable` | `metadata.service.plan_updateable` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__rc_provisionable` | `metadata.service.rc_provisionable` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__service_key_supported` | `metadata.service.service_key_supported` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__description` | `metadata.service.parameters.description` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__displayname` | `metadata.service.parameters.displayname` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__hidden` | `metadata.service.parameters.hidden` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__invalidmessage` | `metadata.service.parameters.invalidmessage` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__layout` | `metadata.service.parameters.layout` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__name` | `metadata.service.parameters.name` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__options_url` | `metadata.service.parameters.options_url` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__pattern` | `metadata.service.parameters.pattern` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__placeholder` | `metadata.service.parameters.placeholder` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__readonly` | `metadata.service.parameters.readonly` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__required` | `metadata.service.parameters.required` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__type` | `metadata.service.parameters.type` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__validation_url` | `metadata.service.parameters.validation_url` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__value` | `metadata.service.parameters.value` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__options__displayname` | `metadata.service.parameters.options.displayname` |
| `ibm_onboarding_catalog_deployment_invalid_metadata__service__parameters__options__value` | `metadata.service.parameters.options.value` |
| `ibm_onboarding_catalog_deployment_invalid_object_provider__email` | `object_provider.email` |
| `ibm_onboarding_catalog_deployment_invalid_object_provider__name` | `object_provider.name` |
| `ibm_onboarding_catalog_deployment_invalid_overview_ui__en__description` | `overview_ui.en.description` |
| `ibm_onboarding_catalog_deployment_invalid_overview_ui__en__display_name` | `overview_ui.en.display_name` |
| `ibm_onboarding_catalog_deployment_invalid_overview_ui__en__long_description` | `overview_ui.en.long_description` |

## Example

### Valid

```hcl
resource "ibm_onboarding_catalog_deployment" "test" {
  active = true
  catalog_plan_id = "01234567-89ab-cdef-0123-456789abcdef"
  catalog_product_id = "01234567-89ab-cdef-0123-456789abcdef"
  disabled = true
  kind = "deployment"
  name = "aaaaaaa"
  product_id = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa:o:01234567-89ab-cdef-0123-456789abcdef"
  object_provider {
  }
}
```
