# ibm_container_ingress_secret_tls_data

Schema validation for the ibm_container_ingress_secret_tls data source.

This rule validates `ibm_container_ingress_secret_tls` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DisablePodSecurityPolicyRequest` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_ingress_secret_tls_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cert_crn` | no | — |
| `cluster` | yes | API: The cluster value. |
| `domain_name` | no | — |
| `expires_on` | no | — |
| `last_updated_timestamp` | no | — |
| `persistence` | no | — |
| `secret_name` | yes | pattern `^([a-z0-9]([-a-z0-9]*[a-z0-9])?(.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*)$` |
| `secret_namespace` | yes | — |
| `status` | no | — |
| `type` | no | — |
| `user_managed` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_ingress_secret_tls_data_invalid_cluster` | `cluster` |
| `ibm_container_ingress_secret_tls_data_invalid_secret_name` | `secret_name` |
| `ibm_container_ingress_secret_tls_data_invalid_secret_namespace` | `secret_namespace` |

## Example

### Valid

```hcl
variable "secret_name" {
  sensitive = true
}

variable "secret_namespace" {
  sensitive = true
}

data "ibm_container_ingress_secret_tls" "test" {
  cluster = "aaaaaaa"
  secret_name = var.secret_name
  secret_namespace = var.secret_namespace
}
```
