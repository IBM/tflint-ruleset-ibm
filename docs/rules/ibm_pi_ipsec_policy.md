# ibm_pi_ipsec_policy

Schema validation for the ibm_pi_ipsec_policy resource.

This rule validates `ibm_pi_ipsec_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_ipsec_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `policy_id` | no | — |

## Example

### Valid

```hcl
resource "ibm_pi_ipsec_policy" "test" {
}
```
