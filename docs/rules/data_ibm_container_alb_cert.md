# ibm_container_alb_cert_data

Schema validation for the ibm_container_alb_cert data source.

This rule validates `ibm_container_alb_cert` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ALB` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_alb_cert_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cert_crn` | no | — |
| `cloud_cert_instance_id` | no | — |
| `cluster_crn` | no | — |
| `cluster_id` | yes | — |
| `domain_name` | no | — |
| `expires_on` | no | — |
| `issuer_name` | no | — |
| `namespace` | no | API: The namespace value. |
| `persistence` | no | — |
| `secret_name` | yes | — |
| `status` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_alb_cert_data_invalid_cluster_id` | `cluster_id` |
| `ibm_container_alb_cert_data_invalid_namespace` | `namespace` |
| `ibm_container_alb_cert_data_invalid_secret_name` | `secret_name` |

## Example

### Valid

```hcl
variable "secret_name" {
  sensitive = true
}

data "ibm_container_alb_cert" "test" {
  cluster_id = "01234567-89ab-cdef-0123-456789abcdef"
  secret_name = var.secret_name
}
```
