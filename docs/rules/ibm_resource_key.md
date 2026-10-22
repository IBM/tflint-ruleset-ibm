# ibm_resource_key

Manages ibm_resource_key. The name of the resource key

This rule validates `ibm_resource_key` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/resource_keys:post` (`resource-controller-resource-controller`).

## Configuration

```hcl
rule "ibm_resource_key" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `credentials` | no | — |
| `credentials_json` | no | — |
| `crn` | no | — |
| `deleted_at` | no | — |
| `deleted_by` | no | — |
| `guid` | no | — |
| `iam_compatible` | no | — |
| `name` | yes | pattern `^([\p{L}\p{N}._: -]{1,256})$` |
| `onetime_credentials` | no | — |
| `parameters` | no | API: Configuration options represented as key-value pairs. Service defined options are passed through to the target resource brokers, whereas platform defined options are not. |
| `resource_alias_id` | no | — |
| `resource_group_id` | no | — |
| `resource_instance_id` | no | — |
| `resource_instance_url` | no | — |
| `role` | no | API: The base IAM service role name (Reader, Writer, or Manager), or the service or custom role CRN. Refer to service’s documentation for supported roles. |
| `source_crn` | no | — |
| `state` | no | — |
| `status` | no | — |
| `tags` | no | — |
| `updated_at` | no | — |
| `updated_by` | no | — |
| `url` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_resource_key_invalid_name` | `name` |
| `ibm_resource_key_invalid_parameters` | `parameters` |
| `ibm_resource_key_invalid_resource_alias_id` | `resource_alias_id` |
| `ibm_resource_key_invalid_resource_instance_id` | `resource_instance_id` |
| `ibm_resource_key_invalid_role` | `role` |
| `ibm_resource_key_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_resource_key" "test" {
  name = "example-name"
}
```
