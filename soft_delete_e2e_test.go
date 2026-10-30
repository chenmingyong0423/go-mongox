//go:build e2e

package mongox

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func TestSoftDeleteE2E(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username: "test", Password: "test", AuthSource: "db-test",
	}))
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, client.Disconnect(context.Background())) })
	db := NewClient(client, &Config{}).NewDatabase("db-test")
	col := NewCollection[softDeleteDoc](db, "soft_delete_"+bson.NewObjectID().Hex())
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cleanupCancel()
		require.NoError(t, col.Collection().Drop(cleanupCtx))
	})
	_, err = col.Collection().InsertMany(ctx, []any{
		bson.M{"name": "missing"},
		bson.M{"name": "null", "deleted_at": nil},
		bson.M{"name": "zero", "deleted_at": time.Time{}},
		bson.M{"name": "deleted", "deleted_at": time.Now()},
	})
	require.NoError(t, err)
	count, err := col.Finder().Count(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(3), count)
	var names []string
	require.NoError(t, col.Finder().DistinctWithParse(ctx, "name", &names))
	require.ElementsMatch(t, []string{"missing", "null", "zero"}, names)
	_, err = col.Finder().Filter(bson.M{"name": "deleted"}).FindOne(ctx)
	require.ErrorIs(t, err, mongo.ErrNoDocuments)
	updated, err := col.Updater().Filter(bson.M{}).Updates(bson.M{"$set": bson.M{"touched": true}}).UpdateMany(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(3), updated.MatchedCount)
	_, err = col.Finder().Filter(bson.M{"name": "deleted"}).Updates(bson.M{"$set": bson.M{"name": "changed"}}).FindOneAndUpdate(ctx)
	require.ErrorIs(t, err, mongo.ErrNoDocuments)

	// Finder scopes apply consistently to reads, counts, distinct and updates.
	count, err = col.Finder().WithDeleted().Count(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(4), count)
	count, err = col.Finder().OnlyDeleted().Count(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(1), count)
	all, err := col.Finder().WithDeleted().Find(ctx)
	require.NoError(t, err)
	require.Len(t, all, 4)
	deleted, err := col.Finder().OnlyDeleted().Find(ctx)
	require.NoError(t, err)
	require.Len(t, deleted, 1)
	require.Equal(t, "deleted", deleted[0].Name)
	included, err := col.Finder().WithDeleted().Filter(bson.M{"name": "deleted"}).FindOne(ctx)
	require.NoError(t, err)
	require.Equal(t, "deleted", included.Name)
	onlyDeleted, err := col.Finder().OnlyDeleted().FindOne(ctx)
	require.NoError(t, err)
	require.Equal(t, "deleted", onlyDeleted.Name)
	_, err = col.Finder().OnlyDeleted().Filter(bson.M{"name": "missing"}).FindOne(ctx)
	require.ErrorIs(t, err, mongo.ErrNoDocuments)
	var deletedNames []string
	require.NoError(t, col.Finder().OnlyDeleted().DistinctWithParse(ctx, "name", &deletedNames))
	require.Equal(t, []string{"deleted"}, deletedNames)
	var allNames []string
	require.NoError(t, col.Finder().WithDeleted().DistinctWithParse(ctx, "name", &allNames))
	require.ElementsMatch(t, []string{"missing", "null", "zero", "deleted"}, allNames)
	scopedUpdate := bson.M{"$set": bson.M{"scope_checked": true}}
	_, err = col.Finder().OnlyDeleted().Filter(bson.M{"name": "deleted"}).Updates(scopedUpdate).FindOneAndUpdate(ctx)
	require.NoError(t, err)
	_, err = col.Finder().WithDeleted().Filter(bson.M{"name": "deleted"}).Updates(scopedUpdate).FindOneAndUpdate(ctx)
	require.NoError(t, err)

	var beforeDelete softDeleteDoc
	require.NoError(t, col.Collection().FindOne(ctx, bson.M{"name": "missing"}).Decode(&beforeDelete))
	d := col.Deleter().Filter(bson.M{"name": "missing"})
	result, err := d.DeleteOne(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(1), result.DeletedCount)
	var stored softDeleteDoc
	require.NoError(t, col.Collection().FindOne(ctx, bson.M{"name": "missing"}).Decode(&stored))
	require.False(t, stored.DeletedAt.IsZero())
	require.Equal(t, beforeDelete.UpdatedAt, stored.UpdatedAt)
	result, err = d.DeleteOne(ctx)
	require.NoError(t, err)
	require.Zero(t, result.DeletedCount)
	var repeated softDeleteDoc
	require.NoError(t, col.Collection().FindOne(ctx, bson.M{"name": "missing"}).Decode(&repeated))
	require.Equal(t, stored.DeletedAt, repeated.DeletedAt)
	count, err = col.Collection().CountDocuments(ctx, bson.M{})
	require.NoError(t, err)
	require.Equal(t, int64(4), count, "soft deletion must retain stored documents")

	// Upsert creates a live document; it does not revive the deleted one.
	updated, err = col.Updater().Filter(bson.M{"name": "missing"}).Updates(bson.M{"$set": bson.M{"name": "missing"}}).Upsert(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(1), updated.UpsertedCount)
	count, err = col.Finder().Filter(bson.M{"name": "missing"}).Count(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(1), count)
	result, err = col.Deleter().Filter(bson.M{}).DeleteMany(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(3), result.DeletedCount)
	count, err = col.Finder().Count(ctx)
	require.NoError(t, err)
	require.Zero(t, count)

	// A timestamp model uses integer zero rather than the BSON date zero.
	type TimestampDoc struct {
		DeletedAt int64 `bson:"deleted_at,omitempty" mongox:"autoDeleteTime:milli"`
	}
	integerCol := NewCollection[TimestampDoc](db, col.Collection().Name()+"_integer")
	t.Cleanup(func() { require.NoError(t, integerCol.Collection().Drop(context.Background())) })
	_, err = integerCol.Collection().InsertMany(ctx, []any{bson.M{}, bson.M{"deleted_at": nil}, bson.M{"deleted_at": int64(0)}, bson.M{"deleted_at": int64(123)}})
	require.NoError(t, err)
	count, err = integerCol.Finder().Count(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(3), count)
	result, err = integerCol.Deleter().Filter(bson.M{}).DeleteMany(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(3), result.DeletedCount)
	count, err = integerCol.Finder().OnlyDeleted().Count(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(4), count)
	restored, err := integerCol.Finder().OnlyDeleted().Filter(bson.M{"deleted_at": int64(123)}).
		Updates(bson.M{"$set": bson.M{"deleted_at": int64(0)}}).
		FindOneAndUpdate(ctx, options.FindOneAndUpdate().SetReturnDocument(options.After))
	require.NoError(t, err)
	require.Zero(t, restored.DeletedAt)
	count, err = integerCol.Finder().Count(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(1), count)
	count, err = integerCol.Finder().OnlyDeleted().Count(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(3), count)

	// Models without a deletion field have no deleted documents.
	plainCol := NewCollection[struct{ Name string }](db, col.Collection().Name()+"_plain")
	t.Cleanup(func() { require.NoError(t, plainCol.Collection().Drop(context.Background())) })
	_, err = plainCol.Collection().InsertOne(ctx, bson.M{"name": "plain"})
	require.NoError(t, err)
	count, err = plainCol.Finder().WithDeleted().Count(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(1), count)
	count, err = plainCol.Finder().OnlyDeleted().Count(ctx)
	require.NoError(t, err)
	require.Zero(t, count)
}
