[![GitHub Repo stars](https://img.shields.io/github/stars/chenmingyong0423/go-mongox)](https://github.com/chenmingyong0423/go-mongox/stargazers)
[![GitHub issues](https://img.shields.io/github/issues/chenmingyong0423/go-mongox)](https://github.com/chenmingyong0423/go-mongox/issues)
[![GitHub License](https://img.shields.io/github/license/chenmingyong0423/go-mongox)](https://github.com/chenmingyong0423/go-mongox/blob/main/LICENSE)
[![GitHub release (with filter)](https://img.shields.io/github/v/release/chenmingyong0423/go-mongox)](https://github.com/chenmingyong0423/go-mongox)
[![Go Report Card](https://goreportcard.com/badge/github.com/chenmingyong0423/go-mongox)](https://goreportcard.com/report/github.com/chenmingyong0423/go-mongox)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/chenmingyong0423/go-mongox)
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)

<h1 align="center">
  go-mongox
</h1>

<p>go-mongox 基于 泛型 对 MongoDB 官方框架进行了二次封装，它通过使用链式调用的方式，让我们能够丝滑地操作文档。同时，其还提供了多种类型的 bson 构造器，帮助我们高效的构建 bson 数据。</p>


> **功能持续更新和改进中，对该框架感兴趣的伙伴，欢迎提出宝贵的意见和参与贡献。**

---

[English](./README.md) | 中文简体

# 介绍
`go-mongox` 框架有两个核心，一个核心是基于泛型的 **collection** 形态，另一个核心是 **builder** 构造器。

- 通过 **collection** 对象，我们可以方便地进行相关的 `MongoDB` 操作，从而减少 `bson` 数据的编写，提高开发效率；
- 通过 **builder** 构造器，我们可以构造出我们所需要的 `bson` 数据。

# 安装

> go get github.com/chenmingyong0423/go-mongox@latest

# Collection 形态
创建一个基于泛型类型 `Post` 的 `Collection` 实例。
```go

// 示例代码，不是最佳的创建方式
func newCollection() *mongo.Collection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:   "test",
		Password:   "test",
		AuthSource: "db-test",
	}))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	collection := client.Database("db-test").Collection("test_post")
	return collection
}

type Post struct {
	Id      string `bson:"_id"`
	Title   string `bson:"title"`
        Author  string `bson:"author"`
	Content string `bson:"content"`
}


// 你需要预先创建一个 *mongo.Collection 对象
mongoCollection := newCollection()
// 使用 Post 结构体作为泛型参数创建一个 collection
postCollection := mongox.NewCollection[Post](mongoCollection)
```
## Creator
`Creator` 是一个创造器，用于执行插入相关的操作。

```go
// 插入一个文档
doc := Post{Id: "1", Title: "go 语言 go-mongox 库的使用教程", Author: "陈明勇", Content: "go-mongox 旨在提供更方便和高效的MongoDB数据操作体验。"}
oneResult, err := postCollection.Creator().InsertOne(context.Background(), doc)
// 携带 option 参数
oneResult, err := postCollection.Creator().OneOptions(options.InsertOne().SetComment("test")).InsertOne(context.Background(), Post{Id: "1", Title: "go 语言 go-mongox 库的使用教程", Author: "陈明勇", Content: "go-mongox 旨在提供更方便和高效的MongoDB数据操作体验。"})

// 插入多个文档
docs := []Post{
	{Id: "2", Title: "go", Author: "陈明勇", Content: "..."},
	{Id: "3", Title: "mongo", Author: "陈明勇", Content: "..."},
}
manyResult, err := postCollection.Creator().InsertMany(context.Background(), docs)
// 携带 option 参数
manyResult, err := postCollection.Creator().ManyOptions(options.InsertMany().SetComment("test")).InsertMany(context.Background(), docs)
```
通过 `Creator()` 方法创建一个创造器，基于该创造器我们可以进行相关的插入文档操作。其中 `OneOptions` 和 `ManyOptions` 方法用于设置 `*options.InsertOneOptions` 和 `*options.InsertManyOptions` 参数。

## Finder
`Finder` 是一个查询器，用于执行查询相关的操作。

```go
// 查询单个文档
post, err := postCollection.Finder().Filter(bsonx.Id("1")).FindOne(context.Background())

// 设置 *options.FindOneOptions 参数
post, err := postCollection.Finder().
	Filter(bsonx.Id("1")).
	OneOptions(options.FindOne().SetProjection(bsonx.M("content", 0))).
	FindOne(context.Background())

// - map 作为 filter 条件
post, err := postCollection.Finder().Filter(map[string]any{"_id": "1"}).FindOne(context.Background())

// - 复杂条件查询
// -- 使用 query 包构造复杂的 bson: bson.D{bson.E{Key: "title", Value: bson.M{"$eq": "go"}}, bson.E{Key: "author", Value: bson.M{"$eq": "陈明勇"}}}
post, err := postCollection.Finder().
	Filter(query.BsonBuilder().Eq("title", "go").Eq("author", "陈明勇").Build()).
	FindOne(context.Background())

// 查询多个文档
// bson.D{bson.E{Key: "_id", Value: bson.M{"$in": []string{"1", "2"}}}}
posts, err := postCollection.Finder().Filter(query.BsonBuilder().InString("_id", []string{"1", "2"}...).Build()).Find(context.Background())

// 设置 *options.FindOptions 参数
// bson.D{bson.E{Key: "_id", Value: bson.M{types.In: []string{"1", "2"}}}}
posts, err := postCollection.Finder().
	Filter(query.BsonBuilder().InString("_id", []string{"1", "2"}...).Build()).
	Options(options.Find().SetProjection(bsonx.M("content", 0))).
	Find(context.Background())
```
通过 `Finder()` 方法创建一个查询器，基于该查询器我们可以进行相关的查询操作。设置查询条件的方法是`Filter`，设置 `*options.FindOneOptions` 参数的方法有：`OneOptions` 和 `Options`。

如果查询条件 `filter` 非常简单，我们可以通过 `bsonx.M` 方法进行构造。如果 `filter` 过于复杂，我们可以使用 `query` 包里的 `BsonBuilder()` 构造器进行构造。 

## Updater
`Updater` 是一个更新器，用于执行更新相关的操作。
```go
// 更新单个文档
// 通过 update 包构建 bson 更新语句
updateResult, err := postCollection.Updater().
	Filter(bsonx.Id("1")).
	Updates(update.BsonBuilder().Set(bsonx.M("title", "golang")).Build()).
	UpdateOne(context.Background())

// - 使用 map 构造更新数据，并设置 *options.UpdateOptions，执行 upsert 操作
updateResult, err := postCollection.Updater().
	Filter(bsonx.Id("4")).
	UpdatesWithOperator(types.Set, map[string]any{"title": "golang"}).Options(options.Update().SetUpsert(true)).
	UpdateOne(context.Background())

// 更新多条文档
updateResult, err := postCollection.Updater().
	Filter(query.BsonBuilder().InString("_id", []string{"2", "3"}...).Build()).
	Updates(update.BsonBuilder().Set(bsonx.M("title", "golang")).Build()).
	UpdateMany(context.Background())
```
通过 `Updater()` 方法创建一个更新器，基于该更新器我们可以进行相关的更新操作。设置文档匹配条件的方法是 `Filter`，设置更新内容的方法有：`Updates` 和 `UpdatesWithOperator`，设置 `*options.UpdateOptions` 参数的方法有：`Options`。

其中 `UpdatesWithOperator` 方法需要我们指定操作符。
## Deleter
`Deleter` 是一个删除器，用于执行删除相关的操作。
```go
// 删除单个文档
deleteResult, err := postCollection.Deleter().Filter(bsonx.Id("1")).DeleteOne(context.Background())
// 携带 option 参数
deleteResult, err := postCollection.Deleter().Filter(bsonx.Id("1")).Options(options.Delete().SetComment("test")).DeleteOne(context.Background())

// 删除多个文档
// - 通过 query 包构造复杂的 bson 语句
deleteResult, err = postCollection.Deleter().Filter(query.BsonBuilder().InString("_id", []string{"2", "3", "4"}...).Build()).DeleteMany(context.Background())
```
通过 `Deleter()` 方法创建一个删除器，基于该删除器我们可以进行相关的删除操作。设置文档匹配条件的方法有：`Filter`，设置 `*options.DeleteOptions` 参数的方法有：`Options`。

## Aggregator
`Aggregator` 是一个聚合器，用于执行聚合相关的操作。
```go
// 聚合操作
// - 使用 aggregation 包构造 pipeline
posts, err = postCollection.
	Aggregator().Pipeline(aggregation.StageBsonBuilder().Project(bsonx.M("content", 0)).Build()).
	Aggregate(context.Background())


// 如果我们通过聚合操作更改字段的名称，那么我们可以使用 AggregationWithCallback 方法，然后通过 callback 函数将结果映射到我们预期的结构体中
type DiffPost struct {
	Id          string `bson:"_id"`
	Title       string `bson:"title"`
	Name        string `bson:"name"` // author → name
	Content     string `bson:"content"`
	Outstanding bool   `bson:"outstanding"`
}
result := make([]*DiffPost, 0)
//  将 author 字段更名为 name，排除 content 字段，添加 outstanding 字段，返回结果为 []*DiffPost
err = postCollection.Aggregator().
	Pipeline(aggregation.StageBsonBuilder().Project(
		bsonx.NewD().
			Add("name", "$author").
			Add("author", 1).
			Add("_id", 1).
			Add("title", 1).
			Add("outstanding", aggregation.BsonBuilder().Eq("$author", "陈明勇").Build()).Build(),
	).Build(),
	).
	AggregateWithCallback(context.Background(), func(ctx context.Context, cursor *mongo.Cursor) error {
		return cursor.All(ctx, &result)
	})
```
通过 `Aggregator()` 方法创建一个聚合器，基于该聚合器我们可以进行相关的聚合操作。执行聚合操作有两个方法：`Aggregation` 和 `AggregationWithCallback`。`AggregationWithCallback` 可以通过回调将查询的结果输出到指定的切片里。

通过 `aggregation.StageBsonBuilder` 可以构造包含 `stage` 阶段的操作符的 `bson` 数据；通过 `aggregation.BsonBuilder` 可以构造包含除了 `stage` 阶段的操作符的 `bson` 数据。
# Builder
`go-mongox` 框架提供了以下几种类型的构造器：

- `universal`: 简单而又通用的 `bson` 数据构造函数。
- `query`: 查询构造器，用于构造查询操作所需的 `bson` 数据。
- `update`: 更新构造器，用于构造更新操作所需的 `bson` 数据。
- `aggregation`: 聚合操作构造器，包含两种，一种是用于构造聚合 `stage` 阶段所需的 `bson` 数据，另一种是用于构造除了 `stage` 阶段以外的数据。

## universal 通用构造
我们可以使用 `bsonx` 包里的一些函数进行 `bson` 数据的构造，例如 `bsonx.M`、`bsonx.Id` 和 `bsonx.D` 等等。
```go
// bson.M{"姓名": "陈明勇"}
m := bsonx.M("姓名", "陈明勇")

// bson.M{"_id": "陈明勇"}
id := bsonx.Id("陈明勇")

// bson.D{bson.E{Key:"姓名", Value:"陈明勇"}, bson.E{Key:"手机号", Value:"1888***1234"}}
d := bsonx.NewD().Add("name", "chenmingyong").Add("telephone", "1888***1234").Build()
d := bsonx.D(bsonx.E("name", "chenmingyong"), bsonx.E("telephone", "1888***1234"))

// bson.E{Key:"姓名", Value:"陈明勇"}
e := bsonx.E("姓名", "陈明勇")

// bson.A{"陈明勇", "1888***1234"}
a := bsonx.A("陈明勇", "1888***1234")

```
特别注意的是，使用 `bsonx.D` 方法构造数据时，传入的参数，需要使用 `bsonx.KV` 方法进行传递，目的是强约束 `key-value` 的类型。
## query
`query` 包可以帮我们构造出查询相关的 `bson` 数据，例如 `$in`、`$gt`、`$and` 等等。

```go
// bson.D{bson.E{Key:"姓名", Value:"陈明勇"}}
d := query.BsonBuilder().Add("name", "chenmingyong").Build()

// bson.D{bson.E{Key:"age", Value:bson.D{{Key:"$gt", Value:18}, bson.E{Key:"$lt", Value:25}}}}
d = query.BsonBuilder().Gt("age", 18).Lt("age", 25).Build()

//  bson.D{bson.E{Key:"age", Value:bson.D{{Key:"$in", Value:[]int{18, 19, 20}}}}}
d = query.BsonBuilder().InInt("age", 18, 19, 20).Build()

// bson.d{bson.E{Key: "$and", Value: []any{bson.D{{Key: "x", Value: bson.D{{Key: "$ne", Value: 0}}}}, bson.D{{Key: "y", Value: bson.D{{Key: "$gt", Value: 0}}}}}}
d = query.BsonBuilder().And(
	query.BsonBuilder().Ne("x", 0).Build(),
	query.BsonBuilder().Gt("y", 0).Build(),
).Build()

// bson.D{bson.E{Key:"qty", Value:bson.D{{Key:"$exists", Value:true}, bson.E{Key:"$nin", Value:[]int{5, 15}}}}}
d = query.BsonBuilder().Exists("qty", true).NinInt("qty", 5, 15).Build()

// elemMatch
// bson.D{bson.E{Key: "result", Value: bson.D{bson.E{Key: "$elemMatch", Value: bson.D{bson.E{Key: "$gte", Value: 80}, bson.E{Key: "$lt", Value: 85}}}}}}
d = query.BsonBuilder().ElemMatch("result", bsonx.NewD().Add("$gte", 80).Add("$lt", 85)).Build()
```
`query` 包提供的方法不止这些，以上只是列举出一些典型的例子，还有更多的用法等着你去探索。

## update
`update` 包可以帮我们构造出更新操作相关的 `bson` 数据，例如 `$set`、`$$inc`、`$push` 等等。

```go
// bson.D{bson.E{Key:"$set", Value:bson.M{"name":"陈明勇"}}}
u := update.BsonBuilder().Set(bsonx.M("name", "陈明勇")).Build()

// bson.D{bson.E{Key:"$inc", Value:bson.D{bson.E{Key:"orders", Value:1}, bson.E{Key:"ratings", Value:-1}}}}
u = update.BsonBuilder().Inc(bsonx.NewD().Add("orders", 1).Add("ratings", -1)).Build()

// bson.D{bson.E{Key:"$push", Value:bson.M{"scores":95}}}
u = update.BsonBuilder().Push(bsonx.M("scores", 95)).Build()

// bson.D{bson.E{Key:"$unset", Value:bson.D{primitive.E{Key:"quantity", Value:""}, bson.E{Key:"instock", Value:""}}}}
u = update.BsonBuilder().Unset("quantity", "instock").Build()
```
`update` 包提供的方法不止这些，以上只是列举出一些典型的例子，还有更多的用法等着你去探索。

## aggregation
`aggregation` 包提供了两个 `builder`：
- `StageBsonBuilder`：用于构造 `stage` 阶段所需的 `bson` 数据
- `BsonBuilder`：用于构造除了 `stage` 阶段以外的 `bson` 数据。

```go
// bson.D{bson.E{Key:"$gt", Value:[]any{"$qty", 250}}}
gt := aggregation.BsonBuilder().Gt("$qty", 250).Build()

// mongo.Pipeline{bson.D{bson.E{Key:"$project", Value:bson.D{bson.E{Key:"name", Value:1}, bson.E{Key:"age", Value:1}, bson.E{Key:"qtyGt250", Value:bson.D{bson.E{Key:"$gt", Value:[]interface {}{"$qty", 250}}}}}}}}
pipeline := aggregation.StageBsonBuilder().Project(bsonx.NewD().Add("name", 1).Add("age", 1).Add("qtyGt250", gt)).Build()

// bson.D{bson.E{Key:"$or", Value:[]interface {}{bson.D{bson.E{Key:"score", Value:bson.D{bson.E{Key:"$gt", Value:70}, bson.E{Key:"$lt", Value:90}}}}, bson.D{bson.E{Key:"views", Value:bson.D{bson.E{Key:"$gte", Value:90}}}}}}}
or := aggregation.BsonBuilder().Or(
	query.BsonBuilder().Gt("score", 70).Lt("score", 90).Build(),
	query.BsonBuilder().Gte("views", 90).Build(),
).Build()

// mongo.Pipeline{bson.D{bson.E{Key:"$match", Value:bson.D{bson.E{Key:"$or", Value:[]any{bson.D{bson.E{Key:"score", Value:bson.D{bson.E{Key:"$gt", Value:70}, bson.E{Key:"$lt", Value:90}}}}, bson.D{bson.E{Key:"views", Value:bson.D{bson.E{Key:"$gte", Value:90}}}}}}}}}, bson.D{bson.E{Key:"$group", Value:bson.D{bson.E{Key:"_id", Value:any(nil)}, bson.E{Key:"count", Value:bson.D{bson.E{Key:"$sum", Value:1}}}}}}}
pipeline = aggregation.StageBsonBuilder().Match(or).Group(nil, bsonx.E("count", aggregation.BsonBuilder().Sum(1).Build())).Build()

// mongo.Pipeline{bson.D{bson.E{Key:"$unwind", Value:"$size"}}}
pipeline = aggregation.StageBsonBuilder().Unwind("$size", nil).Build()

// mongo.Pipeline{bson.D{bson.E{Key:"$unwind", Value:bson.D{bson.E{Key:"path", Value:"$size"}, bson.E{Key:"includeArrayIndex", Value:"arrayIndex"}, bson.E{Key:"preserveNullAndEmptyArrays", Value:true}}}}}
pipeline = aggregation.StageBsonBuilder().Unwind("$size", &types.UnWindOptions{
	IncludeArrayIndex:          "arrayIndex",
	PreserveNullAndEmptyArrays: true,
}).Build()
```

`aggregation` 包提供的方法不止这些，以上只是列举出一些典型的例子，还有更多的用法等着你去探索。
