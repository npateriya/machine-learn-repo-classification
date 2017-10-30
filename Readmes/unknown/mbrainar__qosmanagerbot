# QOS Manager Spark Bot

The QOS Manager bot is a user interface for the
[Event Driven QOS](https://github.com/imapex/edqos_app) app.
The ED QOS app offers a set of APIs that allow the modification of QOS policy
on Cisco's APIC EM EasyQOS module.

## Usage
To use QOS Manager bot, send Spark message to `qosmanager@sparkbot.io`.
The bot will respond with known commands. Typical flow would look like:

1. Set policy scope/tag that you wish to modify.
  * `list policy tags` to display all known policy tags in APIC EM
  * `current policy scope` to see if a policy scope has already been set for the session
  * `set policy scope <tag>` to set the policy scope for the session to `tag`

2. Search for current relevance of an application with `search app <string>`
  * If `string` is an exact match on application name, it will only return one result
  * If `string` is found in multiple application names, it will return a bulleted list
  * Relevance level is also returned with the matches;
  provided that policy scope is set for the session

3. Set the relevance level for an application with `set relevance <app> <relevance>`
  * `app` must exactly match an application in APIC EM;
  if you don't know exact name, use `search app <string>`
  * `relevance` must be either: ["Business-Relevant", "Default", "Business-Irrelevant"]

## Development
Fork us and send a pull request...
