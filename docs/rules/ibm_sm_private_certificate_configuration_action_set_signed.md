# ibm_sm_private_certificate_configuration_action_set_signed

Manages ibm_sm_private_certificate_configuration_action_set_signed. The name that uniquely identifies a configuration

This rule validates `ibm_sm_private_certificate_configuration_action_set_signed` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SecretGroupPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_private_certificate_configuration_action_set_signed" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `certificate` | yes | — |
| `name` | yes | pattern `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_private_certificate_configuration_action_set_signed_invalid_certificate` | `certificate` |
| `ibm_sm_private_certificate_configuration_action_set_signed_invalid_name` | `name` |

## Example

### Valid

```hcl
variable "certificate" {
  sensitive = true
}

resource "ibm_sm_private_certificate_configuration_action_set_signed" "test" {
  certificate = var.certificate
  name = "example-name"
}
```
