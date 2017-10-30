# Cisco ASA Fluent Parser

This project is forked from [rogeriobastos/fluent-plugin-cisco-asa-parser](https://github.com/rogeriobastos/fluent-plugin-cisco-asa-parser)

Because the project is seemed to be finished to maintenance, I'm trying to continue update.

## How to use

1. Clone this repository or Download `parser_cisco_asa.rb`.
2. Put `parser_cisco_asa.rb` to `/etc/td-agent/plugin/`.
3. In configuration file, set type to **cisco_asa**.

### Sample Configuration
Sample Fluentd (td-agent) configuration. This snip contains source only. Please create the match direactive.

```
<source>
  @type tail
  path /var/log/ASA5505.log
  pos_file /var/log/ASA5505.log.pos
  tag net.asa5505.<hostname>
  format cisco_asa
</source>
```

### Sample Logs
I tested ASA Log as below.

- Built Connections

```
Apr  2 23:12:15 gateway %ASA-6-305011: Built dynamic TCP translation from management:192.168.1.102/35484 to outside:220.146.21.68/35484
Apr  2 23:12:15 gateway %ASA-6-302013: Built outbound TCP connection 1906206 for outside:172.217.25.238/443 (172.217.25.238/443) to management:192.168.1.102/35484 (220.146.21.68/35484)
Apr  2 23:12:15 gateway %ASA-6-305011: Built dynamic UDP translation from management:192.168.1.102/56513 to outside:220.146.21.68/56513
Apr  2 23:12:12 gateway %ASA-6-302015: Built inbound UDP connection 1906204 for management:192.168.1.70/55192 (192.168.1.70/55192) to identity:192.168.1.5/161 (192.168.1.5/161)
Apr  2 23:12:12 gateway %ASA-6-302020: Built inbound ICMP connection for faddr 169.46.82.181/0 gaddr 220.146.21.68/0 laddr 220.146.21.68/0
```

## Contact
Please create the issue.

