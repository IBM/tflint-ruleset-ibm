# ibm_pi_workspaces_data

Schema validation for the ibm_pi_workspaces data source.

This rule validates `ibm_pi_workspaces` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pi_workspaces_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `pi_cloud_instance_id` | no | — |
| `workspaces` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pi_workspaces_data_invalid_pi_cloud_instance_id` | `pi_cloud_instance_id` |

## Example

### Valid

```hcl
data "ibm_pi_workspaces" "test" {
}
```
