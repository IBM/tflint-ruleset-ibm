# ibm_ssl_certificate

Schema validation for the ibm_ssl_certificate resource.

This rule validates `ibm_ssl_certificate` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

## Configuration

```hcl
rule "ibm_ssl_certificate" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `administrative_address_same_as_organization_flag` | no | — |
| `administrative_contact_same_as_technical_flag` | no | — |
| `billing_address_same_as_organization_flag` | no | — |
| `billing_contact_same_as_technical_flag` | no | — |
| `certificate_signing_request` | yes | — |
| `order_approver_email_address` | yes | — |
| `renewal_flag` | no | — |
| `server_count` | yes | — |
| `server_type` | yes | — |
| `ssl_type` | yes | — |
| `technical_contact_same_as_org_address_flag` | no | — |
| `validity_months` | yes | — |
| `administrative_contact` (block) | no | nested block |
| `administrative_contact.admin_email_address` | no | — |
| `administrative_contact.admin_fax_number` | no | — |
| `administrative_contact.admin_first_name` | no | — |
| `administrative_contact.admin_last_name` | no | — |
| `administrative_contact.admin_organization_name` | no | — |
| `administrative_contact.admin_phone_number` | no | — |
| `administrative_contact.admin_title` | no | — |
| `administrative_contact.admin_address` (block) | no | nested block |
| `administrative_contact.admin_address.admin_address_line1` | no | — |
| `administrative_contact.admin_address.admin_address_line2` | no | — |
| `administrative_contact.admin_address.admin_city` | no | — |
| `administrative_contact.admin_address.admin_country_code` | no | — |
| `administrative_contact.admin_address.admin_postal_code` | no | — |
| `administrative_contact.admin_address.admin_state` | no | — |
| `billing_contact` (block) | no | nested block |
| `billing_contact.billing_email_address` | no | — |
| `billing_contact.billing_fax_number` | no | — |
| `billing_contact.billing_first_name` | no | — |
| `billing_contact.billing_last_name` | no | — |
| `billing_contact.billing_organization_name` | no | — |
| `billing_contact.billing_phone_number` | no | — |
| `billing_contact.billing_title` | no | — |
| `billing_contact.billing_address` (block) | no | nested block |
| `billing_contact.billing_address.billing_address_line1` | no | — |
| `billing_contact.billing_address.billing_address_line2` | no | — |
| `billing_contact.billing_address.billing_city` | no | — |
| `billing_contact.billing_address.billing_country_code` | no | — |
| `billing_contact.billing_address.billing_postal_code` | no | — |
| `billing_contact.billing_address.billing_state` | no | — |
| `organization_information` (block) | yes | nested block |
| `organization_information.org_fax_number` | no | — |
| `organization_information.org_organization_name` | yes | — |
| `organization_information.org_phone_number` | yes | — |
| `organization_information.org_address` (block) | yes | nested block |
| `organization_information.org_address.org_address_line1` | yes | — |
| `organization_information.org_address.org_address_line2` | no | — |
| `organization_information.org_address.org_city` | yes | — |
| `organization_information.org_address.org_country_code` | yes | — |
| `organization_information.org_address.org_postal_code` | yes | — |
| `organization_information.org_address.org_state` | yes | — |
| `technical_contact` (block) | yes | nested block |
| `technical_contact.tech_email_address` | yes | — |
| `technical_contact.tech_fax_number` | no | — |
| `technical_contact.tech_first_name` | yes | — |
| `technical_contact.tech_last_name` | yes | — |
| `technical_contact.tech_organization_name` | yes | — |
| `technical_contact.tech_phone_number` | yes | — |
| `technical_contact.tech_title` | yes | — |
| `technical_contact.tech_address` (block) | no | nested block |
| `technical_contact.tech_address.tech_address_line1` | no | — |
| `technical_contact.tech_address.tech_address_line2` | no | — |
| `technical_contact.tech_address.tech_city` | no | — |
| `technical_contact.tech_address.tech_country_code` | no | — |
| `technical_contact.tech_address.tech_postal_code` | no | — |
| `technical_contact.tech_address.tech_state` | no | — |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_ssl_certificate_invalid_administrative_address_same_as_organization_flag` | `administrative_address_same_as_organization_flag` |
| `ibm_ssl_certificate_invalid_administrative_contact_same_as_technical_flag` | `administrative_contact_same_as_technical_flag` |
| `ibm_ssl_certificate_invalid_billing_address_same_as_organization_flag` | `billing_address_same_as_organization_flag` |
| `ibm_ssl_certificate_invalid_billing_contact_same_as_technical_flag` | `billing_contact_same_as_technical_flag` |
| `ibm_ssl_certificate_invalid_certificate_signing_request` | `certificate_signing_request` |
| `ibm_ssl_certificate_invalid_order_approver_email_address` | `order_approver_email_address` |
| `ibm_ssl_certificate_invalid_renewal_flag` | `renewal_flag` |
| `ibm_ssl_certificate_invalid_server_count` | `server_count` |
| `ibm_ssl_certificate_invalid_server_type` | `server_type` |
| `ibm_ssl_certificate_invalid_ssl_type` | `ssl_type` |
| `ibm_ssl_certificate_invalid_technical_contact_same_as_org_address_flag` | `technical_contact_same_as_org_address_flag` |
| `ibm_ssl_certificate_invalid_validity_months` | `validity_months` |
| `ibm_ssl_certificate_invalid_administrative_contact__admin_email_address` | `administrative_contact.admin_email_address` |
| `ibm_ssl_certificate_invalid_administrative_contact__admin_fax_number` | `administrative_contact.admin_fax_number` |
| `ibm_ssl_certificate_invalid_administrative_contact__admin_first_name` | `administrative_contact.admin_first_name` |
| `ibm_ssl_certificate_invalid_administrative_contact__admin_last_name` | `administrative_contact.admin_last_name` |
| `ibm_ssl_certificate_invalid_administrative_contact__admin_organization_name` | `administrative_contact.admin_organization_name` |
| `ibm_ssl_certificate_invalid_administrative_contact__admin_phone_number` | `administrative_contact.admin_phone_number` |
| `ibm_ssl_certificate_invalid_administrative_contact__admin_title` | `administrative_contact.admin_title` |
| `ibm_ssl_certificate_invalid_administrative_contact__admin_address__admin_address_line1` | `administrative_contact.admin_address.admin_address_line1` |
| `ibm_ssl_certificate_invalid_administrative_contact__admin_address__admin_address_line2` | `administrative_contact.admin_address.admin_address_line2` |
| `ibm_ssl_certificate_invalid_administrative_contact__admin_address__admin_city` | `administrative_contact.admin_address.admin_city` |
| `ibm_ssl_certificate_invalid_administrative_contact__admin_address__admin_country_code` | `administrative_contact.admin_address.admin_country_code` |
| `ibm_ssl_certificate_invalid_administrative_contact__admin_address__admin_postal_code` | `administrative_contact.admin_address.admin_postal_code` |
| `ibm_ssl_certificate_invalid_administrative_contact__admin_address__admin_state` | `administrative_contact.admin_address.admin_state` |
| `ibm_ssl_certificate_invalid_billing_contact__billing_email_address` | `billing_contact.billing_email_address` |
| `ibm_ssl_certificate_invalid_billing_contact__billing_fax_number` | `billing_contact.billing_fax_number` |
| `ibm_ssl_certificate_invalid_billing_contact__billing_first_name` | `billing_contact.billing_first_name` |
| `ibm_ssl_certificate_invalid_billing_contact__billing_last_name` | `billing_contact.billing_last_name` |
| `ibm_ssl_certificate_invalid_billing_contact__billing_organization_name` | `billing_contact.billing_organization_name` |
| `ibm_ssl_certificate_invalid_billing_contact__billing_phone_number` | `billing_contact.billing_phone_number` |
| `ibm_ssl_certificate_invalid_billing_contact__billing_title` | `billing_contact.billing_title` |
| `ibm_ssl_certificate_invalid_billing_contact__billing_address__billing_address_line1` | `billing_contact.billing_address.billing_address_line1` |
| `ibm_ssl_certificate_invalid_billing_contact__billing_address__billing_address_line2` | `billing_contact.billing_address.billing_address_line2` |
| `ibm_ssl_certificate_invalid_billing_contact__billing_address__billing_city` | `billing_contact.billing_address.billing_city` |
| `ibm_ssl_certificate_invalid_billing_contact__billing_address__billing_country_code` | `billing_contact.billing_address.billing_country_code` |
| `ibm_ssl_certificate_invalid_billing_contact__billing_address__billing_postal_code` | `billing_contact.billing_address.billing_postal_code` |
| `ibm_ssl_certificate_invalid_billing_contact__billing_address__billing_state` | `billing_contact.billing_address.billing_state` |
| `ibm_ssl_certificate_invalid_organization_information__org_fax_number` | `organization_information.org_fax_number` |
| `ibm_ssl_certificate_invalid_organization_information__org_organization_name` | `organization_information.org_organization_name` |
| `ibm_ssl_certificate_invalid_organization_information__org_phone_number` | `organization_information.org_phone_number` |
| `ibm_ssl_certificate_invalid_organization_information__org_address__org_address_line1` | `organization_information.org_address.org_address_line1` |
| `ibm_ssl_certificate_invalid_organization_information__org_address__org_address_line2` | `organization_information.org_address.org_address_line2` |
| `ibm_ssl_certificate_invalid_organization_information__org_address__org_city` | `organization_information.org_address.org_city` |
| `ibm_ssl_certificate_invalid_organization_information__org_address__org_country_code` | `organization_information.org_address.org_country_code` |
| `ibm_ssl_certificate_invalid_organization_information__org_address__org_postal_code` | `organization_information.org_address.org_postal_code` |
| `ibm_ssl_certificate_invalid_organization_information__org_address__org_state` | `organization_information.org_address.org_state` |
| `ibm_ssl_certificate_invalid_technical_contact__tech_email_address` | `technical_contact.tech_email_address` |
| `ibm_ssl_certificate_invalid_technical_contact__tech_fax_number` | `technical_contact.tech_fax_number` |
| `ibm_ssl_certificate_invalid_technical_contact__tech_first_name` | `technical_contact.tech_first_name` |
| `ibm_ssl_certificate_invalid_technical_contact__tech_last_name` | `technical_contact.tech_last_name` |
| `ibm_ssl_certificate_invalid_technical_contact__tech_organization_name` | `technical_contact.tech_organization_name` |
| `ibm_ssl_certificate_invalid_technical_contact__tech_phone_number` | `technical_contact.tech_phone_number` |
| `ibm_ssl_certificate_invalid_technical_contact__tech_title` | `technical_contact.tech_title` |
| `ibm_ssl_certificate_invalid_technical_contact__tech_address__tech_address_line1` | `technical_contact.tech_address.tech_address_line1` |
| `ibm_ssl_certificate_invalid_technical_contact__tech_address__tech_address_line2` | `technical_contact.tech_address.tech_address_line2` |
| `ibm_ssl_certificate_invalid_technical_contact__tech_address__tech_city` | `technical_contact.tech_address.tech_city` |
| `ibm_ssl_certificate_invalid_technical_contact__tech_address__tech_country_code` | `technical_contact.tech_address.tech_country_code` |
| `ibm_ssl_certificate_invalid_technical_contact__tech_address__tech_postal_code` | `technical_contact.tech_address.tech_postal_code` |
| `ibm_ssl_certificate_invalid_technical_contact__tech_address__tech_state` | `technical_contact.tech_address.tech_state` |

## Example

### Valid

```hcl
resource "ibm_ssl_certificate" "test" {
  certificate_signing_request = "example"
  order_approver_email_address = "example"
  server_count = 1
  server_type = "example"
  ssl_type = "example"
  validity_months = 1
  organization_information {
    org_organization_name = "example"
    org_phone_number = "example"
    org_address {
      org_address_line1 = "example"
      org_city = "example"
      org_country_code = "example"
      org_postal_code = "example"
      org_state = "example"
    }
  }
  technical_contact {
    tech_email_address = "example"
    tech_first_name = "example"
    tech_last_name = "example"
    tech_organization_name = "example"
    tech_phone_number = "example"
    tech_title = "example"
  }
}
```
