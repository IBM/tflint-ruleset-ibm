# ibm_cm_offering_instance

Schema validation for the ibm_cm_offering_instance resource.

This rule validates `ibm_cm_offering_instance` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `OfferingInstance` (`resource-catalog-private-catalog`).

## Configuration

```hcl
rule "ibm_cm_offering_instance" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `catalog_id` | yes | API: Catalog ID this instance was created from. |
| `channel` | no | API: Channel to pin the operator subscription to |
| `cluster_all_namespaces` | yes | API: designate to install into all namespaces |
| `cluster_id` | yes | API: Cluster ID |
| `cluster_namespaces` | yes | API: List of target namespaces to install into |
| `cluster_region` | yes | API: Cluster region (e.g., us-south) |
| `crn` | no | — |
| `install_plan` | no | API: Type of install plan (also known as approval strategy) for operator subscriptions. Can be either automatic, which automatically upgrades operators to the latest in a channel, or manual, which requires approval on the cluster |
| `kind_format` | yes | API: the format this instance has (helm, operator, ova...). |
| `label` | yes | API: the label for this instance. |
| `offering_id` | yes | API: Offering ID this instance was created from. |
| `parent_crn` | no | crn format |
| `plan_id` | no | API: The software plan ID used when creating this software instance |
| `resource_group_id` | no | API: Id of the resource group to provision the offering instance into |
| `schematics_workspace_id` | no | — |
| `url` | no | — |
| `version` | yes | API: The version this instance was installed from (semver - not version id) |
| `wait_until_successful` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cm_offering_instance_invalid_catalog_id` | `catalog_id` |
| `ibm_cm_offering_instance_invalid_channel` | `channel` |
| `ibm_cm_offering_instance_invalid_cluster_all_namespaces` | `cluster_all_namespaces` |
| `ibm_cm_offering_instance_invalid_cluster_id` | `cluster_id` |
| `ibm_cm_offering_instance_invalid_cluster_namespaces` | `cluster_namespaces` |
| `ibm_cm_offering_instance_invalid_cluster_region` | `cluster_region` |
| `ibm_cm_offering_instance_invalid_install_plan` | `install_plan` |
| `ibm_cm_offering_instance_invalid_kind_format` | `kind_format` |
| `ibm_cm_offering_instance_invalid_label` | `label` |
| `ibm_cm_offering_instance_invalid_offering_id` | `offering_id` |
| `ibm_cm_offering_instance_invalid_parent_crn` | `parent_crn` |
| `ibm_cm_offering_instance_invalid_plan_id` | `plan_id` |
| `ibm_cm_offering_instance_invalid_resource_group_id` | `resource_group_id` |
| `ibm_cm_offering_instance_invalid_version` | `version` |
| `ibm_cm_offering_instance_invalid_wait_until_successful` | `wait_until_successful` |

## Example

### Valid

```hcl
resource "ibm_cm_offering_instance" "test" {
  catalog_id = "01234567-89ab-cdef-0123-456789abcdef"
  cluster_all_namespaces = true
  cluster_id = "01234567-89ab-cdef-0123-456789abcdef"
  cluster_namespaces = "example"
  cluster_region = "example"
  kind_format = "example"
  label = "example"
  offering_id = "01234567-89ab-cdef-0123-456789abcdef"
  version = "example"
}
```
