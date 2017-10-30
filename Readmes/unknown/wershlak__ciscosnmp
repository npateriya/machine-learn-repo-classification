# Ciscosnmp

Ciscosnmp uses SNMP to manage Cisco devices. I include some simple CLI commands for these tasks but it was mainly used as a building block for an app that performed scheduled backups of several thousand devices at a few hundred sites. I've tried to clean out some old code and proprietary stuff to publish this in case someone might find it useful.

##### NOTE:

Some of the choices here are obviously suboptimal for most users. I'm looking at you SNMPv1 and TFTP! But, the client is always right. Maybe, I'll have a chance to increase support for other protocols as I have time to work on this.

## Installation

Add this line to your application's Gemfile:

```ruby
gem 'ciscosnmp'
```

And then execute:

    $ bundle

Or install it yourself as:

    $ gem install ciscosnmp

##### TFTP server:

If you want to utilize the tftpbackup or tftpupdate functions you'll need a running tftp server on your machine.

To enable TFTP OSX:

    sudo launchctl load -F /System/Library/LaunchDaemons/tftp.plist
    sudo launchctl start com.apple.tftpd

To install/enable TFTP on RHEL7 and Fedora:

    sudo yum install tftp-server
    sudo vi /etc/xinetd.d/tftp
    # Change "disable = yes" to no
    sudo systemctl start xinetd

TFTP uses UDP port 69 so you'll have to adjust any firewall rules accordingly 

## Usage

##### From command line:

    $ tftpbackup -h
    
    Usage: tftpbackup [options] device
        -c COMMUNITY                     SNMP write community - default: public
        -d DIRECTORY                     Backup directory - default: ~/Configs

##### In your project:

    require 'ciscosnmp'
    
    # These are the default options
    config = {
        :address => 'localhost',
        :community => 'public',
        :tftp_server_directory => '/var/lib/tftpboot',
        :config_directory => '~/Configs/'
    }
    
    device = Ciscosnmp::Base.new(config)
    response = device.backup
    
    if response.success?
        puts response.message
    end
  
Most of the base methods catch any exceptions and return a response object so success can be determined reliably.

## Development

After checking out the repo, run `bin/setup` to install dependencies. You can also run `bin/console` for an interactive prompt that will allow you to experiment.

To install this gem onto your local machine, run `bundle exec rake install`. To release a new version, update the version number in `version.rb`, and then run `bundle exec rake release`, which will create a git tag for the version, push git commits and tags, and push the `.gem` file to [rubygems.org](https://rubygems.org).

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/wershlak/ciscosnmp.


## License

The gem is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).

