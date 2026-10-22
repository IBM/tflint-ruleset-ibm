# ibm_is_security_group_rule

Manages ibm_is_security_group_rule. The name for this security group rule.

This rule validates `ibm_is_security_group_rule` resources against the IBM Cloud provider schema and the published IBM Cloud API: required arguments, allowed values, numeric ranges, and format constraints.

OpenAPI schema: `SecurityGroupRulePatch` (`vpc-latest`).

## Configuration

```hcl
rule "ibm_is_security_group_rule" {
  enabled = true
}
```

## Checks

| Argument | Required | Constraints |
| --- | --- | --- |
| `code` | no | range 0–255 |
| `direction` | yes | one of: `inbound`, `outbound`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `group` | yes | — |
| `ip_version` | no | one of: `ipv4`; pattern `^[a-z][a-z0-9]*(_[a-z0-9]+)*$` |
| `local` | no | API: The local IP address or range of local IP addresses to which this rule will allow inbound traffic (or from which, for outbound traffic). Can be specified as an IP address or a CIDR block. Specify a CIDR block of `0.0.0.0/0` to allow traffic to all local IP addresses (or from all local IP addresses, for outbound rules). |
| `name` | no | pattern `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$` |
| `port_max` | no | range 1–65535 |
| `port_min` | no | range 1–65535 |
| `protocol` | no | one of: `tcp`, `udp`, `icmp`, `ah`, `any`, `esp`, `gre`, `icmp_tcp_udp`, `ip_in_ip`, `l2tp`, `number_10`, `number_100`, `number_101`, `number_102`, `number_103`, `number_104`, `number_105`, `number_106`, `number_107`, `number_108`, `number_109`, `number_11`, `number_110`, `number_111`, `number_113`, `number_114`, `number_116`, `number_117`, `number_118`, `number_119`, `number_12`, `number_120`, `number_121`, `number_122`, `number_123`, `number_124`, `number_125`, `number_126`, `number_127`, `number_128`, `number_129`, `number_13`, `number_130`, `number_131`, `number_133`, `number_134`, `number_136`, `number_137`, `number_138`, `number_139`, `number_14`, `number_140`, `number_141`, `number_142`, `number_143`, `number_144`, `number_145`, `number_146`, `number_147`, `number_148`, `number_149`, `number_15`, `number_150`, `number_151`, `number_152`, `number_153`, `number_154`, `number_155`, `number_156`, `number_157`, `number_158`, `number_159`, `number_16`, `number_160`, `number_161`, `number_162`, `number_163`, `number_164`, `number_165`, `number_166`, `number_167`, `number_168`, `number_169`, `number_170`, `number_171`, `number_172`, `number_173`, `number_174`, `number_175`, `number_176`, `number_177`, `number_178`, `number_179`, `number_18`, `number_180`, `number_181`, `number_182`, `number_183`, `number_184`, `number_185`, `number_186`, `number_187`, `number_188`, `number_189`, `number_19`, `number_190`, `number_191`, `number_192`, `number_193`, `number_194`, `number_195`, `number_196`, `number_197`, `number_198`, `number_199`, `number_2`, `number_20`, `number_200`, `number_201`, `number_202`, `number_203`, `number_204`, `number_205`, `number_206`, `number_207`, `number_208`, `number_209`, `number_21`, `number_210`, `number_211`, `number_212`, `number_213`, `number_214`, `number_215`, `number_216`, `number_217`, `number_218`, `number_219`, `number_22`, `number_220`, `number_221`, `number_222`, `number_223`, `number_224`, `number_225`, `number_226`, `number_227`, `number_228`, `number_229`, `number_23`, `number_230`, `number_231`, `number_232`, `number_233`, `number_234`, `number_235`, `number_236`, `number_237`, `number_238`, `number_239`, `number_24`, `number_240`, `number_241`, `number_242`, `number_243`, `number_244`, `number_245`, `number_246`, `number_247`, `number_248`, `number_249`, `number_25`, `number_250`, `number_251`, `number_252`, `number_253`, `number_254`, `number_255`, `number_26`, `number_27`, `number_28`, `number_29`, `number_3`, `number_30`, `number_31`, `number_32`, `number_33`, `number_34`, `number_35`, `number_36`, `number_37`, `number_38`, `number_39`, `number_40`, `number_41`, `number_42`, `number_45`, `number_48`, `number_49`, `number_5`, `number_52`, `number_53`, `number_54`, `number_55`, `number_56`, `number_57`, `number_61`, `number_62`, `number_63`, `number_64`, `number_65`, `number_66`, `number_67`, `number_68`, `number_69`, `number_7`, `number_70`, `number_71`, `number_72`, `number_73`, `number_74`, `number_75`, `number_76`, `number_77`, `number_78`, `number_79`, `number_8`, `number_80`, `number_81`, `number_82`, `number_83`, `number_84`, `number_85`, `number_86`, `number_87`, `number_88`, `number_89`, `number_9`, `number_90`, `number_91`, `number_92`, `number_93`, `number_94`, `number_95`, `number_96`, `number_97`, `number_98`, `number_99`, `rsvp`, `sctp`, `vrrp` |
| `related_crn` | no | — |
| `remote` | no | API: The remote IP addresses or security groups from which this rule will allow traffic (or to which, for outbound rules). Can be specified as an IP address, a CIDR block, or a security group. Specify a CIDR block of `0.0.0.0/0` to allow traffic from any source (or to any destination, for outbound rules). |
| `rule_id` | no | — |
| `type` | no | range 0–254 |
| `icmp` (block) | no | nested block |
| `icmp.code` | no | range 0–255 |
| `icmp.type` | no | range 0–254 |
| `tcp` (block) | no | nested block |
| `tcp.port_max` | no | range 1–65535 |
| `tcp.port_min` | no | range 1–65535 |
| `udp` (block) | no | nested block |
| `udp.port_max` | no | range 1–65535 |
| `udp.port_min` | no | range 1–65535 |

## Attribute rules

Each writable argument has its own rule so checks can be enabled or disabled independently.

| Rule | Argument |
| --- | --- |
| `ibm_is_security_group_rule_invalid_code` | `code` |
| `ibm_is_security_group_rule_invalid_direction` | `direction` |
| `ibm_is_security_group_rule_invalid_group` | `group` |
| `ibm_is_security_group_rule_invalid_ip_version` | `ip_version` |
| `ibm_is_security_group_rule_invalid_local` | `local` |
| `ibm_is_security_group_rule_invalid_name` | `name` |
| `ibm_is_security_group_rule_invalid_port_max` | `port_max` |
| `ibm_is_security_group_rule_invalid_port_min` | `port_min` |
| `ibm_is_security_group_rule_invalid_protocol` | `protocol` |
| `ibm_is_security_group_rule_invalid_remote` | `remote` |
| `ibm_is_security_group_rule_invalid_type` | `type` |
| `ibm_is_security_group_rule_invalid_icmp__code` | `icmp.code` |
| `ibm_is_security_group_rule_invalid_icmp__type` | `icmp.type` |
| `ibm_is_security_group_rule_invalid_tcp__port_max` | `tcp.port_max` |
| `ibm_is_security_group_rule_invalid_tcp__port_min` | `tcp.port_min` |
| `ibm_is_security_group_rule_invalid_udp__port_max` | `udp.port_max` |
| `ibm_is_security_group_rule_invalid_udp__port_min` | `udp.port_min` |

## Example

### Valid

```hcl
resource "ibm_is_security_group_rule" "test" {
  direction = "inbound"
  group = "example"
}
```
