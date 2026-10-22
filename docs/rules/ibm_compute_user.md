# ibm_compute_user

Schema validation for the ibm_compute_user resource.

This rule validates `ibm_compute_user` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_user" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `address1` | yes | — |
| `address2` | no | — |
| `api_key` | no | — |
| `city` | yes | — |
| `company_name` | yes | — |
| `country` | yes | — |
| `email` | yes | — |
| `first_name` | yes | — |
| `has_api_key` | no | — |
| `ibm_id` | no | — |
| `last_name` | yes | — |
| `password` | no | — |
| `permissions` | no | — |
| `state` | yes | — |
| `tags` | no | — |
| `timezone` | yes | — |
| `user_status` | no | — |
| `username` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_user_invalid_address1` | `address1` |
| `ibm_compute_user_invalid_address2` | `address2` |
| `ibm_compute_user_invalid_api_key` | `api_key` |
| `ibm_compute_user_invalid_city` | `city` |
| `ibm_compute_user_invalid_company_name` | `company_name` |
| `ibm_compute_user_invalid_country` | `country` |
| `ibm_compute_user_invalid_email` | `email` |
| `ibm_compute_user_invalid_first_name` | `first_name` |
| `ibm_compute_user_invalid_has_api_key` | `has_api_key` |
| `ibm_compute_user_invalid_last_name` | `last_name` |
| `ibm_compute_user_invalid_password` | `password` |
| `ibm_compute_user_invalid_permissions` | `permissions` |
| `ibm_compute_user_invalid_state` | `state` |
| `ibm_compute_user_invalid_tags` | `tags` |
| `ibm_compute_user_invalid_timezone` | `timezone` |
| `ibm_compute_user_invalid_user_status` | `user_status` |
| `ibm_compute_user_invalid_username` | `username` |

## Example

### Valid

```hcl
resource "ibm_compute_user" "test" {
  address1 = "example"
  city = "example"
  company_name = "example"
  country = "example"
  email = "example"
  first_name = "example"
  last_name = "example"
  state = "example"
  timezone = "example"
}
```
