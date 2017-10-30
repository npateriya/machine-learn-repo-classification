INVITE Networks
===============

## Cisco Wsapi Bundle

This bundle provides an integration to Cisco Systems IOS UC Gateway API.

## Installation

### Step 1: Install vendors

To be updated.

Next, update your vendors by running:

````
$ ./bin/vendors
````

### Step 2: Enable the bundle

Finally, enable the bundle in the kernel:

````
// app/AppKernel.php
public function registerBundles()
{
    $bundles = array(
        // ...
        new Invite\Bundle\Cisco\WsapiBundle\CiscoWsapiBundle(),
    );
}
````

## Examples

To be added.

## Configuration

You can configure the cisco wsapi directly in your `config.yml` file. Here are the defaults:

````
cisco_wsapi:
    general:
        expanded:             false
    object:
        max_nesting_level:    3
        show_data:            true
        show_classinfo:       true
        show_constants:       true
        show_methods:         true
        show_properties:      true
    array:
        max_nesting_level:    8
    processor:
        active:               true
    bool:
        html_color:           #008
        cli_color:            blue
    float:
        html_color:           #800
        cli_color:            red
    int:
        html_color:           #800
        cli_color:            red
    string:
        html_color:           #080
        cli_color:            green
        show_quotes:          true
    css:
        path:                 /Asset/tree.min.css
````
