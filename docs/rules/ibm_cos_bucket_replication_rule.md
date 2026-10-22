# ibm_cos_bucket_replication_rule

Schema validation for the ibm_cos_bucket_replication_rule resource.

This rule validates `ibm_cos_bucket_replication_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cos_bucket_replication_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bucket_crn` | yes | crn format |
| `bucket_location` | yes | — |
| `endpoint_type` | no | one of: `public`, `private`, `direct` |
| `replication_rule` (block) | yes | nested block |
| `replication_rule.deletemarker_replication_status` | no | — |
| `replication_rule.destination_bucket_crn` | yes | crn format |
| `replication_rule.enable` | yes | — |
| `replication_rule.prefix` | no | — |
| `replication_rule.priority` | no | — |
| `replication_rule.rule_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cos_bucket_replication_rule_invalid_bucket_crn` | `bucket_crn` |
| `ibm_cos_bucket_replication_rule_invalid_bucket_location` | `bucket_location` |
| `ibm_cos_bucket_replication_rule_invalid_endpoint_type` | `endpoint_type` |
| `ibm_cos_bucket_replication_rule_invalid_replication_rule__deletemarker_replication_status` | `replication_rule.deletemarker_replication_status` |
| `ibm_cos_bucket_replication_rule_invalid_replication_rule__destination_bucket_crn` | `replication_rule.destination_bucket_crn` |
| `ibm_cos_bucket_replication_rule_invalid_replication_rule__enable` | `replication_rule.enable` |
| `ibm_cos_bucket_replication_rule_invalid_replication_rule__prefix` | `replication_rule.prefix` |
| `ibm_cos_bucket_replication_rule_invalid_replication_rule__priority` | `replication_rule.priority` |
| `ibm_cos_bucket_replication_rule_invalid_replication_rule__rule_id` | `replication_rule.rule_id` |

## Example

### Valid

```hcl
resource "ibm_cos_bucket_replication_rule" "test" {
  bucket_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  bucket_location = "example"
  replication_rule {
    destination_bucket_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
    enable = true
  }
}
```
