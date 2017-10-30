# Cisco CDR Parser

At work I had a recent requirement to be able to download and parse Cisco CDR (Call Detail Records) that are provided by our phone system.

This repo is the accumulation of a few hours of playing with the logs and making something that kinda works. Since we've been working a lot with both Python and Angular recently at work, it only seemed fitting to write this in Python and Angular.

It will provide a web server that hosts a single-page site where you can view the logs, sort by all columns, and filter the logs using search terms.

## This is not optimized in any way and will most likely have to be tweaked if you decide to use this code.

# Usage:
This collection of scripts is comprised of two parts
1. A downloader that will connect to an FTP server and download all non-hidden files from the directory
2. The parser that parses the downloaded files and then renders them using bottle.py

Both parts use libaries that should be standard on most Python installations with the exception of bottle.py, which has been included in this repo

Before you can use the downloader, you will need to edit it and set your FTP details inside the config dictionary. It is also advised that you setup a crontab with this file so that new logs will automatically be downloaded.

Once the downloader has been configured, simply run: python downloader.py and it will download all non-hidden files from the path you specify on your FTP server.

The phonelog can be used right away, however you may want to set the page config to something other than "Company name here". To keep this server running indefinitly, there are a few options below.
* use screen, tmux, or some other terminal multiplexor and simply run python phonelog.py in that session
* More complicated: Configure uWSGI or some other CGI server to handle the requests, this will undoubtedly require some changes to the script.


Copyright (c) 2016 David (alopexc0de) Todd <c0de@c0defox.es>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.