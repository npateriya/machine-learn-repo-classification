# Logstash Output Plugin to Cisco Zeus

This is a plugin for [Logstash](https://github.com/elastic/logstash).
It is a simple wrapper of the [Ruby Zeus Client](https://github.com/CiscoZeus/ruby-zeusclient).

## Developing

First of all you'll need JRuby, since Logstash is written in ruby and runs on the JVM.
You can get it [in their website](http://jruby.org/download) or through [RVM](https://rvm.io/).

### Building

```sh
jruby -S bundle install
jruby -S gem build logstash-output-CiscoZeus.gemspec
```

## Usage

### Installation

```sh
logstash-plugin install logstash-output-CiscoZeus
```

### Configuration

#### Logs

Add this configuration to the output section of your logstash configuration file.

```
output {
  CiscoZeus {
    token => "my_token"
    log_name => "my_desired_log_name"
    endpoint => "myendpoint.ciscozeus.io"
  }
}
```

#### Metrics

The basic configuration is analogous to the Logs plugin

```
output {
  CiscoZeusMetrics {
    token => "my_token"
    metric_name => "my_desired_log_name"
    endpoint => "myendpoint.ciscozeus.io"
  }
}
```

With this configuration, the plugin will forward all the numeric fields to Zeus as metrics.

Optionally, you can use `fields` to select a specific set of fields instead. The plugin will try to do a type cast in case you select a non integer field.

```
output {
  CiscoZeusMetrics {
    token => "my_token"
    metric_name => "my_desired_log_name"
    endpoint => "myendpoint.ciscozeus.io"
    fields => ["foo","bar"]
  }
}
```



## License and copyright

Copyright(C) 2017 - Cisco Systems, Inc.
Apache License, Version 2.0

