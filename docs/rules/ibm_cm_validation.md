# ibm_cm_validation

Schema validation for the ibm_cm_validation resource.

This rule validates `ibm_cm_validation` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Validation` (`resource-catalog-private-catalog`).

## Configuration

```hcl
rule "ibm_cm_validation" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `last_operation` | no | — |
| `mark_version_consumable` | no | — |
| `message` | no | — |
| `override_values` | no | — |
| `region` | no | — |
| `requested` | no | — |
| `revalidate_if_validated` | no | — |
| `state` | no | — |
| `target` | no | — |
| `validated` | no | — |
| `version_locator` | yes | API: A dotted value of `catalogID`.`versionID` |
| `x_auth_refresh_token` | no | — |
| `environment_variables` (block) | no | nested block |
| `environment_variables.name` | no | — |
| `environment_variables.secure` | no | — |
| `environment_variables.value` | no | — |
| `schematics` (block) | no | nested block |
| `schematics.description` | no | — |
| `schematics.name` | no | — |
| `schematics.region` | no | — |
| `schematics.resource_group_id` | no | — |
| `schematics.tags` | no | — |
| `schematics.terraform_version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cm_validation_invalid_mark_version_consumable` | `mark_version_consumable` |
| `ibm_cm_validation_invalid_override_values` | `override_values` |
| `ibm_cm_validation_invalid_region` | `region` |
| `ibm_cm_validation_invalid_revalidate_if_validated` | `revalidate_if_validated` |
| `ibm_cm_validation_invalid_version_locator` | `version_locator` |
| `ibm_cm_validation_invalid_x_auth_refresh_token` | `x_auth_refresh_token` |
| `ibm_cm_validation_invalid_environment_variables__name` | `environment_variables.name` |
| `ibm_cm_validation_invalid_environment_variables__secure` | `environment_variables.secure` |
| `ibm_cm_validation_invalid_environment_variables__value` | `environment_variables.value` |
| `ibm_cm_validation_invalid_schematics__description` | `schematics.description` |
| `ibm_cm_validation_invalid_schematics__name` | `schematics.name` |
| `ibm_cm_validation_invalid_schematics__region` | `schematics.region` |
| `ibm_cm_validation_invalid_schematics__resource_group_id` | `schematics.resource_group_id` |
| `ibm_cm_validation_invalid_schematics__tags` | `schematics.tags` |
| `ibm_cm_validation_invalid_schematics__terraform_version` | `schematics.terraform_version` |

## Example

### Valid

```hcl
resource "ibm_cm_validation" "test" {
  version_locator = "example"
}
```
