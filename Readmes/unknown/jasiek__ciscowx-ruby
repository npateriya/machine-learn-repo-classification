
# ciscowx

This is a very simple Cisco XML service which pulls a simple, textual weather forecast for wherever you are.

![screenshot](media/screenshot.jpg)

##### How to use?

Either use your `services.xml` file to point to an instance I run:

```
<CiscoIPPhoneMenu>
  <Title>Exciting services</Title>
  <MenuItem>
    <Name>Weather</Name>
    <URL>http://ciscowx.herokuapp.com/</URL>
  </MenuItem>
</CiscoIPPhoneMenu>
```
Or deploy your own wherever you want. You'll need the following:

* Ruby 2.2
* Redis (for caching requests to forecast.io)
* MaxMind GeoIP Lite city database (for IP lookups, goes into `vendor/GeoIPCity.dat`.
* forecast.io API key (goes into the `FORECASTIO_KEY` env variable)

##### Compatibility

Tested with Cisco 7960, will probably work with all other Cisco IP phones.

##### License

MIT, see LICENSE file.


