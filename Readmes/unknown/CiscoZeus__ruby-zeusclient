# Ruby Zeus Client [![license](https://img.shields.io/hexpm/l/plug.svg)](http://www.apache.org/licenses/LICENSE-2.0) [![Build Status](https://travis-ci.org/CiscoZeus/ruby-zeusclient.svg)](https://travis-ci.org/CiscoZeus/ruby-zeusclient)

![Alt text](https://github.com/CiscoZeus/ruby-zeusclient/blob/master/icons/zeus-logo.png?raw=true "Zeus Logo")

Ruby client for [Cisco Zeus](http://www.ciscozeus.io/). This allows us to send and recieve data to and from Zeus.

## Installation

Add this line to your application's Gemfile:

```ruby
gem 'zeusclient'
```

And then execute:

    $ bundle

Or install it yourself as:

    $ gem install zeusclient

## Usage

```ruby
require 'zeus/api_client'
zeus_client = Zeus::APIClient.new({
    :access_token => "your_token_here",
    :endpoint => "zeus_endpoint_here"
})
```

List All Metrics

```ruby
result = zeus_client.list_metrics(
  # optional parameters
  {
    regex: "metric.name*" # regex for filtering metrics names
    from_date: "2016-06-12", # YYYY-MM-DD
    to_date: "2017-12-20", # YYYY-MM-DD
    aggregator_function: "sum", # Aggregator function to be applied to the metric values.
                                # Available options. count, min, max, sum, mean, mode, median
    group_interval: "100m",     # Interval for grouping, when applying aggregations via the aggregator_function parameter.
                                # The aggregator_function parameter must be specific for this parameter to work.
                                # Examples: 1000s, 100m, 10h, 1d , 1w. Use 's' for seconds, 'm' for minutes, 'h' for hours,
                                # 'd' for days and 'w' for weeks.
    filter_condition: "column1 > 0", # Filters to be applied to metric values.
                                     # Examples: "column1 > 0", "column1 > 50 AND column2 = 10".
    offset: 100,
    limit: 10
  }
)
p result.code      # 200
p result.success?  # true
p result.data      # => {}
```

Get Metric

```ruby
result = zeus_client.get_metrics( # same optional arguments as list_metrics)
p result.code      # 200
p result.success?  # true
p result.data      # => {}
```

Push Metric

```ruby
result = zeus_client.send_metrics([{point: {value: 1, ...}}, ...])
p result.code      # 200
p result.success?  # true
p result.data      # => {}
```

Delete metric

```ruby
result = zeus_client.delete_metrics()
p result.code      # 200
p result.success?  # true
p result.data      # => {}
```

Get logs

```ruby
result = zeus_client.get_logs(
  "log_name_here",
  #optional parameters
  {
    attribute_name: "cpu", # name of the attribute within the log to be searched
    pattern: "memory.*", # a pattern for name filtering
    from_date: "2016-06-12", # YYYY-MM-DD
    to_date: "2017-12-20", # YYYY-MM-DD
    offset: 200,
    limit: 10
  }
)
p result.code      # 200
p result.success?  # true
p result.data      # => {}
```

Push logs

```ruby
result = zeus_client.send_logs([{},{}, ...])
p result.code      # 200
p result.success?  # true
p result.data      # => {}
```

Get alerts

```ruby
result = zeus_client.get_alerts()
p result.code      # 200
p result.success?  # true
p result.data      # => {}
```

Create alert

```ruby
result = zeus_client.create_alert(
  {
    alert_name: "name of the alert",
    username: "username associated with alert",
    alert_expression: "cpu.value > 80", # expression to match alert against
    # optional params
    alerts_type: "metrics", # Either metric or log
    alert_severity: "S1", # severity of the alert, S1-S5
    metric_name: "name of the metrics associated with the alert",
    emails: "email@provider.com", # email to be notified when alert triggers
    status: "active", # either active or disabled
    frequency: 60, # frequency for alert to be checked
  }
)
p result.code      # 201
p result.success?  # true
p result.data      # => {}
```

Modify alert

```ruby
result = zeus_client.modify_alert(
  alert_id,
  {
    # parameter you wish to modify
    alert_name: "name of the alert",
    username: "username associated with alert",
    alert_expression: "cpu.value > 80", # expression to match alert against
    alerts_type: "metric", # Either metric or log
    alert_severity: "S1", # severity of the alert, S1-S5
    metric_name: "name of the metrics associated with the alert",
    emails: "email@provider.com", # email to be notified when alert triggers
    status: "active", # either active or disabled
    frequency: 60, # frequency for alert to be checked
  }
)
p result.code      # 200
p result.success?  # true
p result.data      # => {}
```

Delete alert

```ruby
result = zeus_client.delete_alert(alert_id)
p result.code      # 204
p result.success?  # true
p result.data      # => {}
```

Get triggered alerts

```ruby
result = zeus_client.triggered_alerts()
p result.code      # 200
p result.success?  # true
p result.data      # => {}
Get triggered alerts
```

Get triggered alerts in the last 24 hours

```ruby
result = zeus_client.triggered_alerts_last_24_hours()
p result.code      # 200
p result.success?  # true
p result.data      # => {}
```

For more details, refer to [this documentation](http://www.rubydoc.info/github/CiscoZeus/ruby-zeusclient/)

## Development

After checking out the repo, run `bin/setup` to install dependencies. Then, run `bin/console` for an interactive prompt that will allow you to experiment.

To install this gem onto your local machine, run `bundle exec rake install`. To release a new version, update the version number in `version.rb`, and then run `bundle exec rake release` to create a git tag for the version, push git commits and tags, and push the `.gem` file to [rubygems.org](https://rubygems.org).

## Contributing

1. Fork it ( https://github.com/CiscoZeus/zeusclient/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
