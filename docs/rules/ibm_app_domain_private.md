# ibm_app_domain_private

Manages ibm_app_domain_private. The name of the domain

This rule validates `ibm_app_domain_private` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_app_domain_private" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | yes | — |
| `org_guid` | yes | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_domain_private_invalid_name` | `name` |
| `ibm_app_domain_private_invalid_org_guid` | `org_guid` |
| `ibm_app_domain_private_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_app_domain_private" "test" {
  name = "example"
  org_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
