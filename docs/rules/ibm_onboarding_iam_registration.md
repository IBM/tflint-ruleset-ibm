# ibm_onboarding_iam_registration

Manages ibm_onboarding_iam_registration. The IAM registration name, which must be the programmatic name of the product.

This rule validates `ibm_onboarding_iam_registration` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `leaf-unique` (`partner-center-sell`).

## Configuration

```hcl
rule "ibm_onboarding_iam_registration" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `additional_policy_scopes` | no | — |
| `enabled` | no | — |
| `env` | no | pattern `^[a-z_.-]+$` |
| `name` | yes | — |
| `parent_ids` | no | — |
| `product_id` | yes | pattern `^[a-zA-Z0-9]{32}:o:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `service_type` | no | one of: `platform_service`, `service` |
| `supported_action_control` | no | — |
| `actions` (block) | no | nested block |
| `actions.api_types` | no | — |
| `actions.id` | no | — |
| `actions.roles` | no | — |
| `actions.description` (block) | no | nested block |
| `actions.description.de` | no | — |
| `actions.description.default` | no | — |
| `actions.description.en` | no | — |
| `actions.description.es` | no | — |
| `actions.description.fr` | no | — |
| `actions.description.it` | no | — |
| `actions.description.ja` | no | — |
| `actions.description.ko` | no | — |
| `actions.description.pt_br` | no | — |
| `actions.description.zh_cn` | no | — |
| `actions.description.zh_tw` | no | — |
| `actions.display_name` (block) | no | nested block |
| `actions.display_name.de` | no | — |
| `actions.display_name.default` | no | — |
| `actions.display_name.en` | no | — |
| `actions.display_name.es` | no | — |
| `actions.display_name.fr` | no | — |
| `actions.display_name.it` | no | — |
| `actions.display_name.ja` | no | — |
| `actions.display_name.ko` | no | — |
| `actions.display_name.pt_br` | no | — |
| `actions.display_name.zh_cn` | no | — |
| `actions.display_name.zh_tw` | no | — |
| `actions.options` (block) | no | nested block |
| `actions.options.hidden` | no | — |
| `display_name` (block) | no | nested block |
| `display_name.de` | no | — |
| `display_name.default` | no | — |
| `display_name.en` | no | — |
| `display_name.es` | no | — |
| `display_name.fr` | no | — |
| `display_name.it` | no | — |
| `display_name.ja` | no | — |
| `display_name.ko` | no | — |
| `display_name.pt_br` | no | — |
| `display_name.zh_cn` | no | — |
| `display_name.zh_tw` | no | — |
| `resource_hierarchy_attribute` (block) | no | nested block |
| `resource_hierarchy_attribute.key` | no | — |
| `resource_hierarchy_attribute.value` | no | API: The amount per unit of time. |
| `supported_anonymous_accesses` (block) | no | nested block |
| `supported_anonymous_accesses.roles` | no | — |
| `supported_anonymous_accesses.attributes` (block) | no | nested block |
| `supported_anonymous_accesses.attributes.account_id` | yes | — |
| `supported_anonymous_accesses.attributes.additional_properties` | yes | — |
| `supported_anonymous_accesses.attributes.service_name` | yes | — |
| `supported_attributes` (block) | no | nested block |
| `supported_attributes.key` | no | — |
| `supported_attributes.description` (block) | no | nested block |
| `supported_attributes.description.de` | no | — |
| `supported_attributes.description.default` | no | — |
| `supported_attributes.description.en` | no | — |
| `supported_attributes.description.es` | no | — |
| `supported_attributes.description.fr` | no | — |
| `supported_attributes.description.it` | no | — |
| `supported_attributes.description.ja` | no | — |
| `supported_attributes.description.ko` | no | — |
| `supported_attributes.description.pt_br` | no | — |
| `supported_attributes.description.zh_cn` | no | — |
| `supported_attributes.description.zh_tw` | no | — |
| `supported_attributes.display_name` (block) | no | nested block |
| `supported_attributes.display_name.de` | no | — |
| `supported_attributes.display_name.default` | no | — |
| `supported_attributes.display_name.en` | no | — |
| `supported_attributes.display_name.es` | no | — |
| `supported_attributes.display_name.fr` | no | — |
| `supported_attributes.display_name.it` | no | — |
| `supported_attributes.display_name.ja` | no | — |
| `supported_attributes.display_name.ko` | no | — |
| `supported_attributes.display_name.pt_br` | no | — |
| `supported_attributes.display_name.zh_cn` | no | — |
| `supported_attributes.display_name.zh_tw` | no | — |
| `supported_attributes.options` (block) | no | nested block |
| `supported_attributes.options.hidden` | no | — |
| `supported_attributes.options.is_empty_value_supported` | no | — |
| `supported_attributes.options.is_string_exists_false_value_supported` | no | — |
| `supported_attributes.options.key` | no | — |
| `supported_attributes.options.operators` | no | — |
| `supported_attributes.options.policy_types` | no | — |
| `supported_attributes.options.supported_patterns` | no | — |
| `supported_attributes.options.resource_hierarchy` (block) | no | nested block |
| `supported_attributes.options.resource_hierarchy.key` (block) | no | nested block |
| `supported_attributes.options.resource_hierarchy.key.key` | no | — |
| `supported_attributes.options.resource_hierarchy.key.value` | no | API: The amount per unit of time. |
| `supported_attributes.options.resource_hierarchy.value` (block) | no | nested block |
| `supported_attributes.options.resource_hierarchy.value.key` | no | — |
| `supported_attributes.ui` (block) | no | nested block |
| `supported_attributes.ui.input_type` | no | — |
| `supported_attributes.ui.input_details` (block) | no | nested block |
| `supported_attributes.ui.input_details.type` | no | — |
| `supported_attributes.ui.input_details.gst` (block) | no | nested block |
| `supported_attributes.ui.input_details.gst.input_option_label` | no | — |
| `supported_attributes.ui.input_details.gst.label_property_name` | no | — |
| `supported_attributes.ui.input_details.gst.query` | no | — |
| `supported_attributes.ui.input_details.gst.value_property_name` | no | — |
| `supported_attributes.ui.input_details.url` (block) | no | nested block |
| `supported_attributes.ui.input_details.url.input_option_label` | no | — |
| `supported_attributes.ui.input_details.url.url_endpoint` | no | — |
| `supported_attributes.ui.input_details.values` (block) | no | nested block |
| `supported_attributes.ui.input_details.values.value` | no | API: The amount per unit of time. |
| `supported_attributes.ui.input_details.values.display_name` (block) | no | nested block |
| `supported_authorization_subjects` (block) | no | nested block |
| `supported_authorization_subjects.roles` | no | — |
| `supported_authorization_subjects.attributes` (block) | no | nested block |
| `supported_authorization_subjects.attributes.resource_type` | no | — |
| `supported_authorization_subjects.attributes.service_name` | no | — |
| `supported_network` (block) | no | nested block |
| `supported_network.environment_attributes` (block) | no | nested block |
| `supported_network.environment_attributes.key` | no | — |
| `supported_network.environment_attributes.values` | no | — |
| `supported_network.environment_attributes.options` (block) | no | nested block |
| `supported_network.environment_attributes.options.hidden` | no | — |
| `supported_network.operations` (block) | no | nested block |
| `supported_network.operations.api_types` (block) | no | nested block |
| `supported_network.operations.api_types.enforcement_method` | no | — |
| `supported_network.operations.api_types.name` | no | — |
| `supported_network.operations.api_types.description` (block) | no | nested block |
| `supported_network.operations.api_types.description.de` | no | — |
| `supported_network.operations.api_types.description.default` | no | — |
| `supported_network.operations.api_types.description.en` | no | — |
| `supported_network.operations.api_types.description.es` | no | — |
| `supported_network.operations.api_types.description.fr` | no | — |
| `supported_network.operations.api_types.description.it` | no | — |
| `supported_network.operations.api_types.description.ja` | no | — |
| `supported_network.operations.api_types.description.ko` | no | — |
| `supported_network.operations.api_types.description.pt_br` | no | — |
| `supported_network.operations.api_types.description.zh_cn` | no | — |
| `supported_network.operations.api_types.description.zh_tw` | no | — |
| `supported_network.operations.api_types.display_name` (block) | no | nested block |
| `supported_network.operations.api_types.display_name.de` | no | — |
| `supported_network.operations.api_types.display_name.default` | no | — |
| `supported_network.operations.api_types.display_name.en` | no | — |
| `supported_network.operations.api_types.display_name.es` | no | — |
| `supported_network.operations.api_types.display_name.fr` | no | — |
| `supported_network.operations.api_types.display_name.it` | no | — |
| `supported_network.operations.api_types.display_name.ja` | no | — |
| `supported_network.operations.api_types.display_name.ko` | no | — |
| `supported_network.operations.api_types.display_name.pt_br` | no | — |
| `supported_network.operations.api_types.display_name.zh_cn` | no | — |
| `supported_network.operations.api_types.display_name.zh_tw` | no | — |
| `supported_network.operations.api_types.event_publishing` (block) | no | nested block |
| `supported_network.operations.api_types.event_publishing.state` | no | — |
| `supported_network.operations.defaults` (block) | no | nested block |
| `supported_network.operations.defaults.enforcement_method` | no | — |
| `supported_network.operations.defaults.event_publishing` (block) | no | nested block |
| `supported_network.operations.defaults.event_publishing.state` | no | — |
| `supported_network.self_managed_allowlist_enforcement` (block) | no | nested block |
| `supported_network.self_managed_allowlist_enforcement.event_publishing` (block) | no | nested block |
| `supported_network.self_managed_allowlist_enforcement.event_publishing.api_types` | no | — |
| `supported_roles` (block) | no | nested block |
| `supported_roles.id` | no | — |
| `supported_roles.description` (block) | no | nested block |
| `supported_roles.description.de` | no | — |
| `supported_roles.description.default` | no | — |
| `supported_roles.description.en` | no | — |
| `supported_roles.description.es` | no | — |
| `supported_roles.description.fr` | no | — |
| `supported_roles.description.it` | no | — |
| `supported_roles.description.ja` | no | — |
| `supported_roles.description.ko` | no | — |
| `supported_roles.description.pt_br` | no | — |
| `supported_roles.description.zh_cn` | no | — |
| `supported_roles.description.zh_tw` | no | — |
| `supported_roles.display_name` (block) | no | nested block |
| `supported_roles.display_name.de` | no | — |
| `supported_roles.display_name.default` | no | — |
| `supported_roles.display_name.en` | no | — |
| `supported_roles.display_name.es` | no | — |
| `supported_roles.display_name.fr` | no | — |
| `supported_roles.display_name.it` | no | — |
| `supported_roles.display_name.ja` | no | — |
| `supported_roles.display_name.ko` | no | — |
| `supported_roles.display_name.pt_br` | no | — |
| `supported_roles.display_name.zh_cn` | no | — |
| `supported_roles.display_name.zh_tw` | no | — |
| `supported_roles.options` (block) | no | nested block |
| `supported_roles.options.access_policy` | yes | — |
| `supported_roles.options.account_type` | no | — |
| `supported_roles.options.policy_type` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_onboarding_iam_registration_invalid_additional_policy_scopes` | `additional_policy_scopes` |
| `ibm_onboarding_iam_registration_invalid_enabled` | `enabled` |
| `ibm_onboarding_iam_registration_invalid_env` | `env` |
| `ibm_onboarding_iam_registration_invalid_name` | `name` |
| `ibm_onboarding_iam_registration_invalid_parent_ids` | `parent_ids` |
| `ibm_onboarding_iam_registration_invalid_product_id` | `product_id` |
| `ibm_onboarding_iam_registration_invalid_service_type` | `service_type` |
| `ibm_onboarding_iam_registration_invalid_supported_action_control` | `supported_action_control` |
| `ibm_onboarding_iam_registration_invalid_actions__api_types` | `actions.api_types` |
| `ibm_onboarding_iam_registration_invalid_actions__id` | `actions.id` |
| `ibm_onboarding_iam_registration_invalid_actions__roles` | `actions.roles` |
| `ibm_onboarding_iam_registration_invalid_actions__description__de` | `actions.description.de` |
| `ibm_onboarding_iam_registration_invalid_actions__description__default` | `actions.description.default` |
| `ibm_onboarding_iam_registration_invalid_actions__description__en` | `actions.description.en` |
| `ibm_onboarding_iam_registration_invalid_actions__description__es` | `actions.description.es` |
| `ibm_onboarding_iam_registration_invalid_actions__description__fr` | `actions.description.fr` |
| `ibm_onboarding_iam_registration_invalid_actions__description__it` | `actions.description.it` |
| `ibm_onboarding_iam_registration_invalid_actions__description__ja` | `actions.description.ja` |
| `ibm_onboarding_iam_registration_invalid_actions__description__ko` | `actions.description.ko` |
| `ibm_onboarding_iam_registration_invalid_actions__description__pt_br` | `actions.description.pt_br` |
| `ibm_onboarding_iam_registration_invalid_actions__description__zh_cn` | `actions.description.zh_cn` |
| `ibm_onboarding_iam_registration_invalid_actions__description__zh_tw` | `actions.description.zh_tw` |
| `ibm_onboarding_iam_registration_invalid_actions__display_name__de` | `actions.display_name.de` |
| `ibm_onboarding_iam_registration_invalid_actions__display_name__default` | `actions.display_name.default` |
| `ibm_onboarding_iam_registration_invalid_actions__display_name__en` | `actions.display_name.en` |
| `ibm_onboarding_iam_registration_invalid_actions__display_name__es` | `actions.display_name.es` |
| `ibm_onboarding_iam_registration_invalid_actions__display_name__fr` | `actions.display_name.fr` |
| `ibm_onboarding_iam_registration_invalid_actions__display_name__it` | `actions.display_name.it` |
| `ibm_onboarding_iam_registration_invalid_actions__display_name__ja` | `actions.display_name.ja` |
| `ibm_onboarding_iam_registration_invalid_actions__display_name__ko` | `actions.display_name.ko` |
| `ibm_onboarding_iam_registration_invalid_actions__display_name__pt_br` | `actions.display_name.pt_br` |
| `ibm_onboarding_iam_registration_invalid_actions__display_name__zh_cn` | `actions.display_name.zh_cn` |
| `ibm_onboarding_iam_registration_invalid_actions__display_name__zh_tw` | `actions.display_name.zh_tw` |
| `ibm_onboarding_iam_registration_invalid_actions__options__hidden` | `actions.options.hidden` |
| `ibm_onboarding_iam_registration_invalid_display_name__de` | `display_name.de` |
| `ibm_onboarding_iam_registration_invalid_display_name__default` | `display_name.default` |
| `ibm_onboarding_iam_registration_invalid_display_name__en` | `display_name.en` |
| `ibm_onboarding_iam_registration_invalid_display_name__es` | `display_name.es` |
| `ibm_onboarding_iam_registration_invalid_display_name__fr` | `display_name.fr` |
| `ibm_onboarding_iam_registration_invalid_display_name__it` | `display_name.it` |
| `ibm_onboarding_iam_registration_invalid_display_name__ja` | `display_name.ja` |
| `ibm_onboarding_iam_registration_invalid_display_name__ko` | `display_name.ko` |
| `ibm_onboarding_iam_registration_invalid_display_name__pt_br` | `display_name.pt_br` |
| `ibm_onboarding_iam_registration_invalid_display_name__zh_cn` | `display_name.zh_cn` |
| `ibm_onboarding_iam_registration_invalid_display_name__zh_tw` | `display_name.zh_tw` |
| `ibm_onboarding_iam_registration_invalid_resource_hierarchy_attribute__key` | `resource_hierarchy_attribute.key` |
| `ibm_onboarding_iam_registration_invalid_resource_hierarchy_attribute__value` | `resource_hierarchy_attribute.value` |
| `ibm_onboarding_iam_registration_invalid_supported_anonymous_accesses__roles` | `supported_anonymous_accesses.roles` |
| `ibm_onboarding_iam_registration_invalid_supported_anonymous_accesses__attributes__account_id` | `supported_anonymous_accesses.attributes.account_id` |
| `ibm_onboarding_iam_registration_invalid_supported_anonymous_accesses__attributes__additional_properties` | `supported_anonymous_accesses.attributes.additional_properties` |
| `ibm_onboarding_iam_registration_invalid_supported_anonymous_accesses__attributes__service_name` | `supported_anonymous_accesses.attributes.service_name` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__key` | `supported_attributes.key` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__description__de` | `supported_attributes.description.de` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__description__default` | `supported_attributes.description.default` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__description__en` | `supported_attributes.description.en` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__description__es` | `supported_attributes.description.es` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__description__fr` | `supported_attributes.description.fr` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__description__it` | `supported_attributes.description.it` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__description__ja` | `supported_attributes.description.ja` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__description__ko` | `supported_attributes.description.ko` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__description__pt_br` | `supported_attributes.description.pt_br` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__description__zh_cn` | `supported_attributes.description.zh_cn` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__description__zh_tw` | `supported_attributes.description.zh_tw` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__display_name__de` | `supported_attributes.display_name.de` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__display_name__default` | `supported_attributes.display_name.default` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__display_name__en` | `supported_attributes.display_name.en` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__display_name__es` | `supported_attributes.display_name.es` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__display_name__fr` | `supported_attributes.display_name.fr` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__display_name__it` | `supported_attributes.display_name.it` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__display_name__ja` | `supported_attributes.display_name.ja` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__display_name__ko` | `supported_attributes.display_name.ko` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__display_name__pt_br` | `supported_attributes.display_name.pt_br` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__display_name__zh_cn` | `supported_attributes.display_name.zh_cn` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__display_name__zh_tw` | `supported_attributes.display_name.zh_tw` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__options__hidden` | `supported_attributes.options.hidden` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__options__is_empty_value_supported` | `supported_attributes.options.is_empty_value_supported` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__options__is_string_exists_false_value_supported` | `supported_attributes.options.is_string_exists_false_value_supported` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__options__key` | `supported_attributes.options.key` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__options__operators` | `supported_attributes.options.operators` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__options__policy_types` | `supported_attributes.options.policy_types` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__options__supported_patterns` | `supported_attributes.options.supported_patterns` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__options__resource_hierarchy__key__key` | `supported_attributes.options.resource_hierarchy.key.key` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__options__resource_hierarchy__key__value` | `supported_attributes.options.resource_hierarchy.key.value` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__options__resource_hierarchy__value__key` | `supported_attributes.options.resource_hierarchy.value.key` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__ui__input_type` | `supported_attributes.ui.input_type` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__ui__input_details__type` | `supported_attributes.ui.input_details.type` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__ui__input_details__gst__input_option_label` | `supported_attributes.ui.input_details.gst.input_option_label` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__ui__input_details__gst__label_property_name` | `supported_attributes.ui.input_details.gst.label_property_name` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__ui__input_details__gst__query` | `supported_attributes.ui.input_details.gst.query` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__ui__input_details__gst__value_property_name` | `supported_attributes.ui.input_details.gst.value_property_name` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__ui__input_details__url__input_option_label` | `supported_attributes.ui.input_details.url.input_option_label` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__ui__input_details__url__url_endpoint` | `supported_attributes.ui.input_details.url.url_endpoint` |
| `ibm_onboarding_iam_registration_invalid_supported_attributes__ui__input_details__values__value` | `supported_attributes.ui.input_details.values.value` |
| `ibm_onboarding_iam_registration_invalid_supported_authorization_subjects__roles` | `supported_authorization_subjects.roles` |
| `ibm_onboarding_iam_registration_invalid_supported_authorization_subjects__attributes__resource_type` | `supported_authorization_subjects.attributes.resource_type` |
| `ibm_onboarding_iam_registration_invalid_supported_authorization_subjects__attributes__service_name` | `supported_authorization_subjects.attributes.service_name` |
| `ibm_onboarding_iam_registration_invalid_supported_network__environment_attributes__key` | `supported_network.environment_attributes.key` |
| `ibm_onboarding_iam_registration_invalid_supported_network__environment_attributes__values` | `supported_network.environment_attributes.values` |
| `ibm_onboarding_iam_registration_invalid_supported_network__environment_attributes__options__hidden` | `supported_network.environment_attributes.options.hidden` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__enforcement_method` | `supported_network.operations.api_types.enforcement_method` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__name` | `supported_network.operations.api_types.name` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__description__de` | `supported_network.operations.api_types.description.de` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__description__default` | `supported_network.operations.api_types.description.default` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__description__en` | `supported_network.operations.api_types.description.en` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__description__es` | `supported_network.operations.api_types.description.es` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__description__fr` | `supported_network.operations.api_types.description.fr` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__description__it` | `supported_network.operations.api_types.description.it` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__description__ja` | `supported_network.operations.api_types.description.ja` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__description__ko` | `supported_network.operations.api_types.description.ko` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__description__pt_br` | `supported_network.operations.api_types.description.pt_br` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__description__zh_cn` | `supported_network.operations.api_types.description.zh_cn` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__description__zh_tw` | `supported_network.operations.api_types.description.zh_tw` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__display_name__de` | `supported_network.operations.api_types.display_name.de` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__display_name__default` | `supported_network.operations.api_types.display_name.default` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__display_name__en` | `supported_network.operations.api_types.display_name.en` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__display_name__es` | `supported_network.operations.api_types.display_name.es` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__display_name__fr` | `supported_network.operations.api_types.display_name.fr` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__display_name__it` | `supported_network.operations.api_types.display_name.it` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__display_name__ja` | `supported_network.operations.api_types.display_name.ja` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__display_name__ko` | `supported_network.operations.api_types.display_name.ko` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__display_name__pt_br` | `supported_network.operations.api_types.display_name.pt_br` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__display_name__zh_cn` | `supported_network.operations.api_types.display_name.zh_cn` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__display_name__zh_tw` | `supported_network.operations.api_types.display_name.zh_tw` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__api_types__event_publishing__state` | `supported_network.operations.api_types.event_publishing.state` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__defaults__enforcement_method` | `supported_network.operations.defaults.enforcement_method` |
| `ibm_onboarding_iam_registration_invalid_supported_network__operations__defaults__event_publishing__state` | `supported_network.operations.defaults.event_publishing.state` |
| `ibm_onboarding_iam_registration_invalid_supported_network__self_managed_allowlist_enforcement__event_publishing__api_types` | `supported_network.self_managed_allowlist_enforcement.event_publishing.api_types` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__id` | `supported_roles.id` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__description__de` | `supported_roles.description.de` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__description__default` | `supported_roles.description.default` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__description__en` | `supported_roles.description.en` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__description__es` | `supported_roles.description.es` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__description__fr` | `supported_roles.description.fr` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__description__it` | `supported_roles.description.it` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__description__ja` | `supported_roles.description.ja` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__description__ko` | `supported_roles.description.ko` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__description__pt_br` | `supported_roles.description.pt_br` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__description__zh_cn` | `supported_roles.description.zh_cn` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__description__zh_tw` | `supported_roles.description.zh_tw` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__display_name__de` | `supported_roles.display_name.de` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__display_name__default` | `supported_roles.display_name.default` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__display_name__en` | `supported_roles.display_name.en` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__display_name__es` | `supported_roles.display_name.es` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__display_name__fr` | `supported_roles.display_name.fr` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__display_name__it` | `supported_roles.display_name.it` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__display_name__ja` | `supported_roles.display_name.ja` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__display_name__ko` | `supported_roles.display_name.ko` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__display_name__pt_br` | `supported_roles.display_name.pt_br` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__display_name__zh_cn` | `supported_roles.display_name.zh_cn` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__display_name__zh_tw` | `supported_roles.display_name.zh_tw` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__options__access_policy` | `supported_roles.options.access_policy` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__options__account_type` | `supported_roles.options.account_type` |
| `ibm_onboarding_iam_registration_invalid_supported_roles__options__policy_type` | `supported_roles.options.policy_type` |

## Example

### Valid

```hcl
resource "ibm_onboarding_iam_registration" "test" {
  name = "example"
  product_id = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa:o:01234567-89ab-cdef-0123-456789abcdef"
}
```
