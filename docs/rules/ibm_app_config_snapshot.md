# ibm_app_config_snapshot

Schema validation for the ibm_app_config_snapshot resource.

This rule validates `ibm_app_config_snapshot` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `CreateGitConfig` (`app-configuration`).

## Configuration

```hcl
rule "ibm_app_config_snapshot" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `action` | no | — |
| `collection` | no | — |
| `collection_id` | yes | API: Collection Id. |
| `created_time` | no | — |
| `environment` | no | — |
| `environment_id` | yes | API: Environment Id. |
| `git_branch` | yes | API: Branch name to which you need to write or update the configuration. Just provide the branch name, do not provide any /(slashes) in the beginning or at the end of the branch name. Note make sure branch exists in your repository. |
| `git_config_id` | yes | API: Git config id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `git_config_name` | yes | API: Git config name. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only |
| `git_file_path` | yes | API: Git file path, this is a path where your configuration file will be written. The path must contain the file name with `json` extension. We only create or update `json` extension file. Note do not provide any /(slashes) in the beginning or at the end of the file path. |
| `git_token` | yes | API: Git token, this needs to be provided with enough permission to write and update the file |
| `git_url` | yes | API: Git url which will be used to connect to the github account. The url must be formed in this format, https://api.github.com/repos/{owner}/{repo_name} for the personal git account. If you are using the organization account then url must be in this format https://github.{organization_name}.com/api/v3/repos/{owner}/{repo_name} . Note do not provide /(slash) in the beginning or at the end of the url. |
| `guid` | yes | — |
| `href` | no | — |
| `updated_time` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_app_config_snapshot_invalid_action` | `action` |
| `ibm_app_config_snapshot_invalid_collection_id` | `collection_id` |
| `ibm_app_config_snapshot_invalid_environment_id` | `environment_id` |
| `ibm_app_config_snapshot_invalid_git_branch` | `git_branch` |
| `ibm_app_config_snapshot_invalid_git_config_id` | `git_config_id` |
| `ibm_app_config_snapshot_invalid_git_config_name` | `git_config_name` |
| `ibm_app_config_snapshot_invalid_git_file_path` | `git_file_path` |
| `ibm_app_config_snapshot_invalid_git_token` | `git_token` |
| `ibm_app_config_snapshot_invalid_git_url` | `git_url` |
| `ibm_app_config_snapshot_invalid_guid` | `guid` |

## Example

### Valid

```hcl
variable "git_token" {
  sensitive = true
}

resource "ibm_app_config_snapshot" "test" {
  collection_id = "01234567-89ab-cdef-0123-456789abcdef"
  environment_id = "01234567-89ab-cdef-0123-456789abcdef"
  git_branch = "example"
  git_config_id = "01234567-89ab-cdef-0123-456789abcdef"
  git_config_name = "aaaaaaa"
  git_file_path = "example"
  git_token = var.git_token
  git_url = "example"
  guid = "example"
}
```
