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

<p>go-mongox is a secondary encapsulation of the MongoDB official framework based on generics. It enables smooth document operations through the use of chain calls. Additionally, it provides various types of BSON builders, facilitating the efficient construction of BSON data.</p>


> **Continuously evolving with ongoing updates and improvements, we welcome valuable feedback and contributions from those interested in this framework. Your input is highly appreciated as we strive to enhance and expand the functionality of the go-mongox framework.**

---

English | [中文简体](./README-zh_CN.md)

# Introduction
The `go-mongox` framework has two core components: one is based on a generic **collection** paradigm, and the other is the **builder** constructor.

- With the **collection** object, we can easily perform relevant MongoDB operations, reducing the need for manual bson data writing and enhancing development efficiency.

- Using the **builder** constructor, we can construct the bson data we need.

# Install

> go get github.com/chenmingyong0423/go-mongox@latest

# Collection Paradigm
Create a `Collection` instance based on the generic type `Post`.
```go

// Example Code, Not the Optimal Creation Method
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


// You need to pre-create a *mongo.Collection object.
mongoCollection := newCollection()
// Creating a Collection Using the Post Struct as a Generic Parameter
postCollection := mongox.NewCollection[Post](mongoCollection)
```
## Creator
`Creator` is a constructor used to perform insertion-related operations.

```go
// insert a document
doc := Post{Id: "1", Title: "go", Author: "chenmingyong", Content: "..."}
oneResult, err := postCollection.Creator().InsertOne(context.Background(), doc)
// with option 
oneResult, err := postCollection.Creator().OneOptions(options.InsertOne().SetComment("test")).InsertOne(context.Background(), Post{Id: "1", Title: "go", Author: "chenmingyong", Content: "..."})

// insert many documents
docs := []Post{
	{Id: "2", Title: "go", Author: "chenmingyong", Content: "..."},
	{Id: "3", Title: "mongo", Author: "chenmingyong", Content: "..."},
}
manyResult, err := postCollection.Creator().InsertMany(context.Background(), docs)
// with option
manyResult, err := postCollection.Creator().ManyOptions(options.InsertMany().SetComment("test")).InsertMany(context.Background(), docs)
```
Create a constructor using the `Creator()` method, based on which we can perform related document insertion operations. The `OneOptions` and `ManyOptions` methods are used to set the `*options.InsertOneOptions` and `*options.InsertManyOptions` parameters, respectively.
## Finder
`Finder` is a query builder used to perform operations related to queries.

```go
// find a document
post, err := postCollection.Finder().Filter(bsonx.Id("1")).FindOne(context.Background())

// set *options.FindOneOptions
post, err := postCollection.Finder().
	Filter(bsonx.Id("1")).
	OneOptions(options.FindOne().SetProjection(bsonx.M("content", 0))).
	FindOne(context.Background())

// - map as filter
post, err := postCollection.Finder().Filter(map[string]any{"_id": "1"}).FindOne(context.Background())

// -- using query package constructs bson: bson.D{bson.E{Key: "title", Value: bson.M{"$eq": "go"}}, bson.E{Key: "author", Value: bson.M{"$eq": "chenmingyong"}}}
post, err := postCollection.Finder().
	Filter(query.BsonBuilder().Eq("title", "go").Eq("author", "chenmingyong").Build()).
	FindOne(context.Background())

// find many documents
// bson.D{bson.E{Key: "_id", Value: bson.M{"$in": []string{"1", "2"}}}}
posts, err := postCollection.Finder().Filter(query.BsonBuilder().InString("_id", []string{"1", "2"}...).Build()).Find(context.Background())

// set *options.FindOptions
// bson.D{bson.E{Key: "_id", Value: bson.M{types.In: []string{"1", "2"}}}}
posts, err := postCollection.Finder().
	Filter(query.BsonBuilder().InString("_id", []string{"1", "2"}...).Build()).
	Options(options.Find().SetProjection(bsonx.M("content", 0))).
	Find(context.Background())
```
Create a query builder using the `Finder()` method, based on which we can perform related query operations. The method for setting query conditions is `Filter`, and methods for setting `*options.FindOneOptions` parameters include `OneOptions` and `Options`.

If the query condition `filter` is straightforward, it can be constructed using the `bsonx.M` method. For more complex `filter` conditions, we can use the `BsonBuilder()` constructor from the `query` package for construction.

## Updater
`Updater` is an updater used to perform operations related to updates.

```go
// update a document
// Build BSON update statements using the `update` package.
updateResult, err := postCollection.Updater().
	Filter(bsonx.Id("1")).
	Updates(update.BsonBuilder().Set(bsonx.M("title", "golang")).Build()).
	UpdateOne(context.Background())

// - Construct update data using a map, set *options.UpdateOptions, and perform an upsert operation.
updateResult, err := postCollection.Updater().
	Filter(bsonx.Id("4")).
	UpdatesWithOperator(types.Set, map[string]any{"title": "golang"}).Options(options.Update().SetUpsert(true)).
	UpdateOne(context.Background())

// update many documents
updateResult, err := postCollection.Updater().
	Filter(query.BsonBuilder().InString("_id", []string{"2", "3"}...).Build()).
	Updates(update.BsonBuilder().Set(bsonx.M("title", "golang")).Build()).
	UpdateMany(context.Background())
```
Create an updater using the `Updater()` method, based on which we can perform related update operations. The method for setting document matching conditions is `Filter`, and methods for setting update content include `Updates` and `UpdatesWithOperator`. The method for setting `*options.UpdateOptions` parameters is `Options`.

The `UpdatesWithOperator` method requires us to specify an operator.
## Deleter
`Deleter` is a deleter used to perform operations related to deletions.

```go
// delete a document
deleteResult, err := postCollection.Deleter().Filter(bsonx.Id("1")).DeleteOne(context.Background())
// with option
deleteResult, err := postCollection.Deleter().Filter(bsonx.Id("1")).Options(options.Delete().SetComment("test")).DeleteOne(context.Background())

// delete many documents
// - Construct complex BSON statements using the `query` package.
deleteResult, err = postCollection.Deleter().Filter(query.BsonBuilder().InString("_id", []string{"2", "3", "4"}...).Build()).DeleteMany(context.Background())
```
Create a deleter using the `Deleter()` method, based on which we can perform related deletion operations. Methods for setting document matching conditions include `Filter`, and the method for setting `*options.DeleteOptions` parameters is `Options`.

## Aggregator
`Aggregator` is an aggregator used to perform operations related to aggregations.

```go
// Construct a pipeline using the `aggregation` package.
posts, err = postCollection.
	Aggregator().Pipeline(aggregation.StageBsonBuilder().Project(bsonx.M("content", 0)).Build()).
	Aggregate(context.Background())


// If we change the field names through aggregation operations, we can use the AggregationWithCallback method and then map the results to the expected struct through the callback function.
type DiffPost struct {
	Id          string `bson:"_id"`
	Title       string `bson:"title"`
	Name        string `bson:"name"` // author → name
	Content     string `bson:"content"`
	Outstanding bool   `bson:"outstanding"`
}
result := make([]*DiffPost, 0)
//  Rename the 'author' field to 'name,' exclude the 'content' field, add the 'outstanding' field, and return the result as []*DiffPost.
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
Create an aggregator using the `Aggregator()` method, based on which we can perform relevant aggregation operations. There are two methods for executing aggregation operations: `Aggregation` and `AggregationWithCallback`. `AggregationWithCallback` allows the query results to be output to a specified slice through a callback.

Use `aggregation.StageBsonBuilder` to construct BSON data containing operators for the `stage` phase, and use `aggregation.BsonBuilder` to construct BSON data containing operators excluding the `stage` phase.
# Builder
The `go-mongox` framework provides the following types of builders:

- `universal`: A simple and versatile `bson` data construction function.
- `query`: A query builder used to construct `bson` data needed for query operations.
- `update`: An update builder used to construct `bson` data needed for update operations.
- `aggregation`: An aggregation operation builder, which includes two types—one for constructing `bson` data needed for the aggregation `stage` phase, and another for constructing data excluding the `stage` phase.

## universal - General Construction
We can use some functions in the `bsonx` package to construct BSON data, such as `bsonx.M`, `bsonx.Id`, and `bsonx.D`, and so on.

```go
// bson.M{"name": "chenmingyong"}
m := bsonx.M("name", "chenmingyong")

// bson.M{"_id": "chenmingyong"}
id := bsonx.Id("chenmingyong")

// bson.D{bson.E{Key:"name", Value:"chenmingyong"}, bson.E{Key:"telephone", Value:"1888***1234"}}
d := bsonx.NewD().Add("name", "chenmingyong").Add("telephone", "1888***1234").Build()
d := bsonx.D(bsonx.E("name", "chenmingyong"), bsonx.E("telephone", "1888***1234"))

// bson.E{Key:"name", Value:"chenmingyong"}
e := bsonx.E("name", "chenmingyong")

// bson.A{"chenmingyong", "1888***1234"}
a := bsonx.A("chenmingyong", "1888***1234")

```

Particularly noteworthy is that when constructing data using the `bsonx.D` method, the parameters need to be passed using the `bsonx.KV` method. This is to enforce strict typing of the `key-value` pairs.

## query
The `query` package can help us construct `bson` data related to queries, such as `$in`, `$gt`, `$and`, and so on.

```go
// bson.D{bson.E{Key:"name", Value:"陈明勇"}}
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
The methods provided by the `query` package are not limited to these. The examples listed above are just a few typical ones, and there are more usages waiting for you to explore.

## update
The `update` package can help us construct `bson` data related to update operations, such as `$set`, `$inc`, `$push`, and so on.

```go
// bson.D{bson.E{Key:"$set", Value:bson.M{"name":"chenmingyong"}}}
u := update.BsonBuilder().Set(bsonx.M("name", "chenmingyong")).Build()

// bson.D{bson.E{Key:"$inc", Value:bson.D{bson.E{Key:"orders", Value:1}, bson.E{Key:"ratings", Value:-1}}}}
u = update.BsonBuilder().Inc(bsonx.NewD().Add("orders", 1).Add("ratings", -1)).Build()

// bson.D{bson.E{Key:"$push", Value:bson.M{"scores":95}}}
u = update.BsonBuilder().Push(bsonx.M("scores", 95)).Build()

// bson.D{bson.E{Key:"$unset", Value:bson.D{primitive.E{Key:"quantity", Value:""}, bson.E{Key:"instock", Value:""}}}}
u = update.BsonBuilder().Unset("quantity", "instock").Build()
```
The methods provided by the `update` package are not limited to these. The examples listed above are just a few typical ones, and there are more usages waiting for you to explore.

## aggregation
The `aggregation` package provides two builders:
- `StageBsonBuilder`: Used to construct `bson` data needed for the `stage` phase.
- `BsonBuilder`: Used to construct data excluding the `stage` phase.

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

The methods provided by the `aggregation` package are not limited to these. The examples listed above are just a few typical ones, and there are more usages waiting for you to explore.
