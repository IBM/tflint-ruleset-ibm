# ibm_cm_version_data

Schema validation for the ibm_cm_version data source.

This rule validates `ibm_cm_version` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cm_version_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `catalog_id` | no | — |
| `configuration` | no | — |
| `created` | no | — |
| `crn` | no | — |
| `deprecate_pending` | no | — |
| `deprecated` | no | — |
| `entitlement` | no | — |
| `flavor` | no | — |
| `iam_permissions` | no | — |
| `image_manifest_url` | no | — |
| `image_pull_key_name` | no | — |
| `install` | no | — |
| `is_consumable` | no | — |
| `kind_id` | no | — |
| `licenses` | no | — |
| `long_description` | no | — |
| `long_description_i18n` | no | — |
| `metadata` | no | — |
| `offering_id` | no | — |
| `outputs` | no | — |
| `package_version` | no | — |
| `pre_install` | no | — |
| `repo_url` | no | — |
| `required_resources` | no | — |
| `rev` | no | — |
| `sha` | no | — |
| `single_instance` | no | — |
| `solution_info` | no | — |
| `source_url` | no | — |
| `state` | no | — |
| `tags` | no | — |
| `tgz_url` | no | — |
| `updated` | no | — |
| `validation` | no | — |
| `version` | no | — |
| `version_id` | no | — |
| `version_loc_id` | yes | — |
| `version_locator` | no | — |
| `whitelisted_accounts` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cm_version_data_invalid_version_loc_id` | `version_loc_id` |

## Example

### Valid

```hcl
data "ibm_cm_version" "test" {
  version_loc_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
