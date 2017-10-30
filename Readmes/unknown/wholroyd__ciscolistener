# CiscoListener
Utility to listen for Cisco Smart Call Home enabled device messages

# Overview
## What is Cisco Smart Call Home?
Cisco Smart Call Home is an award-winning, embedded support feature available on a broad range of Cisco devices. The devices become monitored around the clock to provide proactive network diagnostics, alerts, and remediation through an automatic call home mechanism to Cisco. The feature provides additional monitoring of your network, reduces time to resolution, and automates access to expertise and knowledge - even having Cisco automatically start RMA processes for failed chassis or device components as part of a support contract.

## What is the CiscoListener?
The CiscoListener is a messaging bridge to allow Cisco Smart Call Home devices to provide a health state from their events. Received event data can provide richer health states of the devices you already monitor in other ways and allows you to audit what kind of data is being sent for auditing/security reasons. The service translates the messages from the devices into NT event log events for monitoring systems to pick up and process as necessary.

## How does the Cisco device talk to the listener?
The Cisco devices with are Cisco Smart Call Home capable have a destination profile that defines one or more endpoints to talk to, inlcuding through email or http(s) calls, in plain text or XML format. Simply adding the machine and port combination for where the CiscoListener lives into the destination profile is all that is needed.

## What does the listener do with the event message?
It performs a message verification to ensure the message received is allowed, and then evaluates the message received against one or more rules that determine how the event is logged. Based on the rule that matches, it applies a template that defines priority, severity, and the event’s design.

## What happens to the incoming Cisco events?
1. Message is normalized internally to remove all namespaces, and an attempt to remove illegal characters is performed.
2. Rules are used to match specific conditions using XPath queries to which the related template is applied
3. An event is written to the Application event log with the appropriate properties set on the event

# Getting Started
The first part of the project to take a look at is the `configuration.xml` file sitting within the Build folder. This is the default behavior and NT event log message that will be generated for events with a severity higher than 5. In this file, you will see several placeholders surrounded by dollar signs '$'. These are simply XPath queries for grabbing information out of the incoming messages from the devices to use in the template.

You can use any XPath tool out there to find the pieces you want to look for and you can use the `example_after.xml` file sitting in the root of the project for an example of what a message looks like. I've included the `example_before.xml` to show what the original message looks like prior to having all namespaces removed, if you're curious.

You may also see a `configuration.json` in the folder as well, which is a valid configuration file which the project's serializer can read, but it's not wired up at this time.
