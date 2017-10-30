cisco-vrf-bgp-neigh.go
======================

Parser for cisco command output:

```
show bgp vpnv4 unicast all neighbors
```

Usage
=====

0. Save output of command 'show bgp vpnv4 unicast all neighbors' to a file. For instance, 'output.txt'.
0. Feed that output file to cisco-vrf-bgp-neigh.go:
```
go run cisco-vrf-bgp-neigh.go < output.txt
```

Example
=======

NOTICE: Addresses and ASNs purposely omitted below.

```
C:\tmp\devel\cisco-vrf-bgp-neigh>go run src\main.go < output2.txt
2016/01/19 15:41:10 main: reading from stdin
2016/01/19 15:41:10 main: reading from stdin: done: 1755 lines
2016/01/19 15:41:10 main: found 16 neighbors
Neighbor        VRF            ASN    State       Uptime  Prefixes
1.1.1.1         --             11111  Idle        ?            0
1.1.1.1         --             11111  Established 27w3d      236
1.1.1.1         --             11111  Established 1y8w      4715
1.1.1.1         --             11111  Established 9w4d        94
1.1.1.1         --             11111  Established 42w5d        2
1.1.1.1         --             11111  Idle        ?            0
1.1.1.1         --             11111  Established 1y8w        10
1.1.1.1         --             11111  Established 26w2d     3450
1.1.1.1         --             11111  Established 42w5d       25
1.1.1.1         --             11111  Established 14w2d       78
1.1.1.1         --             11111  Established 44w0d      110
1.1.1.1         --             11111  Established 1y46w       61
1.1.1.1         --             11111  Established 2y38w       77
1.1.1.1         --             11111  Established 19w1d      416
1.1.1.1         --             11111  Established 1y46w      157
1.1.1.1         --             11111  Established 13w0d     3451

C:\tmp\devel\cisco-vrf-bgp-neigh>
```

END

