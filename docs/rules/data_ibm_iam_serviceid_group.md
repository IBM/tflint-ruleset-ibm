# ibm_iam_serviceid_group_data

Reads ibm_iam_serviceid_group. Name of the service ID group.

This rule validates `ibm_iam_serviceid_group` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_iam_serviceid_group_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `created_at` | no | — |
| `created_by` | no | — |
| `crn` | no | — |
| `description` | no | — |
| `entity_tag` | no | — |
| `iam_serviceid_group_id` | yes | — |
| `modified_at` | no | — |
| `name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_iam_serviceid_group_data_invalid_iam_serviceid_group_id` | `iam_serviceid_group_id` |

## Example

### Valid

```hcl
data "ibm_iam_serviceid_group" "test" {
  iam_serviceid_group_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
