[![GoDoc](https://godoc.org/github.com/caboucha/nxgo/nx?status.svg)](http://godoc.org/github.com/caboucha/nxgo/nx)
[![Go Report Card](https://goreportcard.com/badge/github.com/caboucha/nxgo)](https://goreportcard.com/report/github.com/caboucha/nxgo)

About nxgo
===========

nxgo is a Go package for interacting with Cisco Nexus Switch REST API calls.

Usage
=====

1\. Grab the source

    go get github.com/caboucha/nxgo

2\. Get Dependencies

    go get github.com/gorilla/websocket

3\. Set Environment variables to run program

    export NEXUS_HOSTS = "your-nexus-ip-address"
    export NEXUS_USER = "your-nexus-admin-username"
    export NEXUS_PASS = "your-nexus-admin-password"

4\. Import the package in your program

    import "github.com/caboucha/nxgo/nx"

Example
=======

    package main

    import (
        "fmt"
        "github.com/caboucha/nxgo/nx"
    )

    func main() {

        a, errNew := nx.New(nx.ClientOptions{})
        if errNew != nil {
                fmt.Printf("login new client error: %v\n", errNew)
                return
        }

        // Since credentials have not been specified explicitly under ClientOptions,
        // Login() will use env vars: NEXUS_HOSTS=host, NEXUS_USER=username, NEXUS_PASS=pwd
        errLogin := a.Login()
        if errLogin != nil {
                fmt.Printf("login error: %v\n", errLogin)
                return
        }
        defer a.Logout()

        errGet := a.GetInterface("tenant-example", "")
        if errGet != nil {
                fmt.Printf("Get Interface error: %v\n", errAdd)
                return
        }
    }

Documentation
=============

nxgo documentation in GoDoc: https://godoc.org/github.com/caboucha/nxgo/nx

See Also
========

[Cisco NX API Programmability Guide](http://www.cisco.com/c/en/us/td/docs/switches/datacenter/nexus9000/sw/6-x/programmability/guide/b_Cisco_Nexus_9000_Series_NX-OS_Programmability_Guide/b_Cisco_Nexus_9000_Series_NX-OS_Programmability_Configuration_Guide_chapter_0101.pdf)

[Cisco NX API Developer Documentation](https://developer.cisco.com/site/nx-os/docs/apis/nx-api-rest/)
