# ibm_is_image_export_jobs_data

Schema validation for the ibm_is_image_export_jobs data source.

This rule validates `ibm_is_image_export_jobs` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `InstanceByImageAndCatalogOffering` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_image_export_jobs_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `export_jobs` | no | — |
| `image` | yes | API: Identifies an image by a unique property. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_image_export_jobs_data_invalid_image` | `image` |

## Example

### Valid

```hcl
data "ibm_is_image_export_jobs" "test" {
  image = "example"
}
```
