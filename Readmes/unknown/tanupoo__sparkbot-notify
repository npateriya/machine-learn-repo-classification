spartbot-notify
===============

This is a ultra simple Spark Bot submitting a timestamp
to the room you specified.

## How to build it.

Don't forget to add --recursive option when you clone it.

    % git clone --recursive https://github.com/tanupoo/sparkbot-notify.git

It uses *libcurl* to submit the messages.
Typical OS has the library from the initial.
If not, you need to install it first.

Then, simpley type make.

    % cd sparkbot-notify
    % make

That's it.

## How to run it.

- Get your access key and room id.

See https://developer.ciscospark.com/index.html

- Edit your configuration file properly.  

For example,

    % cat config.ini 
    [spark]
    server_url = https://api.ciscospark.com/v1/messages
    verify_peer = false
    auth_key = N2Q1...YmMz
    room_id = Y2lz...MWI0
    dialect = Now, it's %H:%M:%S.
    frequency = 10

- Set the environment variables.

Set the path of your configuration file into CAF_APP_PATH.
Set the file name of the one into CAF_APP_CONFIG_FILE.

    % export CAF_APP_PATH="./"
    % export CAF_APP_CONFIG_FILE="config.ini"

- Then, launch it.

    % spartbot-notify

with the -d option(s), you can see more information.
See sample-test.sh for your reference.

Enjoy.

