# ibm_iam_trusted_profile_identity

Schema validation for the ibm_iam_trusted_profile_identity resource.

This rule validates `ibm_iam_trusted_profile_identity` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ProfileIdentityRequest` (`iam-identity-token-api`).

## Configuration

```hcl
rule "ibm_iam_trusted_profile_identity" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accounts` | no | API: Only valid for the type user. Accounts from which a user can assume the trusted profile. |
| `description` | no | API: Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project' |
| `identifier` | yes | API: Identifier of the identity that can assume the trusted profiles. This can be a user identifier (IAM id), serviceid or crn. Internally it uses account id of the service id for the identifier 'serviceid' and for the identifier 'crn' it uses account id contained in the CRN. |
| `identity_type` | yes | one of: `crn`, `serviceid`, `user` |
| `profile_id` | yes | — |
| `type` | yes | one of: `crn`, `serviceid`, `user` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_trusted_profile_identity_invalid_accounts` | `accounts` |
| `ibm_iam_trusted_profile_identity_invalid_description` | `description` |
| `ibm_iam_trusted_profile_identity_invalid_identifier` | `identifier` |
| `ibm_iam_trusted_profile_identity_invalid_identity_type` | `identity_type` |
| `ibm_iam_trusted_profile_identity_invalid_profile_id` | `profile_id` |
| `ibm_iam_trusted_profile_identity_invalid_type` | `type` |

## Example

### Valid

```hcl
resource "ibm_iam_trusted_profile_identity" "test" {
  identifier = "example"
  identity_type = "crn"
  profile_id = "01234567-89ab-cdef-0123-456789abcdef"
  type = "crn"
}
```
