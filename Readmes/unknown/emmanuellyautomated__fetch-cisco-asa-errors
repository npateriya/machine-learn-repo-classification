Here lies a tool that will trove the [Cisco ASA syslog sites]('http://www.cisco.com/c/en/us/td/docs/security/asa/syslog-guide/syslogs/logmsgs1.html') and returns dictionaries of information found.


###Usage
---
Everything currently lives inside of `fetch_cisco_asa_errors.py` so just run the tests with `python fetch_cisco_asa_errors.py` to see what it does. You should see some output like the following:

```
[

  ...

  {'action': ['Install the same image version on both units as soon as '
             'possible.'],
  'aux_exp': ['ver_num —Version number'],
  'exp': [' The ASA has detected that the peer unit is running a version that '
          'is not identical, but supports Hitless Upgrade and is compatible '
          'with the local unit. The system performance may be degraded because '
          'the image version is not identical, and the ASA may develop a '
          'stability issue if the nonidentical image runs for an extended '
          'period.'],
  'id': ['103007'],
  'msg': ['%ASA-1-103007: (Primary|Secondary) Mate version  ver_num is not '
          'identical with ours  ver_num']},

  ...

]
```

In the future, a more robust version of this tool can be used to inform programs that care about this sort of thing (i.e. Logstash) of changes to the compendium of cisco-asa syslog messages.
