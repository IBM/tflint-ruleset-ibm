# ibm_resource_reclamation_delete

Schema validation for the ibm_resource_reclamation_delete resource.

This rule validates `ibm_resource_reclamation_delete` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/reclamations/{id}/actions/{action_name}:post` (`resource-controller-resource-controller`).

## Configuration

```hcl
rule "ibm_resource_reclamation_delete" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `comment` | no | API: A comment to describe the action. |
| `created_at` | no | — |
| `created_by` | no | — |
| `custom_properties` | no | — |
| `entity_crn` | no | — |
| `entity_id` | no | — |
| `entity_type_id` | no | — |
| `policy_id` | no | — |
| `reclamation_id` | yes | — |
| `request_by` | no | API: The request initiator, if different from the request token. |
| `resource_group_id` | no | — |
| `resource_instance_id` | no | — |
| `state` | no | — |
| `target_time` | no | — |
| `updated_at` | no | — |
| `updated_by` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_resource_reclamation_delete_invalid_comment` | `comment` |
| `ibm_resource_reclamation_delete_invalid_reclamation_id` | `reclamation_id` |
| `ibm_resource_reclamation_delete_invalid_request_by` | `request_by` |

## Example

### Valid

```hcl
resource "ibm_resource_reclamation_delete" "test" {
  reclamation_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
