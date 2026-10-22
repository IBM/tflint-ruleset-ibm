# ibm_schematics_inventory_data

Reads ibm_schematics_inventory. The unique name of your Inventory.

This rule validates `ibm_schematics_inventory` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/v2/inventories/{inventory_id}` (`schematics-schematics`).

## Configuration

```hcl
rule "ibm_schematics_inventory_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `created_by` | no | — |
| `description` | no | — |
| `id` | no | — |
| `inventories_ini` | no | — |
| `inventory_id` | yes | — |
| `location` | no | — |
| `name` | no | — |
| `resource_group` | no | — |
| `resource_queries` | no | — |
| `updated_at` | no | — |
| `updated_by` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_schematics_inventory_data_invalid_inventory_id` | `inventory_id` |
| `ibm_schematics_inventory_data_invalid_location` | `location` |

## Example

### Valid

```hcl
data "ibm_schematics_inventory" "test" {
  inventory_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
