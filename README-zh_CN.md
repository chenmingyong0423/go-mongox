<p align="center">
  <img src="https://raw.githubusercontent.com/chenmingyong0423/go-mongox-doc/main/docs/public/go-mongox-logo.png" width="200" height="200" alt="go mongox"></img>
</p>

[![GitHub Repo stars](https://img.shields.io/github/stars/chenmingyong0423/go-mongox)](https://github.com/chenmingyong0423/go-mongox/stargazers)
[![GitHub issues](https://img.shields.io/github/issues/chenmingyong0423/go-mongox)](https://github.com/chenmingyong0423/go-mongox/issues)
[![GitHub License](https://img.shields.io/github/license/chenmingyong0423/go-mongox)](https://github.com/chenmingyong0423/go-mongox/blob/main/LICENSE)
[![GitHub release (with filter)](https://img.shields.io/github/v/release/chenmingyong0423/go-mongox)](https://github.com/chenmingyong0423/go-mongox)
[![codecov](https://codecov.io/gh/chenmingyong0423/go-mongox/graph/badge.svg?token=H3CROTTDZ1)](https://codecov.io/gh/chenmingyong0423/go-mongox)
[![Go Report Card](https://goreportcard.com/badge/github.com/chenmingyong0423/go-mongox)](https://goreportcard.com/report/github.com/chenmingyong0423/go-mongox)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/chenmingyong0423/go-mongox)
[![All Contributors](https://img.shields.io/badge/all_contributors-4-orange.svg?style=flat-square)](#contributors-)

[English](./README.md) | 中文简体

# go mongox
`mongox` 是一个基于泛型的库，扩展了 `MongoDB` 的官方框架。通过泛型技术，它实现了结构体与 `MongoDB` 集合的绑定，旨在提供类型安全和简化的数据操作。`mongox` 还引入链式调用，让文档操作更流畅，并且提供了丰富的 `BSON` 构建器和内置函数，简化了 `BSON` 数据的构建。此外，它还支持插件化编程和内置多种钩子函数，为数据库操作前后的自定义逻辑提供灵活性，增强了应用的可扩展性和可维护性。

# 功能特性
- 泛型的 `MongoDB` 集合
- 文档的 `CRUD` 操作
- 聚合操作
- 内置基本的 `Model` 结构体，自动化更新默认的 `field` 字段
- 支持 `BSON` 数据的构建
- 内置 `Hooks`
- 支持插件化编程

# 安装
```go
go get github.com/chenmingyong0423/go-mongox/v2
```

# 快速开始
- `go mongox` 指南： [https://go-mongox.dev](https://go-mongox.dev)

# 贡献
[如果有您的加入，go-mongox 将会变得更加强大！](https://go-mongox.dev/contribute.html)

# 贡献者
非常感谢 [您们](https://github.com/chenmingyong0423/go-mongox/graphs/contributors) 为 `go mongox` 框架做出的贡献！

# 社区

加入我们的社区获取支持、分享想法、参与讨论！

[![加入 Discord](https://img.shields.io/badge/Discord-加入聊天-5865F2?logo=discord&logoColor=white)](https://discord.gg/Cx8CMcRurb)

# 版权
© [陈明勇](https://github.com/chenmingyong0423)，2024-至今

这个项目遵循 [Apache License](https://github.com/chenmingyong0423/go-mongox/blob/main/LICENSE) 许可。
