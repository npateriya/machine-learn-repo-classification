ctlcisco
========

透過cdp方式建立 cisco設備拓譜結構

### 程式說明

主要程式有:

* get_cdp_entry.py 自動連線取得cdp資料
* generate_topology.py 利用cdp資料 以BFS方式搜尋建立拓譜資料
* topology_log.py 將搜尋後的資料以json方式儲存紀錄
* diff_topology.py 比對json格式 有差異的部分

#### 使用方式

```bash
$ git clone https://github.com/fajoy/ctlcisco.git
$ cd ctlcisco
$ cp etc/config.ini.example etc/config.ini
$ vim etc/config.ini
$ ./get_cdp_entry.py -h
$ ./generate_topology.py -h
$ ./topology_log.py -h
$ ./diff_topology.py -h
```
