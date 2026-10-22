# ibm_pn_application_chrome

Schema validation for the ibm_pn_application_chrome resource.

This rule validates `ibm_pn_application_chrome` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_pn_application_chrome" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `guid` | yes | — |
| `server_key` | yes | — |
| `web_site_url` | yes | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_pn_application_chrome_invalid_guid` | `guid` |
| `ibm_pn_application_chrome_invalid_server_key` | `server_key` |
| `ibm_pn_application_chrome_invalid_web_site_url` | `web_site_url` |

## Example

### Valid

```hcl
resource "ibm_pn_application_chrome" "test" {
  guid = "example"
  server_key = "example"
  web_site_url = "example"
}
```
