# ibm_cloudant_database_data

Schema validation for the ibm_cloudant_database data source.

This rule validates `ibm_cloudant_database` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `UserContext` (`cloudant-cloudant-gen2`).

## Configuration

```hcl
rule "ibm_cloudant_database_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster` | no | — |
| `committed_update_seq` | no | — |
| `compact_running` | no | — |
| `compacted_seq` | no | — |
| `db` | yes | pattern `^_dbs$|^_global_changes$|^_metadata$|^_nodes$|^_replicator$|^_users$|^[a-z][a-z0-9_$()+/-]*$` |
| `disk_format_version` | no | — |
| `doc_count` | no | — |
| `doc_del_count` | no | — |
| `engine` | no | — |
| `instance_crn` | yes | crn format |
| `props` | no | — |
| `sizes` | no | — |
| `update_seq` | no | — |
| `uuid` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cloudant_database_data_invalid_db` | `db` |
| `ibm_cloudant_database_data_invalid_instance_crn` | `instance_crn` |

## Example

### Valid

```hcl
data "ibm_cloudant_database" "test" {
  db = "_dbs"
  instance_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
