package transaction

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ITransaction interface {
	Auto(ctx context.Context, fn func(ctx context.Context) (interface{}, error)) (interface{}, error)
	Begin(ctx context.Context) (*Tx, error)
}

type ITx interface {
	RollBack() error
	Commit() error
}

var (
	_ ITransaction = (*Transaction)(nil)
	_ ITx          = (*Tx)(nil)
)

type Transaction struct {
	collection *mongo.Collection
	session    *mongo.Session
}
type Tx struct {
	SessionCtx context.Context
	session    *mongo.Session
}

func NewSession(collection *mongo.Collection) (*mongo.Session, error) {
	sess, err := collection.Database().Client().StartSession()
	if err != nil {
		return nil, errors.New("start session failed")
	}
	return sess, nil
}

func NewTransaction(collection *mongo.Collection) *Transaction {
	sess, _ := NewSession(collection)
	return &Transaction{
		collection: collection,
		session:    sess,
	}
}

func NewTx(ctx context.Context, session *mongo.Session) *Tx {
	return &Tx{
		SessionCtx: ctx,
		session:    session,
	}
}

func (t *Transaction) Auto(ctx context.Context, fn func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	// If any error occurs, the transaction will be rolled back
	defer t.session.EndSession(ctx)
	if err := t.session.StartTransaction(); err != nil {
		return nil, err
	}
	var result interface{}
	err := mongo.WithSession(ctx, t.session, func(sc context.Context) error {
		r, err := fn(sc)
		if err != nil {
			t.session.AbortTransaction(ctx)
			return err
		}
		result = r
		return t.session.CommitTransaction(ctx)
	})

	if err != nil {
		t.session.AbortTransaction(ctx)
		return nil, err
	}

	return result, nil
}

func (t *Transaction) Begin(ctx context.Context) (*Tx, error) {
	return NewTx(mongo.NewSessionContext(ctx, t.session), t.session), t.session.StartTransaction()
}

func (tx *Tx) RollBack() error {
	// Ensure that even if SessionCtx has timed out, AbortTransaction can still be executed successfully
	return tx.session.AbortTransaction(context.TODO())
}

func (tx *Tx) Commit() error {
	// Ensure that even if SessionCtx has timed out, CommitTransaction can still be executed successfully
	return tx.session.CommitTransaction(context.TODO())
}
