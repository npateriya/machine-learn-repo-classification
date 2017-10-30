# winston-spark
[Cisco Spark][0] transport for [Winston][1]

[![NPM](https://nodei.co/npm/winston-spark.png?downloads=true&downloadRank=true&stars=true)](https://nodei.co/npm/winston-spark/)

[![npm](https://img.shields.io/npm/v/winston-spark.svg)](https://www.npmjs.com/package/winston-spark) [![dependencies Status](https://david-dm.org/joelee/winston-spark/status.svg)](https://david-dm.org/joelee/winston-spark) [![Known Vulnerabilities](https://snyk.io/test/github/joelee/winston-spark/badge.svg)](https://snyk.io/test/github/joelee/winston-spark) [![GitHub issues](https://img.shields.io/github/issues/joelee/winston-spark.svg)](https://github.com/joelee/winston-spark/issues) 

## Installation

``` sh
$ npm install winston
$ npm install winston-spark
```

## Usage
``` js
var winston = require('winston');
require('winston-spark');

var options = {
  accessToken: '***Your Spark Access Token***',
  roomId: '***Spark Room Id***'
};

winston.add(winston.transports.SparkLogger, options);
```

Valid Options are as the following:

* __accessToken__ Your Spark Access Token. *[required]*
* __roomId__ Spark Room Id. *[required]*
* __level__ Log Level (default: info)
* __hideMeta__ Hide MetaData flag (default: false)

## License
The MIT License (MIT)

Copyright (C) 2017 Joseph Lee

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

---

## Contributor(s)

[Joseph Lee](https://github.com/joelee) [![Twitter Follow](https://img.shields.io/twitter/follow/joe_lee.svg?style=social&label=Follow)](https://twitter.com/joe_lee)


[0]: http://ciscospark.com
[1]: https://github.com/flatiron/winston
