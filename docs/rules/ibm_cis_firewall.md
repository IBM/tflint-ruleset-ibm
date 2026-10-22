# ibm_cis_firewall

Schema validation for the ibm_cis_firewall resource.

This rule validates `ibm_cis_firewall` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v1/{crn}/zones/{zone_identifier}/firewall/lockdowns:post` (`cis`).

## Configuration

```hcl
rule "ibm_cis_firewall" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cis_id` | yes | — |
| `domain_id` | yes | — |
| `firewall_type` | yes | — |
| `access_rule` (block) | no | nested block |
| `access_rule.mode` | yes | — |
| `access_rule.notes` | no | — |
| `access_rule.configuration` (block) | yes | nested block |
| `access_rule.configuration.target` | yes | one of: `ip`, `ip_range` |
| `access_rule.configuration.value` | yes | API: IP addresses or CIDR, if target is "ip", then value should be an IP addresses, otherwise CIDR. |
| `lockdown` (block) | no | nested block |
| `lockdown.description` | no | API: A note that you can use to describe the reason for a Lockdown rule. |
| `lockdown.paused` | no | API: Whether this zone lockdown is currently paused |
| `lockdown.priority` | no | API: Firewall priority |
| `lockdown.urls` | yes | API: URLs to be included in this rule definition. Wildcards are permitted. The URL pattern entered here will be escaped before use. This limits the URL to just simple wildcard patterns. |
| `lockdown.configurations` (block) | yes | nested block |
| `lockdown.configurations.target` | yes | one of: `ip`, `ip_range` |
| `lockdown.configurations.value` | yes | API: IP addresses or CIDR, if target is "ip", then value should be an IP addresses, otherwise CIDR. |
| `ua_rule` (block) | no | nested block |
| `ua_rule.description` | no | API: A note that you can use to describe the reason for a Lockdown rule. |
| `ua_rule.mode` | yes | — |
| `ua_rule.paused` | no | API: Whether this zone lockdown is currently paused |
| `ua_rule.configuration` (block) | yes | nested block |
| `ua_rule.configuration.target` | yes | one of: `ip`, `ip_range` |
| `ua_rule.configuration.value` | yes | API: IP addresses or CIDR, if target is "ip", then value should be an IP addresses, otherwise CIDR. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cis_firewall_invalid_cis_id` | `cis_id` |
| `ibm_cis_firewall_invalid_domain_id` | `domain_id` |
| `ibm_cis_firewall_invalid_firewall_type` | `firewall_type` |
| `ibm_cis_firewall_invalid_access_rule__mode` | `access_rule.mode` |
| `ibm_cis_firewall_invalid_access_rule__notes` | `access_rule.notes` |
| `ibm_cis_firewall_invalid_access_rule__configuration__target` | `access_rule.configuration.target` |
| `ibm_cis_firewall_invalid_access_rule__configuration__value` | `access_rule.configuration.value` |
| `ibm_cis_firewall_invalid_lockdown__description` | `lockdown.description` |
| `ibm_cis_firewall_invalid_lockdown__paused` | `lockdown.paused` |
| `ibm_cis_firewall_invalid_lockdown__priority` | `lockdown.priority` |
| `ibm_cis_firewall_invalid_lockdown__urls` | `lockdown.urls` |
| `ibm_cis_firewall_invalid_lockdown__configurations__target` | `lockdown.configurations.target` |
| `ibm_cis_firewall_invalid_lockdown__configurations__value` | `lockdown.configurations.value` |
| `ibm_cis_firewall_invalid_ua_rule__description` | `ua_rule.description` |
| `ibm_cis_firewall_invalid_ua_rule__mode` | `ua_rule.mode` |
| `ibm_cis_firewall_invalid_ua_rule__paused` | `ua_rule.paused` |
| `ibm_cis_firewall_invalid_ua_rule__configuration__target` | `ua_rule.configuration.target` |
| `ibm_cis_firewall_invalid_ua_rule__configuration__value` | `ua_rule.configuration.value` |

## Example

### Valid

```hcl
resource "ibm_cis_firewall" "test" {
  cis_id = "01234567-89ab-cdef-0123-456789abcdef"
  domain_id = "01234567-89ab-cdef-0123-456789abcdef"
  firewall_type = "example"
}
```
