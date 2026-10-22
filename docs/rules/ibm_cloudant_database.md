# ibm_cloudant_database

Schema validation for the ibm_cloudant_database resource.

This rule validates `ibm_cloudant_database` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DesignDocumentOptions` (`cloudant-cloudant-gen2`).

## Configuration

```hcl
rule "ibm_cloudant_database" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `db` | yes | pattern `^_dbs$|^_global_changes$|^_metadata$|^_nodes$|^_replicator$|^_users$|^[a-z][a-z0-9_$()+/-]*$` |
| `instance_crn` | yes | crn format |
| `partitioned` | no | API: Whether this design document describes partitioned or global indexes. Set this option to `false` for a design document that describes global indexes in a partitioned database. A design document describes either global or partitioned indexes, but not both. By default, for a partitioned database this option is `true` and the design document describes partitioned indexes for queries on a single parti |
| `shards` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cloudant_database_invalid_db` | `db` |
| `ibm_cloudant_database_invalid_instance_crn` | `instance_crn` |
| `ibm_cloudant_database_invalid_partitioned` | `partitioned` |
| `ibm_cloudant_database_invalid_shards` | `shards` |

## Example

### Valid

```hcl
resource "ibm_cloudant_database" "test" {
  db = "_dbs"
  instance_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
