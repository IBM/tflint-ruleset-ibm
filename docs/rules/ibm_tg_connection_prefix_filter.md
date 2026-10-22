# ibm_tg_connection_prefix_filter

Schema validation for the ibm_tg_connection_prefix_filter resource.

This rule validates `ibm_tg_connection_prefix_filter` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PrefixFilterTemplate` (`transit-gateway`).

## Configuration

```hcl
rule "ibm_tg_connection_prefix_filter" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | yes | one of: `permit`, `deny` |
| `before` | no | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `connection_id` | yes | pattern `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$` |
| `created_at` | no | — |
| `filter_id` | no | — |
| `gateway` | yes | — |
| `ge` | no | range 0–32 |
| `le` | no | range 0–32 |
| `prefix` | yes | pattern `^(([0-9]){1,3}\.){3}([0-9]{1,3})/([0-9]{1,2})$` |
| `updated_at` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_tg_connection_prefix_filter_invalid_action` | `action` |
| `ibm_tg_connection_prefix_filter_invalid_before` | `before` |
| `ibm_tg_connection_prefix_filter_invalid_connection_id` | `connection_id` |
| `ibm_tg_connection_prefix_filter_invalid_gateway` | `gateway` |
| `ibm_tg_connection_prefix_filter_invalid_ge` | `ge` |
| `ibm_tg_connection_prefix_filter_invalid_le` | `le` |
| `ibm_tg_connection_prefix_filter_invalid_prefix` | `prefix` |

## Example

### Valid

```hcl
resource "ibm_tg_connection_prefix_filter" "test" {
  action = "permit"
  connection_id = "01234567-89ab-cdef-0123-456789abcdef"
  gateway = "example"
  prefix = "10.0.0.0/24"
}
```
