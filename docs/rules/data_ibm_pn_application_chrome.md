# ibm_pn_application_chrome_data

Schema validation for the ibm_pn_application_chrome data source.

This rule validates `ibm_pn_application_chrome` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pn_application_chrome_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `guid` | yes | — |
| `server_key` | no | — |
| `web_site_url` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pn_application_chrome_data_invalid_guid` | `guid` |

## Example

### Valid

```hcl
data "ibm_pn_application_chrome" "test" {
  guid = "example"
}
```
