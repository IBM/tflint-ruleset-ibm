# ibm_scc_profile_attachment_data

Reads ibm_scc_profile_attachment. The name of the attachment.

This rule validates `ibm_scc_profile_attachment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ScanRequest` (`security-compliance`).

## Configuration

```hcl
rule "ibm_scc_profile_attachment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | no | — |
| `attachment_id` | yes | API: The ID of the profile attachment |
| `attachment_item_id` | no | — |
| `attachment_parameters` | no | — |
| `created_by` | no | — |
| `created_on` | no | — |
| `description` | no | — |
| `instance_id` | yes | — |
| `last_scan` | no | — |
| `name` | no | — |
| `next_scan_time` | no | — |
| `notifications` | no | — |
| `profile_id` | yes | pattern `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$` |
| `schedule` | no | — |
| `scope` | no | — |
| `status` | no | — |
| `updated_by` | no | — |
| `updated_on` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_profile_attachment_data_invalid_attachment_id` | `attachment_id` |
| `ibm_scc_profile_attachment_data_invalid_instance_id` | `instance_id` |
| `ibm_scc_profile_attachment_data_invalid_profile_id` | `profile_id` |

## Example

### Valid

```hcl
data "ibm_scc_profile_attachment" "test" {
  attachment_id = "01234567-89ab-cdef-0123-456789abcdef"
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  profile_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
