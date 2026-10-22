# ibm_iam_user_invite

Schema validation for the ibm_iam_user_invite resource.

This rule validates `ibm_iam_user_invite` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_iam_user_invite" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `access_groups` | no | — |
| `invited_users` | no | — |
| `number_of_invited_users` | no | — |
| `users` | yes | — |
| `classic_infra_roles` (block) | no | nested block |
| `classic_infra_roles.permission_set` | no | one of: `noacess`, `viewonly`, `basicuser`, `superuser` |
| `classic_infra_roles.permissions` | no | — |
| `cloud_foundry_roles` (block) | no | nested block |
| `cloud_foundry_roles.org_roles` | yes | — |
| `cloud_foundry_roles.organization_guid` | yes | — |
| `cloud_foundry_roles.spaces` (block) | no | nested block |
| `cloud_foundry_roles.spaces.space_guid` | yes | — |
| `cloud_foundry_roles.spaces.space_roles` | yes | — |
| `iam_policy` (block) | no | nested block |
| `iam_policy.account_management` | no | — |
| `iam_policy.roles` | yes | — |
| `iam_policy.resources` (block) | no | nested block |
| `iam_policy.resources.attributes` | no | — |
| `iam_policy.resources.region` | no | — |
| `iam_policy.resources.resource` | no | — |
| `iam_policy.resources.resource_group_id` | no | — |
| `iam_policy.resources.resource_instance_id` | no | — |
| `iam_policy.resources.resource_type` | no | — |
| `iam_policy.resources.service` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_user_invite_invalid_access_groups` | `access_groups` |
| `ibm_iam_user_invite_invalid_users` | `users` |
| `ibm_iam_user_invite_invalid_classic_infra_roles__permission_set` | `classic_infra_roles.permission_set` |
| `ibm_iam_user_invite_invalid_classic_infra_roles__permissions` | `classic_infra_roles.permissions` |
| `ibm_iam_user_invite_invalid_cloud_foundry_roles__org_roles` | `cloud_foundry_roles.org_roles` |
| `ibm_iam_user_invite_invalid_cloud_foundry_roles__organization_guid` | `cloud_foundry_roles.organization_guid` |
| `ibm_iam_user_invite_invalid_cloud_foundry_roles__spaces__space_guid` | `cloud_foundry_roles.spaces.space_guid` |
| `ibm_iam_user_invite_invalid_cloud_foundry_roles__spaces__space_roles` | `cloud_foundry_roles.spaces.space_roles` |
| `ibm_iam_user_invite_invalid_iam_policy__account_management` | `iam_policy.account_management` |
| `ibm_iam_user_invite_invalid_iam_policy__roles` | `iam_policy.roles` |
| `ibm_iam_user_invite_invalid_iam_policy__resources__attributes` | `iam_policy.resources.attributes` |
| `ibm_iam_user_invite_invalid_iam_policy__resources__region` | `iam_policy.resources.region` |
| `ibm_iam_user_invite_invalid_iam_policy__resources__resource` | `iam_policy.resources.resource` |
| `ibm_iam_user_invite_invalid_iam_policy__resources__resource_group_id` | `iam_policy.resources.resource_group_id` |
| `ibm_iam_user_invite_invalid_iam_policy__resources__resource_instance_id` | `iam_policy.resources.resource_instance_id` |
| `ibm_iam_user_invite_invalid_iam_policy__resources__resource_type` | `iam_policy.resources.resource_type` |
| `ibm_iam_user_invite_invalid_iam_policy__resources__service` | `iam_policy.resources.service` |

## Example

### Valid

```hcl
resource "ibm_iam_user_invite" "test" {
  users = "example"
}
```
