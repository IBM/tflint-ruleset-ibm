# ibm_app_domain_shared

Manages ibm_app_domain_shared. The name of the domain

This rule validates `ibm_app_domain_shared` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_app_domain_shared" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | yes | — |
| `router_group_guid` | no | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_domain_shared_invalid_name` | `name` |
| `ibm_app_domain_shared_invalid_router_group_guid` | `router_group_guid` |
| `ibm_app_domain_shared_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_app_domain_shared" "test" {
  name = "example"
}
```
