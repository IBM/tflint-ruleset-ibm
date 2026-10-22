# ibm_satellite_attach_host_script_data

Schema validation for the ibm_satellite_attach_host_script data source.

This rule validates `ibm_satellite_attach_host_script` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_attach_host_script_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `coreos_host` | no | — |
| `custom_script` | no | — |
| `description` | no | — |
| `host_link_agent_endpoint` | no | — |
| `host_provider` | no | — |
| `host_script` | no | — |
| `labels` | no | — |
| `location` | yes | — |
| `script_dir` | no | — |
| `script_path` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_attach_host_script_data_invalid_coreos_host` | `coreos_host` |
| `ibm_satellite_attach_host_script_data_invalid_custom_script` | `custom_script` |
| `ibm_satellite_attach_host_script_data_invalid_host_link_agent_endpoint` | `host_link_agent_endpoint` |
| `ibm_satellite_attach_host_script_data_invalid_host_provider` | `host_provider` |
| `ibm_satellite_attach_host_script_data_invalid_labels` | `labels` |
| `ibm_satellite_attach_host_script_data_invalid_location` | `location` |
| `ibm_satellite_attach_host_script_data_invalid_script_dir` | `script_dir` |

## Example

### Valid

```hcl
data "ibm_satellite_attach_host_script" "test" {
  location = "example"
}
```
