# ibm_dns_linked_zone

Manages ibm_dns_linked_zone. The name of the DNS Linked zone.

This rule validates `ibm_dns_linked_zone` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `linked_dnszone` (`dns-svcs`).

## Configuration

```hcl
rule "ibm_dns_linked_zone" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `approval_required_before` | no | API: The expired time of linked zone with state `approval pending` |
| `created_on` | no | — |
| `description` | no | API: Descriptive text of the linked zone |
| `instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `label` | no | API: The label of linked zone |
| `linked_to` | no | API: The owner's instance and zone that the zone is linked to |
| `modified_on` | no | — |
| `name` | yes | pattern `^[-0-9a-zA-Z\.]+$` |
| `owner_instance_id` | yes | pattern `^[-0-9a-z_]+$` |
| `owner_zone_id` | yes | pattern `^[-0-9a-z_]+$` |
| `state` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_dns_linked_zone_invalid_approval_required_before` | `approval_required_before` |
| `ibm_dns_linked_zone_invalid_description` | `description` |
| `ibm_dns_linked_zone_invalid_instance_id` | `instance_id` |
| `ibm_dns_linked_zone_invalid_label` | `label` |
| `ibm_dns_linked_zone_invalid_linked_to` | `linked_to` |
| `ibm_dns_linked_zone_invalid_name` | `name` |
| `ibm_dns_linked_zone_invalid_owner_instance_id` | `owner_instance_id` |
| `ibm_dns_linked_zone_invalid_owner_zone_id` | `owner_zone_id` |

## Example

### Valid

```hcl
resource "ibm_dns_linked_zone" "test" {
  instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  name = "example-name"
  owner_instance_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
  owner_zone_id = "example-namexxxxxxxxxxxxxxxxxxxxxxxx"
}
```
