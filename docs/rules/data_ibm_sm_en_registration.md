# ibm_sm_en_registration_data

Schema validation for the ibm_sm_en_registration data source.

This rule validates `ibm_sm_en_registration` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_sm_en_registration_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `event_notifications_instance_crn` | no | — |

## Example

### Valid

```hcl
data "ibm_sm_en_registration" "test" {
}
```
