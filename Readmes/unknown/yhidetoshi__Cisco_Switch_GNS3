![Alt Text](https://github.com/yhidetoshi/Pictures/raw/master/Cisco_Switch_GNS3/gns3-icon.jpeg)
![Alt Text](https://github.com/yhidetoshi/Pictures/raw/master/Cisco_Switch_GNS3/cisco-icon.jpeg)

# Cisco_Switch_GNS3
study cisco_switch memo

##### 検証環境(GNS3のセットアップ)
- GNS3をダウンロード&インストール
-> https://www.gns3.com/

- iOSをダウンロード
-> http://www.w7cloud.com/GNS3-IOS/c3745-adventerprisek.124-25d.bin


Ciscoコマンドメモ　
=======
|コマンド    |説明         |
|:-----------|:------------|
|# enable|特権EXECへ移動|
|# configure terminal|グローバルコンフィギュレーションモードへ|
|# exit|一つ前のモードへ|
|# ?|ヘルプコマンド|
|# show startup-config|設定変更してから設定保存されたコンフィグ確認|
|# sh running-config|コンフィグ確認|
|# copy running-config startup-config もしくは write memory|コンフィグ保存|
|# write memory もしくは wr|設定保存|
|(config)#username <user_name> password <password>|ユーザとパスワード作成|



|# ?|可能なコマンド表示|


- Interface(0/0)にアドレスを付与
```
# configure terminal
(config)# interface FastEthernet 0/0
(config-if)# ip address 10.0.2.11 255.255.255.0
(config-if)# no shutdown
```

- ssh接続できるようにする
```
(config)# username cisco password cisco
(config-line)# line vty 0 4
(config-line)# login local
(config)#ip ssh version 2
```

#### 色々と試してみる
- vlan-trunk

![Alt Text](https://github.com/yhidetoshi/Pictures/raw/master/Cisco_Switch_GNS3/vlan-trunk-test.png)

```
(trunkポート有り)
R1#ping 192.168.1.11

Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 192.168.1.11, timeout is 2 seconds:
!!!!!
Success rate is 100 percent (5/5), round-trip min/avg/max = 36/72/112 ms

(trunkポート無し)
R1#ping 192.168.1.11

Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 192.168.1.11, timeout is 2 seconds:
.....
Success rate is 0 percent (0/5)
```
