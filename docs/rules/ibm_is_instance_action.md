# ibm_is_instance_action

Schema validation for the ibm_is_instance_action resource.

This rule validates `ibm_is_instance_action` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `LoadBalancerDNSPatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_instance_action" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | yes | one of: `start`, `reboot`, `stop` |
| `force_action` | no | — |
| `instance` | yes | API: Identifies a DNS instance by a unique property. |
| `status` | no | — |
| `status_reasons` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_instance_action_invalid_action` | `action` |
| `ibm_is_instance_action_invalid_force_action` | `force_action` |
| `ibm_is_instance_action_invalid_instance` | `instance` |

## Example

### Valid

```hcl
resource "ibm_is_instance_action" "test" {
  action = "start"
  instance = "example"
}
```
