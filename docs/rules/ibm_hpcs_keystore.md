# ibm_hpcs_keystore

Manages ibm_hpcs_keystore. Name of the target keystore.

This rule validates `ibm_hpcs_keystore` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `keystore` (`uko`).

## Configuration

```hcl
rule "ibm_hpcs_keystore" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `aws_access_key_id` | no | pattern `^[A-Za-z0-9-_]*$` |
| `aws_region` | no | one of: `af_south_1`, `ap_east_1`, `ap_northeast_1`, `ap_northeast_2`, `ap_south_1`, `ap_southeast_1`, `ap_southeast_2`, `aws_cn_global`, `aws_global`, `aws_iso_global`, `aws_iso_b_global`, `aws_us_gov_global`, `ca_central_1`, `cn_north_1`, `cn_northwest_1`, `eu-central-1`, `eu_west_1`, `eu_west_2`, `eu_west_3`, `me_south_1`, `sa_east_1`, `us_east_1`, `us_east_2`, `us_gov_east_1`, `us_gov_west_1`, `us_iso_east_1`, `us_isob_east_1`, `us_west_1`, `us_west_2` |
| `aws_secret_access_key` | no | API: The secret access key used for connecting to this instance of AWS KMS. |
| `azure_environment` | no | one of: `azure`, `azure_china`, `azure_germany`, `azure_us_government` |
| `azure_location` | no | one of: `asia_east`, `asia_southeast`, `australia_central`, `australia_central_2`, `australia_east`, `australia_southeast`, `brazil_south`, `canada_central`, `canada_east`, `china_east`, `china_east_2`, `china_north`, `china_north_2`, `europe_north`, `europe_west`, `france_central`, `france_south`, `germany_central`, `germany_northeast`, `india_central`, `india_south`, `india_west`, `japan_east`, `japan_west`, `korea_central`, `korea_south`, `south_africa_north`, `south_africa_west`, `uk_south`, `uk_west`, `us_central`, `us_dod_central`, `us_dod_east`, `us_east`, `us_east_2`, `us_gov_arizona`, `us_gov_iowa`, `us_gov_texas`, `us_gov_virginia`, `us_north_central`, `us_south_central`, `us_west`, `us_west_2`, `us_west_central` |
| `azure_resource_group` | no | pattern `^[-\w\._\(\)]*[^\.]$` |
| `azure_service_name` | no | pattern `^[A-Za-z0-9-]+$` |
| `azure_service_principal_client_id` | no | pattern `^[-0-9a-zA-Z]+$` |
| `azure_service_principal_password` | no | API: Azure service principal password |
| `azure_subscription_id` | no | pattern `^[-0-9a-zA-Z]+$` |
| `azure_tenant` | no | pattern `^[-0-9a-zA-Z]+$` |
| `created_at` | no | — |
| `created_by` | no | — |
| `description` | no | API: Description of the keystore |
| `dry_run` | no | — |
| `etag` | no | — |
| `google_credentials` | no | pattern `^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$` |
| `google_key_ring` | no | pattern `^[A-Za-z0-9-]*$` |
| `google_location` | no | pattern `^[A-Za-z0-9-]*$` |
| `google_private_key_id` | no | pattern `^[A-Za-z0-9-_]*$` |
| `google_project_id` | no | pattern `^[A-Za-z0-9-_]*$` |
| `groups` | no | pattern `^[A-Za-z0-9][A-Za-z0-9-_ ]+$` |
| `href` | no | — |
| `ibm_api_endpoint` | no | pattern `^[A-Za-z0-9._~:/-]+$` |
| `ibm_api_key` | no | pattern `^[A-Za-z0-9-_&.]*$` |
| `ibm_iam_endpoint` | no | pattern `^[A-Za-z0-9._~:/-]+$` |
| `ibm_instance_id` | no | pattern `^[A-Za-z0-9-]*$` |
| `ibm_key_ring` | no | pattern `^[a-zA-Z0-9-]*$` |
| `ibm_variant` | no | one of: `hpcs`, `internal`, `key_protect` |
| `instance_id` | yes | API: The v4 UUID used to uniquely identify the instance in IBM Cloud. |
| `location` | no | — |
| `name` | no | pattern `^[A-Za-z0-9][A-Za-z0-9 ._-]*$` |
| `region` | yes | one of: `au-syd`, `in-che`, `jp-osa`, `jp-tok`, `kr-seo`, `eu-de`, `eu-gb`, `ca-tor`, `us-south`, `us-south-test`, `us-east`, `br-sao` |
| `type` | yes | one of: `aws_kms`, `azure_key_vault`, `google_kms`, `ibm_cloud_kms`, `cca` |
| `uko_vault` | yes | — |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `vault` (block) | yes | nested block |
| `vault.id` | yes | pattern `^[-0-9a-z]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_hpcs_keystore_invalid_aws_access_key_id` | `aws_access_key_id` |
| `ibm_hpcs_keystore_invalid_aws_region` | `aws_region` |
| `ibm_hpcs_keystore_invalid_aws_secret_access_key` | `aws_secret_access_key` |
| `ibm_hpcs_keystore_invalid_azure_environment` | `azure_environment` |
| `ibm_hpcs_keystore_invalid_azure_location` | `azure_location` |
| `ibm_hpcs_keystore_invalid_azure_resource_group` | `azure_resource_group` |
| `ibm_hpcs_keystore_invalid_azure_service_name` | `azure_service_name` |
| `ibm_hpcs_keystore_invalid_azure_service_principal_client_id` | `azure_service_principal_client_id` |
| `ibm_hpcs_keystore_invalid_azure_service_principal_password` | `azure_service_principal_password` |
| `ibm_hpcs_keystore_invalid_azure_subscription_id` | `azure_subscription_id` |
| `ibm_hpcs_keystore_invalid_azure_tenant` | `azure_tenant` |
| `ibm_hpcs_keystore_invalid_description` | `description` |
| `ibm_hpcs_keystore_invalid_dry_run` | `dry_run` |
| `ibm_hpcs_keystore_invalid_google_credentials` | `google_credentials` |
| `ibm_hpcs_keystore_invalid_google_key_ring` | `google_key_ring` |
| `ibm_hpcs_keystore_invalid_google_location` | `google_location` |
| `ibm_hpcs_keystore_invalid_google_private_key_id` | `google_private_key_id` |
| `ibm_hpcs_keystore_invalid_google_project_id` | `google_project_id` |
| `ibm_hpcs_keystore_invalid_groups` | `groups` |
| `ibm_hpcs_keystore_invalid_ibm_api_endpoint` | `ibm_api_endpoint` |
| `ibm_hpcs_keystore_invalid_ibm_api_key` | `ibm_api_key` |
| `ibm_hpcs_keystore_invalid_ibm_iam_endpoint` | `ibm_iam_endpoint` |
| `ibm_hpcs_keystore_invalid_ibm_instance_id` | `ibm_instance_id` |
| `ibm_hpcs_keystore_invalid_ibm_key_ring` | `ibm_key_ring` |
| `ibm_hpcs_keystore_invalid_ibm_variant` | `ibm_variant` |
| `ibm_hpcs_keystore_invalid_instance_id` | `instance_id` |
| `ibm_hpcs_keystore_invalid_name` | `name` |
| `ibm_hpcs_keystore_invalid_region` | `region` |
| `ibm_hpcs_keystore_invalid_type` | `type` |
| `ibm_hpcs_keystore_invalid_uko_vault` | `uko_vault` |
| `ibm_hpcs_keystore_invalid_vault__id` | `vault.id` |

## Example

### Valid

```hcl
resource "ibm_hpcs_keystore" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  region = "au-syd"
  type = "aws_kms"
  uko_vault = "example"
  vault {
    id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  }
}
```
