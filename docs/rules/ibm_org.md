# ibm_org

Manages ibm_org. Org name, for example myorg@domain

This rule validates `ibm_org` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_org" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `auditors` | no | — |
| `billing_managers` | no | — |
| `managers` | no | — |
| `name` | yes | — |
| `org_quota_definition_guid` | no | — |
| `tags` | no | — |
| `users` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_org_invalid_auditors` | `auditors` |
| `ibm_org_invalid_billing_managers` | `billing_managers` |
| `ibm_org_invalid_managers` | `managers` |
| `ibm_org_invalid_name` | `name` |
| `ibm_org_invalid_org_quota_definition_guid` | `org_quota_definition_guid` |
| `ibm_org_invalid_tags` | `tags` |
| `ibm_org_invalid_users` | `users` |

## Example

### Valid

```hcl
resource "ibm_org" "test" {
  name = "example"
}
```
