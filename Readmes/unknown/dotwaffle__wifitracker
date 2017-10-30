# wifitracker

A bad attempt at writing a MAC address tracker for Cisco WiFi APs, by scraping the WLC SNMP tables.

## Usage

In a lame attempt to document, here's the usage statement:

```
$ wifitracker -h
  -config string
        Path to Configuration File (optional)
  -debug
        Turn on debugging output
  -snmpcommunity string
        SNMP community string (default "public")
  -snmphost string
        SNMP host to query (default "localhost")
  -snmppollinterval duration
        SNMP Polling interval (default 10s)
  -snmpretries int
        SNMP retries (default 1)
  -snmptimeout duration
        SNMP timeout (default 1s)
  -sqldb string
        MySQL Database (default "wifi")
  -sqlhost string
        MySQL Host (default "localhost")
  -sqlpass string
        MySQL Pass (default "pass")
  -sqlport int
        MySQL Port (default 3306)
  -sqltls string
        MySQL TLS (default "false") (true, false, skip-verify)
  -sqluser string
        MySQL User (default "user")
```

Due to the particular flag package I'm using, you can change any of those options by three methods, in order of precedence:

1. Setting the flag, e.g. `wifitracker -debug -snmpcommunity cheese`
2. Setting an environment variable, e.g. `export SNMPCOMMUNITY=cheese DEBUG=true; wifitracker`
3. Setting a configuration file, e.g. `wifitracker -config myConfigFile`

The format of the configuration file is:

```
# this is a comment
snmpcommunity=cheese
debug=true
```

If you plan on using the configuration file format as a Docker env-file, be prepared for the fact that in an env-file all the variable keys should be UPPERCASE, whereas in the configuration file they should be lowercase. I should probably fix that sometime.

## Docker

For those of you with a Docker persuasion, the latest version is always published at [Docker Hub](https://hub.docker.com/r/dotwaffle/wifitracker/) and can be easily pulled with: `docker pull dotwaffle/wifitracker:latest`

Dealing with flags and configuration files in Docker is surprisingly annoying, so in a fit of both rage and madness I decided to encourage you to follow Kelsey Hightower's example, and the example of the [12 Fractured Apps](https://medium.com/@kelseyhightower/12-fractured-apps-1080c73d481c) by using environment variables, by writing your own Dockerfile that uses `FROM dotwaffle/wifitracker:latest` or by doing something at run-time like:

```
$ docker run -dt --name wifitracker --rm -e SNMPCOMMUNITY=cheese -e DEBUG=true dotwaffle/wifitracker
```

You can even use your configuration file as an environment variable list!

```
$ docker run -dt --name wifitracker --rm --env-file wifitracker.conf dotwaffle/wifitracker
```

If you're being super-smart and using Docker swarm, you can even use the Docker "secrets" functionality:

```
$ cat wifitracker.conf | docker secret create wifitracker.conf -
rbkdu2zuekxq8mc7pkvfhasyp
$ docker service create --replicas 1 --name wifitracker --secret="wifitracker.conf" -e "CONFIG=/run/secrets/wifitracker.conf" dotwaffle/wifitracker
v26mkzdhibja655n0qszawmyk
```

Using the secrets functionality will allow nosey parkers to stop reading your precious MySQL password strings.

Final note: You'll notice that I've specified `docker run -d` which detaches the process. You can watch the progress with `docker logs --follow wifitracker`, you can attach to it with `docker attach wifitracker` (detach again with ^p^q) or you can start it and immediately attach by changing the run parameter to `docker run -it` for interactive.

## Sample Data Output

I've supplied `sample-output.sql` which if run against your database and spit out the data for you in a nice to digest format. It'll join the two SQL tables nicely so that it'll show which access point and WiFi channel it was on when the scan occurred. Eventually I'll be making a front-end to this app which will plot live where users are, and move them across a map as they move between access-points. Based on my hatred of coding GUIs, this may take considerably longer than this app took.

Enjoy!
