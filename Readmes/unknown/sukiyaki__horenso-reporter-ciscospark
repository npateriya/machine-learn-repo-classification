horenso-reporter-ciscospark
=====================

[![CircleCI](https://circleci.com/gh/sukiyaki/horenso-reporter-ciscospark/tree/master.svg?style=svg)](https://circleci.com/gh/sukiyaki/horenso-reporter-ciscospark/tree/master)
[![Coverage Status](https://coveralls.io/repos/github/sukiyaki/horenso-reporter-ciscospark/badge.svg?branch=master)](https://coveralls.io/github/sukiyaki/horenso-reporter-ciscospark?branch=master)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

![Screenshot](doc/image.png)


[horenso](https://github.com/Songmu/horenso) meets [Cisco Spark](https://www.ciscospark.com/) ❤️

This plugin is strongly inspired by [horenso-reporter-slack](https://github.com/ariarijp/horenso-reporter-slack).

## Usage

```shell
$ go get github.com/sukiyaki/horenso-reporter-ciscospark/cmd/horenso-reporter-ciscospark
$ HRS_SPARK_TOKEN="YOUR_SPARK_TOKEN" HRS_SPARK_ROOMID="YOUR_ROOM_ID" horenso -r ./horenso-reporter-ciscospark -- [command]
$ HRS_SPARK_TOKEN="YOUR_SPARK_TOKEN" HRS_SPARK_TO_EMAIL="YOUR_EMAIL_ADDR" horenso -r ./horenso-reporter-ciscospark -- [command]
```

### Environment Variables

* `HRS_SPARK_TOKEN`(required)
  * [Cisco Spark API Token](https://developer.ciscospark.com/getting-started.html#authentication)
* `HRS_SPARK_ROOMID`(required when `HRS_SPARK_TO_EMAIL` is not provided)
  * The ID of the recipient
  * This ID could be obtained at [here](https://developer.ciscospark.com/endpoint-rooms-get.html)
* `HRS_SPARK_TO_EMAIL`(required when `HRS_SPARK_ROOMID` is not provided)
  * The email address of the recipient
* `HRS_SPARK_ITEMS`(optional defauls: `all`):
  * Select reporting items by Comma-Separated Values
  * example: Stdout,Stderr,ExitCode
  * supported items
    * `Command`
    * `CommandArgs`
    * `Output`
    * `Stdout`
    * `Stderr`
    * `ExitCode`
    * `Result`
    * `Pid`
    * `StartAt`
    * `EndAt`
    * `Hostname`
    * `SystemTime`
    * `UserTime`
* `HRS_SPARK_NOTIFY_EVERYTHING` (optional defauls: `1`)
  * Set `0` if you don't want to raise the report when the command exited with code: 0

## License

MIT

## Author

[Masaki Tagawa](https://github.com/mochipon) ([sukiyaki project](https://www.sukiyaki.ski))
