# Examples

<a href="https://iris-go.com"> <img align="right" width="115px" src="https://iris-go.com/images/icon.svg?v=a" title="logo created by @merry.dii" /> </a>

<a href="https://travis-ci.org/iris-contrib/examples"><img src="https://img.shields.io/travis/iris-contrib/examples.svg?style=flat-square" alt="Build Status"></a>
<a href="https://github.com/iris-contrib/examples/blob/master/LICENSE"><img src="https://img.shields.io/badge/%20license-MIT%20%20License%20-E91E63.svg?style=flat-square" alt="License"></a>
<a href="https://github.com/kataras/iris/blob/v11/HISTORY.md"><img src="https://img.shields.io/badge/version-11.2%20-blue.svg?style=flat-square" alt="CHANGELOG/HISTORY"></a>

This repository provides easy to understand code snippets on how to get started with web development with the Go programming language using the [Iris](https://github.com/kataras/iris) web framework.

To read the Iris documentation please navigate to [the wiki pages](https://github.com/kataras/iris/wiki) instead.

## Table of Contents

* [Overview](overview)
    * [Hello world!](hello-world/main.go)
    * [Hello WebAssemply!](https://github.com/kataras/iris/blob/master/_examples/webassembly/basic/main.go)
    * [Glimpse](overview/main.go)
    * [Tutorial: Online Visitors](tutorial/online-visitors/main.go)
    * [Tutorial: A Todo MVC Application using Iris and Vue.js](https://hackernoon.com/a-todo-mvc-application-using-iris-and-vue-js-5019ff870064)
    * [Tutorial: URL Shortener using BoltDB](https://medium.com/@kataras/a-url-shortener-service-using-go-iris-and-bolt-4182f0b00ae7)
    * [Tutorial: How to turn your Android Device into a fully featured Web Server](https://twitter.com/ThePracticalDev/status/892022594031017988)
    * [POC: Convert the medium-sized project "Parrot" from native to Iris](https://github.com/iris-contrib/parrot)
    * [POC: Isomorphic react/hot reloadable/redux/css-modules starter kit](https://github.com/kataras/iris-starter-kit)
    * [Tutorial: DropzoneJS Uploader](tutorial/dropzonejs)
    * [Tutorial: Caddy](tutorial/caddy)
    * [Tutorial:Iris Go Framework + MongoDB](https://medium.com/go-language/iris-go-framework-mongodb-552e349eab9c)
    * [Tutorial: API for Apache Kafka](tutorial/api-for-apache-kafka)
* [Structuring](structuring)
    * [Bootstrapper](structuring/bootstrap)
    * [MVC with Repository and Service layer Overview](structuring/mvc-plus-repository-and-service-layers)
    * [Login (MVC with Single Responsibility package)](structuring/login-mvc-single-responsibility-package)
    * [Login (MVC with Datamodels, Datasource, Repository and Service layer)](structuring/login-mvc)
* [HTTP Listening](http-listening)
    * [Common, with address](http-listening/listen-addr/main.go)
        * [public domain address](http-listening/listen-addr-public/main.go) **NEW**
        * [omit server errors](http-listening/listen-addr/omit-server-errors/main.go)
    * [UNIX socket file](http-listening/listen-unix/main.go)
    * [TLS](http-listening/listen-tls/main.go)
    * [Letsencrypt (Automatic Certifications)](http-listening/listen-letsencrypt/main.go)
    * [Notify on shutdown](http-listening/notify-on-shutdown/main.go)
    * Custom TCP Listener
        * [common net.Listener](http-listening/custom-listener/main.go)
        * [SO_REUSEPORT for unix systems](http-listening/custom-listener/unix-reuseport/main.go)
    * Custom HTTP Server
        * [HTTP/3 Quic](http-listening/http3-quic) **NEW**
        * [easy way](http-listening/custom-httpserver/easy-way/main.go)
        * [std way](http-listening/custom-httpserver/std-way/main.go)
        * [multi server instances](http-listening/custom-httpserver/multi/main.go)
    * Graceful Shutdown
        * [using the `RegisterOnInterrupt`](http-listening/graceful-shutdown/default-notifier/main.go)
        * [using a custom notifier](http-listening/graceful-shutdown/custom-notifier/main.go)
* [Configuration](configuration)
    * [Functional](configuration/functional/main.go)
    * [From Configuration Struct](configuration/from-configuration-structure/main.go)
    * [Import from YAML file](configuration/from-yaml-file/main.go)
        * [Share Configuration between multiple instances](configuration/from-yaml-file/shared-configuration/main.go)
    * [Import from TOML file](configuration/from-toml-file/main.go)
* [Routing](routing)
    * [Overview](routing/overview/main.go)
    * [Basic](routing/basic/main.go)
    * [Controllers](mvc)
    * [Custom HTTP Errors](routing/http-errors/main.go)
    * [Dynamic Path](routing/dynamic-path/main.go)
        * [root level wildcard path](routing/dynamic-path/root-wildcard/main.go)
    * [Write your own custom parameter types](routing/macros/main.go)
    * [Reverse routing](routing/reverse/main.go)
    * [Custom Router (high-level)](routing/custom-high-level-router/main.go)
    * [Custom Wrapper](routing/custom-wrapper/main.go) **UPDATED**
    * Custom Context
        * [method overriding](routing/custom-context/method-overriding/main.go)
        * [new implementation](routing/custom-context/new-implementation/main.go)
    * [Route State](routing/route-state/main.go)
    * [Writing a middleware](routing/writing-a-middleware)
        * [per-route](routing/writing-a-middleware/per-route/main.go)
        * [globally](routing/writing-a-middleware/globally/main.go)
* [Versioning](versioning)
    * [How it works](https://github.com/kataras/iris/blob/master/versioning/README.md)
    * [Example](versioning/main.go)
* [Dependency Injection](hero)
    * [Basic](hero/basic/main.go)
    * [Overview](hero/overview)
    * [Sessions](hero/sessions)
    * [Yet another dependency injection example and good practises at general](hero/smart-contract/main.go) **NEW**
* [MVC](mvc)
    * [Hello world](mvc/hello-world/main.go)
    * [Regexp](mvc/regexp/main.go) **NEW**
    * [Session Controller](mvc/session-controller/main.go)
    * [Overview - Plus Repository and Service layers](mvc/overview)
    * [Login showcase - Plus Repository and Service layers](mvc/login)
    * [Singleton](mvc/singleton)
    * [Websocket Controller](mvc/websocket) **UPDATED**
    * [Register Middleware](mvc/middleware)
    * [Vue.js Todo MVC](tutorial/vuejs-todo-mvc)
* [Subdomains](subdomains)
    * [Single](subdomains/single/main.go)
    * [Multi](subdomains/multi/main.go)
    * [Wildcard](subdomains/wildcard/main.go)
    * [WWW](subdomains/www/main.go)
    * [Redirect fast](subdomains/redirect/main.go)
* [Convert `http.Handler/HandlerFunc`](convert-handlers)
    * [From func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)](convert-handlers/negroni-like/main.go)
    * [From http.Handler or http.HandlerFunc](convert-handlers/nethttp/main.go)
    * [From func(http.HandlerFunc) http.HandlerFunc](convert-handlers/real-usecase-raven/writing-middleware/main.go)
* [View](view)
    * [Overview](view/overview/main.go)
    * [Hi](view/template_html_0/main.go)
    * [A simple Layout](view/template_html_1/main.go)
    * [Layouts: `yield` and `render` tmpl funcs](view/template_html_2/main.go)
    * [The `urlpath` tmpl func](view/template_html_3/main.go)
    * [The `url` tmpl func](view/template_html_4/main.go)
    * [Inject Data Between Handlers](view/context-view-data/main.go)
    * [Embedding Templates Into App Executable File](view/embedding-templates-into-app/main.go)
    * [Write to a custom `io.Writer`](view/write-to)
    * [Greeting with Pug (Jade)`](view/template_pug_0)
    * [Pug (Jade) Actions`](view/template_pug_1)
    * [Pug (Jade) Includes`](view/template_pug_2)
    * [Pug (Jade) Extends`](view/template_pug_3)
    * [Jet](/view/template_jet_0) **NEW**
    * [Jet Embedded](view/template_jet_1_embedded) **NEW**
* [Authentication](authentication)
    * [Basic Authentication](authentication/basicauth/main.go)
    * [OAUth2](authentication/oauth2/main.go)
    * [Request Auth(JWT)](experimental-handlers/jwt/main.go)
    * [Sessions](#sessions)
* [File Server](file-server)
    * [Favicon](file-server/favicon/main.go)
    * [Basic](file-server/basic/main.go) **UPDATED**
    * [Embedding Files Into App Executable File](file-server/embedding-files-into-app/main.go) **UPDATED**
    * [Embedding Gziped Files Into App Executable File](file-server/embedding-gziped-files-into-app/main.go) **UPDATED**
    * [Send/Force-Download Files](file-server/send-files/main.go)
    * Single Page Applications
        * [single Page Application](file-server/single-page-application/basic/main.go) **UPDATED**
        * [embedded Single Page Application](file-server/single-page-application/embedded-single-page-application/main.go) **UPDATED**
        * [embedded Single Page Application with other routes](file-server/single-page-application/embedded-single-page-application-with-other-routes/main.go) **UPDATED**
* [How to Read from `context.Request() *http.Request`](http_request)
    * [Read JSON](http_request/read-json/main.go)
        * [Struct Validation](http_request/read-json-struct-validation/main.go)
    * [Read XML](http_request/read-xml/main.go)
    * [Read YAML](http_request/read-yaml/main.go) **NEW**
    * [Read Form](http_request/read-form/main.go)
    * [Read Query](http_request/read-query/main.go) **NEW**
    * [Read Custom per type](http_request/read-custom-per-type/main.go)
    * [Read Custom via Unmarshaler](http_request/read-custom-via-unmarshaler/main.go)
    * [Read Many times](http_request/read-many/main.go)
    * [Upload/Read File](http_request/upload-file/main.go)
    * [Upload multiple files with an easy way](http_request/upload-files/main.go)
    * [Extract referrer from "referer" header or URL query parameter](http_request/extract-referer/main.go)
* [How to Write to `context.ResponseWriter() http.ResponseWriter`](http_responsewriter)
    * [Content Negotiation](http_responsewriter/content-negotiation) **NEW**
    * [Write `valyala/quicktemplate` templates](http_responsewriter/quicktemplate)
    * [Write `shiyanhui/hero` templates](http_responsewriter/herotemplate)
    * [Text, Markdown, HTML, JSON, JSONP, XML, Binary](http_responsewriter/write-rest/main.go)
    * [Write Gzip](http_responsewriter/write-gzip/main.go)
    * [Stream Writer](http_responsewriter/stream-writer/main.go)
    * [Transactions](http_responsewriter/transactions/main.go)
    * [SSE](http_responsewriter/sse/main.go)
    * [SSE (third-party package usage for server sent events)](http_responsewriter/sse-third-party/main.go)
* [ORM](orm)
    * [Using xorm(Mysql, MyMysql, Postgres, Tidb, **SQLite**, MsSql, MsSql, Oracle)](orm/xorm/main.go)
    * [Using gorm](orm/gorm/main.go)
* [Miscellaneous](miscellaneous)
    * [HTTP Method Override](https://github.com/kataras/iris/blob/master/middleware/methodoverride/methodoverride_test.go) **NEW**
    * [Request Logger](http_request/request-logger/main.go)
        * [log requests to a file](http_request/request-logger/request-logger-file/main.go)
    * [Localization and Internationalization](miscellaneous/i18n/main.go)
    * [Recovery](miscellaneous/recover/main.go)
    * [Profiling (pprof)](miscellaneous/pprof/main.go)
    * [Internal Application File Logger](miscellaneous/file-logger/main.go)
    * [Google reCAPTCHA](miscellaneous/recaptcha/main.go) 
* [Experimental Handlers](experimental-handlers)
    * [Casbin wrapper](experimental-handlers/casbin/wrapper/main.go)
    * [Casbin middleware](experimental-handlers/casbin/middleware/main.go)
    * [Cloudwatch](experimental-handlers/cloudwatch/simple/main.go)
    * [CORS](experimental-handlers/cors/simple/main.go)
    * [JWT](experimental-handlers/jwt/main.go)
    * [Newrelic](experimental-handlers/newrelic/simple/main.go)
    * [Prometheus](experimental-handlers/prometheus/simple/main.go)
    * [Secure](experimental-handlers/secure/simple/main.go)
    * [Tollboothic](experimental-handlers/tollboothic/limit-handler/main.go)
    * [Cross-Site Request Forgery Protection](experimental-handlers/csrf/main.go)
* [Automated API Documentation](apidoc)
    * [yaag](apidoc/yaag/main.go)
* [Testing](testing)
 * [Example](testing/httptest/main_test.go)
* [Caching](cache)
    * [Simple](cache/simple/main.go)
    * [Client-Side (304)](cache/client-side/main.go)
* [Cookies](cookies)
    * [Basic](cookies/basic/main.go)
    * [Encode/Decode (securecookie)](cookies/securecookie/main.go)
[Sessions](sessions)
    * [Overview](sessions/overview/main.go)
    * [Middleware](sessions/middleware/main.go)
    * [Secure Cookie](sessions/securecookie/main.go)
    * [Flash Messages](sessions/flash-messages/main.go)
    * [Databases](sessions/database)
        * [Badger](sessions/database/badger/main.go)
        * [BoltDB](sessions/database/boltdb/main.go)
        * [Redis](sessions/database/redis/main.go)
* [Websockets](websocket)
    * [Basic](websocket/basic) **NEW**
        * [Server](websocket/basic/server.go)
        * [Go Client](websocket/basic/go-client/client.go)
        * [Browser Client](websocket/basic/browser/index.html)
        * [Browser NPM Client (browserify)](websocket/basic/browserify/app.js)
    * [Native Messages](websocket/native-messages/main.go) **UPDATED**
    * [TLS Enabled](websocket/secure/README.md)

> Examples are tested using Windows 10, Ubuntu 16.10 with [Microsoft's Visual Studio Code](https://code.visualstudio.com/) and built using the [Go 1.9](https://golang.org/dl).

## Run

1. Install the Go Programming Language, version 1.12 from [here](https://golang.org/dl).
2. Install Iris: `go get github.com/kataras/iris@master`
3. [Download the examples](https://github.com/iris-contrib/examples/archive/master.zip) and copy-paste them to your `$GOPATH/src/github.com/iris-contrib/examples`

And run

```sh
$ cd $GOPATH/src/github.com/iris-contrib/examples/overview
$ go run main.go
```

Do not forget to [star or watch the Iris project](https://github.com/kataras/iris/stargazers).

## Any troubles with examples?

    https://github.com/iris-contrib/examples/issues

## Su, 04 June 2017

This repository is just a minor of the https://github.com/kataras/iris/master/_examples folder.
