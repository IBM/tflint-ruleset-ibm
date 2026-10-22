# ibm_cloudant_data

Reads ibm_cloudant. Resource instance name for example, myobjectstorage

This rule validates `ibm_cloudant` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `IndexAnalysisExclusionReason` (`cloudant-cloudant-gen2`).

## Configuration

```hcl
rule "ibm_cloudant_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `crn` | no | — |
| `extensions` | no | — |
| `guid` | no | — |
| `identifier` | no | — |
| `location` | no | — |
| `name` | no | one of: `alphabetically_comes_after`, `empty_selector`, `excluded_by_user`, `field_mismatch`, `is_partial`, `less_overlap`, `needs_text_search`, `scope_mismatch`, `sort_order_mismatch`, `too_many_fields`, `unfavored_type` |
| `onetime_credentials` | no | — |
| `parameters_json` | no | — |
| `plan` | no | — |
| `resource_controller_url` | no | — |
| `resource_crn` | no | — |
| `resource_group_id` | no | — |
| `resource_group_name` | no | — |
| `resource_name` | no | — |
| `resource_status` | no | — |
| `service` | no | — |
| `status` | no | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cloudant_data_invalid_identifier` | `identifier` |
| `ibm_cloudant_data_invalid_location` | `location` |
| `ibm_cloudant_data_invalid_name` | `name` |
| `ibm_cloudant_data_invalid_resource_group_id` | `resource_group_id` |
| `ibm_cloudant_data_invalid_service` | `service` |

## Example

### Valid

```hcl
data "ibm_cloudant" "test" {
}
```
