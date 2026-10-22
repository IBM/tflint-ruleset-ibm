# ibm_cd_toolchain_tool_cos

Manages ibm_cd_toolchain_tool_cos. Name of the tool.

This rule validates `ibm_cd_toolchain_tool_cos` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Toolchain` (`toolchain`).

## Configuration

```hcl
rule "ibm_cd_toolchain_tool_cos" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `href` | no | — |
| `name` | no | pattern `^([^\x00-\x7F]|[a-zA-Z0-9-._ ])+$` |
| `referent` | no | — |
| `resource_group_id` | no | — |
| `state` | no | — |
| `tool_id` | no | — |
| `toolchain_crn` | no | — |
| `toolchain_id` | yes | pattern `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$` |
| `updated_at` | no | — |
| `parameters` (block) | yes | nested block |
| `parameters.auth_type` | no | — |
| `parameters.bucket_name` | no | — |
| `parameters.cos_api_key` | no | — |
| `parameters.endpoint` | no | — |
| `parameters.hmac_access_key_id` | no | — |
| `parameters.hmac_secret_access_key` | no | — |
| `parameters.instance_crn` | no | crn format |
| `parameters.name` | yes | pattern `^([^\x00-\x7F]|[a-zA-Z0-9-._ ])+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cd_toolchain_tool_cos_invalid_name` | `name` |
| `ibm_cd_toolchain_tool_cos_invalid_toolchain_id` | `toolchain_id` |
| `ibm_cd_toolchain_tool_cos_invalid_parameters__auth_type` | `parameters.auth_type` |
| `ibm_cd_toolchain_tool_cos_invalid_parameters__bucket_name` | `parameters.bucket_name` |
| `ibm_cd_toolchain_tool_cos_invalid_parameters__cos_api_key` | `parameters.cos_api_key` |
| `ibm_cd_toolchain_tool_cos_invalid_parameters__endpoint` | `parameters.endpoint` |
| `ibm_cd_toolchain_tool_cos_invalid_parameters__hmac_access_key_id` | `parameters.hmac_access_key_id` |
| `ibm_cd_toolchain_tool_cos_invalid_parameters__hmac_secret_access_key` | `parameters.hmac_secret_access_key` |
| `ibm_cd_toolchain_tool_cos_invalid_parameters__instance_crn` | `parameters.instance_crn` |
| `ibm_cd_toolchain_tool_cos_invalid_parameters__name` | `parameters.name` |

## Example

### Valid

```hcl
resource "ibm_cd_toolchain_tool_cos" "test" {
  toolchain_id = "01234567-89ab-4cde-8123-456789abcdef"
  parameters {
    name = "example-name"
  }
}
```
