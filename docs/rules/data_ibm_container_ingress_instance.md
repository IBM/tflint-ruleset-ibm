# ibm_container_ingress_instance_data

Schema validation for the ibm_container_ingress_instance data source.

This rule validates `ibm_container_ingress_instance` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `DisablePodSecurityPolicyRequest` (`kubernetes-containers-v1-v2`).

## Configuration

```hcl
rule "ibm_container_ingress_instance_data" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `cluster` | yes | API: The cluster value. |
| `instance_crn` | no | — |
| `instance_name` | yes | — |
| `instance_type` | no | — |
| `is_default` | no | — |
| `secret_group_id` | no | — |
| `secret_group_name` | no | — |
| `status` | no | — |
| `user_managed` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_container_ingress_instance_data_invalid_cluster` | `cluster` |
| `ibm_container_ingress_instance_data_invalid_instance_name` | `instance_name` |

## Example

### Valid

```hcl
data "ibm_container_ingress_instance" "test" {
  cluster = "aaaaaaa"
  instance_name = "example"
}
```
