# ibm_compute_ssl_certificate

Schema validation for the ibm_compute_ssl_certificate resource.

This rule validates `ibm_compute_ssl_certificate` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_ssl_certificate" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `certificate` | yes | — |
| `common_name` | no | — |
| `create_date` | no | — |
| `intermediate_certificate` | no | — |
| `key_size` | no | — |
| `modify_date` | no | — |
| `organization_name` | no | — |
| `private_key` | yes | — |
| `tags` | no | — |
| `validity_begin` | no | — |
| `validity_days` | no | — |
| `validity_end` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_ssl_certificate_invalid_certificate` | `certificate` |
| `ibm_compute_ssl_certificate_invalid_intermediate_certificate` | `intermediate_certificate` |
| `ibm_compute_ssl_certificate_invalid_private_key` | `private_key` |
| `ibm_compute_ssl_certificate_invalid_tags` | `tags` |

## Example

### Valid

```hcl
variable "private_key" {
  sensitive = true
}

resource "ibm_compute_ssl_certificate" "test" {
  certificate = "example"
  private_key = var.private_key
}
```
