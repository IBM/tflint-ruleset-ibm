# ibm_cm_version

Manages ibm_cm_version. Name of version.

This rule validates `ibm_cm_version` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Version` (`resource-catalog-private-catalog`).

## Configuration

```hcl
rule "ibm_cm_version" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `catalog_id` | yes | API: Catalog ID |
| `content` | no | API: Byte array representing the content to be imported. Only supported for OVA images at this time. |
| `created` | no | — |
| `crn` | no | — |
| `deprecate` | no | — |
| `deprecate_pending` | no | — |
| `deprecated` | no | — |
| `entitlement` | no | — |
| `format_kind` | no | — |
| `image_manifest_url` | no | — |
| `image_pull_key_name` | no | — |
| `import_sha` | no | — |
| `include_config` | no | — |
| `install_kind` | no | — |
| `is_consumable` | no | — |
| `is_vsi` | no | — |
| `kind_id` | no | — |
| `label` | no | API: Label for this flavor |
| `long_description` | no | API: Long description for version |
| `long_description_i18n` | no | — |
| `metadata` | no | — |
| `name` | no | API: Programmatic name for this flavor |
| `offering_id` | yes | API: Offering ID |
| `offering_identifier` | no | — |
| `outputs` | no | — |
| `package_version` | no | — |
| `product_kind` | no | — |
| `repo_url` | no | — |
| `repotype` | no | — |
| `rev` | no | — |
| `sha` | no | — |
| `single_instance` | no | — |
| `source_url` | no | — |
| `state` | no | — |
| `tags` | no | API: List of tags associated with this catalog |
| `target_kinds` | no | — |
| `target_version` | no | — |
| `terraform_version` | no | — |
| `tgz_url` | no | — |
| `updated` | no | — |
| `usage` | no | — |
| `validation` | no | — |
| `version` | no | — |
| `version_id` | no | — |
| `version_locator` | no | — |
| `working_directory` | no | — |
| `x_auth_token` | no | — |
| `zipurl` | no | — |
| `configuration` (block) | no | nested block |
| `configuration.default_value` | no | API: The default value. To use a secret when the type is password, specify a JSON encoded value of $ref:#/components/schemas/SecretInstance, prefixed with `cmsm_v1:` |
| `configuration.description` | no | API: Key description |
| `configuration.display_name` | no | API: Display name for configuration type |
| `configuration.hidden` | no | API: Hide values |
| `configuration.key` | no | API: Configuration key |
| `configuration.options` | no | API: An option value of any type |
| `configuration.required` | no | API: Is key required to install |
| `configuration.type` | no | API: Value type (string, boolean, int) |
| `configuration.type_metadata` | no | API: The original type, as found in the source being onboarded |
| `configuration.value_constraint` | no | API: Deprecated - Constraint associated with value, e.g., for string type - regx:[a-z] |
| `configuration.custom_config` (block) | no | nested block |
| `configuration.custom_config.config_constraints` | no | API: Map of constraint parameters that will be passed to the custom widget |
| `configuration.custom_config.grouping` | no | API: Determines where this configuration type is rendered (3 sections today - Target, Resource, and Deployment) |
| `configuration.custom_config.grouping_index` | no | API: Determines the order that this configuration item shows in that particular grouping |
| `configuration.custom_config.original_grouping` | no | API: Original grouping type for this configuration (3 types - Target, Resource, and Deployment) |
| `configuration.custom_config.type` | no | API: ID of the widget type |
| `configuration.custom_config.associations` (block) | no | nested block |
| `configuration.custom_config.associations.parameters` (block) | no | nested block |
| `configuration.custom_config.associations.parameters.name` | no | API: Name of this parameter |
| `configuration.custom_config.associations.parameters.options_refresh` | no | — |
| `configuration.value_constraints` (block) | no | nested block |
| `configuration.value_constraints.description` | no | API: The value to display if the inptu value does not match the specified constraint. |
| `configuration.value_constraints.type` | no | one of: `regex` |
| `configuration.value_constraints.value` | no | API: Contstraint value. For type regex, this is a regular expression in Javascript notation |
| `flavor` (block) | no | nested block |
| `flavor.index` | no | API: Order that this flavor should appear when listed for a single version |
| `flavor.label` | no | API: Label for this flavor |
| `flavor.label_i18n` | no | API: A map of translated strings, by language code |
| `flavor.name` | no | API: Programmatic name for this flavor |
| `iam_permissions` (block) | no | nested block |
| `iam_permissions.role_crns` | no | API: Role CRNs for this permission |
| `iam_permissions.service_name` | no | API: Service name |
| `iam_permissions.resources` (block) | no | nested block |
| `iam_permissions.resources.description` | no | API: Resource description |
| `iam_permissions.resources.name` | no | API: Resource name |
| `iam_permissions.resources.role_crns` | no | API: Role CRNs for this permission |
| `import_metadata` (block) | no | nested block |
| `import_metadata.minimum_provisioned_size` | no | — |
| `import_metadata.file` (block) | no | nested block |
| `import_metadata.file.size` | no | range 1–1000 |
| `import_metadata.images` (block) | no | nested block |
| `import_metadata.images.id` | no | API: Unique ID |
| `import_metadata.images.name` | no | API: Programmatic name for this flavor |
| `import_metadata.images.region` | no | — |
| `import_metadata.operating_system` (block) | no | nested block |
| `import_metadata.operating_system.architecture` | no | — |
| `import_metadata.operating_system.dedicated_host_only` | no | — |
| `import_metadata.operating_system.display_name` | no | API: Display name for configuration type |
| `import_metadata.operating_system.family` | no | — |
| `import_metadata.operating_system.href` | no | — |
| `import_metadata.operating_system.name` | no | API: Programmatic name for this flavor |
| `import_metadata.operating_system.vendor` | no | — |
| `import_metadata.operating_system.version` | no | API: Version of content type |
| `install` (block) | no | nested block |
| `install.delete_script` | no | API: Optional script that if run will remove the installed version |
| `install.instructions` | no | API: Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version |
| `install.instructions_i18n` | no | API: A map of translated strings, by language code |
| `install.scope` | no | API: Optional value indicating if this script is scoped to a namespace or the entire cluster |
| `install.script` | no | API: Optional script that needs to be run post any pre-condition script |
| `install.script_permission` | no | API: Optional iam permissions that are required on the target cluster to run this script |
| `licenses` (block) | no | nested block |
| `licenses.description` | no | API: License description |
| `licenses.id` | no | API: License ID |
| `licenses.name` | no | API: license name |
| `licenses.type` | no | API: type of license e.g., Apache xxx |
| `licenses.url` | no | API: URL for the license text |
| `pre_install` (block) | no | nested block |
| `pre_install.delete_script` | no | API: Optional script that if run will remove the installed version |
| `pre_install.instructions` | no | API: Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version |
| `pre_install.instructions_i18n` | no | API: A map of translated strings, by language code |
| `pre_install.scope` | no | API: Optional value indicating if this script is scoped to a namespace or the entire cluster |
| `pre_install.script` | no | API: Optional script that needs to be run post any pre-condition script |
| `pre_install.script_permission` | no | API: Optional iam permissions that are required on the target cluster to run this script |
| `required_resources` (block) | no | nested block |
| `required_resources.type` | no | one of: `mem`, `disk`, `cores`, `targetVersion`, `nodes`, `terraformVersion` |
| `required_resources.value` | no | API: mem, disk, cores, and nodes can be parsed as an int. targetVersion and terraformVersion will be a semver range value. |
| `solution_info` (block) | no | nested block |
| `solution_info.architecture_diagrams` (block) | no | nested block |
| `solution_info.architecture_diagrams.description` | no | API: Description of this diagram |
| `solution_info.architecture_diagrams.description_i18n` | no | API: A map of translated strings, by language code |
| `solution_info.architecture_diagrams.diagram` (block) | no | nested block |
| `solution_info.architecture_diagrams.diagram.api_url` | no | API: CM API specific URL of the specified media item |
| `solution_info.architecture_diagrams.diagram.caption` | no | API: Caption for this media item |
| `solution_info.architecture_diagrams.diagram.caption_i18n` | no | API: A map of translated strings, by language code |
| `solution_info.architecture_diagrams.diagram.thumbnail_url` | no | API: Thumbnail URL for this media item |
| `solution_info.architecture_diagrams.diagram.type` | no | API: Type of this media item |
| `solution_info.architecture_diagrams.diagram.url` | no | API: URL of the specified media item |
| `solution_info.architecture_diagrams.diagram.url_proxy` (block) | no | nested block |
| `solution_info.architecture_diagrams.diagram.url_proxy.sha` | no | API: SHA256 fingerprint of image |
| `solution_info.architecture_diagrams.diagram.url_proxy.url` | no | API: URL of the specified media item being proxied |
| `solution_info.dependencies` (block) | no | nested block |
| `solution_info.dependencies.catalog_id` | no | API: Optional - If not specified, assumes the Public Catalog |
| `solution_info.dependencies.flavors` | no | API: Optional - List of dependent flavors in the specified range |
| `solution_info.dependencies.id` | no | API: Optional - Offering ID - not required if name is set |
| `solution_info.dependencies.name` | no | API: Optional - Programmatic Offering name |
| `solution_info.dependencies.version` | no | API: Required - Semver value or range |
| `solution_info.features` (block) | no | nested block |
| `solution_info.features.description` | no | API: Feature description |
| `solution_info.features.description_i18n` | no | API: A map of translated strings, by language code |
| `solution_info.features.title` | no | API: Heading |
| `solution_info.features.title_i18n` | no | API: A map of translated strings, by language code |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cm_version_invalid_catalog_id` | `catalog_id` |
| `ibm_cm_version_invalid_content` | `content` |
| `ibm_cm_version_invalid_deprecate` | `deprecate` |
| `ibm_cm_version_invalid_format_kind` | `format_kind` |
| `ibm_cm_version_invalid_import_sha` | `import_sha` |
| `ibm_cm_version_invalid_include_config` | `include_config` |
| `ibm_cm_version_invalid_install_kind` | `install_kind` |
| `ibm_cm_version_invalid_is_vsi` | `is_vsi` |
| `ibm_cm_version_invalid_label` | `label` |
| `ibm_cm_version_invalid_long_description` | `long_description` |
| `ibm_cm_version_invalid_name` | `name` |
| `ibm_cm_version_invalid_offering_id` | `offering_id` |
| `ibm_cm_version_invalid_product_kind` | `product_kind` |
| `ibm_cm_version_invalid_repotype` | `repotype` |
| `ibm_cm_version_invalid_tags` | `tags` |
| `ibm_cm_version_invalid_target_kinds` | `target_kinds` |
| `ibm_cm_version_invalid_target_version` | `target_version` |
| `ibm_cm_version_invalid_terraform_version` | `terraform_version` |
| `ibm_cm_version_invalid_usage` | `usage` |
| `ibm_cm_version_invalid_working_directory` | `working_directory` |
| `ibm_cm_version_invalid_x_auth_token` | `x_auth_token` |
| `ibm_cm_version_invalid_zipurl` | `zipurl` |
| `ibm_cm_version_invalid_configuration__default_value` | `configuration.default_value` |
| `ibm_cm_version_invalid_configuration__description` | `configuration.description` |
| `ibm_cm_version_invalid_configuration__display_name` | `configuration.display_name` |
| `ibm_cm_version_invalid_configuration__hidden` | `configuration.hidden` |
| `ibm_cm_version_invalid_configuration__key` | `configuration.key` |
| `ibm_cm_version_invalid_configuration__options` | `configuration.options` |
| `ibm_cm_version_invalid_configuration__required` | `configuration.required` |
| `ibm_cm_version_invalid_configuration__type` | `configuration.type` |
| `ibm_cm_version_invalid_configuration__type_metadata` | `configuration.type_metadata` |
| `ibm_cm_version_invalid_configuration__value_constraint` | `configuration.value_constraint` |
| `ibm_cm_version_invalid_configuration__custom_config__config_constraints` | `configuration.custom_config.config_constraints` |
| `ibm_cm_version_invalid_configuration__custom_config__grouping` | `configuration.custom_config.grouping` |
| `ibm_cm_version_invalid_configuration__custom_config__grouping_index` | `configuration.custom_config.grouping_index` |
| `ibm_cm_version_invalid_configuration__custom_config__original_grouping` | `configuration.custom_config.original_grouping` |
| `ibm_cm_version_invalid_configuration__custom_config__type` | `configuration.custom_config.type` |
| `ibm_cm_version_invalid_configuration__custom_config__associations__parameters__name` | `configuration.custom_config.associations.parameters.name` |
| `ibm_cm_version_invalid_configuration__custom_config__associations__parameters__options_refresh` | `configuration.custom_config.associations.parameters.options_refresh` |
| `ibm_cm_version_invalid_configuration__value_constraints__description` | `configuration.value_constraints.description` |
| `ibm_cm_version_invalid_configuration__value_constraints__type` | `configuration.value_constraints.type` |
| `ibm_cm_version_invalid_configuration__value_constraints__value` | `configuration.value_constraints.value` |
| `ibm_cm_version_invalid_flavor__index` | `flavor.index` |
| `ibm_cm_version_invalid_flavor__label` | `flavor.label` |
| `ibm_cm_version_invalid_flavor__label_i18n` | `flavor.label_i18n` |
| `ibm_cm_version_invalid_flavor__name` | `flavor.name` |
| `ibm_cm_version_invalid_iam_permissions__role_crns` | `iam_permissions.role_crns` |
| `ibm_cm_version_invalid_iam_permissions__service_name` | `iam_permissions.service_name` |
| `ibm_cm_version_invalid_iam_permissions__resources__description` | `iam_permissions.resources.description` |
| `ibm_cm_version_invalid_iam_permissions__resources__name` | `iam_permissions.resources.name` |
| `ibm_cm_version_invalid_iam_permissions__resources__role_crns` | `iam_permissions.resources.role_crns` |
| `ibm_cm_version_invalid_import_metadata__minimum_provisioned_size` | `import_metadata.minimum_provisioned_size` |
| `ibm_cm_version_invalid_import_metadata__file__size` | `import_metadata.file.size` |
| `ibm_cm_version_invalid_import_metadata__images__id` | `import_metadata.images.id` |
| `ibm_cm_version_invalid_import_metadata__images__name` | `import_metadata.images.name` |
| `ibm_cm_version_invalid_import_metadata__images__region` | `import_metadata.images.region` |
| `ibm_cm_version_invalid_import_metadata__operating_system__architecture` | `import_metadata.operating_system.architecture` |
| `ibm_cm_version_invalid_import_metadata__operating_system__dedicated_host_only` | `import_metadata.operating_system.dedicated_host_only` |
| `ibm_cm_version_invalid_import_metadata__operating_system__display_name` | `import_metadata.operating_system.display_name` |
| `ibm_cm_version_invalid_import_metadata__operating_system__family` | `import_metadata.operating_system.family` |
| `ibm_cm_version_invalid_import_metadata__operating_system__href` | `import_metadata.operating_system.href` |
| `ibm_cm_version_invalid_import_metadata__operating_system__name` | `import_metadata.operating_system.name` |
| `ibm_cm_version_invalid_import_metadata__operating_system__vendor` | `import_metadata.operating_system.vendor` |
| `ibm_cm_version_invalid_import_metadata__operating_system__version` | `import_metadata.operating_system.version` |
| `ibm_cm_version_invalid_install__delete_script` | `install.delete_script` |
| `ibm_cm_version_invalid_install__instructions` | `install.instructions` |
| `ibm_cm_version_invalid_install__instructions_i18n` | `install.instructions_i18n` |
| `ibm_cm_version_invalid_install__scope` | `install.scope` |
| `ibm_cm_version_invalid_install__script` | `install.script` |
| `ibm_cm_version_invalid_install__script_permission` | `install.script_permission` |
| `ibm_cm_version_invalid_licenses__description` | `licenses.description` |
| `ibm_cm_version_invalid_licenses__id` | `licenses.id` |
| `ibm_cm_version_invalid_licenses__name` | `licenses.name` |
| `ibm_cm_version_invalid_licenses__type` | `licenses.type` |
| `ibm_cm_version_invalid_licenses__url` | `licenses.url` |
| `ibm_cm_version_invalid_pre_install__delete_script` | `pre_install.delete_script` |
| `ibm_cm_version_invalid_pre_install__instructions` | `pre_install.instructions` |
| `ibm_cm_version_invalid_pre_install__instructions_i18n` | `pre_install.instructions_i18n` |
| `ibm_cm_version_invalid_pre_install__scope` | `pre_install.scope` |
| `ibm_cm_version_invalid_pre_install__script` | `pre_install.script` |
| `ibm_cm_version_invalid_pre_install__script_permission` | `pre_install.script_permission` |
| `ibm_cm_version_invalid_required_resources__type` | `required_resources.type` |
| `ibm_cm_version_invalid_required_resources__value` | `required_resources.value` |
| `ibm_cm_version_invalid_solution_info__architecture_diagrams__description` | `solution_info.architecture_diagrams.description` |
| `ibm_cm_version_invalid_solution_info__architecture_diagrams__description_i18n` | `solution_info.architecture_diagrams.description_i18n` |
| `ibm_cm_version_invalid_solution_info__architecture_diagrams__diagram__api_url` | `solution_info.architecture_diagrams.diagram.api_url` |
| `ibm_cm_version_invalid_solution_info__architecture_diagrams__diagram__caption` | `solution_info.architecture_diagrams.diagram.caption` |
| `ibm_cm_version_invalid_solution_info__architecture_diagrams__diagram__caption_i18n` | `solution_info.architecture_diagrams.diagram.caption_i18n` |
| `ibm_cm_version_invalid_solution_info__architecture_diagrams__diagram__thumbnail_url` | `solution_info.architecture_diagrams.diagram.thumbnail_url` |
| `ibm_cm_version_invalid_solution_info__architecture_diagrams__diagram__type` | `solution_info.architecture_diagrams.diagram.type` |
| `ibm_cm_version_invalid_solution_info__architecture_diagrams__diagram__url` | `solution_info.architecture_diagrams.diagram.url` |
| `ibm_cm_version_invalid_solution_info__architecture_diagrams__diagram__url_proxy__sha` | `solution_info.architecture_diagrams.diagram.url_proxy.sha` |
| `ibm_cm_version_invalid_solution_info__architecture_diagrams__diagram__url_proxy__url` | `solution_info.architecture_diagrams.diagram.url_proxy.url` |
| `ibm_cm_version_invalid_solution_info__dependencies__catalog_id` | `solution_info.dependencies.catalog_id` |
| `ibm_cm_version_invalid_solution_info__dependencies__flavors` | `solution_info.dependencies.flavors` |
| `ibm_cm_version_invalid_solution_info__dependencies__id` | `solution_info.dependencies.id` |
| `ibm_cm_version_invalid_solution_info__dependencies__name` | `solution_info.dependencies.name` |
| `ibm_cm_version_invalid_solution_info__dependencies__version` | `solution_info.dependencies.version` |
| `ibm_cm_version_invalid_solution_info__features__description` | `solution_info.features.description` |
| `ibm_cm_version_invalid_solution_info__features__description_i18n` | `solution_info.features.description_i18n` |
| `ibm_cm_version_invalid_solution_info__features__title` | `solution_info.features.title` |
| `ibm_cm_version_invalid_solution_info__features__title_i18n` | `solution_info.features.title_i18n` |

## Example

### Valid

```hcl
resource "ibm_cm_version" "test" {
  catalog_id = "01234567-89ab-cdef-0123-456789abcdef"
  offering_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
