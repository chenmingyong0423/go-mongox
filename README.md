<p align="center">
  <img src="https://raw.githubusercontent.com/chenmingyong0423/go-mongox-doc/main/docs/public/go-mongox-logo.png" width="200" height="200" akt="go-mongox"></img>
</p>

[![GitHub Repo stars](https://img.shields.io/github/stars/chenmingyong0423/go-mongox)](https://github.com/chenmingyong0423/go-mongox/stargazers)
[![GitHub issues](https://img.shields.io/github/issues/chenmingyong0423/go-mongox)](https://github.com/chenmingyong0423/go-mongox/issues)
[![GitHub License](https://img.shields.io/github/license/chenmingyong0423/go-mongox)](https://github.com/chenmingyong0423/go-mongox/blob/main/LICENSE)
[![GitHub release (with filter)](https://img.shields.io/github/v/release/chenmingyong0423/go-mongox)](https://github.com/chenmingyong0423/go-mongox)
[![Go Report Card](https://goreportcard.com/badge/github.com/chenmingyong0423/go-mongox)](https://goreportcard.com/report/github.com/chenmingyong0423/go-mongox)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/chenmingyong0423/go-mongox)
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)

English | [中文简体](./README-zh_CN.md)

# go-mongox
`go-mongox` is a generics-based library that extends the official `MongoDB` framework. Utilizing generic programming, it facilitates the binding of structs to `MongoDB` collections, aiming to provide type safety and streamlined data operations. `go-mongox` introduces chainable calls for smoother document handling and offers a rich set of `bson` builders and built-in functions to simplify the construction of `bson` data. Moreover, it supports plugin-based programming and incorporates various hooks, offering flexibility for custom logic before and after database operations, thus enhancing the application's extensibility and maintainability.

# Feature Highlights
- Generic MongoDB Collection
- Support for constructing `bson` data
- `CRUD` operations on documents
- Aggregation operations
- Built-in basic `Model` structure for automated updates of default `field` fields
- Struct tag validation
- Hooks
- Plugin programming support

# Install
```go
go get github.com/chenmingyong0423/go-mongox
```

# Getting Started
- `go-mongox` Guides: [https://go-mongox.dev](https://go-mongox.dev/en)

# Contributing
[With your participation, go-mongox will become even more powerful!](https://go-mongox.dev/en/contribute.html)

# Contributors
[Thank you](https://github.com/chenmingyong0423/go-mongox/graphs/contributors) for contributing to the `go-mongox` framework!

# License
© [Mingyong Chen](https://github.com/chenmingyong0423)，2024-now

This project is licensed under the [Apache License](https://github.com/chenmingyong0423/go-mongox/blob/main/LICENSE).
