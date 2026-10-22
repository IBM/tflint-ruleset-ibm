# ibm_lb_vpx_ha

Schema validation for the ibm_lb_vpx_ha resource.

This rule validates `ibm_lb_vpx_ha` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_lb_vpx_ha" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `primary_id` | yes | — |
| `secondary_id` | yes | — |
| `stay_secondary` | no | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_lb_vpx_ha_invalid_primary_id` | `primary_id` |
| `ibm_lb_vpx_ha_invalid_secondary_id` | `secondary_id` |
| `ibm_lb_vpx_ha_invalid_stay_secondary` | `stay_secondary` |
| `ibm_lb_vpx_ha_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_lb_vpx_ha" "test" {
  primary_id = 1
  secondary_id = 1
}
```
