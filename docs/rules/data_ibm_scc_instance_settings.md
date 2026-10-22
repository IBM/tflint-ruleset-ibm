# ibm_scc_instance_settings_data

Schema validation for the ibm_scc_instance_settings data source.

This rule validates `ibm_scc_instance_settings` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_scc_instance_settings_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `event_notifications` | no | — |
| `instance_id` | yes | pattern `^[a-zA-Z0-9 ,\-_]+$` |
| `object_storage` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_scc_instance_settings_data_invalid_instance_id` | `instance_id` |

## Example

### Valid

```hcl
data "ibm_scc_instance_settings" "test" {
  instance_id = "example-name"
}
```
