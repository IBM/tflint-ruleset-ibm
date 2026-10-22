# ibm_cm_account

Schema validation for the ibm_cm_account resource.

This rule validates `ibm_cm_account` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `TerraformEngines` (`resource-catalog-private-catalog`).

## Configuration

```hcl
rule "ibm_cm_account" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `hide_ibm_cloud_catalog` | no | — |
| `region_filter` | no | — |
| `rev` | no | — |
| `account_filters` (block) | no | nested block |
| `account_filters.include_all` | no | — |
| `account_filters.category_filters` (block) | no | nested block |
| `account_filters.category_filters.category_name` | yes | — |
| `account_filters.category_filters.include` | no | — |
| `account_filters.category_filters.filter` (block) | no | nested block |
| `account_filters.category_filters.filter.filter_terms` | no | — |
| `account_filters.id_filters` (block) | no | nested block |
| `account_filters.id_filters.exclude` (block) | no | nested block |
| `account_filters.id_filters.exclude.filter_terms` | no | — |
| `account_filters.id_filters.include` (block) | no | nested block |
| `account_filters.id_filters.include.filter_terms` | no | — |
| `terraform_engines` (block) | no | nested block |
| `terraform_engines.api_token` | no | API: The api key used to access the engine instance |
| `terraform_engines.name` | no | API: User provided name for the specified engine |
| `terraform_engines.private_endpoint` | no | API: The private endpoint for the engine instance |
| `terraform_engines.public_endpoint` | no | API: The public endpoint for the engine instance |
| `terraform_engines.type` | no | API: The terraform engine type. The only one supported at the moment is terraform-enterprise |
| `terraform_engines.da_creation` (block) | no | nested block |
| `terraform_engines.da_creation.default_private_catalog_id` | no | API: Default private catalog to create the deployable architectures in |
| `terraform_engines.da_creation.enabled` | no | API: Determines whether deployable architectures are auto-created from workspaces in the engine |
| `terraform_engines.da_creation.polling_info` (block) | no | nested block |
| `terraform_engines.da_creation.polling_info.scopes` (block) | no | nested block |
| `terraform_engines.da_creation.polling_info.scopes.name` | no | API: User provided name for the specified engine |
| `terraform_engines.da_creation.polling_info.scopes.type` | no | API: The terraform engine type. The only one supported at the moment is terraform-enterprise |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cm_account_invalid_hide_ibm_cloud_catalog` | `hide_ibm_cloud_catalog` |
| `ibm_cm_account_invalid_region_filter` | `region_filter` |
| `ibm_cm_account_invalid_account_filters__include_all` | `account_filters.include_all` |
| `ibm_cm_account_invalid_account_filters__category_filters__category_name` | `account_filters.category_filters.category_name` |
| `ibm_cm_account_invalid_account_filters__category_filters__include` | `account_filters.category_filters.include` |
| `ibm_cm_account_invalid_account_filters__category_filters__filter__filter_terms` | `account_filters.category_filters.filter.filter_terms` |
| `ibm_cm_account_invalid_account_filters__id_filters__exclude__filter_terms` | `account_filters.id_filters.exclude.filter_terms` |
| `ibm_cm_account_invalid_account_filters__id_filters__include__filter_terms` | `account_filters.id_filters.include.filter_terms` |
| `ibm_cm_account_invalid_terraform_engines__api_token` | `terraform_engines.api_token` |
| `ibm_cm_account_invalid_terraform_engines__name` | `terraform_engines.name` |
| `ibm_cm_account_invalid_terraform_engines__private_endpoint` | `terraform_engines.private_endpoint` |
| `ibm_cm_account_invalid_terraform_engines__public_endpoint` | `terraform_engines.public_endpoint` |
| `ibm_cm_account_invalid_terraform_engines__type` | `terraform_engines.type` |
| `ibm_cm_account_invalid_terraform_engines__da_creation__default_private_catalog_id` | `terraform_engines.da_creation.default_private_catalog_id` |
| `ibm_cm_account_invalid_terraform_engines__da_creation__enabled` | `terraform_engines.da_creation.enabled` |
| `ibm_cm_account_invalid_terraform_engines__da_creation__polling_info__scopes__name` | `terraform_engines.da_creation.polling_info.scopes.name` |
| `ibm_cm_account_invalid_terraform_engines__da_creation__polling_info__scopes__type` | `terraform_engines.da_creation.polling_info.scopes.type` |

## Example

### Valid

```hcl
resource "ibm_cm_account" "test" {
}
```
