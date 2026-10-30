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

## 软删除

模型包含 `DeletedAt` 字段时，`DeleteOne` / `DeleteMany` 自动通过更新删除时间进行软删除。内嵌 `mongox.Model` 时请使用 `bson:",inline"`：

```go
type User struct {
	mongox.Model `bson:",inline"`
	Name string `bson:"name"`
}
```

- `DeletedAt` 支持 `time.Time`、`int`、`int64`；整数默认为秒。自定义字段名或时间精度可使用 `mongox:"autoDeleteTime"`、`autoDeleteTime:second`、`autoDeleteTime:milli`、`autoDeleteTime:nano`，精度标签仅适用于整数。
- 字段缺失、`null` 或类型对应的零值均表示未删除。插入时不会自动设置删除时间；软删除仅设置删除时间。
- `FindOne`、`Find`、`Count`、`Distinct`、`DistinctWithParse`、`FindOneAndUpdate`、`UpdateOne`、`UpdateMany` 和 `Upsert` 默认只匹配未删除文档。重复删除返回的 `DeletedCount` 为零。
- Finder 支持 `WithDeleted()` 查询全部文档（包括已删除），以及 `OnlyDeleted()` 仅查询已删除文档。这两个方法适用于 Finder 的全部操作，包括 `Count`、`Distinct` 和 `FindOneAndUpdate`；连续调用时以后一次为准，设置保留在当前 Finder 上，新建 Finder 恢复默认范围。
- 没有有效软删除字段的模型，`WithDeleted()` 保持原查询条件，`OnlyDeleted()` 不匹配任何文档。
- `Upsert` 不会恢复已删除文档；它可能插入新文档，或因已有文档占用唯一键而失败。
- 没有有效软删除字段的模型保持物理删除行为。不支持的类型或时间精度不会启用软删除；`bson:"-"` 字段不参与处理。
- 聚合管道和通过 `Collection()` 获取的官方集合保持原生行为，需要调用方自行指定过滤条件。永久删除、恢复等特殊操作也可通过官方集合完成。

已有模型内嵌 `mongox.Model` 的应用，升级后其删除和上述查询操作会自动采用软删除语义。

```go
// 包含已删除文档，仍保留业务过滤条件。
users, err := collection.Finder().WithDeleted().Filter(filter).Find(ctx)

// 仅统计已删除文档。
count, err := collection.Finder().OnlyDeleted().Filter(filter).Count(ctx)

// 更新已删除文档，例如恢复 time.Time 类型的删除时间。
user, err := collection.Finder().OnlyDeleted().Filter(filter).
	Updates(bson.M{"$set": bson.M{"deleted_at": time.Time{}}}).
	FindOneAndUpdate(ctx)
```

软删除的执行参数通过 `UpdateOneOptions` / `UpdateManyOptions` 配置，直接使用官方 Update options：

```go
_, err := collection.Deleter().
	Filter(filter).
	UpdateOneOptions(options.UpdateOne().SetHint("name_1")).
	DeleteOne(ctx)
```

`DeleteOne` / `DeleteMany` 的 Delete options 仅用于物理删除，不会转换为 Update options。未配置 Update options 时使用官方默认值；每次链式配置替换此前的配置，传入空参数可清除配置。删除方式仍由模型字段自动确定。

# 贡献
[如果有您的加入，go-mongox 将会变得更加强大！](https://go-mongox.dev/contribute.html)

# 贡献者
非常感谢 [您们](https://github.com/chenmingyong0423/go-mongox/graphs/contributors) 为 `go mongox` 框架做出的贡献！

# 社区

加入我们的社区获取支持、分享想法、参与讨论！

[![加入 Discord](https://img.shields.io/badge/Discord-加入聊天-5865F2?logo=discord&logoColor=white)](https://discord.gg/Cx8CMcRurb)

# 贡献者

衷心感谢所有为本项目做出贡献的开发者！  
[![贡献值](https://contrib.rocks/image?repo=chenmingyong0423/go-mongox)](https://github.com/chenmingyong0423/go-mongox/graphs/contributors)

你们的支持和努力使这个项目变得更加出色！

# 致谢

[![JetBrains logo.](https://resources.jetbrains.com/storage/products/company/brand/logos/jetbrains.svg)](https://jb.gg/OpenSourceSupport)

感谢 JetBrains 通过其 [开源许可支持计划](https://www.jetbrains.com/community/opensource/) 对本项目的支持。

# 版权
© [陈明勇](https://github.com/chenmingyong0423)，2024-至今

这个项目遵循 [Apache License](https://github.com/chenmingyong0423/go-mongox/blob/main/LICENSE) 许可。
