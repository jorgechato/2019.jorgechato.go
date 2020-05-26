# Personal web page API (Work in progress)
[![Build Status](https://travis-ci.com/jorgechato/api.jorgechato.com.svg?token=x3vLcsQVEzf1kfJyx1Uv&branch=master)](https://travis-ci.com/jorgechato/api.jorgechato.com)
[![Docker](https://img.shields.io/badge/docker-image-blue.svg)](https://hub.docker.com/r/jorgechato/api.jorgechato.com/)
[![Go Report Card](https://goreportcard.com/badge/github.com/jorgechato/api.jorgechato.com)](https://goreportcard.com/report/github.com/jorgechato/api.jorgechato.com)
[![Godoc](https://img.shields.io/badge/go-documentation-blue.svg)](https://godoc.org/github.com/jorgechato/api.jorgechato.com)
[![whereisjorge.today](https://img.shields.io/badge/web-whereisjorge.today-orange.svg)](https://whereisjorge.today)

## Must have

- [Go>=14.x](https://golang.org/)

## Description

I've been researching the proper way to structure code for the last 4 years.
Currently I'm into a clean code architecture, more specific testing the
**Hexagonal Architecture**.

I run into [this 101 article](https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/) which describe perfectly the DDD, clean code and
hexagonal architecture.

## Structure

Following the basic *go* structure, **cmd** hosted the entry points for each port.
In the **pkg** folder you can find the source code of the project, where the
magic happen.

The different components are split into packages following the Who + What
principle. In the next example we only have one who (location) and 4 what
(application, domain, infrastructure, repository) as described in the hexagonal
architecture.

```bash
.
├── build # core build misc (dockerfile, k8s.yml...)
├── cmd # entry point
│   ├── http
│   └── lambda
├── pkg # source code
│   └── location # who
│       ├── application # use cases
│       │   └── json
│       ├── domain # business logic
│       ├── infrastructure # ports
│       │   ├── api
│       │   ├── lambda
│       │   └── polarsteps
│       └── repository # adapters
└── scripts # build scripts (.sh...)
```
