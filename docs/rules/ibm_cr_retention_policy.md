# ibm_cr_retention_policy

Schema validation for the ibm_cr_retention_policy resource.

This rule validates `ibm_cr_retention_policy` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `/api/v1/retentions:post` (`container-registry`).

## Configuration

```hcl
rule "ibm_cr_retention_policy" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `images_per_repo` | yes | API: Determines how many images are retained in each repository when the retention policy is processed. The default value -1 denotes 'Unlimited' (all images are retained). Additional images might be retained in each repository if the retention setting 'deletion_threshold_days' is used. |
| `namespace` | yes | API: The namespace to which the retention policy is attached. |
| `retain_untagged` | no | API: Determines whether untagged images are retained when the retention policy is processed. The value is false by default, which means that untagged images can be deleted when the policy runs. |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_cr_retention_policy_invalid_images_per_repo` | `images_per_repo` |
| `ibm_cr_retention_policy_invalid_namespace` | `namespace` |
| `ibm_cr_retention_policy_invalid_retain_untagged` | `retain_untagged` |

## Example

### Valid

```hcl
resource "ibm_cr_retention_policy" "test" {
  images_per_repo = 1
  namespace = "example"
}
```
