# ibm_iam_trusted_profile_link

Manages ibm_iam_trusted_profile_link. Optional name of the Link.

This rule validates `ibm_iam_trusted_profile_link` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CreateProfileLinkRequest` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_trusted_profile_link" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cr_type` | yes | API: The compute resource type. Valid values are VSI, PVS, BMS, IKS_SA, ROKS_SA, CE |
| `created_at` | no | — |
| `entity_tag` | no | — |
| `is_cross_account` | no | API: Flag to indicate that the link provides cross account access. If not provided then the account scope of the CRN must match the Profile's account |
| `link_id` | no | — |
| `modified_at` | no | — |
| `name` | no | API: Optional name of the Link |
| `profile_id` | yes | — |
| `link` (block) | yes | nested block |
| `link.component_name` | no | API: Component name of the compute resource, only required if cr_type is CE |
| `link.component_type` | no | API: Component type of the compute resource, only required if cr_type is CE |
| `link.crn` | no | crn format |
| `link.name` | no | API: Name of the compute resource, only required if cr_type is IKS_SA or ROKS_SA |
| `link.namespace` | no | API: The compute resource namespace, only required if cr_type is IKS_SA or ROKS_SA |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_trusted_profile_link_invalid_cr_type` | `cr_type` |
| `ibm_iam_trusted_profile_link_invalid_is_cross_account` | `is_cross_account` |
| `ibm_iam_trusted_profile_link_invalid_name` | `name` |
| `ibm_iam_trusted_profile_link_invalid_profile_id` | `profile_id` |
| `ibm_iam_trusted_profile_link_invalid_link__component_name` | `link.component_name` |
| `ibm_iam_trusted_profile_link_invalid_link__component_type` | `link.component_type` |
| `ibm_iam_trusted_profile_link_invalid_link__crn` | `link.crn` |
| `ibm_iam_trusted_profile_link_invalid_link__name` | `link.name` |
| `ibm_iam_trusted_profile_link_invalid_link__namespace` | `link.namespace` |

## Example

### Valid

```hcl
resource "ibm_iam_trusted_profile_link" "test" {
  cr_type = "example"
  profile_id = "01234567-89ab-cdef-0123-456789abcdef"
  link {
  }
}
```
