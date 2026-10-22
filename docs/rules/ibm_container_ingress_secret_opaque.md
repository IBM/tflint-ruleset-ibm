# ibm_container_ingress_secret_opaque

Schema validation for the ibm_container_ingress_secret_opaque resource.

This rule validates `ibm_container_ingress_secret_opaque` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Secret` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_ingress_secret_opaque" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster` | yes | API: The cluster value. |
| `last_updated_timestamp` | no | — |
| `persistence` | no | API: The persistence value. |
| `secret_name` | yes | pattern `^([a-z0-9]([-a-z0-9]*[a-z0-9])?(.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*)$` |
| `secret_namespace` | yes | — |
| `status` | no | — |
| `type` | no | — |
| `update_secret` | no | — |
| `user_managed` | no | — |
| `fields` (block) | yes | nested block |
| `fields.crn` | yes | crn format |
| `fields.field_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_ingress_secret_opaque_invalid_cluster` | `cluster` |
| `ibm_container_ingress_secret_opaque_invalid_persistence` | `persistence` |
| `ibm_container_ingress_secret_opaque_invalid_secret_name` | `secret_name` |
| `ibm_container_ingress_secret_opaque_invalid_secret_namespace` | `secret_namespace` |
| `ibm_container_ingress_secret_opaque_invalid_update_secret` | `update_secret` |
| `ibm_container_ingress_secret_opaque_invalid_fields__crn` | `fields.crn` |
| `ibm_container_ingress_secret_opaque_invalid_fields__field_name` | `fields.field_name` |

## Example

### Valid

```hcl
variable "secret_name" {
  sensitive = true
}

variable "secret_namespace" {
  sensitive = true
}

resource "ibm_container_ingress_secret_opaque" "test" {
  cluster = "aaaaaaa"
  secret_name = var.secret_name
  secret_namespace = var.secret_namespace
  fields {
    crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  }
}
```
