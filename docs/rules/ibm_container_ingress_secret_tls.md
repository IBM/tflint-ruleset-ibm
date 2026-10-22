# ibm_container_ingress_secret_tls

Schema validation for the ibm_container_ingress_secret_tls resource.

This rule validates `ibm_container_ingress_secret_tls` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Secret` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_ingress_secret_tls" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cert_crn` | yes | crn format |
| `cluster` | yes | API: The cluster value. |
| `domain_name` | no | — |
| `expires_on` | no | — |
| `last_updated_timestamp` | no | — |
| `persistence` | no | API: The persistence value. |
| `secret_name` | yes | pattern `^([a-z0-9]([-a-z0-9]*[a-z0-9])?(.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*)$` |
| `secret_namespace` | yes | — |
| `status` | no | — |
| `type` | no | — |
| `update_secret` | no | — |
| `user_managed` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_ingress_secret_tls_invalid_cert_crn` | `cert_crn` |
| `ibm_container_ingress_secret_tls_invalid_cluster` | `cluster` |
| `ibm_container_ingress_secret_tls_invalid_persistence` | `persistence` |
| `ibm_container_ingress_secret_tls_invalid_secret_name` | `secret_name` |
| `ibm_container_ingress_secret_tls_invalid_secret_namespace` | `secret_namespace` |
| `ibm_container_ingress_secret_tls_invalid_update_secret` | `update_secret` |

## Example

### Valid

```hcl
variable "secret_name" {
  sensitive = true
}

variable "secret_namespace" {
  sensitive = true
}

resource "ibm_container_ingress_secret_tls" "test" {
  cert_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  cluster = "aaaaaaa"
  secret_name = var.secret_name
  secret_namespace = var.secret_namespace
}
```
