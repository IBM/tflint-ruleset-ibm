# ibm_org_data

Reads ibm_org. Org name, for example myorg@domain

This rule validates `ibm_org` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_org_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `name` | no | — |
| `org` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_org_data_invalid_name` | `name` |
| `ibm_org_data_invalid_org` | `org` |

## Example

### Valid

```hcl
data "ibm_org" "test" {
}
```
