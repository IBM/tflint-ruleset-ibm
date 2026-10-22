# ibm_project_environment

Schema validation for the ibm_project_environment resource.

This rule validates `ibm_project_environment` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ProjectComplianceProfileV1` (`projects`).

## Configuration

```hcl
rule "ibm_project_environment" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `href` | no | — |
| `modified_at` | no | — |
| `project` | no | — |
| `project_environment_id` | no | — |
| `project_id` | yes | pattern `^[\.\-0-9a-zA-Z]+$` |
| `target_account` | no | — |
| `definition` (block) | yes | nested block |
| `definition.description` | no | — |
| `definition.inputs` | no | — |
| `definition.name` | yes | — |
| `definition.authorizations` (block) | no | nested block |
| `definition.authorizations.api_key` | no | API: The IBM Cloud API Key. It can be either raw or pulled from the catalog via a `CRN` or `JSON` blob. |
| `definition.authorizations.method` | no | API: The authorization method. It can authorize by using a trusted profile or an API key in Secrets Manager. |
| `definition.authorizations.trusted_profile_id` | no | API: The trusted profile ID. |
| `definition.compliance_profile` (block) | no | nested block |
| `definition.compliance_profile.attachment_id` | no | pattern `^[\./:a-zA-Z0-9\.\-]+$` |
| `definition.compliance_profile.id` | no | pattern `^[\./:a-zA-Z0-9\.\-]+$` |
| `definition.compliance_profile.instance_id` | no | pattern `^[\./:a-zA-Z0-9\.\-]+$` |
| `definition.compliance_profile.instance_location` | no | one of: `us-south`, `us-east`, `eu-gb`, `eu-de`, `ca-tor` |
| `definition.compliance_profile.profile_name` | no | API: The name of the compliance profile. |
| `definition.compliance_profile.wp_instance_id` | no | pattern `^[\./:a-zA-Z0-9\.\-]+$` |
| `definition.compliance_profile.wp_instance_location` | no | one of: `us-south`, `us-east`, `eu-gb`, `eu-de`, `ca-tor` |
| `definition.compliance_profile.wp_instance_name` | no | API: The name of the compliance profile. |
| `definition.compliance_profile.wp_policy_id` | no | pattern `^[\./:a-zA-Z0-9\.\-]+$` |
| `definition.compliance_profile.wp_policy_name` | no | API: The name of the compliance profile. |
| `definition.compliance_profile.wp_zone_id` | no | pattern `^[\./:a-zA-Z0-9\.\-]+$` |
| `definition.compliance_profile.wp_zone_name` | no | API: The name of the compliance profile. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_project_environment_invalid_project_id` | `project_id` |
| `ibm_project_environment_invalid_definition__description` | `definition.description` |
| `ibm_project_environment_invalid_definition__inputs` | `definition.inputs` |
| `ibm_project_environment_invalid_definition__name` | `definition.name` |
| `ibm_project_environment_invalid_definition__authorizations__api_key` | `definition.authorizations.api_key` |
| `ibm_project_environment_invalid_definition__authorizations__method` | `definition.authorizations.method` |
| `ibm_project_environment_invalid_definition__authorizations__trusted_profile_id` | `definition.authorizations.trusted_profile_id` |
| `ibm_project_environment_invalid_definition__compliance_profile__attachment_id` | `definition.compliance_profile.attachment_id` |
| `ibm_project_environment_invalid_definition__compliance_profile__id` | `definition.compliance_profile.id` |
| `ibm_project_environment_invalid_definition__compliance_profile__instance_id` | `definition.compliance_profile.instance_id` |
| `ibm_project_environment_invalid_definition__compliance_profile__instance_location` | `definition.compliance_profile.instance_location` |
| `ibm_project_environment_invalid_definition__compliance_profile__profile_name` | `definition.compliance_profile.profile_name` |
| `ibm_project_environment_invalid_definition__compliance_profile__wp_instance_id` | `definition.compliance_profile.wp_instance_id` |
| `ibm_project_environment_invalid_definition__compliance_profile__wp_instance_location` | `definition.compliance_profile.wp_instance_location` |
| `ibm_project_environment_invalid_definition__compliance_profile__wp_instance_name` | `definition.compliance_profile.wp_instance_name` |
| `ibm_project_environment_invalid_definition__compliance_profile__wp_policy_id` | `definition.compliance_profile.wp_policy_id` |
| `ibm_project_environment_invalid_definition__compliance_profile__wp_policy_name` | `definition.compliance_profile.wp_policy_name` |
| `ibm_project_environment_invalid_definition__compliance_profile__wp_zone_id` | `definition.compliance_profile.wp_zone_id` |
| `ibm_project_environment_invalid_definition__compliance_profile__wp_zone_name` | `definition.compliance_profile.wp_zone_name` |

## Example

### Valid

```hcl
resource "ibm_project_environment" "test" {
  project_id = "example-name"
  definition {
    name = "example"
  }
}
```
