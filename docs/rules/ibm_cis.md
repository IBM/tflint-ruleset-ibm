# ibm_cis

Manages ibm_cis. A name for the resource instance

This rule validates `ibm_cis` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `purge_tags_body` (`cis`).

## Configuration

```hcl
rule "ibm_cis" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `guid` | no | — |
| `location` | yes | — |
| `name` | yes | one of: `url` |
| `parameters` | no | — |
| `plan` | yes | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `service` | no | — |
| `status` | no | — |
| `tags` | no | pattern `^[A-Za-z0-9:_ .-]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_invalid_location` | `location` |
| `ibm_cis_invalid_name` | `name` |
| `ibm_cis_invalid_parameters` | `parameters` |
| `ibm_cis_invalid_plan` | `plan` |
| `ibm_cis_invalid_resource_group_id` | `resource_group_id` |
| `ibm_cis_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_cis" "test" {
  location = "example"
  name = "url"
  plan = "example"
}
```
