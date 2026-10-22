# ibm_satellite_link

Schema validation for the ibm_satellite_link resource.

This rule validates `ibm_satellite_link` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_satellite_link" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `crn` | yes | crn format |
| `description` | no | — |
| `last_change` | no | — |
| `location` | yes | — |
| `performance` | no | — |
| `satellite_link_host` | no | — |
| `status` | no | — |
| `ws_endpoint` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_satellite_link_invalid_crn` | `crn` |
| `ibm_satellite_link_invalid_location` | `location` |
| `ibm_satellite_link_invalid_ws_endpoint` | `ws_endpoint` |

## Example

### Valid

```hcl
resource "ibm_satellite_link" "test" {
  crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  location = "example"
}
```
