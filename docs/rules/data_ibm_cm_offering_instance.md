# ibm_cm_offering_instance_data

Schema validation for the ibm_cm_offering_instance data source.

This rule validates `ibm_cm_offering_instance` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cm_offering_instance_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `_rev` | no | — |
| `catalog_id` | no | — |
| `channel` | no | — |
| `cluster_all_namespaces` | no | — |
| `cluster_id` | no | — |
| `cluster_namespaces` | no | — |
| `cluster_region` | no | — |
| `crn` | no | — |
| `install_plan` | no | — |
| `instance_identifier` | yes | — |
| `kind_format` | no | — |
| `label` | no | — |
| `offering_id` | no | — |
| `parent_crn` | no | — |
| `plan_id` | no | — |
| `resource_group_id` | no | — |
| `schematics_workspace_id` | no | — |
| `url` | no | — |
| `version` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cm_offering_instance_data_invalid_instance_identifier` | `instance_identifier` |

## Example

### Valid

```hcl
data "ibm_cm_offering_instance" "test" {
  instance_identifier = "example"
}
```
