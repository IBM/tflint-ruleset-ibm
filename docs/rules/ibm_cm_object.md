# ibm_cm_object

Manages ibm_cm_object. The programmatic name of this object.

This rule validates `ibm_cm_object` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CatalogObjectPrototype` (`resource-catalog-private-catalog`).

## Configuration

```hcl
rule "ibm_cm_object" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `catalog_id` | yes | API: The id of the catalog containing this offering |
| `catalog_name` | no | — |
| `created` | no | — |
| `crn` | no | — |
| `data` | no | API: Map of data values for this object |
| `kind` | yes | API: Kind of object |
| `label` | no | API: Display name in the requested language |
| `label_i18n` | no | API: A map of translated strings, by language code |
| `name` | yes | API: The programmatic name of this object |
| `object_id` | no | — |
| `parent_id` | no | API: The parent for this specific object |
| `publish` | no | — |
| `rev` | no | — |
| `short_description` | no | API: Short description in the requested language |
| `short_description_i18n` | no | API: A map of translated strings, by language code |
| `state` | no | — |
| `tags` | no | API: List of tags associated with this catalog |
| `updated` | no | — |
| `url` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cm_object_invalid_catalog_id` | `catalog_id` |
| `ibm_cm_object_invalid_data` | `data` |
| `ibm_cm_object_invalid_kind` | `kind` |
| `ibm_cm_object_invalid_label` | `label` |
| `ibm_cm_object_invalid_label_i18n` | `label_i18n` |
| `ibm_cm_object_invalid_name` | `name` |
| `ibm_cm_object_invalid_parent_id` | `parent_id` |
| `ibm_cm_object_invalid_short_description` | `short_description` |
| `ibm_cm_object_invalid_short_description_i18n` | `short_description_i18n` |
| `ibm_cm_object_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_cm_object" "test" {
  catalog_id = "01234567-89ab-cdef-0123-456789abcdef"
  kind = "example"
  name = "example"
}
```
