# ibm_container_addons

Schema validation for the ibm_container_addons resource.

This rule validates `ibm_container_addons` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CommonClusterAddon` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_addons" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster` | yes | API: The cluster value. |
| `manage_all_addons` | no | — |
| `managed_addons` | no | — |
| `resource_group_id` | no | — |
| `addons` (block) | yes | nested block |
| `addons.name` | yes | API: The name value. |
| `addons.parameters_json` | no | — |
| `addons.version` | no | API: The version value. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_addons_invalid_cluster` | `cluster` |
| `ibm_container_addons_invalid_manage_all_addons` | `manage_all_addons` |
| `ibm_container_addons_invalid_resource_group_id` | `resource_group_id` |
| `ibm_container_addons_invalid_addons__name` | `addons.name` |
| `ibm_container_addons_invalid_addons__parameters_json` | `addons.parameters_json` |
| `ibm_container_addons_invalid_addons__version` | `addons.version` |

## Example

### Valid

```hcl
resource "ibm_container_addons" "test" {
  cluster = "aaaaaaa"
  addons {
    name = "aaaaaaa"
  }
}
```
