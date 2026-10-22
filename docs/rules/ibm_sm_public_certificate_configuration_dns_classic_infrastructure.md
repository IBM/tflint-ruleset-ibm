# ibm_sm_public_certificate_configuration_dns_classic_infrastructure

Manages ibm_sm_public_certificate_configuration_dns_classic_infrastructure. A human-readable unique name to assign to your configuration.

This rule validates `ibm_sm_public_certificate_configuration_dns_classic_infrastructure` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PublicCertificateConfigurationDNSClassicInfrastructurePrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_public_certificate_configuration_dns_classic_infrastructure" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `classic_infrastructure_password` | yes | API: Your classic infrastructure API key. For information about viewing and accessing your classic infrastructure API key, see the [docs](https://cloud.ibm.com/docs/account?topic=account-classic_keys). |
| `classic_infrastructure_username` | yes | API: The username that is associated with your classic infrastructure account. In most cases, your classic infrastructure username is your `&lt;account_id&gt;_&lt;email_address&gt;`. For more information, see the [docs](https://cloud.ibm.com/docs/account?topic=account-classic_keys). |
| `config_type` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `name` | yes | pattern `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$` |
| `secret_type` | no | — |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sm_public_certificate_configuration_dns_classic_infrastructure_invalid_classic_infrastructure_password` | `classic_infrastructure_password` |
| `ibm_sm_public_certificate_configuration_dns_classic_infrastructure_invalid_classic_infrastructure_username` | `classic_infrastructure_username` |
| `ibm_sm_public_certificate_configuration_dns_classic_infrastructure_invalid_name` | `name` |

## Example

### Valid

```hcl
variable "classic_infrastructure_password" {
  sensitive = true
}

resource "ibm_sm_public_certificate_configuration_dns_classic_infrastructure" "test" {
  classic_infrastructure_password = var.classic_infrastructure_password
  classic_infrastructure_username = "aaaaaaa"
  name = "example-name"
}
```
