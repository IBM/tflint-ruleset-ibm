# ibm_sds_host

Manages ibm_sds_host. Unique name of the host.

This rule validates `ibm_sds_host` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `Host` (`block-storage`).

## Configuration

```hcl
rule "ibm_sds_host" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `created_at` | no | — |
| `href` | no | — |
| `name` | no | pattern `^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `nqn` | yes | pattern `^nqn\.\d{4}-\d{2}\.[a-z0-9-]+(?:\.[a-z0-9-]+)*:[a-zA-Z0-9.\-:]+$` |
| `sds_endpoint` | no | — |
| `volume_mappings` (block) | no | nested block |
| `volume_mappings.href` | yes | uri format |
| `volume_mappings.id` | yes | pattern `^[-0-9a-z_]+$` |
| `volume_mappings.status` | yes | pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `volume_mappings.subsystem_nqn` | no | pattern `^nqn\.\d{4}-\d{2}\.[a-z0-9-]+(?:\.[a-z0-9-]+)*:[a-zA-Z0-9.\-:]+$` |
| `volume_mappings.gateways` (block) | no | nested block |
| `volume_mappings.host` (block) | no | nested block |
| `volume_mappings.host.id` | yes | pattern `^[-0-9a-z_]+$` |
| `volume_mappings.host.name` | yes | pattern `^[a-z][a-z0-9]*(-[a-z0-9]+)*$` |
| `volume_mappings.host.nqn` | yes | pattern `^nqn\.\d{4}-\d{2}\.[a-z0-9-]+(?:\.[a-z0-9-]+)*:[a-zA-Z0-9.\-:]+$` |
| `volume_mappings.namespace` (block) | no | nested block |
| `volume_mappings.namespace.id` | no | range 1–32 |
| `volume_mappings.namespace.uuid` | no | pattern `^[-0-9a-z_]+$` |
| `volume_mappings.storage_identifier` (block) | no | nested block |
| `volume_mappings.storage_identifier.namespace_id` | yes | — |
| `volume_mappings.storage_identifier.namespace_uuid` | yes | — |
| `volume_mappings.storage_identifier.subsystem_nqn` | yes | pattern `^nqn\.\d{4}-\d{2}\.[a-z0-9-]+(?:\.[a-z0-9-]+)*:[a-zA-Z0-9.\-:]+$` |
| `volume_mappings.storage_identifier.gateways` (block) | yes | nested block |
| `volume_mappings.volume` (block) | no | nested block |
| `volume_mappings.volume.id` | yes | pattern `^[-0-9a-z_]+$` |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_sds_host_invalid_name` | `name` |
| `ibm_sds_host_invalid_nqn` | `nqn` |
| `ibm_sds_host_invalid_sds_endpoint` | `sds_endpoint` |
| `ibm_sds_host_invalid_volume_mappings__href` | `volume_mappings.href` |
| `ibm_sds_host_invalid_volume_mappings__id` | `volume_mappings.id` |
| `ibm_sds_host_invalid_volume_mappings__status` | `volume_mappings.status` |
| `ibm_sds_host_invalid_volume_mappings__subsystem_nqn` | `volume_mappings.subsystem_nqn` |
| `ibm_sds_host_invalid_volume_mappings__host__id` | `volume_mappings.host.id` |
| `ibm_sds_host_invalid_volume_mappings__host__name` | `volume_mappings.host.name` |
| `ibm_sds_host_invalid_volume_mappings__host__nqn` | `volume_mappings.host.nqn` |
| `ibm_sds_host_invalid_volume_mappings__namespace__id` | `volume_mappings.namespace.id` |
| `ibm_sds_host_invalid_volume_mappings__namespace__uuid` | `volume_mappings.namespace.uuid` |
| `ibm_sds_host_invalid_volume_mappings__storage_identifier__namespace_id` | `volume_mappings.storage_identifier.namespace_id` |
| `ibm_sds_host_invalid_volume_mappings__storage_identifier__namespace_uuid` | `volume_mappings.storage_identifier.namespace_uuid` |
| `ibm_sds_host_invalid_volume_mappings__storage_identifier__subsystem_nqn` | `volume_mappings.storage_identifier.subsystem_nqn` |
| `ibm_sds_host_invalid_volume_mappings__volume__id` | `volume_mappings.volume.id` |

## Example

### Valid

```hcl
resource "ibm_sds_host" "test" {
  nqn = "nqn.2014-08.org.nvmexpress:uuid"
}
```
