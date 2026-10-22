# ibm_service_key

Manages ibm_service_key. The name of the service key

This rule validates `ibm_service_key` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_service_key" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `credentials` | no | — |
| `name` | yes | — |
| `parameters` | no | — |
| `service_instance_guid` | yes | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_service_key_invalid_name` | `name` |
| `ibm_service_key_invalid_parameters` | `parameters` |
| `ibm_service_key_invalid_service_instance_guid` | `service_instance_guid` |
| `ibm_service_key_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_service_key" "test" {
  name = "example"
  service_instance_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
