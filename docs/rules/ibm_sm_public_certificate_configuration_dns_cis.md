# ibm_sm_public_certificate_configuration_dns_cis

Manages ibm_sm_public_certificate_configuration_dns_cis. A human-readable unique name to assign to your configuration.

This rule validates `ibm_sm_public_certificate_configuration_dns_cis` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PublicCertificateConfigurationDNSCloudInternetServicesPrototype` (`secrets-manager-secrets-manager-v2`).

## Configuration

```hcl
rule "ibm_sm_public_certificate_configuration_dns_cis" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cloud_internet_services_apikey` | no | API: An IBM Cloud API key that can list domains in your Cloud Internet Services instance and add DNS records. To grant Secrets Manager the ability to view the Cloud Internet Services instance and all of its domains, the API key must be assigned the Reader service role on Internet Services (`internet-svcs`). In order to add DNS records you need to assign the Manager role. If you want to manage specific  |
| `cloud_internet_services_crn` | yes | pattern `^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$` |
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
| `ibm_sm_public_certificate_configuration_dns_cis_invalid_cloud_internet_services_apikey` | `cloud_internet_services_apikey` |
| `ibm_sm_public_certificate_configuration_dns_cis_invalid_cloud_internet_services_crn` | `cloud_internet_services_crn` |
| `ibm_sm_public_certificate_configuration_dns_cis_invalid_name` | `name` |

## Example

### Valid

```hcl
resource "ibm_sm_public_certificate_configuration_dns_cis" "test" {
  cloud_internet_services_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  name = "example-name"
}
```
