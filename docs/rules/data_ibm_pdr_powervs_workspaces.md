# ibm_pdr_powervs_workspaces_data

Schema validation for the ibm_pdr_powervs_workspaces data source.

This rule validates `ibm_pdr_powervs_workspaces` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `PhaDeploymentRequest` (`dr-automation-powervs`).

## Configuration

```hcl
rule "ibm_pdr_powervs_workspaces_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `accept_language` | no | — |
| `dr_standby_workspace_description` | no | — |
| `dr_standby_workspaces` | no | — |
| `dr_workspace_description` | no | — |
| `dr_workspaces` | no | — |
| `instance_id` | yes | — |
| `location_id` | yes | pattern `^[a-zA-Z0-9._:-]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pdr_powervs_workspaces_data_invalid_accept_language` | `accept_language` |
| `ibm_pdr_powervs_workspaces_data_invalid_instance_id` | `instance_id` |
| `ibm_pdr_powervs_workspaces_data_invalid_location_id` | `location_id` |

## Example

### Valid

```hcl
data "ibm_pdr_powervs_workspaces" "test" {
  instance_id = "01234567-89ab-cdef-0123-456789abcdef"
  location_id = "example-name"
}
```
