# TFLint Ruleset for terraform-provider-ibm

[![Build Status](https://github.com/IBM/tflint-ruleset-ibm/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/IBM/tflint-ruleset-ibm/actions)
[![GitHub release](https://img.shields.io/github/release/IBM/tflint-ruleset-ibm.svg)](https://github.com/IBM/tflint-ruleset-ibm/releases/latest)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-blue.svg)](LICENSE)

TFLint ruleset plugin for the IBM Cloud Terraform Provider

This ruleset focuses on possible errors about IBM Cloud resources and data sources. In other words, the rules warn against code that might fail when running `terraform apply`.

> **Scope: this is not a security scanner.** Every rule checks a configuration against the provider schema and the IBM Cloud API (required arguments, allowed values, ranges and formats). No rule judges security posture: open security groups such as `0.0.0.0/0`, broad IAM policies, missing encryption keys, public endpoints or hardcoded `ibmcloud_api_key` values all pass. Run a policy scanner such as [Trivy](https://github.com/aquasecurity/trivy) or [Checkov](https://github.com/bridgecrewio/checkov) alongside this ruleset.

## Requirements

- TFLint v0.46+
- Go 1.26 or later (only to build from source; `go.mod` pins the toolchain)

## Installation

After configuring [TFLint](https://github.com/terraform-linters/tflint/blob/master/docs/user-guide/config.md), you can install the plugin by adding a config to `.tflint.hcl` and running `tflint --init`:

```hcl
plugin "ibm" {
    enabled = true
    version = "0.1.0"
    source  = "github.com/IBM/tflint-ruleset-ibm"

    signing_key = <<-KEY
    (paste the IBM release signing public key here; see "Verifying releases")
    KEY
}
```

Always pin an exact `version`. With `signing_key` set, `tflint --init` refuses a release whose `checksums.txt.sig` does not verify, so a tampered download cannot be installed.

For more information about configuration options for this ruleset plugin, see [Plugin Configuration](docs/configuration.md).

## Getting Started

Terraform is a great tool for Infrastructure as Code. However, many of these tools don't validate provider-specific issues. For example, see the following configuration file:

```hcl
resource "ibm_is_vpc" "example" {
  name                      = "My VPC"      # invalid name
  address_prefix_management = "automatic"   # invalid value
}
```

`My VPC` is not a valid VPC name, and `automatic` is not an allowed value for `address_prefix_management`. An error will occur when you run `terraform apply`. But `terraform validate` and `terraform plan` cannot find this possible error in advance. That's because it's an IBM Cloud provider-specific issue and it's valid as the Terraform Language.

The goal of this ruleset is to find such errors:

![demo](docs/assets/demo.gif)

By running TFLint with this ruleset in advance, you can fix the problem before the error occurs in production CI/CD pipelines.

## Rules

There are 1500+ type-level rules and 10,000+ argument rules available, covering every managed resource and data source in the IBM Cloud Terraform provider. See the [Rules documentation page](docs/rules/README.md) for a complete list.

Each Terraform type has a schema rule (for example `ibm_is_vpc`) plus one rule per writable argument (for example `ibm_is_vpc_invalid_name`) so checks can be enabled or disabled independently. Data-source rules are named `<type>_data` so they do not collide with the managed-resource rule of the same type.

## Verifying releases

Releases are built by the `release` workflow from a tag on `main`, with pinned build tooling and reproducible build flags. Each release provides:

- `checksums.txt` and its detached GPG signature `checksums.txt.sig`. The public key is published with each release and in the repository's security advisories page; put it in `signing_key` as shown above.
- An SBOM for each archive.
- A GitHub build provenance attestation for every archive and binary.

To check a downloaded archive by hand:

```
$ gpg --verify checksums.txt.sig checksums.txt
$ sha256sum --check --ignore-missing checksums.txt
$ gh attestation verify tflint-ruleset-ibm_linux_amd64.zip --repo IBM/tflint-ruleset-ibm
```

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

`make install` replaces any copy already in `~/.tflint.d/plugins`. Build from a release tag you have verified rather than from an arbitrary branch.

Note that if you install the plugin with `make install`, you must omit the `version` and `source` attributes in `.tflint.hcl`:

```hcl
plugin "ibm" {
    enabled = true
}
```

## Regenerating rules

Rules are generated from a local checkout of the IBM Cloud Terraform provider (treated as a disposable data source; never vendored) and from published IBM Cloud OpenAPI documents:

```
$ make extract-all
$ make overlay
$ make generate-all
$ make test build
```

End-to-end:

```
$ make all
```

`make coverage` compares every provider map entry to extracted schema, generated rules, tests, fixtures, registration, and docs. It writes [COVERAGE.md](COVERAGE.md) and fails if anything is missing or stale.

Set `PACKAGES=vpc` to limit extract/generate to one service package when you need a tight loop.

## Add a new rule

Do not hand-edit files matching `rules/*_gen.go` or `docs/rules/*.md`. Change the extractor, the OpenAPI overlay, or the generator, then regenerate. Follow the instructions in [CONTRIBUTING.md](CONTRIBUTING.md) and open a new pull request.

## License

This project is licensed under the Mozilla Public License 2.0. See [LICENSE](LICENSE) and [NOTICE](NOTICE).
