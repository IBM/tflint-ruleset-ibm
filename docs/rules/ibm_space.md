# ibm_space

Manages ibm_space. The name for the space

This rule validates `ibm_space` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_space" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `auditors` | no | — |
| `developers` | no | — |
| `managers` | no | — |
| `name` | yes | — |
| `org` | yes | — |
| `space_quota` | no | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_space_invalid_auditors` | `auditors` |
| `ibm_space_invalid_developers` | `developers` |
| `ibm_space_invalid_managers` | `managers` |
| `ibm_space_invalid_name` | `name` |
| `ibm_space_invalid_org` | `org` |
| `ibm_space_invalid_space_quota` | `space_quota` |
| `ibm_space_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_space" "test" {
  name = "example"
  org = "example"
}
```
