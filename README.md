# User/Track API

This is the main Resonate User/Track API written in Go, not yet in production.
It uses [Twirp](https://github.com/twitchtv/twirp), a RPC framework for service-to-service communication emphasizing simplicity and minimalism. Learn more about
Twirp at its [website](https://twitchtv.github.io/twirp/docs/intro.html).
It also uses [go-pg](https://github.com/go-pg/pg) PostgreSQL ORM.
Its structure is based on the Twirp starter kit [Twisk](https://github.com/ribice/twisk).


## Project Structure

The project structure mostly follows [THIS](https://github.com/golang-standards/project-layout) example repository and [Twirp best practices](https://twitchtv.github.io/twirp/docs/best_practices.html), except for the services that live in `internal/server/<service>`.

## Getting started

### Prerequisites
- [Go](https://golang.org/) 1.7 or higher.
- [PostgreSQL](https://www.postgresql.org/) 9.4 or higher.

### Dev database setup

* Create user and database as follow (as found in the local config file in `./conf.local.yaml`):

username = "resonate-dev-user"

password = "password"

dbname = "resonate-dev"

Add following postgres extensions: "hstore" and "uuid-ossp"

* Run migrations from `./cmd/migration`

```sh
$ go run *.go
```

### Dependencies

[Dep](https://github.com/golang/dep) is used as dependency management tool.
`vendor/` folder contains project dependencies and should be in sync with `Gopkg.toml` and `Gopkg.lock`.

### Various tools installation for development

* [Install Protocol Buffers v3](https://developers.google.com/protocol-buffers/docs/gotutorial),
the `protoc` compiler that is used to auto-generate code. The simplest way to do
this is to download pre-compiled binaries for your platform from here:
https://github.com/google/protobuf/releases

It is also available in MacOS through Homebrew:

```sh
$ brew install protobuf
```

* Install [retool](https://github.com/twitchtv/retool). It helps manage go tools like commands and linters.
protoc-gen-go and protoc-gen-twirp plugins were installed into `_tools` folder using retool.

Build the generators and tool dependencies:
```sh
$ retool build
```

Then, to run the `protoc` command and autogenerate Go code for the server interface, make sure to prefix with `retool do`, for example:
```sh
$ retool do protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/user/service.proto
```

### Running/building the server

First, put this repo into `$GOPATH/src`

Then, run the server
```sh
$ go run ./cmd/api/main.go
```

Alternatively, you can build and run an executable binary
```sh
$ cd ./cmd/api/
$ go build
$ ./api
```

### Implementing new services

1. Under `proto` folder, create a new one named after your service.

2. Define your proto file. If you are not familiar with Protobufs, you can read more about it [here](https://developers.google.com/protocol-buffers/docs/proto3). You can use already existing proto files (eg `rpc/user/service.proto`) as a template.

3. Run `make twirp -B`.

4. Implement the service interface from `service.twirp.go` in `internal/server/tenant`.

5. Wire up everything in `cmd/api/main.go`.

## Documentation

Check out `doc/` folder for API documentation.
But we'll be transitioning to [OpenAPI Specification and Swagger](https://swagger.io/docs/specification/about/).

## Testing

We use Ginkgo and Gomega for testing.

You need to create the testing database and run migrations manually before running tests.

* Create user and database as follow:

username = "resonate-testing-user"

password = ""

dbname = "resonate-testing"

Add following extensions: "hstore" and "uuid-ossp" (TODO: add them on initial migration)

* Run migrations from `./cmd/migration`

```sh
$ go run *.go testing
```

* Run tests `./internal/server/<service>`

```sh
$ go test
```

Or run all tests using ginkgo CLI from `./`

```sh
$ ginkgo -r
```

## Roadmap

- Switch from API Blueprint to OpenAPI and autogenerated API documentation using SwaggerUI
- Implement logging interfaces for every service
- Add JWT based authentication and role-based access control (see `internal/iam`) to existing services

## Contributing

Please check out the [Contributing guide](CONTRIBUTING.md) for guidelines about how to proceed.

We expect contributors to abide by our underlying [code of conduct](CODE_OF_CONDUCT.md).
