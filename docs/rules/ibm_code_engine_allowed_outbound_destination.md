# ibm_code_engine_allowed_outbound_destination

Manages ibm_code_engine_allowed_outbound_destination. The name of the allowed outbound destination.

This rule validates `ibm_code_engine_allowed_outbound_destination` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `AllowedOutboundDestinationPrototype` (`codeengine-v2`).

## Configuration

```hcl
rule "ibm_code_engine_allowed_outbound_destination" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cidr_block` | no | pattern `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$` |
| `entity_tag` | no | — |
| `etag` | no | — |
| `isolation_policy` | no | one of: `dedicated`, `shared` |
| `name` | yes | API: The name of the allowed outbound destination. |
| `private_path_service_gateway_crn` | no | API: The CRN of the Private Path service. |
| `project_id` | yes | pattern `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$` |
| `status` | no | — |
| `status_details` | no | — |
| `type` | yes | one of: `cidr_block`, `private_path_service_gateway`; pattern `^(cidr_block|private_path_service_gateway)$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_code_engine_allowed_outbound_destination_invalid_cidr_block` | `cidr_block` |
| `ibm_code_engine_allowed_outbound_destination_invalid_isolation_policy` | `isolation_policy` |
| `ibm_code_engine_allowed_outbound_destination_invalid_name` | `name` |
| `ibm_code_engine_allowed_outbound_destination_invalid_private_path_service_gateway_crn` | `private_path_service_gateway_crn` |
| `ibm_code_engine_allowed_outbound_destination_invalid_project_id` | `project_id` |
| `ibm_code_engine_allowed_outbound_destination_invalid_type` | `type` |

## Example

### Valid

```hcl
resource "ibm_code_engine_allowed_outbound_destination" "test" {
  name = "example"
  project_id = "01234567-89ab-cdef-0123-456789abcdef"
  type = "cidr_block"
}
```
