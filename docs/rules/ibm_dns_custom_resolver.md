# ibm_dns_custom_resolver

Manages ibm_dns_custom_resolver. Name of the custom resolver

This rule validates `ibm_dns_custom_resolver` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `custom_resolver_update_input` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_custom_resolver" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `allow_disruptive_updates` | no | API: Whether a disruptive update is allowed for the custom resolver. |
| `created_on` | no | — |
| `custom_resolver_id` | no | — |
| `description` | no | API: Descriptive text of the custom resolver. |
| `enabled` | no | API: Whether the custom resolver is enabled. |
| `health` | no | — |
| `high_availability` | no | — |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `modified_on` | no | — |
| `name` | yes | API: Name of the custom resolver. |
| `profile` | no | one of: `essential`, `advanced`, `premier` |
| `rules` | no | — |
| `locations` (block) | no | nested block |
| `locations.enabled` | no | API: Whether the custom resolver is enabled. |
| `locations.subnet_crn` | yes | crn format |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_custom_resolver_invalid_allow_disruptive_updates` | `allow_disruptive_updates` |
| `ibm_dns_custom_resolver_invalid_description` | `description` |
| `ibm_dns_custom_resolver_invalid_enabled` | `enabled` |
| `ibm_dns_custom_resolver_invalid_high_availability` | `high_availability` |
| `ibm_dns_custom_resolver_invalid_instance_id` | `instance_id` |
| `ibm_dns_custom_resolver_invalid_name` | `name` |
| `ibm_dns_custom_resolver_invalid_profile` | `profile` |
| `ibm_dns_custom_resolver_invalid_locations__enabled` | `locations.enabled` |
| `ibm_dns_custom_resolver_invalid_locations__subnet_crn` | `locations.subnet_crn` |

## Example

### Valid

```hcl
resource "ibm_dns_custom_resolver" "test" {
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  name = "example"
}
```
