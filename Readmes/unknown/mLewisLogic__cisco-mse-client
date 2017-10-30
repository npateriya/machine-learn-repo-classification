# Ruby Cisco MSE (Mobility Services Engine) client


## Usage
```ruby
client = CiscoMSE::Client.create(
  username: 'YOUR_USERNAME',
  password: 'YOUR_PASSWORD'
)

```


## Testing

### from your project
In your spec_helper.rb, include the following snippet.
```ruby
CiscoMSE::Client.stub!
```
This will prevent `cisco-mse-client` from reaching out to live servers.
Instead, endpoints will return back [local stubbed data](lib/cisco-mse-client/stub.rb).

### the cisco-mse-client gem
* Copy spec/_creds.stub.rb to spec/_creds.rb
* Update it to include a valid token
* bundle exec rspec


## License
(c) Mike Lewis 2013
Released under the [MIT License](http://www.opensource.org/licenses/MIT).
