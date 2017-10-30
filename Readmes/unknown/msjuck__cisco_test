# [중부]DNSN-ACSM-IMS-L3SW-2 CONSISTENCY LOG해소를 위한 SUP 절체 작업

-------------
```
1. 작성자	: 문영민	
2. 작업자	: 문영민
3. 승인자	: 최홍준M
3. Version	: 1.1	
4. 시스템	: DNSN-ACSM-IMS-L3SW-1, 2	
5. 작성일	: 2017/09/27
6. 작업개요	: [중부]DNSN-ACSM-IMS-L3SW-2 CONSISTENCY LOG해소를 위한 SUP 절체 작업
7. 작업목적	: FM-6-FM_CONSISTENCY_CHECK_LOG_STATUS: Consistency Checker
             found inconsistency 에러 지속 발생(서비스영향X)
8. 작업일자	: 2017-10-11 02:00 ~ 2017-10-11 05:00
```
-------------
[둔산] DNSN-ACSM-IMS-L3SW-1
-------------
<pre><code>
* 해당 작업 후 서비스 이상 유무 확인 
conf t
!
interface Vlan32
 standby 32 priority 115
 standby 32 preempt
!(HSRP 절체 후 preempt 삭제)
 no standby 32 preempt
!
interface Vlan82
 standby 82 priority 115
 standby 82 preempt
!(HSRP 절체 후 preempt 삭제)
 no standby 82 preempt
!
end
!
wr
</code></pre>
