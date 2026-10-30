package mongox

import (
	"context"
	"testing"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/finder"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/x/mongo/driver/drivertest"
)

func TestFinderDeletedScopeCommands(t *testing.T) {
	for _, scope := range []string{"default", "with deleted", "only deleted", "only then with", "with then only"} {
		for _, method := range []string{"FindOne", "Find", "Count", "Distinct", "DistinctWithParse", "FindOneAndUpdate"} {
			t.Run(scope+"/"+method, func(t *testing.T) {
				db, deployment, commands := softDeleteDatabase(t)
				col := NewCollection[softDeleteDoc](db, "docs")
				original := bson.M{"name": "alice"}
				f := col.Finder().Filter(original)
				switch scope {
				case "with deleted":
					f = f.WithDeleted()
				case "only deleted":
					f = f.OnlyDeleted()
				case "only then with":
					f = f.OnlyDeleted().WithDeleted()
				case "with then only":
					f = f.WithDeleted().OnlyDeleted()
				}
				for range 2 {
					runFinderScopeCommand(t, f, deployment, method)
					command := (*commands)[len(*commands)-1]
					filter := finderScopeCommandFilter(command, method)
					switch scope {
					case "default":
						requireLiveFilter(t, filter)
					case "with deleted", "only then with":
						var actual bson.M
						require.NoError(t, bson.Unmarshal(filter, &actual))
						require.Equal(t, original, actual)
					default:
						requireOnlyDeletedFilter(t, filter)
					}
					if method == "FindOneAndUpdate" {
						require.Equal(t, bson.TypeDateTime, command.Lookup("update", "$set", "updated_at").Type)
					}
				}
				require.Len(t, *commands, 2, "reusing a finder must not accumulate filters")
				require.Equal(t, bson.M{"name": "alice"}, original)
				// A new finder retains the default scope on the same collection.
				runFinderScopeCommand(t, col.Finder().Filter(original), deployment, "FindOne")
				requireLiveFilter(t, (*commands)[2].Lookup("filter").Document())
			})
		}
	}
}

func runFinderScopeCommand(t *testing.T, f finder.IFinder[softDeleteDoc], deployment *drivertest.MockDeployment, method string) {
	t.Helper()
	ctx := context.Background()
	var err error
	switch method {
	case "FindOne":
		deployment.AddResponses(cursorResponse(bson.M{"name": "alice"}))
		_, err = f.FindOne(ctx)
	case "Find":
		deployment.AddResponses(cursorResponse(bson.M{"name": "alice"}))
		_, err = f.Find(ctx)
	case "Count":
		deployment.AddResponses(cursorResponse(bson.M{"n": int32(1)}))
		var count int64
		count, err = f.Count(ctx)
		require.Equal(t, int64(1), count)
	case "Distinct", "DistinctWithParse":
		deployment.AddResponses(bson.D{{Key: "ok", Value: 1}, {Key: "values", Value: bson.A{"alice"}}})
		if method == "Distinct" {
			err = f.Distinct(ctx, "name").Err()
		} else {
			var values []string
			err = f.DistinctWithParse(ctx, "name", &values)
			require.Equal(t, []string{"alice"}, values)
		}
	case "FindOneAndUpdate":
		deployment.AddResponses(bson.D{{Key: "ok", Value: 1}, {Key: "value", Value: bson.M{"name": "alice"}}})
		_, err = f.Updates(bson.M{"$set": bson.M{"name": "alice"}}).FindOneAndUpdate(ctx)
	default:
		t.Fatalf("unknown finder method %q", method)
	}
	require.NoError(t, err)
}

func finderScopeCommandFilter(command bson.Raw, method string) bson.Raw {
	switch method {
	case "Count":
		return command.Lookup("pipeline").Array().Index(0).Document().Lookup("$match").Document()
	case "Distinct", "DistinctWithParse", "FindOneAndUpdate":
		return command.Lookup("query").Document()
	default:
		return command.Lookup("filter").Document()
	}
}

func requireOnlyDeletedFilter(t *testing.T, filter bson.Raw) {
	t.Helper()
	clauses, err := filter.Lookup("$and").Array().Values()
	require.NoError(t, err)
	require.Len(t, clauses, 2)
	original := clauses[0].Document()
	if original.Lookup("$and").Type != 0 {
		requireOnlyDeletedFilter(t, original)
	} else {
		require.Equal(t, "alice", original.Lookup("name").StringValue())
	}
	negated, err := clauses[1].Document().Lookup("$nor").Array().Values()
	require.NoError(t, err)
	require.Len(t, negated, 1)
	values := negated[0].Document().Lookup("deleted_at", "$in").Array()
	require.Equal(t, bson.TypeNull, values.Index(0).Type)
	require.Equal(t, time.Time{}.UnixMilli(), values.Index(1).DateTime())
}

func TestFinderDeletedScopeNilFilter(t *testing.T) {
	for _, onlyDeleted := range []bool{false, true} {
		db, _, commands := softDeleteDatabase(t)
		f := NewCollection[softDeleteDoc](db, "docs").Finder().Filter(nil).WithDeleted()
		if onlyDeleted {
			f = f.OnlyDeleted()
		}
		ctx := context.Background()
		_, err := f.FindOne(ctx)
		require.ErrorIs(t, err, mongo.ErrNilDocument)
		_, err = f.Find(ctx)
		require.ErrorIs(t, err, mongo.ErrNilDocument)
		_, err = f.Count(ctx)
		require.ErrorIs(t, err, mongo.ErrNilDocument)
		require.ErrorIs(t, f.Distinct(ctx, "name").Err(), mongo.ErrNilDocument)
		var values []string
		require.ErrorIs(t, f.DistinctWithParse(ctx, "name", &values), mongo.ErrNilDocument)
		_, err = f.Updates(bson.M{"$set": bson.M{"name": "alice"}}).FindOneAndUpdate(ctx)
		require.ErrorIs(t, err, mongo.ErrNilDocument)
		require.Empty(t, *commands, "the driver must still reject nil before sending a command")
	}
}

func TestFinderOnlyDeletedRawFilter(t *testing.T) {
	encoded, err := bson.Marshal(bson.M{"name": "alice"})
	require.NoError(t, err)
	for _, filter := range []any{encoded, bson.Raw(encoded)} {
		db, deployment, commands := softDeleteDatabase(t)
		f := NewCollection[softDeleteDoc](db, "docs").Finder().Filter(filter).OnlyDeleted()
		runFinderScopeCommand(t, f, deployment, "FindOne")
		requireOnlyDeletedFilter(t, (*commands)[0].Lookup("filter").Document())
	}
}

func TestFinderDeletedScopeWithoutDeletedAt(t *testing.T) {
	for _, onlyDeleted := range []bool{false, true} {
		db, deployment, commands := softDeleteDatabase(t)
		original := bson.M{"name": "alice"}
		f := NewCollection[struct{ Name string }](db, "docs").Finder().Filter(original).WithDeleted()
		if onlyDeleted {
			f = f.OnlyDeleted()
		}
		deployment.AddResponses(cursorResponse([]any{}...))
		_, err := f.Find(context.Background())
		require.NoError(t, err)
		filter := (*commands)[0].Lookup("filter").Document()
		if onlyDeleted {
			clauses := filter.Lookup("$and").Array()
			var actual bson.M
			require.NoError(t, bson.Unmarshal(clauses.Index(0).Document(), &actual))
			require.Equal(t, original, actual)
			negated := clauses.Index(1).Document().Lookup("$nor").Array()
			require.Len(t, negated.Index(0).Document(), 5, "negating an empty BSON document matches no documents")
		} else {
			var actual bson.M
			require.NoError(t, bson.Unmarshal(filter, &actual))
			require.Equal(t, original, actual)
		}
	}
}

func TestFinderOnlyDeletedCustomIntegerField(t *testing.T) {
	type timestampDoc struct {
		Removed int64 `bson:"removed_at" mongox:"autoDeleteTime:milli"`
	}
	db, deployment, commands := softDeleteDatabase(t)
	deployment.AddResponses(cursorResponse(bson.M{"n": int32(1)}))
	count, err := NewCollection[timestampDoc](db, "docs").Finder().OnlyDeleted().Count(context.Background())
	require.NoError(t, err)
	require.Equal(t, int64(1), count)
	filter := finderScopeCommandFilter((*commands)[0], "Count")
	guard := filter.Lookup("$and").Array().Index(1).Document().Lookup("$nor").Array().Index(0).Document()
	values := guard.Lookup("removed_at", "$in").Array()
	require.Equal(t, bson.TypeNull, values.Index(0).Type)
	require.Equal(t, int64(0), values.Index(1).Int64())
}
