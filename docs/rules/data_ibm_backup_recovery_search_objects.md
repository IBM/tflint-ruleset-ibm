# ibm_backup_recovery_search_objects_data

Schema validation for the ibm_backup_recovery_search_objects data source.

This rule validates `ibm_backup_recovery_search_objects` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_backup_recovery_search_objects_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster_identifiers` | no | — |
| `environments` | no | — |
| `external_filters` | no | — |
| `include_deleted_objects` | no | — |
| `include_helios_tag_info_for_objects` | no | — |
| `is_deleted` | no | — |
| `is_protected` | no | — |
| `last_run_status_list` | no | — |
| `might_have_snapshot_tag_ids` | no | — |
| `might_have_tag_ids` | no | — |
| `must_have_snapshot_tag_ids` | no | — |
| `must_have_tag_ids` | no | — |
| `object_count` | no | — |
| `object_ids` | no | — |
| `objects` | no | — |
| `os_types` | no | — |
| `pagination_cookie` | no | — |
| `protection_group_ids` | no | — |
| `protection_types` | no | — |
| `request_initiator_type` | no | — |
| `search_string` | no | — |
| `source_ids` | no | — |
| `source_uuids` | no | — |
| `tag_categories` | no | — |
| `tag_names` | no | — |
| `tag_search_name` | no | — |
| `tag_sub_categories` | no | — |
| `tag_types` | no | — |
| `x_ibm_tenant_id` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_search_objects_data_invalid_cluster_identifiers` | `cluster_identifiers` |
| `ibm_backup_recovery_search_objects_data_invalid_environments` | `environments` |
| `ibm_backup_recovery_search_objects_data_invalid_external_filters` | `external_filters` |
| `ibm_backup_recovery_search_objects_data_invalid_include_deleted_objects` | `include_deleted_objects` |
| `ibm_backup_recovery_search_objects_data_invalid_include_helios_tag_info_for_objects` | `include_helios_tag_info_for_objects` |
| `ibm_backup_recovery_search_objects_data_invalid_is_deleted` | `is_deleted` |
| `ibm_backup_recovery_search_objects_data_invalid_is_protected` | `is_protected` |
| `ibm_backup_recovery_search_objects_data_invalid_last_run_status_list` | `last_run_status_list` |
| `ibm_backup_recovery_search_objects_data_invalid_might_have_snapshot_tag_ids` | `might_have_snapshot_tag_ids` |
| `ibm_backup_recovery_search_objects_data_invalid_might_have_tag_ids` | `might_have_tag_ids` |
| `ibm_backup_recovery_search_objects_data_invalid_must_have_snapshot_tag_ids` | `must_have_snapshot_tag_ids` |
| `ibm_backup_recovery_search_objects_data_invalid_must_have_tag_ids` | `must_have_tag_ids` |
| `ibm_backup_recovery_search_objects_data_invalid_object_count` | `object_count` |
| `ibm_backup_recovery_search_objects_data_invalid_object_ids` | `object_ids` |
| `ibm_backup_recovery_search_objects_data_invalid_os_types` | `os_types` |
| `ibm_backup_recovery_search_objects_data_invalid_pagination_cookie` | `pagination_cookie` |
| `ibm_backup_recovery_search_objects_data_invalid_protection_group_ids` | `protection_group_ids` |
| `ibm_backup_recovery_search_objects_data_invalid_protection_types` | `protection_types` |
| `ibm_backup_recovery_search_objects_data_invalid_request_initiator_type` | `request_initiator_type` |
| `ibm_backup_recovery_search_objects_data_invalid_search_string` | `search_string` |
| `ibm_backup_recovery_search_objects_data_invalid_source_ids` | `source_ids` |
| `ibm_backup_recovery_search_objects_data_invalid_source_uuids` | `source_uuids` |
| `ibm_backup_recovery_search_objects_data_invalid_tag_categories` | `tag_categories` |
| `ibm_backup_recovery_search_objects_data_invalid_tag_names` | `tag_names` |
| `ibm_backup_recovery_search_objects_data_invalid_tag_search_name` | `tag_search_name` |
| `ibm_backup_recovery_search_objects_data_invalid_tag_sub_categories` | `tag_sub_categories` |
| `ibm_backup_recovery_search_objects_data_invalid_tag_types` | `tag_types` |
| `ibm_backup_recovery_search_objects_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_search_objects" "test" {
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
