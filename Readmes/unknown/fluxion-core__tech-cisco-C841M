# tech-cisco-C841M
How to setup Cisco routers, etc...

###### rev
01.00.0000 : Create draft


# Overview : 概要
+ Ciscoの新しい事業戦略でリリースされた、SOHO向けルーターです。
+ 新しい試みなのでハード、ソフト、サポート、すべてで粗が目立ちます。
+ [Cisco IOS] ベースでネットワーク構築するStartupとしては、悪くない選択だと考えます。

+ It is a router for SOHO released in Cisco's new business strategy.
+ It's a new attempt, so hard, soft, support, coarse stand out at everything.
+ I think that it is a not bad choice as a startup to build a network on a Cisco IOS basis.

[Cisco IOS]: https://ja.wikipedia.org/wiki/Cisco_IOS "Cisco Internetwork Operating System"


# Premise : 前提
## Target device : 対象装置
| Item : 項目 | Spec : 仕様       | Comment : 備考 |
|:------------|:-----------------|:---------------|
| Model       | C841M-4X-JSEC/K9 | [DataSheet]    |

[DataSheet]: http://www.cisco.com/web/JP/product/hs/routers/c800isr/prodlit/pdf/data_sheet_c78-732678.pdf "C841M-4X-JSEC/K9 Data Sheet"


# Caution : 注意
## Needs backup (Maintain network) : バックアップ環境を用意して取り組むこと
+ 前提を知らずにハマる可能性が高く、ネットワーク接続できないと解決自体が困難。
  - [Cisco IOS] の暗黙の前提を知らずにハマる ([Configurationモード]って1つじゃないの？)
  - サポートにつながらなくてハマる (この製品は**極めて特殊**な形態でのサポート契約になっている(後述))
  - 初期不良でハマる (書き込みでは交換となった事例あり)

+ Without understanding of implicit premise, you will get hurt.
  - [Cisco IOS]'s implicit premise
  - How to get support? Its too difficult!!
  - Initial failure?
 
## Select connection : ルーターへの接続方式とその準備を行うこと
+ ルーターへの接続(ログイン)は以下3つが用意され(必要機材、スキル等で異なる)
  - CCP Express (Webベースの設定機能) : 一番簡単ですが、細かい設定は未サポート。
  - Telnet : LAN接続で可能。
  - Serial : 専用ケーブル([USB-RJ45 Cable])での接続。

+ CCP Express (3.3以降)を使用し、詳細はCCP ExpressからのCLI実行とするのがお勧め
  - ミスをしても出荷状態に戻せばCCP Expressが使用可能である(**誤って構成をリセットしない限り**)
  - CCP ExpressのCLIがあればTelnet接続とする必要性は薄い
  - 3.2以降ではIOS、CCP ExpressのUpdateも簡単に行えるようになっている(Good!)
  
+ 専用ケーブルは必ず用意しておくべし
  - このルータを買う人なら一緒に用意しておくのが吉。転ばぬ先の杖。
  - Amazonで互換品が\1,680程度である

+ Telnet接続の際は前提条件に注意せよ
  - telnetは出荷状態だと**10.10.10.xのネットワークからしか接続できない**罠がある
  - モノによってはTelent接続のためにはPassword設定済みであることが前提条件となっている
  - Password設定にはSerial接続が必要となり、鶏と卵の関係になる。
  - 結局、専用ケーブルを用意しないと何もはじめられない。

+ Connection interface as follows:
  - CCP Express (Web based configuration)
  - Telnet (via LAN)
  - Serial (via [USB-RJ45 Cable]

+ Recommended to use CCP Express (3.3 or later) with CLI execution
  - If make a mistake, can initialize CCP Express. (Factory Reset)
  - With CLI execution in CCP Express, telnet is unnecessary.
  - 3.2 or later, Easy to update IOS, CCP Express yourself (Good!)

+ Dedicated cables must be prepared
  - It is good to have them together if you buy this router. The stagnant cane.
  - Compatible items are about \ 1,680 on Amazon.

+ Note in Telnet
  - Can be connected only from 10.10.10.x (On Factory Setting)
  - Needs password enable (maybe needs serial connection)

[Configurationモード]: http://www.infraexpert.com/study/ciscoios1.html "Mode of IOS"
[USB-RJ45 Cable]: https://www.amazon.co.jp/CISCO%E4%BA%92%E6%8F%9B%E3%82%B1%E3%83%BC%E3%83%96%E3%83%AB-FTDI-chipset-USB-RJ45-%E3%82%B3%E3%83%B3%E3%82%BD%E3%83%BC%E3%83%AB%E3%82%B1%E3%83%BC%E3%83%96%E3%83%AB/dp/B00JPFOTOY/ref=sr_1_1?s=computers&ie=UTF8&qid=1489781124&sr=1-1&keywords=Lontion+Industrial


# Points of basic setting : 基本設定のポイント
+ CCP Expressのクイックセットアップウイザードに従って進めれば良い
+ ただし、セキュリティ設定をデフォルトで有効にすると、ほとんどのポートは閉じられている状態になる
+ Mail関係はIMAP、POP3、SMTPの該当ポートを解放する必要がある

1. [セキュリティ] > [ポリシー]で新規追加
1. [ポート]で該当ポートを追加して保存
1. この時、[アプリケーション]で選択しないこと!! (ポート指定しても無視される)

+ Follow the CCP Express quick setup wizard
+ Note: If enable security setting by default, most ports will be closed.
+ For example, you release corresponding ports of IMAP, POP3, SMTP.

1. [Security] > [Policy] Add new policy
1. [Port] Add release ports
1. Do not select in [Application] at this time!! (Even if specifying a port, it ignored)

# ref : 参考情報
+ [Cisco Start]
+ [Cisco Start C841M]
+ [導入事例1]

[Cisco Start]: http://www.cisco.com/c/m/ja_jp/solutions/cisco-start/index.html
[Cisco Start C841M]: http://www.cisco.com/c/m/ja_jp/solutions/cisco-start/product_841mj.html
[導入事例1]: https://kentai-shiroma.blogspot.jp/2015/10/vpnciscostartc841m-4x-jsec.html

--------------------------------------------------------------------------------
EOF
