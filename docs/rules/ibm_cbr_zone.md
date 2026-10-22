# ibm_cbr_zone

Manages ibm_cbr_zone. The name of the zone.

This rule validates `ibm_cbr_zone` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ServiceRefValue` (`context-based-restrictions`).

## Configuration

```hcl
rule "ibm_cbr_zone" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `account_id` | yes | pattern `^[a-zA-Z0-9\-]+$` |
| `address_count` | no | — |
| `created_at` | no | — |
| `created_by_id` | no | — |
| `crn` | no | — |
| `description` | no | API: The description of the zone. |
| `etag` | no | — |
| `excluded_count` | no | — |
| `href` | no | — |
| `last_modified_at` | no | — |
| `last_modified_by_id` | no | — |
| `name` | yes | pattern `^[a-zA-Z0-9 \-_]+$` |
| `transaction_id` | no | pattern `^[a-zA-Z0-9 ,\-_]+$` |
| `x_correlation_id` | no | pattern `^[a-zA-Z0-9 ,\-_]+$` |
| `addresses` (block) | no | nested block |
| `addresses.type` | yes | — |
| `addresses.value` | no | — |
| `addresses.ref` (block) | no | nested block |
| `addresses.ref.account_id` | yes | pattern `^[a-zA-Z0-9\-]+$` |
| `addresses.ref.location` | no | pattern `^[0-9a-z\-]+$` |
| `addresses.ref.service_instance` | no | pattern `^[0-9a-z\-/]+$` |
| `addresses.ref.service_name` | no | pattern `^[0-9a-z\-]+$` |
| `addresses.ref.service_type` | no | pattern `^[0-9a-z_]+$` |
| `excluded` (block) | no | nested block |
| `excluded.type` | yes | — |
| `excluded.value` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cbr_zone_invalid_account_id` | `account_id` |
| `ibm_cbr_zone_invalid_description` | `description` |
| `ibm_cbr_zone_invalid_name` | `name` |
| `ibm_cbr_zone_invalid_transaction_id` | `transaction_id` |
| `ibm_cbr_zone_invalid_x_correlation_id` | `x_correlation_id` |
| `ibm_cbr_zone_invalid_addresses__type` | `addresses.type` |
| `ibm_cbr_zone_invalid_addresses__value` | `addresses.value` |
| `ibm_cbr_zone_invalid_addresses__ref__account_id` | `addresses.ref.account_id` |
| `ibm_cbr_zone_invalid_addresses__ref__location` | `addresses.ref.location` |
| `ibm_cbr_zone_invalid_addresses__ref__service_instance` | `addresses.ref.service_instance` |
| `ibm_cbr_zone_invalid_addresses__ref__service_name` | `addresses.ref.service_name` |
| `ibm_cbr_zone_invalid_addresses__ref__service_type` | `addresses.ref.service_type` |
| `ibm_cbr_zone_invalid_excluded__type` | `excluded.type` |
| `ibm_cbr_zone_invalid_excluded__value` | `excluded.value` |

## Example

### Valid

```hcl
resource "ibm_cbr_zone" "test" {
  account_id = "example-name"
  name = "example-name"
}
```
