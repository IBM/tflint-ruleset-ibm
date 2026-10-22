---
name: Bug report
about: Report a false positive, missed error, or plugin crash
labels: bug
---

<!--- Please keep this note for the community --->

### Community Note

* Please vote on this issue by adding a 👍 [reaction](https://blog.github.com/2016-03-10-add-reactions-to-pull-requests-issues-and-comments/) to the original issue to help the community and maintainers prioritize this request
* Please do not leave "+1" or other comments that do not add relevant new information or questions

<!--- Thank you for keeping this note for the community --->

### TFLint and plugin versions

```
tflint --version
```

### .tflint.hcl

```hcl
plugin "ibm" {
  enabled = true
}
```

### Terraform configuration

<!--- Remove API keys (for example ibmcloud_api_key), passwords, tokens, private keys and account IDs before posting. tflint output can echo argument values. --->

```hcl
# minimal reproduction, secrets removed
```

### Expected behavior

### Actual behavior

Include the full `tflint` output.
