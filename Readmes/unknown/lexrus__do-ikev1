## do-ikev1

[![Install on DigitalOcean](http://installer.71m.us/button.svg)](http://installer.71m.us/install?url=https://github.com/lexrus/do-ikev1)

> 这是一个 [do-install-button](https://github.com/seven1m/do-install-button) 配置，能使用 [DigitalOcean] 快速部署 Cisco IPSec(又名 IPSec IKEv1) 服务器。配置脚本基于 [vpn-deploy-playbook](https://github.com/lexrus/vpn-deploy-playbook) 实现。虽然用的时候完全不需要知道它们是啥，但是万一遇到问题可以让男朋友传送过去研究一下帮你解决。

### 基本要求

1. 一个 [DigitalOcean] 帐号 - 如果你用我的[推广链接](https://www.digitalocean.com/?refcode=3eb5cf371fc9)注册，帐号里马上就能有 10 刀，如果开 5 刀的 VPS(512MB 内存)，那就可以免费用两个月了哦
2. [DigitalOcean] 的帐号里有余额或者绑了信用卡可用来建 VPS
3. [DigitalOcean] 里加过 public key - 怎么生成 key 可以看[这里](https://gitcafe.com/GitCafe/Help/wiki/如何安装和设置-Git#2创建-ssh-秘钥)

[DigitalOcean]: https://www.digitalocean.com/?refcode=3eb5cf371fc9

### 插播广告

为了在 iPhone (iOS 8+) 上更方便地开关你的 IKEv1 微屁恩连接，你可以购买我开发的 [VPN On](https://itunes.apple.com/app/vpn-on/id951344279)。虽然我已经公开了[它的 Swift 源码](https://github.com/lexrus/VPNOn)，但是编译它是一件相当麻烦的事情。

[<img src="https://cloud.githubusercontent.com/assets/219689/5575342/963e0ee8-9013-11e4-8091-7ece67d64729.png" width="135" height="40" alt="AppStore"/>](https://itunes.apple.com/app/vpn-on/id951344279)

### 安装

点击安装按钮即可开启轻松愉快的安装过程：

[![Install on DigitalOcean](http://installer.71m.us/button.svg)](http://installer.71m.us/install?url=https://github.com/lexrus/do-ikev1)

如果只装 IPSec IKEv1，用 512MB 完全够用。Region 选 Singapore 1 和 San Francisco 1 连接速度都不错。

### 修改默认帐号

安装成功后，默认的帐号是:

```
Account: a_secret_vpn_username
Password: a_secret_vpn_password
Secret(PSK): a_long_long_psk
```
如果想改的话，用 ssh 登录这台服务器，修改 /etc/ipsec.secret 后，`ipsec restart` 重启 IKEv1 服务。

### 赞助

欢迎请我喝咖啡，我的 [PayPal](https://www.paypal.com) 和 [支付宝](https://www.alipay.com) 帐号都是: `lexrus@gmail.com`

### 协议

本配置遵循:
```
『有问题问你男朋友』协议
2015 年 1 月修订版
Lex Tang (https://twitter.com/lexrus)

使用者不得任意修改源码中任何一个字符。
使用过程中遇到问题先问谷歌，然后知乎，再问男友。
因正常使用造成的任何损失由男朋友承担。
```

