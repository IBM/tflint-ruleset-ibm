# ibm_is_ssh_keys_data

Schema validation for the ibm_is_ssh_keys data source.

This rule validates `ibm_is_ssh_keys` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_is_ssh_keys_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `keys` | no | — |

## Example

### Valid

```hcl
data "ibm_is_ssh_keys" "test" {
}
```
