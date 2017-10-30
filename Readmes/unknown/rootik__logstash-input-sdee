# Logstash Cisco SDEE/CIDEE input plugin
[![Gem Version](https://badge.fury.io/rb/logstash-input-sdee.svg)](https://badge.fury.io/rb/logstash-input-sdee)

This plugin is based off [logstash-input-http_poller](https://github.com/logstash-plugins/logstash-input-http_poller) by @maximede.

This [Logstash](https://github.com/elasticsearch/logstash) input plugin allows you to call a Cisco SDEE/CIDEE HTTP API, decode the output of it into event(s), and send them on their merry way.

The idea behind this plugins came from a need to gather events from Cisco security devices and feed them to ELK stack.

This plugin is tested on:
* Hardware: Cisco ASA 5585-X IPS SSP-10
* IPS Version: 7.3(2)E4
* logstash 2.3.4
* Java JRE 1.8.0-60

It is fully free and fully open source. The license is Apache 2.0, meaning you are pretty much free to use it however you want in whatever way.

## TODO
- Add request options like minThreatRating and maxThreatRating, virtualSensors, errorSeverities, includeStatusCategories, excludeStatusCategories and so on.
- Gather sensor status and health events

## Config Example

For config examples see `sdee.conf` in `examples` in this repo.

## SSL notes

You need to import host SSL certificate in Java trust store to be able to connect to Cisco IPS device.
I have had no luck with `truststore` nor `ssl_certificate_validation` configuration options with Java 1.8.0-60 like mentioned in [logstash-mixin-http_client](https://github.com/logstash-plugins/logstash-mixin-http_client) documentation,
so I'm using system trustStore - `$JAVA_HOME/lib/security/cacerts`. Note, default Java trustStore password is `changeit`.
Take the following steps:

* Get server certificate from IPS device:
```sh
echo | openssl s_client -connect ciscoips:443 2>&1 | sed -ne '/-BEGIN CERTIFICATE-/,/-END CERTIFICATE-/p' > cert.pem
```

* Import it into Java cacerts:
```sh
$JAVA_HOME/bin/keytool -keystore $JAVA_HOME/lib/security/cacerts -importcert -alias ciscoips -file cert.pem
```

* Verify if import was successful:
```sh
$JAVA_HOME/bin/keytool -keystore $JAVA_HOME/lib/security/cacerts -list
```

* Modify your logstash input config for SSL connection:
```ruby
input {
  sdee { 
    interval => 60  
    http => { 
      truststore_password => "changeit" 
      url => "https://10.0.2.1"  
      auth => {
        user => "cisco"
        password => "p@ssw0rd"
      }
    }
  }
}
```

* Test logstash

## Documentation

Logstash provides infrastructure to automatically generate documentation for this plugin. We use the asciidoc format to write documentation so any comments in the source code will be first converted into asciidoc and then into html. All plugin documentation are placed under one [central location](http://www.elasticsearch.org/guide/en/logstash/current/).

- For formatting code or config example, you can use the asciidoc `[source,ruby]` directive
- For more asciidoc formatting tips, see the excellent reference here https://github.com/elasticsearch/docs#asciidoc-guide

## Need Help?

Need help? Try #logstash on freenode IRC or the https://discuss.elastic.co/c/logstash discussion forum.

## Developing

### 1. Plugin Developement and Testing

#### Code
- To get started, you'll need JRuby with the Bundler gem installed.

- Create a new plugin or clone and existing from the GitHub [logstash-plugins](https://github.com/logstash-plugins) organization. We also provide [example plugins](https://github.com/logstash-plugins?query=example).

- Install dependencies
```sh
bundle install
```

#### Test

- Update your dependencies

```sh
bundle install
```

- Run tests

```sh
bundle exec rspec
```

### 2. Running your unpublished Plugin in Logstash

#### 2.1 Run in a local Logstash clone

- Edit Logstash `Gemfile` and add the local plugin path, for example:
```ruby
gem "logstash-input-sdee", :path => "/your/local/logstash-input-sdee"
```
- Install plugin
```sh
bin/logstash-plugin install --no-verify
```
- Run Logstash with your plugin
```sh
bin/logstash -e 'input {sdee { interval => 60 http => { url => "http://ciscoips" auth => {user => "cisco" password => "p@ssw0rd"}} session_file => "/tmp/session.db" }}'
```
At this point any modifications to the plugin code will be applied to this local Logstash setup. After modifying the plugin, simply rerun Logstash.

Enable signature 2000 (ICMP Echo Reply) on your Cisco IPS device and ping some host on a network, monitored by IPS.
You should see output like this:

```ruby
{
             "@timestamp" => "2015-09-21T12:21:26.000Z",
               "timezone" => "EEST",
              "tz_offset" => "180",
               "event_id" => "6824288790867",
               "severity" => "informational",
                 "vendor" => "Cisco",
                "host_id" => "sensor1",
               "app_name" => "sensorApp",
        "app_instance_id" => "26957",
            "description" => "ICMP Echo Reply",
                 "sig_id" => "2000",
            "sig_version" => "S666",
               "sig_type" => "other",
            "sig_created" => "20001127",
              "subsig_id" => "0",
            "sig_details" => "ICMP Echo Reply",
        "interface_group" => "vs0",
                   "vlan" => "0",
          "attacker_addr" => "10.0.0.1",
      "attacker_locality" => "OUT",
            "target_addr" => "10.0.1.1",
       "target_os_source" => "learned",
         "target_os_type" => "windows-nt-2k-xp",
    "target_os_relevance" => "relevant",
          "alert_details" => "InterfaceAttributes:  context='single_vf' physical='Unknown' backplane='PortChannel0/0' ; ",
            "risk_rating" => "35",
            "risk_target" => "medium",
          "risk_attacker" => "relevant",
          "threat_rating" => "35",
              "interface" => "PortChannel0/0",
                   "host" => "10.0.2.1",
                   "tags" => "SDEE",
                "message" => "IdsAlert: 'ICMP Echo Reply' Attacker: '10.0.0.1' Target: '10.0.1.1' SigId: '2000'",
               "@version" => "1"
}
```

#### 2.2 Run in an installed Logstash

You can use the same **2.1** method to run your plugin in an installed Logstash by editing its `Gemfile` and pointing the `:path` to your local plugin development directory or you can build the gem and install it using:

- Build your plugin gem
```sh
gem build logstash-input-sdee.gemspec
```
- Install the plugin from the Logstash home
```sh
bin/logstash-plugin install /your/local/plugin/logstash-input-sdee.gem
```
- Start Logstash and proceed to test the plugin

## Contributing

All contributions are welcome: ideas, patches, documentation, bug reports, complaints, and even something you drew up on a napkin.

Programming is not a required skill. Whatever you've seen about open source and maintainers or community members  saying "send patches or die" - you will not see that here.

It is more important to the community that you are able to contribute.

For more information about contributing, see the [CONTRIBUTING](https://github.com/elasticsearch/logstash/blob/master/CONTRIBUTING.md) file.

## Reference
[Cisco Intrusion Detection Event Exchange (CIDEE) Specification](http://www.cisco.com/c/en/us/td/docs/security/ips/specs/CIDEE_Specification.html)
