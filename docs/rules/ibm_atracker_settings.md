# ibm_atracker_settings

Schema validation for the ibm_atracker_settings resource.

This rule validates `ibm_atracker_settings` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SettingsPrototype` (`atracker`).

## Configuration

```hcl
rule "ibm_atracker_settings" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `api_version` | no | — |
| `default_targets` | no | pattern `^[a-zA-Z0-9 -]` |
| `message` | no | — |
| `metadata_region_backup` | no | pattern `^[a-zA-Z0-9 -_]` |
| `metadata_region_primary` | yes | pattern `^[a-zA-Z0-9 -_]` |
| `permitted_target_regions` | no | pattern `^[a-zA-Z0-9 -_]` |
| `private_api_endpoint_only` | yes | API: If you set this true then you cannot access api through public network. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_atracker_settings_invalid_default_targets` | `default_targets` |
| `ibm_atracker_settings_invalid_metadata_region_backup` | `metadata_region_backup` |
| `ibm_atracker_settings_invalid_metadata_region_primary` | `metadata_region_primary` |
| `ibm_atracker_settings_invalid_permitted_target_regions` | `permitted_target_regions` |
| `ibm_atracker_settings_invalid_private_api_endpoint_only` | `private_api_endpoint_only` |

## Example

### Valid

```hcl
resource "ibm_atracker_settings" "test" {
  metadata_region_primary = "example-name"
  private_api_endpoint_only = true
}
```
