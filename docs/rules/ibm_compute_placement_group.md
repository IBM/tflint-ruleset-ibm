# ibm_compute_placement_group

Manages ibm_compute_placement_group. Name

This rule validates `ibm_compute_placement_group` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_placement_group" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `datacenter` | yes | — |
| `name` | yes | — |
| `pod` | yes | — |
| `rule` | no | one of: `SPREAD` |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_placement_group_invalid_datacenter` | `datacenter` |
| `ibm_compute_placement_group_invalid_name` | `name` |
| `ibm_compute_placement_group_invalid_pod` | `pod` |
| `ibm_compute_placement_group_invalid_rule` | `rule` |
| `ibm_compute_placement_group_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_compute_placement_group" "test" {
  datacenter = "example"
  name = "example"
  pod = "example"
}
```
