# ibm_app_config_segment_data

Reads ibm_app_config_segment. Segment name.

This rule validates `ibm_app_config_segment` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Segment` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_segment_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_time` | no | — |
| `description` | no | — |
| `features` | no | — |
| `guid` | yes | — |
| `href` | no | — |
| `includes` | no | — |
| `name` | no | — |
| `properties` | no | — |
| `rules` | no | — |
| `segment_id` | yes | API: Segment id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `tags` | no | — |
| `updated_time` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_segment_data_invalid_guid` | `guid` |
| `ibm_app_config_segment_data_invalid_includes` | `includes` |
| `ibm_app_config_segment_data_invalid_segment_id` | `segment_id` |

## Example

### Valid

```hcl
data "ibm_app_config_segment" "test" {
  guid = "example"
  segment_id = "01234567-89ab-cdef-0123-456789abcdef"
}
```
