# ibm_service_instance_data

Reads ibm_service_instance. Service instance name for example, speech_to_text

This rule validates `ibm_service_instance` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_service_instance_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `credentials` | no | — |
| `name` | yes | — |
| `service_keys` | no | — |
| `service_plan_guid` | no | — |
| `space_guid` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_service_instance_data_invalid_name` | `name` |
| `ibm_service_instance_data_invalid_space_guid` | `space_guid` |

## Example

### Valid

```hcl
data "ibm_service_instance" "test" {
  name = "example"
  space_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
