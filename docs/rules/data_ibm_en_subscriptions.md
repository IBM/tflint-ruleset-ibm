# ibm_en_subscriptions_data

Schema validation for the ibm_en_subscriptions data source.

This rule validates `ibm_en_subscriptions` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_en_subscriptions_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `instance_guid` | yes | — |
| `search_key` | no | — |
| `subscriptions` | no | — |
| `total_count` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_en_subscriptions_data_invalid_instance_guid` | `instance_guid` |
| `ibm_en_subscriptions_data_invalid_search_key` | `search_key` |

## Example

### Valid

```hcl
data "ibm_en_subscriptions" "test" {
  instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
