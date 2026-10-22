# ibm_cdn

Schema validation for the ibm_cdn resource.

This rule validates `ibm_cdn` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_cdn" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `bucket_name` | no | — |
| `cache_key_query_rule` | no | one of: `include-all`, `ignore-all`, `ignore: space separated query-args`, `include: space separated query-args` |
| `certificate_type` | no | one of: `SHARED_SAN_CERT`, `WILDCARD_CERT` |
| `cname` | no | — |
| `file_extension` | no | — |
| `header` | no | — |
| `host_name` | yes | — |
| `http_port` | no | range 1–65535; port format |
| `https_port` | no | range 1–65535; port format |
| `origin_address` | yes | — |
| `origin_type` | no | one of: `HOST_SERVER`, `OBJECT_STORAGE` |
| `path` | no | — |
| `performance_configuration` | no | — |
| `protocol` | no | one of: `HTTP`, `HTTPS`, `HTTP_AND_HTTPS` |
| `respect_headers` | no | — |
| `status` | no | — |
| `vendor_name` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cdn_invalid_bucket_name` | `bucket_name` |
| `ibm_cdn_invalid_cache_key_query_rule` | `cache_key_query_rule` |
| `ibm_cdn_invalid_certificate_type` | `certificate_type` |
| `ibm_cdn_invalid_cname` | `cname` |
| `ibm_cdn_invalid_file_extension` | `file_extension` |
| `ibm_cdn_invalid_header` | `header` |
| `ibm_cdn_invalid_host_name` | `host_name` |
| `ibm_cdn_invalid_http_port` | `http_port` |
| `ibm_cdn_invalid_https_port` | `https_port` |
| `ibm_cdn_invalid_origin_address` | `origin_address` |
| `ibm_cdn_invalid_origin_type` | `origin_type` |
| `ibm_cdn_invalid_path` | `path` |
| `ibm_cdn_invalid_performance_configuration` | `performance_configuration` |
| `ibm_cdn_invalid_protocol` | `protocol` |
| `ibm_cdn_invalid_respect_headers` | `respect_headers` |
| `ibm_cdn_invalid_vendor_name` | `vendor_name` |

## Example

### Valid

```hcl
resource "ibm_cdn" "test" {
  host_name = "example"
  origin_address = "example"
}
```
