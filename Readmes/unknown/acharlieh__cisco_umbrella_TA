# cisco_umbrella_TA
Tech Add-on for Splunk developed to sourcetype Cisco Umbrella DNS Activity Logs.

## Usage
1. For Cisco Umbrella customers, setup Log Management to export logs to Amazon S3 as described in the [Umbrella Customer Docs](https://support.umbrella.com/hc/en-us/articles/231248448-Cisco-Umbrella-Log-Management-in-Amazon-S3)
2. Install this app into your Splunk environment on your Search Heads, Indexers, and/or Heavy Forwarders as applicable in your environment, restarting as appropriate.
3. Follow most of the [umbrella documentation for Splunk Integration](https://support.umbrella.com/hc/en-us/articles/230650987-Configuring-Splunk-for-use-with-Cisco-Umbrella-Log-Management-in-AWS-S3) realizing they were written for setting up an earlier version of the [Splunk Add On for Amazon Web Services](https://splunkbase.splunk.com/app/1876/) to pull your logs from S3. Instead of using the recommended default `aws:s3` as a sourcetype, use `cisco:umbrella` instead.
4. Enjoy correctly timestamped logs with [field extractions](https://support.umbrella.com/hc/en-us/articles/231248508-Log-Management-Export-Format)!

## Notes
* Contributions Welcome. Licensed under Apache 2. Trademarks belong to their respective owners. etc etc.
* I threw this app together one weekend that as I was playing with these Umbrella logs as part of a trial / demo worked my way through the instructions. 
* I think there might be some value to normalize to the CIM model but also I think there's might be some work to be done with Cisco folks to clean up data as it is sent to S3.
* Questions/Comments, I can often be found as "teddybfez" in the [Splunk Usergroups Slack](https://splunk-usergroups.slack.com). Request an invite through the form on the [Nebraska User Group site](http://splunk402.com/chat/)
