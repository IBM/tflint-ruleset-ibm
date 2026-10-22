# ibm_compute_ssh_key

Schema validation for the ibm_compute_ssh_key resource.

This rule validates `ibm_compute_ssh_key` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_ssh_key" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `fingerprint` | no | — |
| `label` | yes | — |
| `notes` | no | — |
| `public_key` | yes | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_ssh_key_invalid_label` | `label` |
| `ibm_compute_ssh_key_invalid_notes` | `notes` |
| `ibm_compute_ssh_key_invalid_public_key` | `public_key` |
| `ibm_compute_ssh_key_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_compute_ssh_key" "test" {
  label = "example"
  public_key = "example"
}
```
