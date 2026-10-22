# ibm_cbr_zone_addresses

Schema validation for the ibm_cbr_zone_addresses resource.

This rule validates `ibm_cbr_zone_addresses` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `ServiceRefValue` (`context-based-restrictions`).

## Configuration

```hcl
rule "ibm_cbr_zone_addresses" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `transaction_id` | no | pattern `^[a-zA-Z0-9 ,\-_]+$` |
| `x_correlation_id` | no | pattern `^[a-zA-Z0-9 ,\-_]+$` |
| `zone_id` | yes | pattern `^[a-zA-Z0-9\-]+$` |
| `addresses` (block) | yes | nested block |
| `addresses.type` | yes | — |
| `addresses.value` | no | — |
| `addresses.ref` (block) | no | nested block |
| `addresses.ref.account_id` | yes | pattern `^[a-zA-Z0-9\-]+$` |
| `addresses.ref.location` | no | pattern `^[0-9a-z\-]+$` |
| `addresses.ref.service_instance` | no | pattern `^[0-9a-z\-/]+$` |
| `addresses.ref.service_name` | no | pattern `^[0-9a-z\-]+$` |
| `addresses.ref.service_type` | no | pattern `^[0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cbr_zone_addresses_invalid_transaction_id` | `transaction_id` |
| `ibm_cbr_zone_addresses_invalid_x_correlation_id` | `x_correlation_id` |
| `ibm_cbr_zone_addresses_invalid_zone_id` | `zone_id` |
| `ibm_cbr_zone_addresses_invalid_addresses__type` | `addresses.type` |
| `ibm_cbr_zone_addresses_invalid_addresses__value` | `addresses.value` |
| `ibm_cbr_zone_addresses_invalid_addresses__ref__account_id` | `addresses.ref.account_id` |
| `ibm_cbr_zone_addresses_invalid_addresses__ref__location` | `addresses.ref.location` |
| `ibm_cbr_zone_addresses_invalid_addresses__ref__service_instance` | `addresses.ref.service_instance` |
| `ibm_cbr_zone_addresses_invalid_addresses__ref__service_name` | `addresses.ref.service_name` |
| `ibm_cbr_zone_addresses_invalid_addresses__ref__service_type` | `addresses.ref.service_type` |

## Example

### Valid

```hcl
resource "ibm_cbr_zone_addresses" "test" {
  zone_id = "example-name"
  addresses {
    type = "example"
  }
}
```
