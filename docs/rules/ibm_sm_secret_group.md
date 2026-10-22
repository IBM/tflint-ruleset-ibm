# ibm_sm_secret_group

Manages ibm_sm_secret_group. The name of your secret group.

This rule validates `ibm_sm_secret_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SecretGroupPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_secret_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `description` | no | API: An extended description of your secret group. To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group. |
| `name` | yes | pattern `[A-Za-z0-9][A-Za-z0-9]*(?:_*-*\\.*[A-Za-z0-9]+)*$` |
| `secret_group_id` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_secret_group_invalid_description` | `description` |
| `ibm_sm_secret_group_invalid_name` | `name` |

## Example

### Valid

```hcl
resource "ibm_sm_secret_group" "test" {
  name = "example-name"
}
```
