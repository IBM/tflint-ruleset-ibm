# ibm_compute_provisioning_hook

Manages ibm_compute_provisioning_hook. Provision hook name

This rule validates `ibm_compute_provisioning_hook` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_provisioning_hook" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | yes | — |
| `tags` | no | — |
| `uri` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_provisioning_hook_invalid_name` | `name` |
| `ibm_compute_provisioning_hook_invalid_tags` | `tags` |
| `ibm_compute_provisioning_hook_invalid_uri` | `uri` |

## Example

### Valid

```hcl
resource "ibm_compute_provisioning_hook" "test" {
  name = "example"
  uri = "example"
}
```
