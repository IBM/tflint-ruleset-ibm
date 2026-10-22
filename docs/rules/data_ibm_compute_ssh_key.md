# ibm_compute_ssh_key_data

Schema validation for the ibm_compute_ssh_key data source.

This rule validates `ibm_compute_ssh_key` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_compute_ssh_key_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `fingerprint` | no | — |
| `label` | yes | — |
| `most_recent` | no | — |
| `notes` | no | — |
| `public_key` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_compute_ssh_key_data_invalid_label` | `label` |
| `ibm_compute_ssh_key_data_invalid_most_recent` | `most_recent` |

## Example

### Valid

```hcl
data "ibm_compute_ssh_key" "test" {
  label = "example"
}
```
