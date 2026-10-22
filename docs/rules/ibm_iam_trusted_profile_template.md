# ibm_iam_trusted_profile_template

Manages ibm_iam_trusted_profile_template. The name of the trusted profile template.

This rule validates `ibm_iam_trusted_profile_template` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ProfileIdentityRequest` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_trusted_profile_template" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `committed` | no | — |
| `created_at` | no | — |
| `created_by_id` | no | — |
| `crn` | no | — |
| `description` | no | API: Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project' |
| `entity_tag` | no | — |
| `id` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `name` | no | — |
| `template_id` | no | — |
| `version` | no | — |
| `policy_template_references` (block) | no | nested block |
| `policy_template_references.id` | yes | — |
| `policy_template_references.version` | yes | API: Version of the the template |
| `profile` (block) | no | nested block |
| `profile.description` | no | API: Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project' |
| `profile.name` | yes | — |
| `profile.identities` (block) | no | nested block |
| `profile.identities.accounts` | no | API: Only valid for the type user. Accounts from which a user can assume the trusted profile. |
| `profile.identities.description` | no | API: Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project' |
| `profile.identities.iam_id` | yes | — |
| `profile.identities.identifier` | yes | API: Identifier of the identity that can assume the trusted profiles. This can be a user identifier (IAM id), serviceid or crn. Internally it uses account id of the service id for the identifier 'serviceid' and for the identifier 'crn' it uses account id contained in the CRN. |
| `profile.identities.type` | yes | one of: `user`, `serviceid`, `crn` |
| `profile.rules` (block) | no | nested block |
| `profile.rules.expiration` | no | API: Number of seconds counted since January 1st, 1970, until the IAM access token will expire. |
| `profile.rules.name` | no | — |
| `profile.rules.realm_name` | no | — |
| `profile.rules.type` | yes | one of: `user`, `serviceid`, `crn` |
| `profile.rules.conditions` (block) | yes | nested block |
| `profile.rules.conditions.claim` | yes | — |
| `profile.rules.conditions.operator` | yes | — |
| `profile.rules.conditions.value` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_trusted_profile_template_invalid_account_id` | `account_id` |
| `ibm_iam_trusted_profile_template_invalid_committed` | `committed` |
| `ibm_iam_trusted_profile_template_invalid_description` | `description` |
| `ibm_iam_trusted_profile_template_invalid_name` | `name` |
| `ibm_iam_trusted_profile_template_invalid_template_id` | `template_id` |
| `ibm_iam_trusted_profile_template_invalid_policy_template_references__id` | `policy_template_references.id` |
| `ibm_iam_trusted_profile_template_invalid_policy_template_references__version` | `policy_template_references.version` |
| `ibm_iam_trusted_profile_template_invalid_profile__description` | `profile.description` |
| `ibm_iam_trusted_profile_template_invalid_profile__name` | `profile.name` |
| `ibm_iam_trusted_profile_template_invalid_profile__identities__accounts` | `profile.identities.accounts` |
| `ibm_iam_trusted_profile_template_invalid_profile__identities__description` | `profile.identities.description` |
| `ibm_iam_trusted_profile_template_invalid_profile__identities__iam_id` | `profile.identities.iam_id` |
| `ibm_iam_trusted_profile_template_invalid_profile__identities__identifier` | `profile.identities.identifier` |
| `ibm_iam_trusted_profile_template_invalid_profile__identities__type` | `profile.identities.type` |
| `ibm_iam_trusted_profile_template_invalid_profile__rules__expiration` | `profile.rules.expiration` |
| `ibm_iam_trusted_profile_template_invalid_profile__rules__name` | `profile.rules.name` |
| `ibm_iam_trusted_profile_template_invalid_profile__rules__realm_name` | `profile.rules.realm_name` |
| `ibm_iam_trusted_profile_template_invalid_profile__rules__type` | `profile.rules.type` |
| `ibm_iam_trusted_profile_template_invalid_profile__rules__conditions__claim` | `profile.rules.conditions.claim` |
| `ibm_iam_trusted_profile_template_invalid_profile__rules__conditions__operator` | `profile.rules.conditions.operator` |
| `ibm_iam_trusted_profile_template_invalid_profile__rules__conditions__value` | `profile.rules.conditions.value` |

## Example

### Valid

```hcl
resource "ibm_iam_trusted_profile_template" "test" {
}
```
