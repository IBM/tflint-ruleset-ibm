# ibm_backup_recovery_search_indexed_object_data

Schema validation for the ibm_backup_recovery_search_indexed_object data source.

This rule validates `ibm_backup_recovery_search_indexed_object` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Object` (`baas-backup-recovery`).

## Configuration

```hcl
rule "ibm_backup_recovery_search_indexed_object_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cassandra_objects` | no | — |
| `couchbase_objects` | no | — |
| `emails` | no | — |
| `exchange_objects` | no | — |
| `files` | no | — |
| `hbase_objects` | no | — |
| `hdfs_objects` | no | — |
| `hive_objects` | no | — |
| `include_tenants` | no | — |
| `might_have_snapshot_tag_ids` | no | — |
| `might_have_tag_ids` | no | — |
| `mongo_objects` | no | — |
| `ms_group_items` | no | — |
| `must_have_snapshot_tag_ids` | no | — |
| `must_have_tag_ids` | no | — |
| `object_count` | no | — |
| `object_type` | yes | — |
| `one_drive_items` | no | — |
| `pagination_cookie` | no | — |
| `protection_group_ids` | no | — |
| `public_folder_items` | no | — |
| `sfdc_records` | no | — |
| `sharepoint_items` | no | — |
| `snapshot_tags` | no | — |
| `storage_domain_ids` | no | — |
| `tags` | no | — |
| `teams_items` | no | — |
| `tenant_id` | no | — |
| `uda_objects` | no | — |
| `use_cached_data` | no | — |
| `x_ibm_tenant_id` | yes | — |
| `cassandra_params` (block) | no | nested block |
| `cassandra_params.cassandra_object_types` | yes | — |
| `cassandra_params.search_string` | yes | — |
| `cassandra_params.source_ids` | no | — |
| `couchbase_params` (block) | no | nested block |
| `couchbase_params.couchbase_object_types` | yes | — |
| `couchbase_params.search_string` | yes | — |
| `couchbase_params.source_ids` | no | — |
| `email_params` (block) | no | nested block |
| `email_params.attendees_addresses` | no | — |
| `email_params.bcc_recipient_addresses` | no | — |
| `email_params.cc_recipient_addresses` | no | — |
| `email_params.created_end_time_secs` | no | — |
| `email_params.created_start_time_secs` | no | — |
| `email_params.due_date_end_time_secs` | no | — |
| `email_params.due_date_start_time_secs` | no | — |
| `email_params.email_address` | no | — |
| `email_params.email_subject` | no | — |
| `email_params.first_name` | no | — |
| `email_params.folder_names` | no | — |
| `email_params.has_attachment` | no | — |
| `email_params.last_modified_end_time_secs` | no | — |
| `email_params.last_modified_start_time_secs` | no | — |
| `email_params.last_name` | no | — |
| `email_params.middle_name` | no | — |
| `email_params.organizer_address` | no | — |
| `email_params.received_end_time_secs` | no | — |
| `email_params.received_start_time_secs` | no | — |
| `email_params.recipient_addresses` | no | — |
| `email_params.sender_address` | no | — |
| `email_params.source_environment` | no | — |
| `email_params.task_status_types` | no | — |
| `email_params.types` | no | — |
| `email_params.o365_params` (block) | no | nested block |
| `email_params.o365_params.domain_ids` | no | — |
| `email_params.o365_params.mailbox_ids` | no | — |
| `exchange_params` (block) | no | nested block |
| `exchange_params.search_string` | yes | — |
| `file_params` (block) | no | nested block |
| `file_params.object_ids` | no | — |
| `file_params.search_string` | no | — |
| `file_params.source_environments` | no | — |
| `file_params.source_ids` | no | — |
| `file_params.types` | no | — |
| `hbase_params` (block) | no | nested block |
| `hbase_params.hbase_object_types` | yes | — |
| `hbase_params.search_string` | yes | — |
| `hbase_params.source_ids` | no | — |
| `hdfs_params` (block) | no | nested block |
| `hdfs_params.hdfs_types` | yes | — |
| `hdfs_params.search_string` | yes | — |
| `hdfs_params.source_ids` | no | — |
| `hive_params` (block) | no | nested block |
| `hive_params.hive_object_types` | yes | — |
| `hive_params.search_string` | yes | — |
| `hive_params.source_ids` | no | — |
| `mongodb_params` (block) | no | nested block |
| `mongodb_params.mongo_db_object_types` | yes | — |
| `mongodb_params.search_string` | yes | — |
| `mongodb_params.source_ids` | no | — |
| `ms_groups_params` (block) | no | nested block |
| `ms_groups_params.mailbox_params` (block) | no | nested block |
| `ms_groups_params.mailbox_params.attendees_addresses` | no | — |
| `ms_groups_params.mailbox_params.bcc_recipient_addresses` | no | — |
| `ms_groups_params.mailbox_params.cc_recipient_addresses` | no | — |
| `ms_groups_params.mailbox_params.created_end_time_secs` | no | — |
| `ms_groups_params.mailbox_params.created_start_time_secs` | no | — |
| `ms_groups_params.mailbox_params.due_date_end_time_secs` | no | — |
| `ms_groups_params.mailbox_params.due_date_start_time_secs` | no | — |
| `ms_groups_params.mailbox_params.email_address` | no | — |
| `ms_groups_params.mailbox_params.email_subject` | no | — |
| `ms_groups_params.mailbox_params.first_name` | no | — |
| `ms_groups_params.mailbox_params.folder_names` | no | — |
| `ms_groups_params.mailbox_params.has_attachment` | no | — |
| `ms_groups_params.mailbox_params.last_modified_end_time_secs` | no | — |
| `ms_groups_params.mailbox_params.last_modified_start_time_secs` | no | — |
| `ms_groups_params.mailbox_params.last_name` | no | — |
| `ms_groups_params.mailbox_params.middle_name` | no | — |
| `ms_groups_params.mailbox_params.organizer_address` | no | — |
| `ms_groups_params.mailbox_params.received_end_time_secs` | no | — |
| `ms_groups_params.mailbox_params.received_start_time_secs` | no | — |
| `ms_groups_params.mailbox_params.recipient_addresses` | no | — |
| `ms_groups_params.mailbox_params.sender_address` | no | — |
| `ms_groups_params.mailbox_params.source_environment` | no | — |
| `ms_groups_params.mailbox_params.task_status_types` | no | — |
| `ms_groups_params.mailbox_params.types` | no | — |
| `ms_groups_params.o365_params` (block) | no | nested block |
| `ms_groups_params.o365_params.domain_ids` | no | — |
| `ms_groups_params.o365_params.group_ids` | no | — |
| `ms_groups_params.o365_params.site_ids` | no | — |
| `ms_groups_params.o365_params.teams_ids` | no | — |
| `ms_groups_params.o365_params.user_ids` | no | — |
| `ms_groups_params.site_params` (block) | no | nested block |
| `ms_groups_params.site_params.category_types` | no | — |
| `ms_groups_params.site_params.creation_end_time_secs` | no | — |
| `ms_groups_params.site_params.creation_start_time_secs` | no | — |
| `ms_groups_params.site_params.include_files` | no | — |
| `ms_groups_params.site_params.include_folders` | no | — |
| `ms_groups_params.site_params.owner_names` | no | — |
| `ms_groups_params.site_params.search_string` | no | — |
| `ms_groups_params.site_params.size_bytes_lower_limit` | no | — |
| `ms_groups_params.site_params.size_bytes_upper_limit` | no | — |
| `ms_groups_params.site_params.o365_params` (block) | no | nested block |
| `ms_groups_params.site_params.o365_params.domain_ids` | no | — |
| `ms_groups_params.site_params.o365_params.group_ids` | no | — |
| `ms_groups_params.site_params.o365_params.site_ids` | no | — |
| `ms_groups_params.site_params.o365_params.teams_ids` | no | — |
| `ms_groups_params.site_params.o365_params.user_ids` | no | — |
| `ms_teams_params` (block) | no | nested block |
| `ms_teams_params.category_types` | no | — |
| `ms_teams_params.channel_names` | no | — |
| `ms_teams_params.creation_end_time_secs` | no | — |
| `ms_teams_params.creation_start_time_secs` | no | — |
| `ms_teams_params.owner_names` | no | — |
| `ms_teams_params.search_string` | no | — |
| `ms_teams_params.size_bytes_lower_limit` | no | — |
| `ms_teams_params.size_bytes_upper_limit` | no | — |
| `ms_teams_params.types` | no | — |
| `ms_teams_params.channel_params` (block) | no | nested block |
| `ms_teams_params.channel_params.channel_email` | no | — |
| `ms_teams_params.channel_params.channel_id` | no | — |
| `ms_teams_params.channel_params.channel_name` | no | — |
| `ms_teams_params.channel_params.include_private_channels` | no | — |
| `ms_teams_params.channel_params.include_public_channels` | no | — |
| `ms_teams_params.o365_params` (block) | no | nested block |
| `ms_teams_params.o365_params.domain_ids` | no | — |
| `ms_teams_params.o365_params.group_ids` | no | — |
| `ms_teams_params.o365_params.site_ids` | no | — |
| `ms_teams_params.o365_params.teams_ids` | no | — |
| `ms_teams_params.o365_params.user_ids` | no | — |
| `one_drive_params` (block) | no | nested block |
| `one_drive_params.category_types` | no | — |
| `one_drive_params.creation_end_time_secs` | no | — |
| `one_drive_params.creation_start_time_secs` | no | — |
| `one_drive_params.include_files` | no | — |
| `one_drive_params.include_folders` | no | — |
| `one_drive_params.owner_names` | no | — |
| `one_drive_params.search_string` | no | — |
| `one_drive_params.size_bytes_lower_limit` | no | — |
| `one_drive_params.size_bytes_upper_limit` | no | — |
| `one_drive_params.o365_params` (block) | no | nested block |
| `one_drive_params.o365_params.domain_ids` | no | — |
| `one_drive_params.o365_params.group_ids` | no | — |
| `one_drive_params.o365_params.site_ids` | no | — |
| `one_drive_params.o365_params.teams_ids` | no | — |
| `one_drive_params.o365_params.user_ids` | no | — |
| `public_folder_params` (block) | no | nested block |
| `public_folder_params.bcc_recipient_addresses` | no | — |
| `public_folder_params.cc_recipient_addresses` | no | — |
| `public_folder_params.has_attachment` | no | — |
| `public_folder_params.received_end_time_secs` | no | — |
| `public_folder_params.received_start_time_secs` | no | — |
| `public_folder_params.recipient_addresses` | no | — |
| `public_folder_params.search_string` | no | — |
| `public_folder_params.sender_address` | no | — |
| `public_folder_params.types` | no | — |
| `sfdc_params` (block) | no | nested block |
| `sfdc_params.mutation_types` | yes | — |
| `sfdc_params.object_name` | yes | — |
| `sfdc_params.query_string` | no | — |
| `sfdc_params.snapshot_id` | yes | — |
| `sharepoint_params` (block) | no | nested block |
| `sharepoint_params.category_types` | no | — |
| `sharepoint_params.creation_end_time_secs` | no | — |
| `sharepoint_params.creation_start_time_secs` | no | — |
| `sharepoint_params.include_files` | no | — |
| `sharepoint_params.include_folders` | no | — |
| `sharepoint_params.owner_names` | no | — |
| `sharepoint_params.search_string` | no | — |
| `sharepoint_params.size_bytes_lower_limit` | no | — |
| `sharepoint_params.size_bytes_upper_limit` | no | — |
| `sharepoint_params.o365_params` (block) | no | nested block |
| `sharepoint_params.o365_params.domain_ids` | no | — |
| `sharepoint_params.o365_params.group_ids` | no | — |
| `sharepoint_params.o365_params.site_ids` | no | — |
| `sharepoint_params.o365_params.teams_ids` | no | — |
| `sharepoint_params.o365_params.user_ids` | no | — |
| `uda_params` (block) | no | nested block |
| `uda_params.search_string` | yes | — |
| `uda_params.source_ids` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_backup_recovery_search_indexed_object_data_invalid_include_tenants` | `include_tenants` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_might_have_snapshot_tag_ids` | `might_have_snapshot_tag_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_might_have_tag_ids` | `might_have_tag_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_must_have_snapshot_tag_ids` | `must_have_snapshot_tag_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_must_have_tag_ids` | `must_have_tag_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_object_count` | `object_count` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_object_type` | `object_type` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_pagination_cookie` | `pagination_cookie` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_protection_group_ids` | `protection_group_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_snapshot_tags` | `snapshot_tags` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_storage_domain_ids` | `storage_domain_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_tags` | `tags` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_tenant_id` | `tenant_id` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_use_cached_data` | `use_cached_data` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_x_ibm_tenant_id` | `x_ibm_tenant_id` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_cassandra_params__cassandra_object_types` | `cassandra_params.cassandra_object_types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_cassandra_params__search_string` | `cassandra_params.search_string` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_cassandra_params__source_ids` | `cassandra_params.source_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_couchbase_params__couchbase_object_types` | `couchbase_params.couchbase_object_types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_couchbase_params__search_string` | `couchbase_params.search_string` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_couchbase_params__source_ids` | `couchbase_params.source_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__attendees_addresses` | `email_params.attendees_addresses` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__bcc_recipient_addresses` | `email_params.bcc_recipient_addresses` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__cc_recipient_addresses` | `email_params.cc_recipient_addresses` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__created_end_time_secs` | `email_params.created_end_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__created_start_time_secs` | `email_params.created_start_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__due_date_end_time_secs` | `email_params.due_date_end_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__due_date_start_time_secs` | `email_params.due_date_start_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__email_address` | `email_params.email_address` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__email_subject` | `email_params.email_subject` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__first_name` | `email_params.first_name` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__folder_names` | `email_params.folder_names` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__has_attachment` | `email_params.has_attachment` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__last_modified_end_time_secs` | `email_params.last_modified_end_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__last_modified_start_time_secs` | `email_params.last_modified_start_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__last_name` | `email_params.last_name` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__middle_name` | `email_params.middle_name` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__organizer_address` | `email_params.organizer_address` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__received_end_time_secs` | `email_params.received_end_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__received_start_time_secs` | `email_params.received_start_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__recipient_addresses` | `email_params.recipient_addresses` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__sender_address` | `email_params.sender_address` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__source_environment` | `email_params.source_environment` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__task_status_types` | `email_params.task_status_types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__types` | `email_params.types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__o365_params__domain_ids` | `email_params.o365_params.domain_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_email_params__o365_params__mailbox_ids` | `email_params.o365_params.mailbox_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_exchange_params__search_string` | `exchange_params.search_string` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_file_params__object_ids` | `file_params.object_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_file_params__search_string` | `file_params.search_string` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_file_params__source_environments` | `file_params.source_environments` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_file_params__source_ids` | `file_params.source_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_file_params__types` | `file_params.types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_hbase_params__hbase_object_types` | `hbase_params.hbase_object_types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_hbase_params__search_string` | `hbase_params.search_string` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_hbase_params__source_ids` | `hbase_params.source_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_hdfs_params__hdfs_types` | `hdfs_params.hdfs_types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_hdfs_params__search_string` | `hdfs_params.search_string` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_hdfs_params__source_ids` | `hdfs_params.source_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_hive_params__hive_object_types` | `hive_params.hive_object_types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_hive_params__search_string` | `hive_params.search_string` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_hive_params__source_ids` | `hive_params.source_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_mongodb_params__mongo_db_object_types` | `mongodb_params.mongo_db_object_types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_mongodb_params__search_string` | `mongodb_params.search_string` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_mongodb_params__source_ids` | `mongodb_params.source_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__attendees_addresses` | `ms_groups_params.mailbox_params.attendees_addresses` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__bcc_recipient_addresses` | `ms_groups_params.mailbox_params.bcc_recipient_addresses` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__cc_recipient_addresses` | `ms_groups_params.mailbox_params.cc_recipient_addresses` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__created_end_time_secs` | `ms_groups_params.mailbox_params.created_end_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__created_start_time_secs` | `ms_groups_params.mailbox_params.created_start_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__due_date_end_time_secs` | `ms_groups_params.mailbox_params.due_date_end_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__due_date_start_time_secs` | `ms_groups_params.mailbox_params.due_date_start_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__email_address` | `ms_groups_params.mailbox_params.email_address` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__email_subject` | `ms_groups_params.mailbox_params.email_subject` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__first_name` | `ms_groups_params.mailbox_params.first_name` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__folder_names` | `ms_groups_params.mailbox_params.folder_names` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__has_attachment` | `ms_groups_params.mailbox_params.has_attachment` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__last_modified_end_time_secs` | `ms_groups_params.mailbox_params.last_modified_end_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__last_modified_start_time_secs` | `ms_groups_params.mailbox_params.last_modified_start_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__last_name` | `ms_groups_params.mailbox_params.last_name` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__middle_name` | `ms_groups_params.mailbox_params.middle_name` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__organizer_address` | `ms_groups_params.mailbox_params.organizer_address` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__received_end_time_secs` | `ms_groups_params.mailbox_params.received_end_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__received_start_time_secs` | `ms_groups_params.mailbox_params.received_start_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__recipient_addresses` | `ms_groups_params.mailbox_params.recipient_addresses` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__sender_address` | `ms_groups_params.mailbox_params.sender_address` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__source_environment` | `ms_groups_params.mailbox_params.source_environment` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__task_status_types` | `ms_groups_params.mailbox_params.task_status_types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__mailbox_params__types` | `ms_groups_params.mailbox_params.types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__o365_params__domain_ids` | `ms_groups_params.o365_params.domain_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__o365_params__group_ids` | `ms_groups_params.o365_params.group_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__o365_params__site_ids` | `ms_groups_params.o365_params.site_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__o365_params__teams_ids` | `ms_groups_params.o365_params.teams_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__o365_params__user_ids` | `ms_groups_params.o365_params.user_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__site_params__category_types` | `ms_groups_params.site_params.category_types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__site_params__creation_end_time_secs` | `ms_groups_params.site_params.creation_end_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__site_params__creation_start_time_secs` | `ms_groups_params.site_params.creation_start_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__site_params__include_files` | `ms_groups_params.site_params.include_files` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__site_params__include_folders` | `ms_groups_params.site_params.include_folders` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__site_params__owner_names` | `ms_groups_params.site_params.owner_names` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__site_params__search_string` | `ms_groups_params.site_params.search_string` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__site_params__size_bytes_lower_limit` | `ms_groups_params.site_params.size_bytes_lower_limit` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__site_params__size_bytes_upper_limit` | `ms_groups_params.site_params.size_bytes_upper_limit` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__site_params__o365_params__domain_ids` | `ms_groups_params.site_params.o365_params.domain_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__site_params__o365_params__group_ids` | `ms_groups_params.site_params.o365_params.group_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__site_params__o365_params__site_ids` | `ms_groups_params.site_params.o365_params.site_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__site_params__o365_params__teams_ids` | `ms_groups_params.site_params.o365_params.teams_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_groups_params__site_params__o365_params__user_ids` | `ms_groups_params.site_params.o365_params.user_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__category_types` | `ms_teams_params.category_types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__channel_names` | `ms_teams_params.channel_names` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__creation_end_time_secs` | `ms_teams_params.creation_end_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__creation_start_time_secs` | `ms_teams_params.creation_start_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__owner_names` | `ms_teams_params.owner_names` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__search_string` | `ms_teams_params.search_string` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__size_bytes_lower_limit` | `ms_teams_params.size_bytes_lower_limit` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__size_bytes_upper_limit` | `ms_teams_params.size_bytes_upper_limit` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__types` | `ms_teams_params.types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__channel_params__channel_email` | `ms_teams_params.channel_params.channel_email` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__channel_params__channel_id` | `ms_teams_params.channel_params.channel_id` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__channel_params__channel_name` | `ms_teams_params.channel_params.channel_name` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__channel_params__include_private_channels` | `ms_teams_params.channel_params.include_private_channels` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__channel_params__include_public_channels` | `ms_teams_params.channel_params.include_public_channels` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__o365_params__domain_ids` | `ms_teams_params.o365_params.domain_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__o365_params__group_ids` | `ms_teams_params.o365_params.group_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__o365_params__site_ids` | `ms_teams_params.o365_params.site_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__o365_params__teams_ids` | `ms_teams_params.o365_params.teams_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_ms_teams_params__o365_params__user_ids` | `ms_teams_params.o365_params.user_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_one_drive_params__category_types` | `one_drive_params.category_types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_one_drive_params__creation_end_time_secs` | `one_drive_params.creation_end_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_one_drive_params__creation_start_time_secs` | `one_drive_params.creation_start_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_one_drive_params__include_files` | `one_drive_params.include_files` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_one_drive_params__include_folders` | `one_drive_params.include_folders` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_one_drive_params__owner_names` | `one_drive_params.owner_names` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_one_drive_params__search_string` | `one_drive_params.search_string` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_one_drive_params__size_bytes_lower_limit` | `one_drive_params.size_bytes_lower_limit` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_one_drive_params__size_bytes_upper_limit` | `one_drive_params.size_bytes_upper_limit` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_one_drive_params__o365_params__domain_ids` | `one_drive_params.o365_params.domain_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_one_drive_params__o365_params__group_ids` | `one_drive_params.o365_params.group_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_one_drive_params__o365_params__site_ids` | `one_drive_params.o365_params.site_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_one_drive_params__o365_params__teams_ids` | `one_drive_params.o365_params.teams_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_one_drive_params__o365_params__user_ids` | `one_drive_params.o365_params.user_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_public_folder_params__bcc_recipient_addresses` | `public_folder_params.bcc_recipient_addresses` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_public_folder_params__cc_recipient_addresses` | `public_folder_params.cc_recipient_addresses` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_public_folder_params__has_attachment` | `public_folder_params.has_attachment` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_public_folder_params__received_end_time_secs` | `public_folder_params.received_end_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_public_folder_params__received_start_time_secs` | `public_folder_params.received_start_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_public_folder_params__recipient_addresses` | `public_folder_params.recipient_addresses` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_public_folder_params__search_string` | `public_folder_params.search_string` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_public_folder_params__sender_address` | `public_folder_params.sender_address` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_public_folder_params__types` | `public_folder_params.types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sfdc_params__mutation_types` | `sfdc_params.mutation_types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sfdc_params__object_name` | `sfdc_params.object_name` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sfdc_params__query_string` | `sfdc_params.query_string` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sfdc_params__snapshot_id` | `sfdc_params.snapshot_id` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sharepoint_params__category_types` | `sharepoint_params.category_types` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sharepoint_params__creation_end_time_secs` | `sharepoint_params.creation_end_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sharepoint_params__creation_start_time_secs` | `sharepoint_params.creation_start_time_secs` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sharepoint_params__include_files` | `sharepoint_params.include_files` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sharepoint_params__include_folders` | `sharepoint_params.include_folders` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sharepoint_params__owner_names` | `sharepoint_params.owner_names` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sharepoint_params__search_string` | `sharepoint_params.search_string` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sharepoint_params__size_bytes_lower_limit` | `sharepoint_params.size_bytes_lower_limit` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sharepoint_params__size_bytes_upper_limit` | `sharepoint_params.size_bytes_upper_limit` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sharepoint_params__o365_params__domain_ids` | `sharepoint_params.o365_params.domain_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sharepoint_params__o365_params__group_ids` | `sharepoint_params.o365_params.group_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sharepoint_params__o365_params__site_ids` | `sharepoint_params.o365_params.site_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sharepoint_params__o365_params__teams_ids` | `sharepoint_params.o365_params.teams_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_sharepoint_params__o365_params__user_ids` | `sharepoint_params.o365_params.user_ids` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_uda_params__search_string` | `uda_params.search_string` |
| `ibm_backup_recovery_search_indexed_object_data_invalid_uda_params__source_ids` | `uda_params.source_ids` |

## Example

### Valid

```hcl
data "ibm_backup_recovery_search_indexed_object" "test" {
  object_type = "example"
  x_ibm_tenant_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
