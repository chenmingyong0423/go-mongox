// Copyright 2025 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mongox

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/deleter"
	"github.com/chenmingyong0423/go-mongox/v2/finder"
	"github.com/chenmingyong0423/go-mongox/v2/operation"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/event"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/x/mongo/driver/drivertest"
)

type softDeleteDoc struct {
	Model `bson:",inline"`
	Name  string `bson:"name"`
}

// Exercise the actual collection/operation/driver path and inspect its wire
// commands. No running MongoDB is required for these tests.
func softDeleteDatabase(t *testing.T) (*Database, *drivertest.MockDeployment, *[]bson.Raw) {
	t.Helper()
	deployment := drivertest.NewMockDeployment()
	var commands []bson.Raw
	opts := options.Client().SetRetryReads(false).SetRetryWrites(false).SetMonitor(&event.CommandMonitor{
		Started: func(_ context.Context, ev *event.CommandStartedEvent) {
			commands = append(commands, append(bson.Raw(nil), ev.Command...))
		},
	})
	opts.Deployment = deployment
	client, err := mongo.Connect(opts)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, client.Disconnect(context.Background())) })
	return NewClient(client, &Config{}).NewDatabase("test"), deployment, &commands
}

func cursorResponse(docs ...any) bson.D {
	return bson.D{{Key: "ok", Value: 1}, {Key: "cursor", Value: bson.D{
		{Key: "id", Value: int64(0)}, {Key: "ns", Value: "test.docs"}, {Key: "firstBatch", Value: docs},
	}}}
}

func writeResponse(n int32) bson.D {
	return bson.D{{Key: "ok", Value: 1}, {Key: "n", Value: n}, {Key: "nModified", Value: n}}
}

func requireLiveFilter(t *testing.T, filter bson.Raw) {
	t.Helper()
	clauses, err := filter.Lookup("$and").Array().Values()
	require.NoError(t, err)
	require.Len(t, clauses, 2)
	original := clauses[0].Document()
	if original.Lookup("$and").Type != 0 {
		requireLiveFilter(t, original)
	} else {
		require.Equal(t, "alice", original.Lookup("name").StringValue())
	}
	values, err := clauses[1].Document().Lookup("deleted_at", "$in").Array().Values()
	require.NoError(t, err)
	require.Len(t, values, 2)
	require.Equal(t, bson.TypeNull, values[0].Type)
	require.Equal(t, time.Time{}.UnixMilli(), values[1].DateTime())
}

func TestSoftDeleteCommands(t *testing.T) {
	for _, many := range []bool{false, true} {
		t.Run(map[bool]string{false: "one", true: "many"}[many], func(t *testing.T) {
			db, deployment, commands := softDeleteDatabase(t)
			col := NewCollection[softDeleteDoc](db, "docs")
			filter := bson.M{"name": "alice"}
			var order []string
			var started time.Time
			db.RegisterPlugin("inspect-delete", func(_ context.Context, op *operation.OpContext, _ ...any) error {
				started = op.StartTime
				require.Equal(t, started, op.Updates.(bson.M)["$set"].(bson.M)["deleted_at"])
				order = append(order, "before-delete")
				return nil
			}, operation.OpTypeBeforeDelete)
			db.RegisterPlugin("reject-update-hook", func(context.Context, *operation.OpContext, ...any) error {
				t.Fatal("soft deletion must not execute update callbacks")
				return nil
			}, operation.OpTypeBeforeUpdate)
			db.RegisterPlugin("inspect-delete-updates", func(_ context.Context, op *operation.OpContext, _ ...any) error {
				require.Equal(t, bson.M{"$set": bson.M{"deleted_at": started, "deleted_by": "local-hook"}}, op.Updates)
				return nil
			}, operation.OpTypeAfterDelete)
			d := col.Deleter().Filter(filter).RegisterBeforeHooks(func(_ context.Context, op *deleter.OpContext, _ ...any) error {
				require.Equal(t, bson.M{"$set": bson.M{"deleted_at": started}}, op.Updates)
				// Replace the document to verify synchronization beyond shared map mutations.
				op.Updates = bson.M{"$set": bson.M{"deleted_at": started, "deleted_by": "local-hook"}}
				encoded, err := bson.Marshal(op.Filter)
				require.NoError(t, err)
				requireLiveFilter(t, encoded)
				order = append(order, "local-before")
				return nil
			}).RegisterAfterHooks(func(_ context.Context, op *deleter.OpContext, _ ...any) error {
				require.Equal(t, bson.M{"$set": bson.M{"deleted_at": started, "deleted_by": "local-hook"}}, op.Updates)
				require.IsType(t, &mongo.DeleteResult{}, op.Result)
				order = append(order, "local-after")
				return nil
			})
			// The second response models a repeat delete matching no live documents.
			count := int32(1)
			if many {
				count = 2
			}
			deployment.AddResponses(writeResponse(count), writeResponse(0))
			for _, want := range []int64{int64(count), 0} {
				var result *mongo.DeleteResult
				var err error
				collation := &options.Collation{Locale: "en"}
				if many {
					result, err = d.UpdateManyOptions(options.UpdateMany().SetComment("old"), options.UpdateMany().SetComment("soft-delete").SetHint("name_1").SetLet(bson.M{"v": 1}).SetCollation(collation)).DeleteMany(context.Background(), options.DeleteMany().SetComment("physical-delete"))
				} else {
					result, err = d.UpdateOneOptions(options.UpdateOne().SetComment("old"), options.UpdateOne().SetComment("soft-delete").SetHint("name_1").SetLet(bson.M{"v": 1}).SetCollation(collation)).DeleteOne(context.Background(), options.DeleteOne().SetComment("physical-delete"))
				}
				require.NoError(t, err)
				require.Equal(t, want, result.DeletedCount)
				require.True(t, result.Acknowledged)
				command := (*commands)[len(*commands)-1]
				require.Equal(t, "docs", command.Lookup("update").StringValue())
				require.Equal(t, "soft-delete", command.Lookup("comment").StringValue())
				require.Equal(t, int32(1), command.Lookup("let", "v").Int32())
				update := command.Lookup("updates").Array().Index(0).Document()
				multi, _ := update.Lookup("multi").BooleanOK()
				require.Equal(t, many, multi)
				require.Equal(t, "name_1", update.Lookup("hint").StringValue())
				require.Equal(t, "en", update.Lookup("collation", "locale").StringValue())
				requireLiveFilter(t, update.Lookup("q").Document())
				require.Equal(t, started.UnixMilli(), update.Lookup("u", "$set", "deleted_at").DateTime())
				require.Equal(t, "local-hook", update.Lookup("u", "$set", "deleted_by").StringValue())
				require.Equal(t, bson.Type(0), update.Lookup("u", "$set", "updated_at").Type)
				require.Equal(t, bson.Type(0), update.Lookup("u", "$set", "created_at").Type)
			}
			require.Equal(t, bson.M{"name": "alice"}, filter)
			require.Equal(t, []string{"before-delete", "local-before", "local-after", "before-delete", "local-before", "local-after"}, order)
		})
	}
}

func TestSoftDeleteReadCommands(t *testing.T) {
	for _, method := range []string{"FindOne", "Find", "Count", "Distinct", "DistinctWithParse", "FindOneAndUpdate"} {
		t.Run(method, func(t *testing.T) {
			db, deployment, commands := softDeleteDatabase(t)
			col := NewCollection[softDeleteDoc](db, "docs")
			filter := bson.M{"name": "alice"}
			f := col.Finder().Filter(filter)
			var err error
			switch method {
			case "FindOne":
				deployment.AddResponses(cursorResponse(bson.M{"name": "alice"}))
				_, err = f.FindOne(context.Background())
			case "Find":
				deployment.AddResponses(cursorResponse(bson.M{"name": "alice"}))
				_, err = f.Find(context.Background())
			case "Count":
				deployment.AddResponses(cursorResponse(bson.M{"n": int32(1)}))
				var n int64
				n, err = f.Count(context.Background())
				require.Equal(t, int64(1), n)
			case "Distinct", "DistinctWithParse":
				deployment.AddResponses(bson.D{{Key: "ok", Value: 1}, {Key: "values", Value: bson.A{"alice"}}})
				if method == "Distinct" {
					err = f.Distinct(context.Background(), "name").Err()
				} else {
					var values []string
					err = f.DistinctWithParse(context.Background(), "name", &values)
					require.Equal(t, []string{"alice"}, values)
				}
			case "FindOneAndUpdate":
				deployment.AddResponses(bson.D{{Key: "ok", Value: 1}, {Key: "value", Value: bson.M{"name": "alice"}}})
				_, err = f.Updates(bson.M{"$set": bson.M{"name": "alice"}}).FindOneAndUpdate(context.Background())
			}
			require.NoError(t, err)
			require.Len(t, *commands, 1)
			command := (*commands)[0]
			switch method {
			case "Count":
				requireLiveFilter(t, command.Lookup("pipeline").Array().Index(0).Document().Lookup("$match").Document())
			case "Distinct", "DistinctWithParse", "FindOneAndUpdate":
				requireLiveFilter(t, command.Lookup("query").Document())
			default:
				requireLiveFilter(t, command.Lookup("filter").Document())
			}
			require.Equal(t, bson.M{"name": "alice"}, filter)
		})
	}
}

func TestSoftDeleteUpdateCommands(t *testing.T) {
	for _, method := range []string{"UpdateOne", "UpdateMany", "Upsert"} {
		t.Run(method, func(t *testing.T) {
			db, deployment, commands := softDeleteDatabase(t)
			col := NewCollection[softDeleteDoc](db, "docs")
			filter := bson.D{{Key: "name", Value: "alice"}}
			u := col.Updater().Filter(filter).Updates(bson.M{"$set": bson.M{"name": "alice"}})
			deployment.AddResponses(writeResponse(1))
			var err error
			switch method {
			case "UpdateOne":
				_, err = u.UpdateOne(context.Background())
			case "UpdateMany":
				_, err = u.UpdateMany(context.Background())
			case "Upsert":
				_, err = u.Upsert(context.Background())
			}
			require.NoError(t, err)
			update := (*commands)[0].Lookup("updates").Array().Index(0).Document()
			requireLiveFilter(t, update.Lookup("q").Document())
			require.Equal(t, bson.TypeDateTime, update.Lookup("u", "$set", "updated_at").Type)
			require.Equal(t, bson.Type(0), update.Lookup("u", "$set", "deleted_at").Type)
			if method == "Upsert" {
				require.True(t, update.Lookup("upsert").Boolean())
			}
			require.Equal(t, bson.D{{Key: "name", Value: "alice"}}, filter)
		})
	}
}

func TestPhysicalDeleteWithoutDeletedAt(t *testing.T) {
	for _, many := range []bool{false, true} {
		db, deployment, commands := softDeleteDatabase(t)
		col := NewCollection[struct {
			Name string `bson:"name"`
		}](db, "docs")
		deployment.AddResponses(writeResponse(1))
		db.RegisterPlugin("supply-updates", func(_ context.Context, op *operation.OpContext, _ ...any) error {
			require.False(t, op.SoftDelete)
			op.Updates = bson.M{"$set": bson.M{"from": "database-hook"}}
			return nil
		}, operation.OpTypeBeforeDelete)
		d := col.Deleter().Filter(bson.M{"name": "alice"}).RegisterBeforeHooks(func(_ context.Context, op *deleter.OpContext, _ ...any) error {
			require.NotNil(t, op.Updates)
			op.Updates = bson.M{"$set": bson.M{"from": "local-hook"}}
			return nil
		})
		var err error
		if many {
			_, err = d.UpdateManyOptions(&options.UpdateManyOptionsBuilder{Opts: []func(*options.UpdateManyOptions) error{
				func(*options.UpdateManyOptions) error { return errors.New("unused update options") },
			}}).DeleteMany(context.Background(), options.DeleteMany().SetComment("physical-delete"))
		} else {
			_, err = d.UpdateOneOptions(&options.UpdateOneOptionsBuilder{Opts: []func(*options.UpdateOneOptions) error{
				func(*options.UpdateOneOptions) error { return errors.New("unused update options") },
			}}).DeleteOne(context.Background(), options.DeleteOne().SetComment("physical-delete"))
		}
		require.NoError(t, err)
		require.Equal(t, "docs", (*commands)[0].Lookup("delete").StringValue())
		require.Equal(t, "physical-delete", (*commands)[0].Lookup("comment").StringValue())
		require.Equal(t, "alice", (*commands)[0].Lookup("deletes").Array().Index(0).Document().Lookup("q", "name").StringValue())
	}
}

func TestSoftDeleteWithClearedUpdates(t *testing.T) {
	for _, method := range []string{"DeleteOne", "DeleteMany"} {
		for _, hook := range []string{"database", "local"} {
			t.Run(method+"/"+hook, func(t *testing.T) {
				db, deployment, commands := softDeleteDatabase(t)
				col := NewCollection[softDeleteDoc](db, "docs")
				// A physical-delete regression would succeed and be visible in commands.
				deployment.AddResponses(writeResponse(1))
				db.RegisterPlugin("clear-updates", func(_ context.Context, op *operation.OpContext, _ ...any) error {
					require.True(t, op.SoftDelete)
					require.NotNil(t, op.Updates)
					if hook == "database" {
						op.Updates = nil
					}
					return nil
				}, operation.OpTypeBeforeDelete)
				afterCalled := false
				d := col.Deleter().Filter(bson.M{"name": "alice"}).RegisterBeforeHooks(func(_ context.Context, op *deleter.OpContext, _ ...any) error {
					if hook == "local" {
						require.NotNil(t, op.Updates)
						op.Updates = nil
					} else {
						require.Nil(t, op.Updates)
					}
					return nil
				}).RegisterAfterHooks(func(context.Context, *deleter.OpContext, ...any) error {
					afterCalled = true
					return nil
				})
				var result *mongo.DeleteResult
				var err error
				if method == "DeleteOne" {
					result, err = d.DeleteOne(context.Background())
				} else {
					result, err = d.DeleteMany(context.Background())
				}
				require.ErrorIs(t, err, mongo.ErrNilDocument)
				require.Nil(t, result)
				require.Empty(t, *commands, "nil updates must be rejected without sending a delete command")
				require.False(t, afterCalled)
			})
		}
	}
}

func TestSoftDeleteOptionsEvaluatedByDriver(t *testing.T) {
	for _, method := range []string{"DeleteOne", "DeleteMany"} {
		t.Run(method, func(t *testing.T) {
			db, deployment, commands := softDeleteDatabase(t)
			col := NewCollection[softDeleteDoc](db, "docs")
			calls := 0
			comment := "first"
			var setterErr error
			var run func(any) (*mongo.DeleteResult, error)
			if method == "DeleteOne" {
				opts := &options.UpdateOneOptionsBuilder{Opts: []func(*options.UpdateOneOptions) error{
					func(update *options.UpdateOneOptions) error {
						calls++
						update.Comment = comment
						return setterErr
					},
				}}
				d := col.Deleter().UpdateOneOptions(opts)
				require.Zero(t, calls, "configuration must only store the options")
				run = func(filter any) (*mongo.DeleteResult, error) {
					return d.Filter(filter).DeleteOne(context.Background(), &options.DeleteOneOptionsBuilder{Opts: []func(*options.DeleteOneOptions) error{
						func(*options.DeleteOneOptions) error { return errors.New("unused delete options") },
					}})
				}
			} else {
				opts := &options.UpdateManyOptionsBuilder{Opts: []func(*options.UpdateManyOptions) error{
					func(update *options.UpdateManyOptions) error {
						calls++
						update.Comment = comment
						return setterErr
					},
				}}
				d := col.Deleter().UpdateManyOptions(opts)
				require.Zero(t, calls, "configuration must only store the options")
				run = func(filter any) (*mongo.DeleteResult, error) {
					return d.Filter(filter).DeleteMany(context.Background(), &options.DeleteManyOptionsBuilder{Opts: []func(*options.DeleteManyOptions) error{
						func(*options.DeleteManyOptions) error { return errors.New("unused delete options") },
					}})
				}
			}
			// The driver marshals the filter before evaluating option setters.
			_, err := run(bson.M{"invalid": make(chan int)})
			require.Error(t, err)
			require.Zero(t, calls)
			require.Empty(t, *commands)
			for i, value := range []string{"first", "second"} {
				comment = value
				deployment.AddResponses(writeResponse(1))
				_, err = run(bson.M{})
				require.NoError(t, err)
				require.Equal(t, i+1, calls, "evaluate setters once per operation")
				require.Equal(t, value, (*commands)[i].Lookup("comment").StringValue())
			}
			setterErr = errors.New("invalid update option")
			_, err = run(bson.M{})
			require.ErrorIs(t, err, setterErr)
			require.Equal(t, 3, calls)
			require.Len(t, *commands, 2, "option errors must prevent command execution")
		})
	}
}

func TestSoftDeleteUpdateOptionsReplaceAndClear(t *testing.T) {
	for _, method := range []string{"DeleteOne", "DeleteMany"} {
		t.Run(method, func(t *testing.T) {
			db, deployment, commands := softDeleteDatabase(t)
			col := NewCollection[softDeleteDoc](db, "docs")
			d := col.Deleter().Filter(bson.M{})
			var run func() (*mongo.DeleteResult, error)
			var clear func()
			if method == "DeleteOne" {
				d.UpdateManyOptions(&options.UpdateManyOptionsBuilder{Opts: []func(*options.UpdateManyOptions) error{
					func(*options.UpdateManyOptions) error { return errors.New("unused many options") },
				}})
				d.UpdateOneOptions(options.UpdateOne().SetHint("old_index"))
				d.UpdateOneOptions(options.UpdateOne().SetComment("replacement"))
				run = func() (*mongo.DeleteResult, error) { return d.DeleteOne(context.Background()) }
				clear = func() { d.UpdateOneOptions() }
			} else {
				d.UpdateOneOptions(&options.UpdateOneOptionsBuilder{Opts: []func(*options.UpdateOneOptions) error{
					func(*options.UpdateOneOptions) error { return errors.New("unused one options") },
				}})
				d.UpdateManyOptions(options.UpdateMany().SetHint("old_index"))
				d.UpdateManyOptions(options.UpdateMany().SetComment("replacement"))
				run = func() (*mongo.DeleteResult, error) { return d.DeleteMany(context.Background()) }
				clear = func() { d.UpdateManyOptions() }
			}
			deployment.AddResponses(writeResponse(1), writeResponse(1))
			_, err := run()
			require.NoError(t, err)
			require.Equal(t, "replacement", (*commands)[0].Lookup("comment").StringValue())
			require.Zero(t, (*commands)[0].Lookup("updates").Array().Index(0).Document().Lookup("hint").Type)
			clear()
			_, err = run()
			require.NoError(t, err)
			require.Zero(t, (*commands)[1].Lookup("comment").Type)
		})
	}
}

func TestSoftDeleteErrors(t *testing.T) {
	db, deployment, commands := softDeleteDatabase(t)
	col := NewCollection[softDeleteDoc](db, "docs")
	_, err := col.Deleter().DeleteMany(context.Background())
	require.ErrorIs(t, err, mongo.ErrNilDocument)
	require.Empty(t, *commands)
	wantErr := errors.New("invalid update options")
	_, err = col.Deleter().Filter(bson.M{}).UpdateOneOptions(&options.UpdateOneOptionsBuilder{Opts: []func(*options.UpdateOneOptions) error{
		func(*options.UpdateOneOptions) error { return wantErr },
	}}).DeleteOne(context.Background())
	require.ErrorIs(t, err, wantErr)
	require.Empty(t, *commands)
	deployment.AddResponses(bson.D{{Key: "ok", Value: 0}, {Key: "code", Value: int32(121)}, {Key: "errmsg", Value: "validation failed"}})
	afterCalled := false
	result, err := col.Deleter().Filter(bson.M{}).RegisterAfterHooks(func(context.Context, *deleter.OpContext, ...any) error {
		afterCalled = true
		return nil
	}).DeleteMany(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.False(t, afterCalled)
}

func TestNilFilterHandledByDriver(t *testing.T) {
	var filter any
	db, _, commands := softDeleteDatabase(t)
	col := NewCollection[softDeleteDoc](db, "docs")
	ctx := context.Background()
	f := col.Finder().Filter(filter)
	_, err := f.FindOne(ctx)
	require.ErrorIs(t, err, mongo.ErrNilDocument)
	_, err = f.Find(ctx)
	require.ErrorIs(t, err, mongo.ErrNilDocument)
	_, err = f.Count(ctx)
	require.ErrorIs(t, err, mongo.ErrNilDocument)
	require.ErrorIs(t, f.Distinct(ctx, "name").Err(), mongo.ErrNilDocument)
	var names []string
	require.ErrorIs(t, f.DistinctWithParse(ctx, "name", &names), mongo.ErrNilDocument)
	updates := bson.M{"$set": bson.M{"name": "alice"}}
	_, err = f.Updates(updates).FindOneAndUpdate(ctx)
	require.ErrorIs(t, err, mongo.ErrNilDocument)
	u := col.Updater().Filter(filter).Updates(updates)
	_, err = u.UpdateOne(ctx)
	require.ErrorIs(t, err, mongo.ErrNilDocument)
	_, err = u.UpdateMany(ctx)
	require.ErrorIs(t, err, mongo.ErrNilDocument)
	_, err = u.Upsert(ctx)
	require.ErrorIs(t, err, mongo.ErrNilDocument)
	d := col.Deleter().Filter(filter)
	_, err = d.DeleteOne(ctx)
	require.ErrorIs(t, err, mongo.ErrNilDocument)
	_, err = d.DeleteMany(ctx)
	require.ErrorIs(t, err, mongo.ErrNilDocument)
	require.Empty(t, *commands, "the driver rejects an untyped nil filter before sending a command")
}

func TestSoftDeleteFilterReplacementAndReuse(t *testing.T) {
	db, deployment, commands := softDeleteDatabase(t)
	col := NewCollection[softDeleteDoc](db, "docs")
	original := bson.M{"name": "alice"}
	f := col.Finder().Filter(original).RegisterBeforeHooks(func(_ context.Context, op *finder.OpContext[softDeleteDoc], _ ...any) error {
		// A local hook replaces the top-level filter while keeping its live scope.
		op.Filter = bson.D{{Key: "$and", Value: bson.A{op.Filter, bson.M{"tenant": "test"}}}}
		return nil
	})
	deployment.AddResponses(cursorResponse(bson.M{"name": "alice"}), cursorResponse(bson.M{"name": "alice"}))
	for range 2 {
		_, err := f.FindOne(context.Background())
		require.NoError(t, err)
	}
	require.Equal(t, (*commands)[0].Lookup("filter").Document(), (*commands)[1].Lookup("filter").Document())
	require.Equal(t, bson.M{"name": "alice"}, original)
	clauses := (*commands)[0].Lookup("filter", "$and").Array()
	requireLiveFilter(t, clauses.Index(0).Document())
	require.Equal(t, "test", clauses.Index(1).Document().Lookup("tenant").StringValue())
}
