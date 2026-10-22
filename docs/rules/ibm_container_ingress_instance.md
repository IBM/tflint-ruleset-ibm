# ibm_container_ingress_instance

Schema validation for the ibm_container_ingress_instance resource.

This rule validates `ibm_container_ingress_instance` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/ingress/v2/security/port80:patch` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_ingress_instance" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster` | yes | API: The cluster value. |
| `instance_crn` | yes | crn format |
| `instance_name` | no | — |
| `instance_type` | no | — |
| `is_default` | no | — |
| `secret_group_id` | no | — |
| `secret_group_name` | no | — |
| `status` | no | — |
| `user_managed` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_ingress_instance_invalid_cluster` | `cluster` |
| `ibm_container_ingress_instance_invalid_instance_crn` | `instance_crn` |
| `ibm_container_ingress_instance_invalid_is_default` | `is_default` |
| `ibm_container_ingress_instance_invalid_secret_group_id` | `secret_group_id` |

## Example

### Valid

```hcl
resource "ibm_container_ingress_instance" "test" {
  cluster = "aaaaaaa"
  instance_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
