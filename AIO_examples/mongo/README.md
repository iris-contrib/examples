## Folder information

This folder contains an example integration of Iris with MongoDb using the mgo driver, based on the [Iris commanad line tool](https://github.com/kataras/iris/tree/v6/iris).

### Additionally

Elements on this example:

* Types of handler in main.go (function, api, custom in registerAPI)
* Basic Mongo Integration
* Manage index
* Models example
* Sessions example
* Package with basic hashing functions like create, validate and random token generator

## Notes
If you want to start manually the server, you should run the server from the backend folder, no outside. or use iris-cli
```sh
$cd mongo/backend

$go run main.go
```

or

```sh
$iris run backend/main.go
```

Packages needed to integrate into your code
```sh
go get -u gopkg.in/kataras/iris.v6/iris
go get -u github.com/iris-contrib/middleware/
go get -u gopkg.in/mgo.v2/bson
go get -u gopkg.in/mgo.v2
go get -u golang.org/x/crypto/bcrypt
```

#### Contribution by [ivancduran](https://github.com/ivancduran)
