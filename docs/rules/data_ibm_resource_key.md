# ibm_resource_key_data

Reads ibm_resource_key. The name of the resource key

This rule validates `ibm_resource_key` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/resource_keys/{id}:patch` (`resource-controller-resource-controller`).

## Configuration

```hcl
rule "ibm_resource_key_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `credentials` | no | — |
| `credentials_json` | no | — |
| `crn` | no | — |
| `most_recent` | no | — |
| `name` | yes | pattern `^([\p{L}\p{N}._: -]{1,256})$` |
| `onetime_credentials` | no | — |
| `resource_alias_id` | no | — |
| `resource_instance_id` | no | — |
| `role` | no | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_resource_key_data_invalid_most_recent` | `most_recent` |
| `ibm_resource_key_data_invalid_name` | `name` |
| `ibm_resource_key_data_invalid_resource_alias_id` | `resource_alias_id` |
| `ibm_resource_key_data_invalid_resource_instance_id` | `resource_instance_id` |

## Example

### Valid

```hcl
data "ibm_resource_key" "test" {
  name = "example-name"
}
```
