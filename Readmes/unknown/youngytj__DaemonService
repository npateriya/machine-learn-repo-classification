# 简介
[![Platform](https://img.shields.io/badge/Language-Python%202.7-blue.svg)](https://www.python.org/)
[![Build](https://travis-ci.org/youngytj/DaemonService.svg?branch=master)](https://travis-ci.org/youngytj/DaemonService)
[![Gitter](https://img.shields.io/badge/chat-on%20gitter-blue.svg)](https://gitter.im/tianjyan/Lobby)

自用的Spark Bot的样例程序

# 开发准备
* 注册[Cisco Spark](https://developer.ciscospark.com/#)开发者账号；
* 创建[Bots](https://developer.ciscospark.com/apps.html)，创建完请记得保存Token到`config.json`文件中；
* 创建[WebHook](https://developer.ciscospark.com/resource-webhooks.html)<sup>①</sup>；
* 注册并下载[ngrok](https://ngrok.com/)，注册完成后记得保存Token；

*Note*：1. WebHook的创建可以等ngrok运行后获得二级域名以后。

# 运行ngrok
    ./ngrok authtoken TOKEN // TOKEN是注册时候保存的TOKEN
    ./ngrok http 8080 //运行完成后可获取二级运行

# 真机环境中运行
* 安装[Python 2.7.13](https://www.python.org/downloads/release/python-2713/)；
* 安装依赖：`pip install -r requirements.txt`；
* 运行：`python docker.py`。

# 阿里云环境(Ubuntu 16.04)运行
* 安装依赖：`pip install -r requirements.txt`；
* 运行：`python daemon.py start`。

# Docker环境下运行
* 创建镜像：`docker build -t spark-bots:0.1 .`；
* 保存镜像：`docker save spark-bots:0.1 | gzip > spark-bots.tar.zip`；
* 导入镜像：`docker load -i spark-bots.tar.zip`；
* 运行镜像：`docker run -d -p 8080:8080 spark-bots:0.1`。

# 扩展
继承`bots/abcbot.py`可以扩展自己的Bot