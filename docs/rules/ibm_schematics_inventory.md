# ibm_schematics_inventory

Manages ibm_schematics_inventory. The unique name of your Inventory definition.

This rule validates `ibm_schematics_inventory` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/v2/inventories/{inventory_id}:put` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_inventory" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `created_by` | no | — |
| `description` | no | API: The description of your Inventory definition. The description can be up to 2048 characters long in size. |
| `inventories_ini` | no | API: Input inventory of host and host group for the playbook, in the `.ini` file format. |
| `location` | no | one of: `us-south`, `us-east`, `eu-gb`, `eu-de` |
| `name` | no | API: The unique name of your Inventory definition. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. |
| `resource_group` | no | API: Resource-group id for the shared dataset |
| `resource_queries` | no | API: Input resource query definitions that is used to dynamically generate the inventory of host and host group for the playbook |
| `updated_at` | no | — |
| `updated_by` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_inventory_invalid_description` | `description` |
| `ibm_schematics_inventory_invalid_inventories_ini` | `inventories_ini` |
| `ibm_schematics_inventory_invalid_location` | `location` |
| `ibm_schematics_inventory_invalid_name` | `name` |
| `ibm_schematics_inventory_invalid_resource_group` | `resource_group` |
| `ibm_schematics_inventory_invalid_resource_queries` | `resource_queries` |

## Example

### Valid

```hcl
resource "ibm_schematics_inventory" "test" {
}
```
