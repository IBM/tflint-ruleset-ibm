# ibm_sm_public_certificate_action_validate_manual_dns

Schema validation for the ibm_sm_public_certificate_action_validate_manual_dns resource.

This rule validates `ibm_sm_public_certificate_action_validate_manual_dns` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SecretLocks` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_public_certificate_action_validate_manual_dns" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `secret_id` | yes | pattern `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_public_certificate_action_validate_manual_dns_invalid_secret_id` | `secret_id` |

## Example

### Valid

```hcl
variable "secret_id" {
  sensitive = true
}

resource "ibm_sm_public_certificate_action_validate_manual_dns" "test" {
  secret_id = var.secret_id
}
```
