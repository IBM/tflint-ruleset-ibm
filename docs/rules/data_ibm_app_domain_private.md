# ibm_app_domain_private_data

Reads ibm_app_domain_private. The name of the private domain

This rule validates `ibm_app_domain_private` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_app_domain_private_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_domain_private_data_invalid_name` | `name` |

## Example

### Valid

```hcl
data "ibm_app_domain_private" "test" {
  name = "example"
}
```
