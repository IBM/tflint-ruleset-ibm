# ibm_cm_object_data

Reads ibm_cm_object. The programmatic name of this object.

This rule validates `ibm_cm_object` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `params:put:/{object_id}/artifacts/{artifact_id}` (`resource-catalog-global-catalog`).

## Configuration

```hcl
rule "ibm_cm_object_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `catalog_id` | yes | API: Catalog ID |
| `catalog_name` | no | — |
| `catalog_object_id` | no | — |
| `created` | no | — |
| `crn` | no | — |
| `data` | no | — |
| `kind` | no | — |
| `label` | no | — |
| `label_i18n` | no | — |
| `name` | no | — |
| `object_id` | yes | — |
| `parent_id` | no | — |
| `publish` | no | — |
| `rev` | no | — |
| `short_description` | no | — |
| `short_description_i18n` | no | — |
| `state` | no | — |
| `tags` | no | — |
| `updated` | no | — |
| `url` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cm_object_data_invalid_catalog_id` | `catalog_id` |
| `ibm_cm_object_data_invalid_object_id` | `object_id` |

## Example

### Valid

```hcl
data "ibm_cm_object" "test" {
  catalog_id = "01234567-89ab-cdef-0123-456789abcdef"
  object_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
