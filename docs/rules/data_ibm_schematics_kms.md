# ibm_schematics_kms_data

Schema validation for the ibm_schematics_kms data source.

This rule validates `ibm_schematics_kms` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/settings/kms:put` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_kms_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `encryption_scheme` | no | — |
| `location` | no | API: The location to integrate kms instance. For example, location can be `US` and `EU`. |
| `primary_crk` | no | — |
| `resource_group` | no | — |
| `secondary_crk` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_kms_data_invalid_location` | `location` |

## Example

### Valid

```hcl
data "ibm_schematics_kms" "test" {
}
```
