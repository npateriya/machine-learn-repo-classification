# Ruby-ImcSdk README

SDK for Cisco IMC Server (standalone)

## Installation
SDK could be installed using one of the following ways:
### From Rubygems.org
    1. gem install ImcSdk
    
### From Github
    1. git clone https://github.com/CiscoUcs/imcsdk-ruby
    2. cd imcsdk-ruby
    3. make install

## Un-installation

### From Rubygems.org
    1. gem uninstall ImcSdk
### From Github
    1. make clean

## Usage

add `require 'ImcSdk'`
in your application file to start using the methods.
For using the apis defined in the SDK additionally add  
* `require 'ImcSdk/apis/admin/<admin api file name>'`
* `require 'ImcSdk/apis/server/<server api file name>'`

Admin api file names:
* ipmi.rb
* ldap.rb
* ntp.rb
* user.rb

Server Api file names:
* bios.rb
* power.rb
* remotepresence.rb
* serveractions.rb
* storage.rb

## Documentation
http://www.rubydoc.info/gems/ImcSdk

## Development

* Configure the `test_helper.rb` file, `test_helper.rb` file is located under `ImcSdk/test` folder.
* Add an appropriate CIMC IP, Username and Password in the `test_helper.rb` file.
* Run `rake test` from ImcSdk folder to run all the unit tests in test folder.
* For running tests in a single file run as follows
    `rake test TEST=test/ImcSdk_ipmi_test.rb`


## Community

We are on Slack - slack requires registration, but the ucspython team is open invitation to anyone to register here : https://ucspython.herokuapp.com/

## License
***
Copyright 2017 Cisco Systems, Inc.

Licensed under the Apache License, Version 2.0 (the "License");   
you may not use this file except in compliance with the License.   
You may obtain a copy of the License at   

    http://www.apache.org/licenses/LICENSE-2.0   

Unless required by applicable law or agreed to in writing, software   
distributed under the License is distributed on an "AS IS" BASIS,   
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.   
See the License for the specific language governing permissions and   
limitations under the License.   

## Authors
***
Name: Amit Mandal   
email: amimanda@cisco.com 

