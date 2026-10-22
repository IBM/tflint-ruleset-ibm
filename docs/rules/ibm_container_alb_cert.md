# ibm_container_alb_cert

Schema validation for the ibm_container_alb_cert resource.

This rule validates `ibm_container_alb_cert` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ALB` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_alb_cert" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cert_crn` | yes | crn format |
| `cloud_cert_instance_id` | no | — |
| `cluster_id` | yes | — |
| `domain_name` | no | — |
| `expires_on` | no | — |
| `issuer_name` | no | — |
| `namespace` | no | API: The namespace value. |
| `persistence` | no | — |
| `region` | no | API: The region value. |
| `secret_name` | yes | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_alb_cert_invalid_cert_crn` | `cert_crn` |
| `ibm_container_alb_cert_invalid_cluster_id` | `cluster_id` |
| `ibm_container_alb_cert_invalid_namespace` | `namespace` |
| `ibm_container_alb_cert_invalid_persistence` | `persistence` |
| `ibm_container_alb_cert_invalid_region` | `region` |
| `ibm_container_alb_cert_invalid_secret_name` | `secret_name` |

## Example

### Valid

```hcl
variable "secret_name" {
  sensitive = true
}

resource "ibm_container_alb_cert" "test" {
  cert_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  cluster_id = "01234567-89ab-cdef-0123-456789abcdef"
  secret_name = var.secret_name
}
```
