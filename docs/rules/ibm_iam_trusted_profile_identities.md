# ibm_iam_trusted_profile_identities

Schema validation for the ibm_iam_trusted_profile_identities resource.

This rule validates `ibm_iam_trusted_profile_identities` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ProfileIdentityResponse` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_trusted_profile_identities" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `if_match` | no | — |
| `profile_id` | yes | — |
| `identities` (block) | no | nested block |
| `identities.accounts` | no | API: Only valid for the type user. Accounts from which a user can assume the trusted profile. |
| `identities.description` | no | API: Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project' |
| `identities.iam_id` | yes | API: IAM ID of the identity |
| `identities.identifier` | yes | API: Identifier of the identity that can assume the trusted profiles. This can be a user identifier (IAM id), serviceid or crn. Internally it uses account id of the service id for the identifier 'serviceid' and for the identifier 'crn' it uses account id contained in the CRN. |
| `identities.type` | yes | one of: `user`, `serviceid`, `crn` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_trusted_profile_identities_invalid_profile_id` | `profile_id` |
| `ibm_iam_trusted_profile_identities_invalid_identities__accounts` | `identities.accounts` |
| `ibm_iam_trusted_profile_identities_invalid_identities__description` | `identities.description` |
| `ibm_iam_trusted_profile_identities_invalid_identities__iam_id` | `identities.iam_id` |
| `ibm_iam_trusted_profile_identities_invalid_identities__identifier` | `identities.identifier` |
| `ibm_iam_trusted_profile_identities_invalid_identities__type` | `identities.type` |

## Example

### Valid

```hcl
resource "ibm_iam_trusted_profile_identities" "test" {
  profile_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
