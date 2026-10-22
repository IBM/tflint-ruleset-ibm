# ibm_container_addons_data

Schema validation for the ibm_container_addons data source.

This rule validates `ibm_container_addons` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DisablePodSecurityPolicyRequest` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_addons_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `addons` | no | — |
| `cluster` | yes | API: The cluster value. |
| `resource_group_id` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_addons_data_invalid_cluster` | `cluster` |

## Example

### Valid

```hcl
data "ibm_container_addons" "test" {
  cluster = "aaaaaaa"
}
```
