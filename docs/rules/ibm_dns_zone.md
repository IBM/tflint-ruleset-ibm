# ibm_dns_zone

Manages ibm_dns_zone. Zone name

This rule validates `ibm_dns_zone` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/instances/{instance_id}/dnszones:post` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_zone" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_on` | no | — |
| `description` | no | API: The text describing the purpose of a DNS zone |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `label` | no | API: The label of a DNS zone |
| `modified_on` | no | — |
| `name` | yes | API: Name of DNS zone |
| `state` | no | — |
| `zone_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_zone_invalid_description` | `description` |
| `ibm_dns_zone_invalid_instance_id` | `instance_id` |
| `ibm_dns_zone_invalid_label` | `label` |
| `ibm_dns_zone_invalid_name` | `name` |

## Example

### Valid

```hcl
resource "ibm_dns_zone" "test" {
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  name = "aaaaaaa"
}
```
