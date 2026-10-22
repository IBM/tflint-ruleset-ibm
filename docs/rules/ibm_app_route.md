# ibm_app_route

Schema validation for the ibm_app_route resource.

This rule validates `ibm_app_route` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_app_route" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `domain_guid` | yes | — |
| `host` | no | — |
| `path` | no | — |
| `port` | no | range 1–65535; port format |
| `space_guid` | yes | — |
| `tags` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_route_invalid_domain_guid` | `domain_guid` |
| `ibm_app_route_invalid_host` | `host` |
| `ibm_app_route_invalid_path` | `path` |
| `ibm_app_route_invalid_port` | `port` |
| `ibm_app_route_invalid_space_guid` | `space_guid` |
| `ibm_app_route_invalid_tags` | `tags` |

## Example

### Valid

```hcl
resource "ibm_app_route" "test" {
  domain_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
  space_guid = "crn:v1:bluemix:public:cloud-object-storage:global:a/0123456789abcdef0123456789abcdef:01234567-89ab-cdef-0123-456789abcdef::"
}
```
