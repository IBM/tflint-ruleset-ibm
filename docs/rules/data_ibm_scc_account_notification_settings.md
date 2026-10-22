# ibm_scc_account_notification_settings_data

ibm_scc_account_notification_settings data source is registered with an empty schema (deprecated or schema-less).

This rule validates `ibm_scc_account_notification_settings` data sources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_scc_account_notification_settings_data" {
  enabled = true
}
```

## Checks

This data source exposes no user-configurable arguments beyond provider defaults.
