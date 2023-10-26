<h1 align="center">
  go-mongox
</h1>

<p>go-mongox is a secondary encapsulation framework based on the official MongoDB framework, aiming to provide a more convenient and efficient MongoDB data manipulation experience.</p>


> **The functions are continuously updated and improved, and the partners who are interested in the framework are welcome to put forward valuable comments and participate in contributions.**

---

# 介绍
The `go-mongox` framework has two cores, one core is the generics-based **collection** form, and the other core is the **builder** constructor.


- With the **collection** object, we can easily conduct related MongoDB operations, thereby reducing the writing of bson data and improving development efficiency.
- With the **builder** constructor, we can construct the 'bson' data we need from a 'key-value' pair or a 'map' and a struct object.

# Install

> go get github.com/chenmingyong0423/go-mongox@latest

# Quick Start
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
	// You need to create a *mongo.Collection object in advance
	mongoCollection := newCollection()

	// Create a collection using the Post struct as a generic parameter
	postCollection := mongox.NewCollection[Post](mongoCollection)

	// Insert a data
	insertOneResult, err := postCollection.Creator().InsertOne(context.Background(), Post{
		Id:      "666",
		Title:   "go-mongox",
		Content: "go-mongox is a...",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("id: %v\n", insertOneResult.InsertedID) // id: 666

	// query data
	// - Constructs a bson query condition statement with _id "666"
	filter := builder.NewBsonBuilder().Id("666").Build()
	post, err := postCollection.Finder().Filter(filter).FindOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("new data: %v\n", post) // new data: &{666 go-mongox go-mongox is a...}

	// Update data by _id
	// - Constructs an update statement for the bson data to update the value of the content field
	updates := builder.NewBsonBuilder().Set("content", "go-mongox is a very useful framework").Build()
	updateResult, err := postCollection.Updater().Filter(filter).Updates(updates).UpdateOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("ModifiedCount: %d\n", updateResult.ModifiedCount) // ModifiedCount: 1

	// Query the updated data
	updatedPost, err := postCollection.Finder().Filter(filter).FindOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated data: %v\n", updatedPost) // Updated data: &{666 go-mongox go-mongox is a very useful framework}

	// delete data by _id
	deleteResult, err := postCollection.Deleter().Filter(filter).DeleteOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("DeletedCount: %d", deleteResult.DeletedCount) // DeletedCount: 1
}

// Sample code is not the best way to create it
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

> More detailed documentation will be updated later
