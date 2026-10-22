# Configuration

This plugin can take additional configuration in the `plugin` block.

```hcl
plugin "ibm" {
    // Plugin common attributes
    enabled = true
    version = "0.1.0"
    source  = "github.com/IBM/tflint-ruleset-ibm"

    region     = "us-south"
    deep_check = false
}
```

## `region`

Default: `us-south`

Hint used by future deep-check rules. Schema-validation rules ignore this value.

## `deep_check`

Default: `false`

Reserved for future checks that need read access to a target IBM Cloud account. No rule uses it today, so setting it to `true` adds no checks. Schema-validation rules do not need IBM Cloud credentials.

Never put IBM Cloud API keys in `.tflint.hcl`. When deep checks are added they will read credentials from the environment (for example `IC_API_KEY`) and need only read-only IAM access.

## Disabling a rule

Disabling a type rule (for example `ibm_is_vpc`) turns off every schema check for that type, and `plugin "ibm" { enabled = false }` turns off the whole ruleset. In CI, keep `.tflint.hcl` under code review so checks are not switched off silently.

```hcl
rule "ibm_is_vpc" {
    enabled = false
}

rule "ibm_is_vpc_invalid_name" {
    enabled = false
}

rule "ibm_is_vpc_data" {
    enabled = false
}
```
