# Examples

<a href="https://travis-ci.org/iris-contrib/examples"><img src="https://img.shields.io/travis/iris-contrib/adaptors.svg?style=flat-square" alt="Build Status"></a>
<a href="https://github.com/iris-contrib/examples/blob/master/LICENSE"><img src="https://img.shields.io/badge/%20license-MIT%20%20License%20-E91E63.svg?style=flat-square" alt="License"></a>
<a href="https://github.com/kataras/iris/blob/master/HISTORY.md"><img src="https://img.shields.io/badge/version-7%20-blue.svg?style=flat-square" alt="CHANGELOG/HISTORY"></a>



This repository provides easy to understand code snippets on how to get started with web development with the Go programming language using the [Iris](https://github.com/kataras/iris) web framework.

It doesn't contains "best ways" neither explains all its features. It's just a simple, practical cookbook for young Gophers!

## Table of contents

* [Level: Beginner]()
    * [Overview](overview/main.go)
    * [Listening](listening)
        * [Common, with address](listening/listen-addr/main.go)
        * [UNIX socket file](listening/listen-unix/main.go)
        * [TLS](listening/listen-tls)
        * [Letsencrypt (Automatic Certifications)](listening/listen-letsencrypt/main.go)
        * [Custom TCP Listener](listening/custom-listener/main.go)
    * [Configuration](configuration)
        * [Basic way](configuration/basic/main.go)
        * [Functional way](configuration/functional/main.go)
        * [Import from YAML file](configuration/from-yaml-file/main.go)
        * [Import from TOML file](configuration/from-toml-file/main.go)
    * [Routing](routing)
        * [Overview](routing/main.go)
        * [Basic](routing/basic/main.go)
        * [Dynamic Path](routing/dynamic-path/main.go)
        * [Reverse routing](routing/reverse/main.go)
    * [Transform any third-party handler to iris-compatible handler](convert-handlers)
        * [From func(http.ResponseWriter, *http.Request, http.HandlerFunc)](convert-handlers/negroni-like/main.go)
        * [From http.Handler or http.HandlerFunc](convert-handlers/nethttp/main.go)
    * [Internal Application File Logger](file-logger/main.go)
    * [Custom HTTP Errors](http-errors/main.go)
    * [Write JSON](write-json/main.go)
    * [Read JSON](read-json/main.go)
    * [Read Form](read-form/main.go)
    * [Favicon](favicon/main.go)
    * [File Server](file-server/main.go)
    * [Send Files](send-files/main.go)
    * [Stream Writer](stream-writer/main.go)
    * [Send An E-mail](e-mail/main.go)
    * [Upload/Read Files](upload-files/main.go)
    * [Recovery](recover/main.go)
    * [Profiling (pprof)](pprof/main.go)
    * [Request Logger](request-logger/main.go)
    * [Basic Authentication](basicauth/main.go)
* [Level: Intermediate]()
    * [Transactions](transactions/main.go)
    * [HTTP Testing](httptest/main_test.go)
    * [Watch & Compile Typescript source files](typescript/main.go)
    * [Cloud Editor](cloud-editor/main.go)
    * [Serve Embedded Files](serve-embedded-files/main.go)
    * [HTTP Access Control](cors/main.go)
    * [Cache Markdown](cache-markdown/main.go)
    * [Localization and Internationalization](i18n/main.go)
    * [Graceful Shutdown](graceful-shutdown)
        * [Basic and simple](graceful-shutdown/basic/main.go)
        * [Custom Host](graceful-shutdown/custom-host/main.go)
    * [Custom HTTP Server](custom-httpserver)
        * [Iris way](custom-httpserver/iris-way/main.go)
        * [Standar way](custom-httpserver/std-way/main.go)
        * [More than one server](custom-httpserver/multi/main.go)
    * [Custom Context](custom-context)
        * [Method Overriding](custom-context/method-overriding/main.go)
    * [Route State](route-state/main.go)
    * [View Engine](view)
        * [Overview](view/overview/main.go)
        * [Hi](view/template_html_0/main.go)
        * [Showcase one simple Layout](view/template_html_1/main.go)
        * [Layouts `yield` and `render` tmpl funcs](view/template_html_2/main.go)
        * [Showcase of the `urlpath` tmpl func](view/template_html_3/main.go)
        * [Showcase of the `url` tmpl func](view/template_html_4/main.go)
        * [Inject Data Between Handlers](view/context-view-data/main.go)
        * [Embedding Templates Into App Executable File](view/embedding-templates-into-app)
    * [Sessions](sessions)
        * [Overview](sessions/overview/main.go)
        * [Encoding & Decoding the Session ID: Secure Cookie](sessions/securecookie/main.go)
        * [Standalone](sessions/standalone/main.go)
        * [With A Back-End Database](sessions/database/main.go)
        * [Password Hashing](sessions/password-hashing/main.go)
    * [Flash Messages](flash-messages/main.go)
    * [Websockets](websockets)
        * [Ridiculous Simple](websockets/ridiculous-simple/main.go)
        * [Overview](websockets/overview/main.go)
        * [Connection List](websockets/connectionlist/main.go)
        * [Native Messages](websockets/naive-messages/main.go)
        * [Secure](websockets/secure/main.go)
        * [Custom Go Client](websockets/custom-go-client/main.go)
    * [Subdomains](subdomains)
        * [Single](subdomains/single/main.go)
        * [Multi](subdomains/multi/main.go)
        * [Wildcard](subdomains/wildcard/main.go)
* [Level: Advanced]()
    * [Online Visitors](online-visitors/main.go)
    * [URL Shortener using BoltDB](url-shortener/main.go)

> Do not forget to [star or watch the project](https://github.com/kataras/iris/stargazers) in order to stay updated with the latest tech trends, it takes some seconds for the sake of go!


> Examples are tested using Windows 7, Ubuntu 16.04 with [LiteIDE](https://github.com/visualfc/liteide).

## Run

```sh
$ cd $GOPATH/src/github.com/iris-contrib/examples/overview
$ go run main.go
```

## Support 
http://support.iris-go.com

### Older versions

- [Iris version 6/http2](https://github.com/kataras/iris/tree/v6) examples, click [here](https://github.com/kataras/iris/tree/v6/_examples).
- [Iris version 5/fasthttp](https://github.com/kataras/iris/tree/5.0.0) examples, click [here](https://github.com/iris-contrib/examples/tree/5.0.0).
- [Iris version 4/fasthttp](https://github.com/kataras/iris/tree/4.0.0) examples, click [here](https://github.com/iris-contrib/examples/tree/4.0.0).

## Su, 04 June 2017

This repository is just a minor of the github.com/kataras/iris/[version or master]/_examples folder.
