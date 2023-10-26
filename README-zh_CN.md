<h1 align="center">
  go-mongox
</h1>

<p>go-mongox 是一个基于 MongoDB 官方框架进行二次封装的一个框架，旨在提供更方便高效的 MongoDB 数据操作体验。</p>


> **功能持续更新和改进中，对该框架感兴趣的伙伴，欢迎提出宝贵的意见和参与贡献。**

---

# 介绍
`go-mongox` 框架有两个核心，一个核心是基于泛型的 **collection** 形态，另一个核心是 **builder** 构造器。

- 通过 **collection** 对象，我们可以方便地进行相关的 MongoDB 操作，从而减少 bson 数据的编写，提高开发效率；
- 通过 **builder** 构造器，我们可以通过 `key-value` 键值对或者 `map` 以及结构体对象，构造出我们所需要的 `bson` 数据。

# 安装

> go get github.com/chenmingyong0423/go-mongox@latest

# 快速开始
```go
package main

import (
	"context"
	"fmt"

	"github.com/chenmingyong0423/go-mongox"
	"github.com/chenmingyong0423/go-mongox/builder"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Post struct {
	Id      string `bson:"_id"`
	Title   string `bson:"title"`
	Content string `bson:"content"`
}

func main() {
	// 你需要预先创建一个 *mongo.Collection 对象
	mongoCollection := newCollection()

	// 使用 Post 结构体作为泛型参数创建一个 collection
	postCollection := mongox.NewCollection[Post](mongoCollection)

	// 插入一条数据
	insertOneResult, err := postCollection.Creator().InsertOne(context.Background(), Post{
		Id:      "666",
		Title:   "go-mongox",
		Content: "go-mongox 是一个...",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("id: %v\n", insertOneResult.InsertedID) // id: 666

	// 查询数据
	// - 构造 _id 为 "666" 的 bson 查询条件语句
	filter := builder.NewBsonBuilder().Id("666").Build()
	post, err := postCollection.Finder().Filter(filter).FindOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("新数据: %v\n", post) // 新数据: &{666 go-mongox go-mongox 是一个...}

	// 根据 _id 更新数据
	// - 构造 bson 数据的更新语句，更新 content 字段
	updates := builder.NewBsonBuilder().Set("content", "go-mongox 是一个非常好用的框架").Build()
	updateResult, err := postCollection.Updater().Filter(filter).Updates(updates).UpdateOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("ModifiedCount: %d\n", updateResult.ModifiedCount) // ModifiedCount: 1

	// 查询更新后的数据
	updatedPost, err := postCollection.Finder().Filter(filter).FindOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("更新后的数据: %v\n", updatedPost) // 更新后的数据: &{666 go-mongox go-mongox 是一个非常好用的框架}

	// 根据 _id 删除数据
	deleteResult, err := postCollection.Deleter().Filter(filter).DeleteOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("DeletedCount: %d", deleteResult.DeletedCount) // DeletedCount: 1
}

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

```

> 后续会更新更详细的介绍文档
