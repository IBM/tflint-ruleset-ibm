# ibm_hpcs_keystore_data

Reads ibm_hpcs_keystore. Name of the target keystore.

This rule validates `ibm_hpcs_keystore` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `keystore` (`uko`).

## Configuration

```hcl
rule "ibm_hpcs_keystore_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `aws_access_key_id` | no | — |
| `aws_region` | no | — |
| `aws_secret_access_key` | no | — |
| `azure_environment` | no | — |
| `azure_location` | no | — |
| `azure_resource_group` | no | — |
| `azure_service_name` | no | — |
| `azure_service_principal_client_id` | no | — |
| `azure_service_principal_password` | no | — |
| `azure_subscription_id` | no | — |
| `azure_tenant` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `description` | no | — |
| `google_credentials` | no | — |
| `google_key_ring` | no | — |
| `google_location` | no | — |
| `google_private_key_id` | no | — |
| `google_project_id` | no | — |
| `groups` | no | — |
| `href` | no | — |
| `ibm_api_endpoint` | no | — |
| `ibm_api_key` | no | — |
| `ibm_iam_endpoint` | no | — |
| `ibm_instance_id` | no | — |
| `ibm_key_ring` | no | — |
| `ibm_variant` | no | — |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `keystore_id` | yes | — |
| `location` | no | — |
| `name` | no | — |
| `region` | yes | one of: `au-syd`, `in-che`, `jp-osa`, `jp-tok`, `kr-seo`, `eu-de`, `eu-gb`, `ca-tor`, `us-south`, `us-south-test`, `us-east`, `br-sao` |
| `type` | no | — |
| `uko_vault` | yes | — |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `vault` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_hpcs_keystore_data_invalid_instance_id` | `instance_id` |
| `ibm_hpcs_keystore_data_invalid_keystore_id` | `keystore_id` |
| `ibm_hpcs_keystore_data_invalid_region` | `region` |
| `ibm_hpcs_keystore_data_invalid_uko_vault` | `uko_vault` |

## Example

### Valid

```hcl
data "ibm_hpcs_keystore" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  keystore_id = "01234567-89ab-cdef-0123-456789abcdef"
  region = "au-syd"
  uko_vault = "example"
}
```
